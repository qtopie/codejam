#!/usr/bin/env node

import { mkdir, readFile, writeFile, chmod } from "node:fs/promises";
import os from "node:os";
import path from "node:path";
import { stdin, stdout } from "node:process";
import readline from "node:readline";

const LEETCODE_TOML_PATH = path.join(os.homedir(), ".leetcode", "leetcode.toml");
const COOKIES_JSON_PATH = path.join(os.homedir(), ".leetcode", "leetcode-cookies.json");

function parseArgs(argv) {
  const result = {
    site: "leetcode.com",
    output: COOKIES_JSON_PATH,
  };

  for (let i = 0; i < argv.length; i += 1) {
    const arg = argv[i];
    if (arg === "--site") {
      result.site = argv[i + 1];
      i += 1;
      continue;
    }
    if (arg === "--output") {
      result.output = argv[i + 1];
      i += 1;
      continue;
    }
    if (arg === "--help" || arg === "-h") {
      result.help = true;
      continue;
    }
  }

  return result;
}

function printHelp() {
  console.log(`Usage: node scripts/parse-leetcode-curl.mjs [--site leetcode.com|leetcode.cn] [--output <path>]

Parses a curl command (from browser DevTools "Copy as cURL").
If you pipe data to it, it processes immediately. 
Otherwise, it will prompt you to paste the command.

Examples:
  pbpaste | node scripts/parse-leetcode-curl.mjs
  node scripts/parse-leetcode-curl.mjs  # Will prompt for input
`);
}

async function readInput() {
  if (!stdin.isTTY) {
    // Piped input
    let result = "";
    for await (const chunk of stdin) {
      result += chunk;
    }
    return result;
  }

  // Interactive input
  return new Promise((resolve) => {
    const rl = readline.createInterface({
      input: stdin,
      output: stdout,
      terminal: true,
    });

    console.log("\x1b[36m%s\x1b[0m", "--- LeetCode Cookie Parser ---");
    console.log("Please paste the cURL command from your browser below.");
    console.log("Tip: After pasting, if it doesn't finish automatically, press Enter twice or Ctrl+D.");
    console.log("----------------------------");

    let buffer = [];
    
    rl.on("line", (line) => {
      buffer.push(line);
      
      // Heuristic: If we have content and the line doesn't end with a backslash,
      // it's likely the end of a standard "Copy as cURL" command.
      const trimmed = line.trim();
      if (buffer.length > 0 && trimmed.length > 0 && !trimmed.endsWith("\\")) {
        // Wait a tiny bit to make sure no more lines are coming (in case of slow paste)
        setTimeout(() => {
          rl.close();
        }, 100);
      }
    });

    rl.on("close", () => {
      resolve(buffer.join("\n"));
    });
  });
}

