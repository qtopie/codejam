# 字符串

- **模式匹配**: KMP、Z-algorithm、Boyer–Moore、Rabin–Karp（滚动哈希）。
- **字典结构**: Trie、Aho–Corasick（多模式）。
- **后缀结构**: 后缀数组/树、LCP、后缀自动机。
- **其他经典**: Manacher（回文）、最小表示法。

## 现有代码实现

| 文件 | 说明 |
|------|------|
| [`calc-string-distance.go`](calc-string-distance.go) | 字符串编辑距离计算（递归/DP） |
| [`str-shift.go`](str-shift.go) | 字符串旋转/移位包含判断 |
| [`replace-space.go`](replace-space.go) | 空格替换为 `%20`（原地双指针思想） |
