# Go 刷题语法复习

刷 LeetCode / 竞赛时最常用的 Go 语法速查，覆盖输入输出、类型转换、字符串、异常、随机数、堆与位运算。

## 1. 输入输出

### fmt 包（简单场景）

```go
var n int
fmt.Scan(&n)            // 读一个整数

var a, b int
fmt.Scanf("%d %d\n", &a, &b)

fmt.Println(res)        // 输出并换行
fmt.Printf("%d\n", res) // 格式化输出
```

### bufio（竞赛/大数据量，推荐）

```go
in := bufio.NewScanner(os.Stdin)
in.Buffer(make([]byte, 1e6), 1e6) // 大输入时扩容缓冲
in.Split(bufio.ScanWords)         // 按空白分词，读起来最方便

for in.Scan() {
    s := in.Text()
    n, _ := strconv.Atoi(s)
    // ...处理
}
out := bufio.NewWriter(os.Stdout)
fmt.Fprintln(out, ans) // 或 out.WriteString(...)
out.Flush()            // 记得 Flush！
```

> 读数字最常用组合：`bufio.Scanner` + `strconv.Atoi`；写多行用 `bufio.Writer`，结束必须 `Flush()`。

## 2. 类型转换与 strconv

| 场景 | 写法 |
|------|------|
| 数字 → 字符串（十进制） | `strconv.Itoa(n)` |
| 字符串 → 数字 | `n, err := strconv.Atoi(s)` |
| 64 位数字 → 字符串 | `strconv.FormatInt(int64(n), 10)` |
| 字符串 → int64 | `n, _ := strconv.ParseInt(s, 10, 64)` |
| 字符串 → 浮点 | `f, _ := strconv.ParseFloat(s, 64)` |
| 字符/字节 → 数字 | `int(s[i]-'0')`，`int(s[i]-'a')` |

`Itoa` = Integer to ASCII，`Atoi` = ASCII to Integer，只需记这两个。

```go
s := strconv.Itoa(123)          // "123"
n, _ := strconv.Atoi("456")     // 456
c := int('9' - '0')             // 9，数字字符转数字
```

**整型之间转换**：显式转换，`int64(a)`, `int(b)`, `byte(x)`, `rune(x)`。不同位宽计算前必须先统一类型。

**数组/切片排序**：

```go
sort.Ints(nums)                 // int 切片升序
sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] }) // 降序
sort.Slice(people, func(i, j int) bool { // 多字段
    if people[i].age != people[j].age {
        return people[i].age < people[j].age
    }
    return people[i].name < people[j].name
})
sort.SearchInts(nums, x)        // 二分找 x 的下标（找不到返回插入点）
```

## 3. 字符串：byte vs rune

- **byte** = `uint8`（1 字节），适合 **ASCII**。
- **rune** = `int32`（4 字节），代表一个 **Unicode 码点**，适合 **中文等多字节字符**。

```go
s := "abc中"
fmt.Println(len(s))         // 6（字节数！ASCII 1 字节 + 中文字符 3 字节）
fmt.Println(len([]rune(s))) // 4（字符个数）

// 按字节遍历（ASCII 安全，中文会乱）
for i := 0; i < len(s); i++ { _ = s[i] }

// 按字符遍历（推荐！中文安全）
for _, r := range s { _ = r } // r 是 rune
```

**常见坑**：
- `s[i]` 是 byte；`s[i]` 与 `'a'` 比较 OK，但 `'中'` 是 rune，不能与 byte 直接比较。
- 需要按"字符"处理（翻转、去重、回文）时先转 `[]rune(s)`。

## 4. 常见字符串操作（strings 包）

```go
strings.Split(s, ",")        // 切分
strings.Join(parts, "/")     // 拼接
strings.Contains(s, sub)     // 包含
strings.HasPrefix(s, pre)    // 前缀
strings.HasSuffix(s, suf)    // 后缀
strings.ToLower(s)           // 转小写
strings.ToUpper(s)           // 转大写
strings.TrimSpace(s)         // 去首尾空白
strings.Trim(s, "/")         // 去首尾指定字符
strings.Index(s, sub)        // 子串位置，找不到返回 -1
strings.Repeat("ab", 3)      // "ababab"
strings.TrimPrefix(s, pre)   // 去前缀
strings.Count(s, "a")        // 统计出现次数
```

**字符串构建**：拼接少直接 `+`；循环拼接用 `strings.Builder`（性能远优于 `+`）。

```go
var b strings.Builder
for i := 0; i < 1000; i++ {
    b.WriteByte(byte(i % 26 + 'a'))
}
res := b.String()
```

**字符串 ↔ 字节/字符数组**：

```go
b := []byte(s)   // string -> []byte（可修改）
s2 := string(b)  // []byte -> string
r := []rune(s)   // string -> []rune（按字符）
s3 := string(r)  // []rune -> string
sort.Slice(r, func(i, j int) bool { return r[i] < r[j] }) // 按字符排序
```

## 5. 异常处理：error / panic / recover

### 常规错误（error）

```go
n, err := strconv.Atoi(s)
if err != nil {
    // 处理错误
}
```

