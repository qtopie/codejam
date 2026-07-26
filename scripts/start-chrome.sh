#!/bin/bash

set -euo pipefail

OS_NAME="${OS_NAME:-$(uname -s)}"
PORT="${PORT:-9222}"
WORKSPACE_DIR="${WORKSPACE_DIR:-$HOME/workspace/.copilot}"
USER_DATA_DIR="${USER_DATA_DIR:-$WORKSPACE_DIR/data/chrome-debug-profile}"
LOG_DIR="${LOG_DIR:-$WORKSPACE_DIR/logs}"
LOG_FILE="$LOG_DIR/chrome-debug.log"

resolve_macos_app_bin() {
  local app_path="$1"

  case "$app_path" in
    *.app)
      app_path="$app_path/Contents/MacOS/$(basename "$app_path" .app)"
      ;;
  esac

  if [[ -x "$app_path" ]]; then
    printf '%s\n' "$app_path"
    return 0
  fi

  return 1
}

resolve_chrome_bin() {
  local candidate

  if [[ -n "${CHROME_BIN:-}" ]]; then
    case "$OS_NAME" in
      Darwin)
        if resolve_macos_app_bin "$CHROME_BIN"; then
          return 0
        fi
        ;;
      *)
        if [[ -x "$CHROME_BIN" ]]; then
          printf '%s\n' "$CHROME_BIN"
          return 0
        fi
        ;;
    esac

    echo "CHROME_BIN does not point to an executable: $CHROME_BIN" >&2
    return 1
  fi

  case "$OS_NAME" in
    Darwin)
      for candidate in \
        "/Applications/Google Chrome.app" \
        "$HOME/Applications/Google Chrome.app" \
        "/Applications/Chromium.app" \
        "$HOME/Applications/Chromium.app" \
        "/Applications/Microsoft Edge.app" \
        "$HOME/Applications/Microsoft Edge.app"
      do
        if resolve_macos_app_bin "$candidate"; then
          return 0
        fi
      done
      ;;
    Linux)
      for candidate in \
        google-chrome \
        google-chrome-stable \
        chromium \
        chromium-browser \
        chrome \
        microsoft-edge \
        microsoft-edge-stable
      do
        if command -v "$candidate" >/dev/null 2>&1; then
          command -v "$candidate"
          return 0
        fi
      done

      for candidate in \
        /snap/bin/chromium \
        /usr/bin/google-chrome \
        /usr/bin/chromium \
        /usr/bin/microsoft-edge \
        /usr/bin/microsoft-edge-stable
      do
        if [[ -x "$candidate" ]]; then
          printf '%s\n' "$candidate"
          return 0
        fi
      done
      ;;
    *)
      echo "Unsupported OS: $OS_NAME" >&2
      return 1
      ;;
  esac

  echo "Could not find Chrome, Chromium, or Microsoft Edge for $OS_NAME. Set CHROME_BIN to the browser executable." >&2
  return 1
}

CHROME_BIN="$(resolve_chrome_bin)"

mkdir -p "$USER_DATA_DIR" "$LOG_DIR"

"$CHROME_BIN" \
  --remote-debugging-port="$PORT" \
  --user-data-dir="$USER_DATA_DIR" \
  --no-first-run \
  --no-default-browser-check \
  "about:blank" >>"$LOG_FILE" 2>&1 &

echo "Browser started on port $PORT"
echo "Binary: $CHROME_BIN"
echo "Profile: $USER_DATA_DIR"
echo "Log: $LOG_FILE"
