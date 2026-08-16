# 数论与计算几何

- **基础数论**: GCD/扩展欧几里得、快速幂、模逆、CRT、筛法。
- **快速变换**: FFT/NTT、多项式卷积（进阶）。
- **计算几何**: 叉积/方向判断、凸包（Graham/Andrew）、扫描线、线段相交、最近点对。
- **位运算**: 位掩码、低位操作（Brian Kernighan 算法 `n & (n-1)`）。

## 现有代码实现

| 文件 | 说明 |
|------|------|
| [`bit-operator.go`](bit-operator.go) | 常用位操作（获取、设置、清除、取反位） |
| [`num-of-ones.go`](num-of-ones.go) | 二进制中 1 的个数统计（Hamming Weight） |
