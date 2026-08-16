# Java 刷题语法与常用 API 极速速查表

> 适用于 LeetCode / 算法面试快速唤醒记忆。特别针对从 Go / Python / C++ / JS 切换到 Java，或久未写 Java 的开发者。

---

## 目录
1. [基础类型与常用常数](#1-基础类型与常用常数)
2. [长度与容量的混淆规避（高频坑）](#2-长度与容量的混淆规避高频坑)
3. [数组 (Array)](#3-数组-array)
4. [字符串 (String & StringBuilder)](#4-字符串-string--stringbuilder)
5. [列表 (List / ArrayList)](#5-列表-list--arraylist)
6. [队列与双端队列 (Queue & Deque)](#6-队列与双端队列-queue--deque)
7. [栈 (Stack)](#7-栈-stack)
8. [堆 / 优先队列 (PriorityQueue)](#8-堆--优先队列-priorityqueue)
9. [哈希表与有序表 (HashMap & TreeMap)](#9-哈希表与有序表-hashmap--treemap)
10. [哈希集合与有序集合 (HashSet & TreeSet)](#10-哈希集合与有序集合-hashset--treeset)
11. [排序与自定义比较器 (Sort & Comparator)](#11-排序与自定义比较器-sort--comparator)
12. [位运算核心技巧与 API (Bit Manipulation)](#12-位运算核心技巧与-api-bit-manipulation)
13. [常用数学与算法技巧 (Math & Algo Tricks)](#13-常用数学与算法技巧-math--algo-tricks)
14. [刷题必背避坑清单 (Gotchas)](#14-刷题必背避坑清单-gotchas)

---

## 1. 基础类型与常用常数

| 基础类型 | 包装类型 | 极值与常用常数 |
| :--- | :--- | :--- |
| `int` (32-bit) | `Integer` | `Integer.MAX_VALUE` ($2^{31}-1 \approx 2.14 \times 10^9$)<br>`Integer.MIN_VALUE` ($-2^{31}$) |
| `long` (64-bit) | `Long` | `Long.MAX_VALUE`<br>`Long.MIN_VALUE`<br>字面量必须加 `L`（如 `0L`, `10000000000L`） |
| `char` (16-bit) | `Character` | `'a' ~ 'z'` (97~122), `'A' ~ 'Z'` (65~90), `'0' ~ '9'` (48~57) |
| `double` (64-bit) | `Double` | `Double.MAX_VALUE`, `Double.POSITIVE_INFINITY` |
| `boolean` | `Boolean` | `true`, `false` |

### 类型互转
```java
// String <-> int / long
int val = Integer.parseInt("123");
long lVal = Long.parseLong("12345678901");
String s = String.valueOf(123); // 或 Integer.toString(123)

// char <-> int
int digit = c - '0';       // '5' -> 5
char ch = (char) ('0' + 5); // 5 -> '5'
int lowerIdx = c - 'a';    // 'c' -> 2
```

---

## 2. 长度与容量的混淆规避（高频坑）

| 数据结构 | 获取长度/大小的方式 | 易错示例 |
| :--- | :--- | :--- |
| **数组** (`int[]`, `int[][]`) | `arr.length` （**属性**，无括号） | ❌ `arr.length()` ❌ `arr.size()` |
| **字符串** (`String`) | `str.length()` （**方法**，有括号） | ❌ `str.length` ❌ `str.size()` |
| **集合** (`List`, `Set`, `Map`, `Queue`) | `coll.size()` （**方法**） | ❌ `coll.length` ❌ `coll.length()` |

---

## 3. 数组 (Array)

### 声明与初始化
```java
int[] nums = new int[n];                 // 默认全 0
int[][] matrix = new int[m][n];          // m 行 n 列，默认全 0
int[] primes = {2, 3, 5, 7, 11};         // 静态初始化
int[][] grid = new int[][]{{1, 2}, {3, 4}};
```

### 常用操作 (`java.util.Arrays`)
```java
// 填充
Arrays.fill(nums, -1);
for (int[] row : matrix) Arrays.fill(row, -1); // 二维数组批量填充

// 复制
int[] copy1 = Arrays.copyOf(nums, nums.length);
int[] subArray = Arrays.copyOfRange(nums, fromIdx, toIdx); // [from, to) 左闭右开
System.arraycopy(src, srcPos, dest, destPos, length);      // 原生高性能拷贝

// 排序与二分
Arrays.sort(nums);                           // 原地升序（基本类型双轴快排）
int idx = Arrays.binarySearch(nums, target); // 返回索引；若未找到返回 (-(insertion point) - 1)

// 转字符串打印 (调试必备)
System.out.println(Arrays.toString(nums));        // 一维数组
System.out.println(Arrays.deepToString(matrix));   // 二维/多维数组
```

---

## 4. 字符串 (String & StringBuilder)

> 💡 **关键机制**：Java 中 `String` 是**不可变对象**。拼接/频繁修改必须使用 `StringBuilder`！

### `String` 常用 API
```java
String s = "Hello World";

// 长度与字符
int len = s.length();
char c = s.charAt(i);
char[] chars = s.toCharArray(); // 转字符数组（原地修改更方便）

// 截取与查找
String sub = s.substring(start, end); // [start, end) 左闭右开
int idx = s.indexOf("or");            // 返回首个匹配索引，无匹配返回 -1
int lastIdx = s.lastIndexOf('o');
boolean has = s.contains("Wor");
boolean start = s.startsWith("He");
boolean end = s.endsWith("ld");

// 分割与替换
String[] parts = s.split("\\s+");     // 正则分割，\\s+ 代表连续空格
String clean = s.trim();              // 去除两端首尾空白
String rep = s.replace('o', 'a');     // 替换所有匹配字符

// 比较（严禁使用 ==）
boolean eq = s1.equals(s2);
boolean eqIgnoreCase = s1.equalsIgnoreCase(s2);
int cmp = s1.compareTo(s2);           // 字典序比较 (<0, 0, >0)
```

### `Character` 工具类
```java
Character.isLetter(c);        // 是否为字母
Character.isDigit(c);         // 是否为数字 (0-9)
Character.isLetterOrDigit(c); // 是否为字母或数字（回文串题常用）
Character.isLowerCase(c);
Character.isUpperCase(c);
Character.toLowerCase(c);
Character.toUpperCase(c);
```

### `StringBuilder` 核心操作
```java
StringBuilder sb = new StringBuilder();
sb.append("abc").append(123).append('d');
sb.deleteCharAt(sb.length() - 1);       // 弹出最后一个字符 (回溯/路径拼接高频)
sb.delete(start, end);                  // 删除 [start, end)
sb.insert(offset, "xyz");
sb.reverse();                           // 反转字符串
sb.setCharAt(idx, 'x');                 // 修改某位
String res = sb.toString();
```

---

## 5. 列表 (List / ArrayList)

```java
List<Integer> list = new ArrayList<>();

// 增删查改
list.add(10);
list.add(0, 5);               // 在索引 0 处插入 5 (O(n))
int val = list.get(i);
list.set(i, 99);              // 修改索引 i 的值
list.remove(list.size() - 1); // 删除末尾 (O(1))

// 注意：删除包装类型时按值还是按索引！
list.remove(0);                       // 删除索引为 0 的元素
list.remove(Integer.valueOf(10));     // 删除值为 10 的首个元素

// 判空与包含
boolean empty = list.isEmpty();
int size = list.size();
boolean exists = list.contains(10);

// 回溯嵌套 List 必备深拷贝
List<List<Integer>> ans = new ArrayList<>();
List<Integer> path = new ArrayList<>();
ans.add(new ArrayList<>(path)); // ⚠️ 必须 new ArrayList<>(path)，不可直接 ans.add(path)

// List 与 Array 互转
Integer[] arr = list.toArray(new Integer[0]);
List<Integer> fromArr = new ArrayList<>(Arrays.asList(1, 2, 3));
```

---

## 6. 队列与双端队列 (Queue & Deque)

> 💡 **最佳实践**：无论是单向队列还是双端队列/单调队列，**统一使用 `ArrayDeque`** 实现（性能优于 `LinkedList`，无并发锁开销）。

```java
// 单向队列 (BFS 常用)
Queue<TreeNode> queue = new ArrayDeque<>();
queue.offer(root);          // 入队 (推荐用 offer，不用 add)
TreeNode node = queue.poll(); // 出队并返回 (若空返回 null；remove 会抛异常)
TreeNode peek = queue.peek(); // 查看队首
boolean empty = queue.isEmpty();
int size = queue.size();

// 双端队列 (滑动窗口最大值 / 单调队列 / 0-1 BFS)
Deque<Integer> deque = new ArrayDeque<>();
deque.offerFirst(x); // 队头插入
deque.offerLast(x);  // 队尾插入 (等价于 offer)
int a = deque.pollFirst(); // 队头弹出 (等价于 poll)
int b = deque.pollLast();  // 队尾弹出
int head = deque.peekFirst();
int tail = deque.peekLast();
```

---

## 7. 栈 (Stack)

> ⚠️ **避坑警告**：不要使用历史遗留类 `java.util.Stack`（继承自 `Vector`，所有方法带 `synchronized`，性能低且违背接口隔离）。  
> 刷题**推荐使用 `Deque` 接口 + `ArrayDeque`** 作为栈。

```java
Deque<Integer> stack = new ArrayDeque<>();

stack.push(10);          // 入栈
int top = stack.peek();  // 查看栈顶
int popped = stack.pop();// 出栈并返回
boolean empty = stack.isEmpty();
int size = stack.size();
```

---

## 8. 堆 / 优先队列 (PriorityQueue)

> 💡 默认是**小顶堆**（队首为最小值）。

```java
// 1. 小顶堆 (默认)
PriorityQueue<Integer> minHeap = new PriorityQueue<>();

// 2. 大顶堆 (自定义 Comparator)
PriorityQueue<Integer> maxHeap = new PriorityQueue<>((a, b) -> Integer.compare(b, a));
// 或：PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Collections.reverseOrder());

// 3. 针对自定义对象/数组排序 (例如带权图节点 [nodeId, distance])
// 按 distance 升序的小顶堆
PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> Integer.compare(a[1], b[1]));

// 核心操作
pq.offer(val);          // 插入 (O(log k))
int min = pq.poll();    // 弹出堆顶 (O(log k))
int top = pq.peek();    // 查看堆顶 (O(1))
int size = pq.size();
boolean empty = pq.isEmpty();
```

---

## 9. 哈希表与有序表 (HashMap & TreeMap)

### `HashMap`
```java
Map<String, Integer> map = new HashMap<>();

// 存取与默认值
map.put("apple", 3);
int count = map.getOrDefault("banana", 0);
boolean hasKey = map.containsKey("apple");
boolean hasVal = map.containsValue(3);
map.remove("apple");

// 刷题计数最简写法：
map.put(key, map.getOrDefault(key, 0) + 1);
// 或使用 merge:
map.merge(key, 1, Integer::sum);

// 遍历
for (Map.Entry<String, Integer> entry : map.entrySet()) {
    String k = entry.getKey();
    int v = entry.getValue();
}
for (String k : map.keySet()) { ... }
for (int v : map.values()) { ... }
```

### `TreeMap` (基于红黑树，Key 有序，操作 $O(\log n)$)
```java
TreeMap<Integer, String> treeMap = new TreeMap<>();
treeMap.put(10, "A");
treeMap.put(20, "B");
treeMap.put(30, "C");

int first = treeMap.firstKey();       // 最小 Key: 10
int last = treeMap.lastKey();         // 最大 Key: 30
Integer floor = treeMap.floorKey(25);     // <= 25 的最大 Key: 20
Integer ceil = treeMap.ceilingKey(25);   // >= 25 的最小 Key: 30
Integer lower = treeMap.lowerKey(20);     // < 20 的最大 Key: 10
Integer higher = treeMap.higherKey(20);   // > 20 的最小 Key: 30
```

---

## 10. 哈希集合与有序集合 (HashSet & TreeSet)

```java
// HashSet
Set<Integer> set = new HashSet<>();
set.add(1);
set.remove(1);
boolean has = set.contains(1);
set.size();
set.isEmpty();

// TreeSet (有序集合，支持二分范围查找)
TreeSet<Integer> ts = new TreeSet<>();
ts.add(10);
ts.add(20);
Integer le = ts.floor(15);   // <= 15: 10
Integer ge = ts.ceiling(15); // >= 15: 20
Integer lt = ts.lower(10);   // < 10: null
Integer gt = ts.higher(10);  // > 10: 20
```

---

## 11. 排序与自定义比较器 (Sort & Comparator)

### 常用排序模板
```java
// 1. 基本一维数组排序
Arrays.sort(nums); // 升序

// 2. 二维区间数组排序 (如合并区间 / 会议室问题)
// 先按左端点升序；左端点相同按右端点降序
Arrays.sort(intervals, (a, b) -> {
    if (a[0] != b[0]) {
        return Integer.compare(a[0], b[0]);
    }
    return Integer.compare(b[1], a[1]);
});

// 3. List 排序
Collections.sort(list); // 升序
list.sort((a, b) -> Integer.compare(b, a)); // 降序
```

> ⚠️ **避坑**：禁止在 Comparator 中直接写 `(a, b) -> a - b`！  
> 当 `a = -2147483648`，`b = 1` 时，`a - b` 会发生**整型下溢**变正数，导致排序彻底错乱。请一律使用 `Integer.compare(a, b)`。

---

## 12. 位运算核心技巧与 API (Bit Manipulation)

### 1. 基础运算符对照表
| 运算符 | 名称 | 规则 / 语法 | 示例 |
| :--- | :--- | :--- | :--- |
| `&` | 按位与 (AND) | 两位均为 1 则为 1，否则为 0 | `5 & 3` $\to$ `101 & 011 = 001 (1)` |
| `\|` | 按位或 (OR) | 两位只要有 1 则为 1，全 0 为 0 | `5 \| 3` $\to$ `101 \| 011 = 111 (7)` |
| `^` | 按位异或 (XOR) | 相同为 0，不同为 1（无进位加法） | `5 ^ 3` $\to$ `101 ^ 011 = 110 (6)` |
| `~` | 按位取反 (NOT) | 0 变 1，1 变 0（`~x = -x - 1`） | `~5` $\to$ `-6` |
| `<<` | 左移 | 各二进制位左移，低位补 0（相当于乘 $2^k$） | `3 << 2` $\to$ `12` |
| `>>` | 算术右移 | 各二进制位右移，**高位补符号位**（负数补 1） | `-8 >> 2` $\to$ `-2` |
| `>>>` | 无符号右移 | 各二进制位右移，**高位一律补 0**（忽略符号位） | `-8 >>> 2` $\to$ `1073741822` |

---

### 2. 单 bit 基础操控（$k$ 从 0 开始计数，从右至左）

> 💡 **位索引规则**：$k$ 从 **0** 开始计数（即 0-indexed，从右向左），$k = 0$ 表示二进制的**最低位（第 0 位，即 $2^0$ 对应的个位）**。

#### 原理解析：`(n >> k) & 1`
1. **右移对齐**：`n >> k` 将目标第 $k$ 位移动到最右端的第 0 位。
2. **掩码提取**：`& 1`（即 `& 000...0001`）将左边所有高位清零，仅保留最右侧的值，得到 `0` 或 `1`。

**示例**（以 $n = 6$，二进制 `0b110` 为例）：
- $k = 0$（第 0 位）：`6 >> 0 = 0b110`，`0b110 & 1 = 0`（第 0 位为 0）
- $k = 1$（第 1 位）：`6 >> 1 = 0b011`，`0b011 & 1 = 1`（第 1 位为 1）
- $k = 2$（第 2 位）：`6 >> 2 = 0b001`，`0b001 & 1 = 1`（第 2 位为 1）

```java
// 1. 获取第 k 位 (返回 0 或 1，k 从 0 开始)
int bit = (n >> k) & 1;
boolean isOne = (n & (1 << k)) != 0; // 等价判定：第 k 位是否为 1

// 2. 将第 k 位置为 1 (Set Bit)
n |= (1 << k);

// 3. 将第 k 位置为 0 (Clear Bit)
n &= ~(1 << k);

// 4. 将第 k 位取反 (Toggle Bit)
n ^= (1 << k);
```

---

### 3. 经典位运算技巧与神技 (Classic Bit Tricks)
```java
// 1. 消除二进制最低位的 1 (Brian Kernighan 算法)
// 原理：n - 1 会将最低位的 1 变为 0，其后的 0 全部变为 1，按位与后消除该位 1
n = n & (n - 1);

// 判定是否为 2 的幂次方 (Power of Two)
boolean isPowerOfTwo = n > 0 && (n & (n - 1)) == 0;

// 判定是否为 4 的幂次方
boolean isPowerOfFour = n > 0 && (n & (n - 1)) == 0 && (n & 0xaaaaaaaa) == 0;

// 2. 提取最低位的 1 (lowbit 运算 / 树状数组核心)
// 原理：利用计算机补码特性 -n = ~n + 1
int lowbit = n & (-n); // 6 (0110) & -6 (1010) -> 2 (0010)

// 3. 异或运算核心性质 (Single Number 题型核心)
// 性质：a ^ a = 0; a ^ 0 = a; 且满足交换律与结合律
int single = 0;
for (int num : nums) single ^= num; // 数组中唯一出现一次的数

// 4. 不用临时变量交换两数
a ^= b;
b ^= a;
a ^= b;
```

---

### 4. `Integer` 与 `Long` 内置高性能位操作 API
> 💡 源码基于底层的查表法或 CPU 原语指令，性能极高。

```java
int n = 0b00101100; // 44

// 1. 计算二进制中 1 的个数 (Hamming Weight / 汉明重量)
int count = Integer.bitCount(n); // 3

// 2. 前导 0 与后导 0 的个数
int lz = Integer.numberOfLeadingZeros(n);  // 32 位整型高位连续 0 的个数
int tz = Integer.numberOfTrailingZeros(n); // 低位连续 0 的个数 (即 lowbit 对应的 2 的幂指数)

// 3. 保留最高位 / 最低位的 1 (其余位全置 0)
int highest = Integer.highestOneBit(n); // 32 (0b00100000)
int lowest = Integer.lowestOneBit(n);   // 4  (0b00000100)

// 4. 进制转换与调试
String binStr = Integer.toBinaryString(n); // "101100" (无前导零)
int parsed = Integer.parseInt("101100", 2); // 44
```

---

### 5. 集合与状压 DP 常用操作 (Bitmask & State Compression)

在状压 DP 或回溯中，通常用一个 `int` 的二进制位表示一个元素大小不超过 30 的集合（第 $i$ 位为 1 表示集合包含元素 $i$）：

```java
// 全集 (包含 0 ~ n-1 共 n 个元素)
int allSet = (1 << n) - 1;

// 空集
int emptySet = 0;

// 集合基本运算
int union = maskA | maskB;           // 并集 A ∪ B
int intersect = maskA & maskB;       // 交集 A ∩ B
int diff = maskA & (~maskB);         // 差集 A \ B
int symDiff = maskA ^ maskB;         // 对称差 (A ∪ B) \ (A ∩ B)
int complement = allSet ^ maskA;     // 补集 ∁ A

// 判断元素 i 是否在集合中
boolean contains = ((mask >> i) & 1) == 1;

// 子集遍历核心模板 (Submask Enumeration)
// 高效枚举 mask 的所有非空子集（时间复杂度：所有 mask 子集枚举之和为 O(3^n)）
for (int sub = mask; sub > 0; sub = (sub - 1) & mask) {
    // 处理子集 sub
}
```

---

## 13. 常用数学与算法技巧 (Math & Algo Tricks)

### `Math` 常用函数
```java
Math.max(a, b);
Math.min(a, b);
Math.abs(x);
Math.pow(base, exp); // 返回 double
Math.sqrt(x);        // 返回 double
Math.ceil(x);        // 向上取整 (double)
Math.floor(x);       // 向下取整 (double)
Math.round(x);       // 四舍五入 (long/int)
```

### 二分查找标准中点防溢出
```java
int mid = left + (right - left) / 2; // 避免 (left + right) / 2 整型溢出
```

### 取模防负数技巧
```java
int MOD = 1_000_000_007;
int ans = (int) (((long) a % MOD + MOD) % MOD);
```

### 最大公约数 (GCD) 辗转相除法
```java
int gcd(int a, int b) {
    return b == 0 ? a : gcd(b, a % b);
}
```

---

## 14. 刷题必背避坑清单 (Gotchas)

1. **对象比较必须用 `.equals()`**：
   - 对于 `Integer`, `String`，使用 `==` 会比较对象内存地址（`Integer` 仅在 `[-128, 127]` 命中缓存时 `==` 为 true，超出即 false）。
2. **位运算优先级极低（天坑！）**：
   - `if (n & 1 == 0)` 会被编译器解析成 `if (n & (1 == 0))`，直接类型不兼容报错！
   - `int a = 1 << 2 + 1` 会被解析成 `1 << 3 = 8`，而不是 `4 + 1 = 5`！
   - **规则**：涉及位运算一律加括号：`if ((n & 1) == 0)`、`int a = (1 << 2) + 1`。
3. **64 位移位必须加 `L`**：
   - 对 `long` 进行超过 31 位的左移时，必须写 `1L << k`。如果写 `1 << 40`，字面量默认按 32 位 `int` 计算（实际只左移了 `40 % 32 = 8` 位）。
4. **回溯算法的结果保存必须深拷贝**：
   - `ans.add(new ArrayList<>(path))` 而不是 `ans.add(path)`。
5. **二分与排序防溢出**：
   - 中点计算写 `left + (right - left) / 2`。
   - 自定义 Comparator 用 `Integer.compare(a, b)` 替代 `a - b`。
6. **字符串频繁拼接与修改**：
   - 不要在循环里用 `s += "a"`（每次创建新 String，复杂度飙升至 $O(n^2)$），请用 `StringBuilder`。