function extractCookies(curlCommand) {
  const cookies = {};
  
  // Try to find cookies in -b or --cookie
  const cookieMatch = curlCommand.match(/(?:-b|--cookie)\s+(['"])(.*?)\1/);
  if (cookieMatch) {
    const cookieStr = cookieMatch[2];
    cookieStr.split(";").forEach(part => {
      const [key, ...valueParts] = part.trim().split("=");
      if (key && valueParts.length > 0) {
        cookies[key.trim()] = valueParts.join("=").trim();
      }
    });
  }

  // Try to find cookies in -H 'cookie: ...'
  const headerCookieMatches = curlCommand.matchAll(/-H\s+(['"])cookie:\s*(.*?)\1/gi);
  for (const match of headerCookieMatches) {
    const cookieStr = match[2];
    cookieStr.split(";").forEach(part => {
      const [key, ...valueParts] = part.trim().split("=");
      if (key && valueParts.length > 0) {
        cookies[key.trim()] = valueParts.join("=").trim();
      }
    });
  }

  // Try to find x-csrftoken header
  const csrfMatch = curlCommand.match(/-H\s+(['"])x-csrftoken:\s*(.*?)\1/i);
  if (csrfMatch) {
    cookies["csrftoken"] = csrfMatch[2];
  }

  return cookies;
}

// Logic for updating leetcode.toml (reused from get-leetcode-cookies.mjs)
function buildTomlSnippet({ csrf, session, site }) {
  return [
    "[cookies]",
    `csrf = "${csrf}"`,
    `session = "${session}"`,
    `site = "${site}"`,
  ].join("\n");
}

function upsertCookiesKeys(sectionBody, keyValues) {
  const lines = sectionBody ? sectionBody.split("\n") : [];
  const keyOrder = ["csrf", "session", "site"];
  const seenKeys = new Set();
  const updatedLines = [];

  for (const line of lines) {
    const match = line.match(/^(\s*)(csrf|session|site)\s*=/);
    if (!match) {
      updatedLines.push(line);
      continue;
    }

    const key = match[2];
    if (seenKeys.has(key)) continue;
    seenKeys.add(key);
    updatedLines.push(`${match[1]}${key} = "${keyValues[key]}"`);
  }

  for (const key of keyOrder) {
    if (!seenKeys.has(key)) {
      updatedLines.push(`${key} = "${keyValues[key]}"`);
    }
  }

  return updatedLines.join("\n").replace(/^\n+|\n+$/g, "");
}

function replaceCookiesSection(existingToml, payload) {
  const keyValues = { csrf: payload.csrf, session: payload.session, site: payload.site };
  const cookiesSectionRegex = /((?:^|\n)\[cookies\][^\n]*\n?)([\s\S]*?)(?=\n\s*\[|$)/;
  const match = existingToml.match(cookiesSectionRegex);

  if (match) {
    const fullMatch = match[0];
    const header = match[1];
    const body = match[2];
    const updatedBody = upsertCookiesKeys(body, keyValues);
    return existingToml.replace(fullMatch, `${header}${updatedBody}\n`);
  }

  const appendedSection = `${buildTomlSnippet(payload)}\n`;
  return existingToml.trim() ? `${existingToml.trimEnd()}\n\n${appendedSection}` : appendedSection;
}

async function updateLeetCodeToml(payload) {
  let existingToml = "";
  try {
    existingToml = await readFile(LEETCODE_TOML_PATH, "utf8");
  } catch (error) {
    if (error.code !== "ENOENT") throw error;
  }
  const updatedToml = replaceCookiesSection(existingToml, payload);
  await writeFile(LEETCODE_TOML_PATH, updatedToml, { mode: 0o600 });
}

async function main() {
  const args = parseArgs(process.argv.slice(2));
  if (args.help) {
    printHelp();
    return;
  }

  const curlCommand = await readInput();
  if (!curlCommand.trim()) {
    throw new Error("Empty input. Please provide a curl command.");
  }

  const cookies = extractCookies(curlCommand);
  const csrf = cookies["csrftoken"];
  const session = cookies["LEETCODE_SESSION"];

  if (!csrf || !session) {
    console.error("Available cookies found:", Object.keys(cookies).join(", "));
    if (!csrf) console.error("Error: Could not find 'csrftoken'");
    if (!session) console.error("Error: Could not find 'LEETCODE_SESSION'");
    process.exit(1);
  }

  const payload = {
    site: args.site,
    csrf,
    session,
    capturedAt: new Date().toISOString(),
  };

  await mkdir(path.dirname(args.output), { recursive: true });
  await writeFile(args.output, `${JSON.stringify(payload, null, 2)}\n`, { mode: 0o600 });
  await chmod(args.output, 0o600);

  await updateLeetCodeToml(payload);

  console.log(`Successfully extracted cookies!`);
  console.log(`CSRF: ${csrf.substring(0, 8)}...`);
  console.log(`Session: ${session.substring(0, 8)}...`);
  console.log(`Saved to ${args.output} and ${LEETCODE_TOML_PATH}`);
}

main().catch(err => {
  console.error(err.message);
  process.exit(1);
});