### panic / recover（边界情况兜底）

```go
func safeDivide(a, b int) (res int) {
    defer func() {
        if r := recover(); r != nil {
            res = -1 // 返回兜底值
        }
    }()
    return a / b // b == 0 时触发 panic，被 recover 捕获
}
```

**要点**：
- `defer` + `recover` 必须在**同一个函数**内才生效。
- `recover` 只在 `defer` 的函数中调用才有意义，捕获后程序继续执行。
- LeetCode 环境通常直接忽略错误 `_`，但**出栈/索引前判断长度**更关键。

## 6. 随机数：math/rand

```go
import "math/rand"

rand.Seed(time.Now().UnixNano()) // Go 1.20 前需要，之后版本自动随机种子
n := rand.Intn(100)              // [0, 100) 随机整数
n2 := rand.Intn(100) + 1         // [1, 100]
f := rand.Float64()              // [0, 1) 随机浮点
perm := rand.Perm(10)            // 0..9 的一个随机排列（洗牌）
rand.Shuffle(len(nums), func(i, j int) {
    nums[i], nums[j] = nums[j], nums[i]
}) // 原地洗牌
```

**用途**：随机快排、随机化算法（如随机选择第 k 大）、蒙特卡洛。

## 7. 堆（container/heap）与堆排序

Go 的 `container/heap` 需要自己实现 `sort.Interface` 接口（`Len / Less / Swap`）+ `Push / Pop`。**`Pop` 注意取末尾元素**。

```go
// 小顶堆（PriorityQueue 也同理，元素换成结构体即可）
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] } // 小顶堆；改成 > 即大顶堆
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]   // 取末尾！
    *h = old[:n-1]
    return x
}

// 使用
h := &IntHeap{}
heap.Init(h)          // 建堆 O(n)
heap.Push(h, 3)       // 插入 O(log n)
x := heap.Pop(h).(int) // 弹出堆顶 O(log n)
top := (*h)[0]        // 查看堆顶 O(1)
```

**要点**：
- `Push` 追加到末尾，`Pop` 取末尾元素（`container/heap` 内部会先交换）。
- 大顶堆只需把 `Less` 改成 `h[i] > h[j]`。
- 求 Top-K：维护大小为 k 的小顶堆（保留最大的 k 个）。
- **堆排序**：`heap.Init` 建堆 + 不断 `heap.Pop`，即升序排序（用大顶堆得降序）。

```go
h := &IntHeap{5, 1, 3}
heap.Init(h)
sorted := make([]int, 0, len(*h))
for h.Len() > 0 {
    sorted = append(sorted, heap.Pop(h).(int)) // 小顶堆 → 升序
}
```

## 8. 位运算（bit 操作）

| 操作 | Go 写法 | 说明 |
|------|---------|------|
| 与 / 或 / 异或 | `a & b`, `a \| b`, `a ^ b` | |
| 取反 | `^a` | 按位取反 |
| 左移 / 右移 | `a << k`, `a >> k` | 左移乘 2，右移整除 2 |
| 判奇偶 | `x & 1 == 0` | 偶数为 0 |
| 取最低位 1 | `x & -x` | lowbit，树状数组核心 |
| 去掉最低位 1 | `x & (x - 1)` | 统计 1 个数常用 |
| 检查第 k 位 | `(x >> k) & 1` | |
| 置第 k 位为 1 | `x \| (1 << k)` | |
| 第 k 位取反 | `x ^ (1 << k)` | |

**常见算法模板**：

```go
// 统计二进制中 1 的个数
count := 0
for x := n; x > 0; x &= x - 1 { count++ }

// 判断 2 的幂
isPow2 := n > 0 && n&(n-1) == 0

// 枚举所有子集（位掩码）
for mask := 0; mask < 1<<n; mask++ {
    for i := 0; i < n; i++ {
        if mask>>i&1 == 1 { /* 选中第 i 个元素 */ }
    }
}

// 异或找只出现一次的数（其余成对）
single := 0
for _, v := range nums { single ^= v }
```

## 9. 其他高频点

- **取最大值/最小值**：Go 无内置 `max/min`（1.21+ 有 `max/min` 内置函数），手写 `if a > b { a }`。
- **切片初始化**：`make([]int, n)`（含 n 个零值）、`make([]int, 0, n)`（空但预留容量）。
- **二维切片**：`make([][]int, m); for i := range dp { dp[i] = make([]int, n) }`。
- **map**：`m := make(map[int]int)`；判断 key 存在 `if v, ok := m[k]; ok {}`；`m[k]++` 自动加（缺省为 0）。
- **深拷贝切片**：`copy(dst, src)` 或 `append([]int(nil), src...)`。
- **反转切片**：`for i, j := 0, n-1; i < j; i, j = i+1, j-1 { s[i], s[j] = s[j], s[i] }`。

## 参考

- 相关笔记：[栈笔记](./stack.md)、[背包问题](./knapsack.md)
- Go 文档：[strconv](https://pkg.go.dev/strconv)、[strings](https://pkg.go.dev/strings)、[container/heap](https://pkg.go.dev/container/heap)、[math/rand](https://pkg.go.dev/math/rand)
