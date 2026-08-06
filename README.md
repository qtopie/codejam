# 我的算法练习记录 (CodeJam)

算法与数据结构学习仓库，包含 **LeetCode 刷题记录**、**经典算法专题练习** 与 **编程之美（Joy of Coding）** 练习。

## 📌 内容索引

### LeetCode 刷题
| 目录 | 说明 |
|------|------|
| [`leetcode/`](leetcode/) | LeetCode 题解（前 200 题） |
| [`310.minimum-height-trees/`](310.minimum-height-trees/) | LeetCode 310. Minimum Height Trees 题解 |

### 算法专题
| 目录 | 说明 |
|------|------|
| [`backtracking/`](backtracking/) | 回溯搜索、剪枝、状态空间树 |
| [`divide-and-conquer/`](divide-and-conquer/) | 分治（归并排序 / 快速排序） |
| [`dynamic-programming/`](dynamic-programming/) | 动态规划（背包、LCS、爬楼梯等） |
| [`monotonic-stack/`](monotonic-stack/) | 单调栈 |
| [`prefix-sum/`](prefix-sum/) | 前缀和 |
| [`sliding-window/`](sliding-window/) | 滑动窗口 |
| [`sorting/`](sorting/) | 排序算法（堆排序等） |
| [`tree/`](tree/) | 树的遍历、B/B+ 树笔记 |
| [`joy-of-coding/`](joy-of-coding/) | 编程之美练习（位运算 / 链表 / 数据结构 / 剑指 Offer） |

### 文档
| 文件 | 说明 |
|------|------|
| [`docs/binary-search.md`](docs/binary-search.md) | LeetCode 前 200 题 Binary Search 题单 |
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
