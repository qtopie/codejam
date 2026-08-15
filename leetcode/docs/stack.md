# 栈 Stack 笔记

> 基于五道经典题总结：
> - [Valid Parentheses](https://leetcode.cn/problems/valid-parentheses/) (Easy) — 括号匹配
> - [Simplify Path](https://leetcode.cn/problems/simplify-path/) (Medium) — 路径简化
> - [Min Stack](https://leetcode.cn/problems/min-stack/) (Medium) — 辅助栈
> - [Evaluate Reverse Polish Notation](https://leetcode.cn/problems/evaluate-reverse-polish-notation/) (Medium) — 后缀表达式
> - [Basic Calculator](https://leetcode.cn/problems/basic-calculator/) (Hard) — 中缀表达式 + 括号

## 1. 栈的本质

栈是 **LIFO（后进先出）** 的线性结构，只允许在栈顶操作。它天然适合处理**具有"最近关联性"**或**"嵌套结构"**的问题：

- 配对/匹配（括号、标签闭合）
- 表达式的括号嵌套与求值
- 历史回退/撤销（路径导航、函数调用栈）
- 维护"到目前为止"的某种极值（配合辅助栈）

Go 中用切片模拟栈：

```go
stack := make([]int, 0)
stack = append(stack, x)        // push
top := stack[len(stack)-1]      // peek
stack = stack[:len(stack)-1]    // pop
empty := len(stack) == 0
```

## 2. 题型一：括号匹配（Valid Parentheses）

**思路**：遍历字符，遇到**左括号入栈**；遇到**右括号**则与栈顶左括号配对，匹配则弹栈，不匹配直接返回 false。最后栈为空说明全部匹配。

```go
func isValid(s string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
```

**要点**：
- 用 Map 存右括号 → 左括号，避免写多个 if。
- 出栈前**先判空**：遇到右括号时栈为空说明没有可配对的左括号。

## 3. 题型二：路径简化（Simplify Path）

**思路**：按 `/` 切分路径得到各段，用栈维护当前目录层级：

- 段为 `"."` 或空串 → 忽略
- 段为 `".."` → **弹栈**（回到上一级，栈空则忽略）
- 其他 → 入栈（进入子目录）

最后用 `/` 拼接栈中元素。

**要点**：栈在这里体现"回退"语义 —— `..` 是撤销一步，恰好符合 LIFO。

## 4. 题型三：辅助栈（Min Stack）

**思路**：两个栈，`stack` 存数据，`minStack` 单调不增地记录**每个时刻的最小值**。

- push：`stack` 入栈；若新值 <= 当前最小值则同步入 `minStack`。
- pop：`stack` 弹栈；若弹出的值 == `minStack` 栈顶，`minStack` 也弹栈。
- getMin：直接返回 `minStack` 栈顶，O(1)。

**要点**：辅助栈与主栈**同生共死**（同进同出），保证任意时刻栈顶就是当前栈中最小值。这是"栈 + 单调性"的结合，为单调栈打基础。

## 5. 题型四：后缀表达式求值（Evaluate Reverse Polish Notation）

**思路**：遍历 tokens，遇到**数字入栈**；遇到**运算符**则弹出两个数，按运算顺序计算，结果**再入栈**。遍历结束栈顶即答案。

```go
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, t := range tokens {
		if isOp(t) {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, apply(a, b, t))
		} else {
			n, _ := strconv.Atoi(t)
			stack = append(stack, n)
		}
	}
	return stack[0]
}
```

**要点**：
- 后缀表达式**天然没有括号和优先级歧义**，从左到右一遍扫描即可，是表达式求值的基础形态。
- 减法/除法注意**操作数顺序**：先弹出的是右操作数（b），后弹出的是左操作数（a），即 `a - b`、`a / b`。

## 6. 题型五：中缀表达式 + 括号（Basic Calculator）

**思路**：中缀表达式需要处理**优先级**与**括号**，用两个信息：

1. **符号栈/变量**：记录当前作用域的符号位（sign = +1 / -1）。
2. **结果栈**：遇到 `(` 把当前累计结果和符号压栈保存，进入括号内新的作用域；遇到 `)` 弹栈恢复外层，并把括号内结果按外层符号累加。

核心循环逻辑：

```go
res, sign := 0, 1
stack := []int{}
for i := 0; i < len(s); i++ {
	switch {
	case s[i] >= '0' && s[i] <= '9':
		num := 0
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			i++
		}
		i--
		res += sign * num
	case s[i] == '+':
		sign = 1
	case s[i] == '-':
		sign = -1
	case s[i] == '(':
		stack = append(stack, res, sign) // 保存外层结果与符号
		res, sign = 0, 1                 // 重置内层
	case s[i] == ')':
		prevRes, prevSign := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		res = prevRes + prevSign*res // 内层结果带上外层符号并累加
	}
}
```

**要点**：
- **括号本质是作用域的入栈/出栈**：进括号保存现场，出括号恢复现场。
- `num = num*10 + digit` 是连续数字拼接的固定写法。
- 关键理解：`res` 存的是"当前括号作用域内已经计算完的和"，`sign` 表示下一个数要加还是减。
- 进阶：若有 `*` `/`（Basic Calculator II），需在乘除时先弹出栈顶运算再压回，配合优先级处理。

## 7. 延伸：单调栈（Monotonic Stack）

**定义**：栈中元素保持单调（栈底到栈顶单调递增或递减）。核心价值是**在 O(n) 内为每个元素找到下一个更大/更小元素**，把"暴力两两比较 O(n²)"降到 O(n)。

典型题：[Next Greater Element I](https://leetcode.cn/problems/next-greater-element-i/)（496）、[Largest Rectangle in Histogram](https://leetcode.cn/problems/largest-rectangle-in-histogram/)（84）。

### 下一个更大元素（Next Greater Element）

**思路**：**递减栈**（栈顶最小）。从左到右遍历，元素 `x` 不断弹出所有比它小的栈顶（这些元素的下一个更大元素就是 `x`），然后 `x` 入栈。弹出的瞬间就"结算"了答案。

```go
func nextGreaterElement(nums []int) []int {
	ans := make([]int, len(nums))
	stack := []int{} // 存索引，保证栈内数值单调递减
	for i, x := range nums {
		for len(stack) > 0 && x > nums[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = x // 结算：x 是 nums[j] 的下一个更大元素
		}
		stack = append(stack, i)
	}
	for _, j := range stack { // 栈内剩余元素没有更大值
		ans[j] = -1
	}
	return ans
}
```

**要点**：
- 栈内常存**索引**而非值（方便回写答案位置）。
- 递减栈求"下一个更大"，递增栈求"下一个更小"，方向可记忆为：**栈内保留的正是还没找到答案的元素**，它们按待处理顺序压在栈中。

### 柱状图中最大矩形（Largest Rectangle in Histogram）

**思路**：**递增栈**（栈底到栈顶高度递增）。遍历每个柱子，当新柱高 `h[i]` 小于栈顶柱高时，**弹出栈顶并以其高度结算矩形面积**——左右边界分别是栈内前一个元素（左边第一个比它矮的）和当前 `i`（右边第一个比它矮的）。最后末尾补一个高度 0 的哨兵柱，强制清空栈。

```go
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0) // 哨兵：强制最后清栈
	stack := []int{-1}           // 哨兵索引，方便左边界计算
	maxArea := 0
	for i, h := range heights {
		for len(stack) > 1 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			area := height * (i - left - 1) // 宽 = 右边界 - 左边界 - 1
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}
```

**要点**：
- 单调栈上结算的边界：**左边界 = 栈内上一个元素**（更矮），**右边界 = 当前元素**（更矮）。
- 哨兵技巧（数组尾部补 0 / 栈底放 -1）省去大量边界判断。
- 该题是 [Maximal Rectangle](https://leetcode.cn/problems/maximal-rectangle/)（85）与接雨水（42）的核心子问题。

**单调栈总结**：
1. **求下一个更大/更小** → 递减/递增栈，入栈时结算被弹出的元素。
2. **求左右第一个更矮/更高** → 递增/递减栈，弹出时用栈内前一个元素作左边界、当前元素作右边界。
3. 常用哨兵（补 0、栈底放 -1）统一边界处理。

## 8. 总结：什么时候用栈

| 特征 | 示例 |
|------|------|
| 后出现的先处理（最近相关性） | 括号匹配、`..` 回退 |
| 嵌套结构 / 作用域 | 括号求值、函数调用 |
| 逆序处理 | RPN、后缀/前缀表达式 |
| 需要 O(1) 查极值且数据在变 | Min Stack |
| 求下一个更大/更小元素 | 单调栈（Next Greater Element） |
| 求最大矩形/蓄水量 | 单调栈（Largest Rectangle、接雨水） |

**通用模板**：
1. 确定什么"入栈"、什么"出栈"（触发条件）。
2. 出栈前必判空。
3. 栈里存什么？—— 值，或 (值, 辅助信息) 元组。
4. 画图模拟一个例子，确认出入栈顺序与题意一致。

## 参考

- LeetCode 相关题目：
  - [Valid Parentheses #20](https://leetcode.cn/problems/valid-parentheses/) / [Simplify Path #71](https://leetcode.cn/problems/simplify-path/) / [Min Stack #155](https://leetcode.cn/problems/min-stack/)
  - [Eval RPN #150](https://leetcode.cn/problems/evaluate-reverse-polish-notation/) / [Basic Calculator #224](https://leetcode.cn/problems/basic-calculator/)
  - 单调栈：[Next Greater Element I #496](https://leetcode.cn/problems/next-greater-element-i/) / [Largest Rectangle #84](https://leetcode.cn/problems/largest-rectangle-in-histogram/) / [Trapping Rain Water #42](https://leetcode.cn/problems/trapping-rain-water/)
