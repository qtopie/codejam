# 背包问题 Knapsack Problem

## 0, 1背包问题

给定一组``\:n``个物品，每种物品都有自己的重量（``w_i``）和价值``v_i``，在限定的总重量/总容量``\:C``内，选择其中若干个（也即每种物品可以选0个或1个），设计选择方案使得物品的总价值最高。

即给定正整数``\:\{(w_i, v_i)\}_1 \leq i \leq n``、给定正整数``\:C``，求解0-1规划问题：

```math
max \sum_{i=0}^n{x_iv_i},\: \text{s.t.} \sum_i^n{x_iw_i} \leq C,\: x_i \leq \{0, 1\}
```

> `s.t.` is abbrev for `subject to`.

### 解法分析

* 暴力解法： 每个物品都有0/1两种可能是否选入到背包中, 可以通过枚举方式解决，但是需要``\:2^n``次计算.

* 使用动态规划, 空间换时间(过滤不必要的计算). 从最低重量开始(最轻物品未知, 为了方便编写程序我们一般从0开始) 到C考察可以满足重量的集合. 同时考察的范围也是从第一个(假设物品从左到右边编号排序)到最后一个, 比较每个物品加入后对当前容量的影响.

以最后一个物品为例, 是否加入到背包中, 取决于加入它后的价值会变得更大, 还是不加入它的价值最大. 这两种情况都会动态的改变背包容量和物品项集合, 但同时可以利用前面计算的结果.

* 递推公式

```math
m(i, W) = max\{m(i -1, W),\: m(i-1, W - w_i) + v_i\}
```

## 无界（完全）背包问题 Unbounded Knapsack 

将0,1背包问题延伸, 每个物品可以取0，1或者多个, 求此时背包物品的最大价值.

* 递推公式

```math
F[X] = max \left\{ \begin{array}{l}
    F[X-w_1] + v_1 \\
    F[X-w_2] + v_2 \\
    ...\\
    F[X-w_n] + v_n \\
 \end{array} \right.
```

直观解释，对于给定重量，（因为可以重复取），在不超重的情况下尽可能取性价比最高的，而实际上会利用之前给定重量计算的最优组合.

### Memory Function

*重复利用数组，二位降为一维（空间复杂度）。结合分治Top-Down和bottom-up累积计算*

无界背包和01背包的递推公式不同


## 参考
- Introduction to the design and analysis of algorithms (3rd) page 292
- https://simplecodehints.com/blog/knapsack-problem-dynamic-programming-algorithm/#unbounded-knapsack-problem