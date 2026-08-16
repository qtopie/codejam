# 我的算法练习记录 (CodeJam)

算法与数据结构学习仓库，包含 **LeetCode 刷题记录**、**经典算法专题练习** 与 **编程之美（Joy of Coding）** 练习。

## 📌 内容索引

### LeetCode 刷题
| 目录 | 说明 |
|------|------|
| [`leetcode/`](leetcode/) | LeetCode 题解（前 200 题） |
| [`310.minimum-height-trees/`](310.minimum-height-trees/) | LeetCode 310. Minimum Height Trees 题解 |

### 算法知识库（[经典算法知识纲要](https://qtopie.github.io/notes/codejam/classic-algorithms/)）
| 目录 | 说明 |
|------|------|
| [`algorithms/core-concepts/`](algorithms/core-concepts/) | 复杂度分析、主定理、正确性证明 |
| [`algorithms/data-structures/`](algorithms/data-structures/) | 数组/链表、栈/队列、哈希、树/堆、并查集 |
| [`algorithms/paradigms/`](algorithms/paradigms/) | 分治、DP、贪心、回溯、滑动窗口、前缀和、单调栈、排序 |
| [`algorithms/graph/`](algorithms/graph/) | 图论：最短路、MST、拓扑、流与匹配 |
| [`algorithms/string/`](algorithms/string/) | 字符串：KMP、Trie、后缀结构 |
| [`algorithms/math/`](algorithms/math/) | 数论、位运算与计算几何 |
| [`algorithms/dp-optimization/`](algorithms/dp-optimization/) | DP 优化（进阶） |

### 文档
| 文件 | 说明 |
|------|------|
| [`docs/binary-search.md`](docs/binary-search.md) | LeetCode 前 200 题 Binary Search 题单 |
| [`docs/linked-list.md`](docs/linked-list.md) | LeetCode 经典 Linked List 题单与核心解题技巧总结 |
| [`docs/be-a-good-engineer.md`](docs/be-a-good-engineer.md) | 软件工程师能力模型与面试准备脉络 |
| [`docs/system-design.md`](docs/system-design.md) | 系统设计 |
| [`docs/project-layout.md`](docs/project-layout.md) | 项目结构规范 |
| [`docs/references/code-conventions.md`](docs/references/code-conventions.md) | Go 编码规范 |
| [`docs/testing/`](docs/testing/) | 测试规范与 Harness 工程指南 |
| [`docs/rfcs/`](docs/rfcs/) · [`docs/bugs/`](docs/bugs/) | RFC 与 Bug RCA 模板 |

### 工程化脚手架
| 目录/文件 | 说明 |
|------|------|
| [`AGENTS.md`](AGENTS.md) | Agent 协作规则总控（Spec-First + Harness 门禁） |
| `specs/` · `harness/` · `testings/` | Spec 驱动开发 (SDD) 与 Harness 测试工程 |
| [`scripts/`](scripts/) | 校验工具链（`check.sh` / `check-harness.sh` 等） |

## 🚀 快速开始

```bash
# 运行 Harness 校验与测试
./scripts/check.sh

# 运行全部 Go 单元测试
go test -v -race ./...
```

## 📚 基础知识

> 待补充：算法导论笔记、常见数据结构总结。
