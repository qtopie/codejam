# Binary Search — LeetCode Top 200 题单

前 200 题中与 **Binary Search（二分查找）** 直接相关的题目，按经典程度与难度整理。

> 🔗 在线笔记：<https://qtopie.github.io/notes/codejam/binary-search/>

> ✅ = 本仓库已有实现

## 核心二分查找（模板题）

| # | 题目 | 难度 | 考点 |
|---|------|------|------|
| 35 | Search Insert Position | 🟢 Easy | 基础二分模板（插入位置/下界） |
| 69 | Sqrt(x) ✅ | 🟢 Easy | 二分逼近整数平方根，注意溢出 |
| 34 | Find First and Last Position of Element in Sorted Array ✅ | 🟡 Medium | 二分求左/右边界（lower_bound / upper_bound） |
| 33 | Search in Rotated Sorted Array ✅ | 🟡 Medium | 旋转数组中的二分，判断哪半有序 |
| 74 | Search a 2D Matrix ✅ | 🟡 Medium | 二维矩阵映射为一维二分 / 行首二分 |
| 167 | Two Sum II - Input Array Is Sorted | 🟡 Medium | 有序数组（双指针为主，二分亦可） |

## 进阶二分

| # | 题目 | 难度 | 考点 |
|---|------|------|------|
| 81 | Search in Rotated Sorted Array II | 🟡 Medium | 含重复元素的旋转数组，需处理 `nums[l]==nums[mid]` |
| 153 | Find Minimum in Rotated Sorted Array | 🟡 Medium | 二分找旋转点 / 最小值 |
| 154 | Find Minimum in Rotated Sorted Array II | 🔴 Hard | 含重复元素的找最小值 |
| 162 | Find Peak Element | 🟡 Medium | 根据相邻比较二分找峰值 |

## 二分答案 / 分治

| # | 题目 | 难度 | 考点 |
|---|------|------|------|
| 50 | Pow(x, n) | 🟡 Medium | 快速幂（分治/二分指数） |
| 4 | Median of Two Sorted Arrays | 🔴 Hard | 对较小数组二分，O(log(min(m,n))) |

## 说明

- **33 / 34 / 35 / 69**：LeetCode 二分专题的「四件套」，建议优先掌握。
- **81 / 153 / 154**：旋转数组系列，理解「哪一半有序」的判断逻辑是核心。
- **50** 通常归类为 Math/Divide & Conquer，但本质是指数二分，常与二分题一起刷。
- 超出 200 但同属于二分高频题的补充：`278. First Bad Version`（找第一个坏版本）、`287. Find the Duplicate Number`（快慢指针/二分）、`222. Count Complete Tree Nodes`（树高二分）。

## 本地练习入口

- 基础模板参考：`joy-of-coding/list/binary-search.go`
- 仓库已有实现：`leetcode/33.search-in-rotated-sorted-array/`、`leetcode/34.find-first-and-last-position-of-element-in-sorted-array/`、`leetcode/sqrtx.go`、`leetcode/search-a-2d-matrix.go`
