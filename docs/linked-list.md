# Linked List — 链表常见算法题与核心技巧总结

本篇文档基于 LeetCode 经典链表高频题目，系统性总结链表数据结构的核心解题范式、关键技巧与易错陷阱。

> 🔗 在线笔记：<https://qtopie.github.io/notes/codejam/linked-list/>  
> ✅ = 本仓库已有实现

---

## 目录
1. [高频题单总览](#高频题单总览)
2. [链表六大核心解题技巧与模板](#链表六大核心解题技巧与模板)
   - [技巧 1: 虚拟头节点（Dummy Head / Sentinel）](#技巧-1-虚拟头节点dummy-head--sentinel)
   - [技巧 2: 快慢指针（Fast & Slow Pointers）](#技巧-2-快慢指针fast--slow-pointers)
   - [技巧 3: 前后定距双指针（Fixed-Distance Pointers）](#技巧-3-前后定距双指针fixed-distance-pointers)
   - [技巧 4: 链表局部与分组反转（List Reversal）](#技巧-4-链表局部与分组反转list-reversal)
   - [技巧 5: 双链表拆分与缝合（Partition & Stitching）](#技巧-5-双链表拆分与缝合partition--stitching)
   - [技巧 6: 哈希表 + 双向链表（Hash + Doubly Linked List）](#技巧-6-哈希表--双向链表hash--doubly-linked-list)
3. [题目分类深度剖析与考点](#题目分类深度剖析与考点)
4. [链表避坑指南与易错点](#链表避坑指南与易错点)
5. [本地练习入口与仓库实现](#本地练习入口与仓库实现)

---

## 高频题单总览

| # | 题目 | 难度 | 核心技巧 / 考点 | 仓库状态 |
|---|------|------|-----------------|----------|
| 141 | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) | 🟢 Easy | Floyd 判圈法（快慢指针） | 待补充 |
| 2 | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | 🟡 Medium | 虚拟头节点、高精度加法模拟、进位传递 | ✅ `leetcode/2.add-two-numbers/` |
| 21 | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | 🟢 Easy | 虚拟头节点、双指针归并 | 待补充 |
| 138 | [Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/) | 🟡 Medium | 哈希映射 / 节点原地交织复制与拆分 | 待补充 |
| 92 | [Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/) | 🟡 Medium | 局部区间反转、头插法 / 三指针穿针引线 | 待补充 |
| 25 | [Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/) | 🔴 Hard | 分组反转、长度前瞻探测、区间指针缝合 | ✅ `leetcode/25.reverse-nodes-in-k-group/` |
| 19 | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | 🟡 Medium | 虚拟头节点、定距双指针（前驱定位） | 待补充 |
| 82 | [Remove Duplicates from Sorted List II](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/) | 🟡 Medium | 虚拟头节点、前驱指针、重复区间批量跳过 | ✅ `leetcode/remove-duplicates-from-sorted-list-ii.go` |
| 61 | [Rotate List](https://leetcode.com/problems/rotate-list/) | 🟡 Medium | 闭环成环、取模计算有效偏移、断环截断 | ✅ `leetcode/rotate-list.go` |
| 86 | [Partition List](https://leetcode.com/problems/partition-list/) | 🟡 Medium | 双虚拟头节点、双链表分流与合并、尾指针置空防环 | ✅ `leetcode/86.partition-list/` |
| 146 | [LRU Cache](https://leetcode.com/problems/lru-cache/) | 🟡 Medium | 哈希表 + 双向链表（带头尾哨兵）、$O(1)$ 驱逐与提升 | 待补充 |

---

## 链表六大核心解题技巧与模板

### 技巧 1: 虚拟头节点（Dummy Head / Sentinel）
- **核心思想**：当操作可能涉及原链表的**头节点被删除、被交换或在头部新增节点**时，创建 `dummy := &ListNode{Next: head}`。
- **作用**：
  1. 消除对 `head` 为空或对头节点特殊处理的 `if/else` 分支。
  2. 统一所有节点（包括第 1 个节点）都具有前驱节点 `pre`。
  3. 最终统一返回 `dummy.Next`。

```go
// 典型通用模板
dummy := &ListNode{Next: head}
cur := dummy
for cur.Next != nil {
    // 统一处理 cur.Next 节点
}
return dummy.Next
```

---

### 技巧 2: 快慢指针（Fast & Slow Pointers）
- **核心思想**：定义步长不同的两个指针（如 `slow` 每次走 1 步，`fast` 每次走 2 步）。
- **应用场景**：
  1. **判断是否有环**（LC 141）：若有环，`fast` 必定在环内追上 `slow`（相对速度差为 1，不会跳过）。
  2. **寻找环起点**（LC 142）：相遇后将一个指针置于 `head`，两指针同速前进，再次相遇即为环入口。
  3. **寻找链表中点**（LC 876 / 归并排序 LC 148）：`fast` 走到末尾时，`slow` 恰在中间。

```go
// 环检测模板 (LC 141)
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

---

### 技巧 3: 前后定距双指针（Fixed-Distance Pointers）
- **核心思想**：寻找倒数第 $n$ 个节点时，无需遍历两遍求长度。
- **步骤**：
  1. 让 `fast` 先向前走 $n$ 步（若要定位待删节点的前驱，配合 `dummy` 让 `fast` 先走 $n+1$ 步）。
  2. 然后 `slow` 和 `fast` 同时同速向前走。
  3. 当 `fast == nil` 时，`slow` 正好停在目标位置（或目标前驱）。

```go
// 删除倒数第 N 个节点 (LC 19)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    fast, slow := dummy, dummy
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}
```

---

### 技巧 4: 链表局部与分组反转（List Reversal）
- **迭代反转模板（三指针迭代）**：
  保存 `next := cur.Next`，反转指针 `cur.Next = pre`，推进 `pre = cur; cur = next`。
- **区间反转（LC 92）**：
  定位到反转区间的前驱 `pre`，使用**头插法**连续将后续节点移动到 `pre` 之后。
- **K个一组反转（LC 25）**：
  1. 探测剩余节点数是否 $\ge k$，若不足则保持原样返回；
  2. 对这 $k$ 个节点执行局部反转；
  3. 递归或迭代连接上一组的尾部与下一组的头部。

```go
// LC 92 区间反转（头插法）
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{Next: head}
    pre := dummy
    for i := 0; i < left-1; i++ {
        pre = pre.Next
    }
    cur := pre.Next
    for i := 0; i < right-left; i++ {
        // 把nxt插入过来到prev和cur之间
        nxt := cur.Next
        cur.Next = nxt.Next
        nxt.Next = pre.Next
        pre.Next = nxt
    }
    return dummy.Next
}
```

---

### 技巧 5: 双链表拆分与缝合（Partition & Stitching）
- **核心思想**：遇到根据某个条件（如节点值大小、奇偶性）重新分组链表时，维护两套独立的 `dummy` 与 `tail` 指针。
- **关键细节**：
  1. 分别将原节点接入对应的小链表/大链表。
  2. 遍历结束后，**必须将大链表的尾节点 `next` 置为 `nil`**（否则可能保留原有后继指针导致链表成环）。
  3. 将小链表的尾部连接到大链表的首节点（`dummy2.Next`）。

```go
// LC 86 分隔链表
func partition(head *ListNode, x int) *ListNode {
    smallDummy, largeDummy := &ListNode{}, &ListNode{}
    small, large := smallDummy, largeDummy
    for head != nil {
        if head.Val < x {
            small.Next = head
            small = small.Next
        } else {
            large.Next = head
            large = large.Next
        }
        head = head.Next
    }
    large.Next = nil // 关键：断开原有后续，防止环
    small.Next = largeDummy.Next
    return smallDummy.Next
}
```

---

### 技巧 6: 哈希表 + 双向链表（Hash + Doubly Linked List）
- **核心思想**：常用于需要 $O(1)$ 查找与 $O(1)$ 插入/移动/删除的场景（如 **LRU Cache**）。
- **结构设计**：
  1. 双向链表节点包含 `key, val, prev, next`。
  2. 使用伪头 `head` 和伪尾 `tail` 哨兵简化空链表插入删除。
  3. 哈希表 `map[int]*DLinkedNode` 映射 key 到链表节点，支持 $O(1)$ 定位。

```go
// LC 146 LRU Cache 核心架构
type DLinkedNode struct {
    key, value int
    prev, next *DLinkedNode
}

type LRUCache struct {
    size, capacity int
    cache          map[int]*DLinkedNode
    head, tail     *DLinkedNode
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        cache:    map[int]*DLinkedNode{},
        head:     &DLinkedNode{},
        tail:     &DLinkedNode{},
        capacity: capacity,
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}
```

---

## 题目分类深度剖析与考点

### 1. 模拟与高精度运算（Simulation & High Precision）
- **LC 2. Add Two Numbers**
  - **题意**：逆序存储的两数相加，返回新的链表。
  - **技巧**：使用 `carry` 维护进位；循环条件为 `l1 != nil || l2 != nil || carry > 0`，精简代码无需在循环外单独处理最高位进位。

### 2. 归并与双指针合并（Merge & Two Pointers）
- **LC 21. Merge Two Sorted Lists**
  - **题意**：合并两个有序链表。
  - **技巧**：Dummy 节点配合双指针比较头部较小者，循环结束将非空的一条剩余部分直接挂接到末尾（无需逐个循环追加）。

### 3. 重复元素处理（Deduplication）
- **LC 82. Remove Duplicates from Sorted List II**
  - **题意**：删除排序链表中所有出现重复的节点（只要重复过一个不留）。
  - **技巧**：`cur := dummy`，每次检查 `cur.Next` 与 `cur.Next.Next` 是否值相同。若相同，记录该重复值 `val`，用循环跳过所有值为 `val` 的节点并更新 `cur.Next`；若不同才向前移动 `cur`。

### 4. 旋转与位移（Rotation & Shift）
- **LC 61. Rotate List**
  - **题意**：将链表每个节点向右移动 $k$ 个位置。
  - **技巧**：
    1. 先遍历得到链表长度 $n$，同时找到尾节点；
    2. 计算有效移动步数 $k = k \pmod n$，若 $k == 0$ 直接返回；
    3. 将原尾节点连接到原头节点形成闭环；
    4. 从原尾节点出发再走 $n - k$ 步到达新尾节点，断开 `newTail.Next = nil` 并返回新的头节点。

### 5. 深拷贝与复杂指针（Deep Copy with Random Pointer）
- **LC 138. Copy List with Random Pointer**
  - **解法一（哈希表）**：空间 $O(N)$，`map[*Node]*Node` 记录原节点到新节点的映射。
  - **解法二（空间优化 $O(1)$ 原地拆分）**：
    1. 复制每个节点插入原节点后面：`A -> A' -> B -> B'`；
    2. 设置新节点的 random：`cur.Next.Random = cur.Random.Next`（注意判空）；
    3. 拆分链表还原原链表并提取新链表。

### 6. 分组与局部反转（Reversal & Sub-reversal）
- **LC 92. Reverse Linked List II**：头插法一次遍历反转区间 $[left, right]$。
- **LC 25. Reverse Nodes in k-Group**：先通过计数探测剩余是否有 $k$ 个节点，有则调用子反转函数，无则保持原序，连接各区间。

### 7. 缓存与综合数据结构（Cache & Design）
- **LC 146. LRU Cache**：
  - `Get(key)`：哈希命中则将节点移动到双向链表头部（`moveToHead`），返回值；未命中返回 -1。
  - `Put(key, value)`：若存在则更新值并移至头部；若不存在则新建节点插入头部，若超出 `capacity`，删除尾部真实节点（`removeTail`）并在哈希表中清除对应键。

---

## 链表避坑指南与易错点

| 序号 | 常见错误 | 原因与表现 | 解决方案 |
|------|----------|------------|----------|
| 1 | **空指针异常（NPE）** | 访问了 `nil.Next`（如在 `fast = fast.Next.Next` 时未检查 `fast.Next`） | 循环条件严格把控，如 `for fast != nil && fast.Next != nil` |
| 2 | **链表成环（Memory Limit / TLE）** | 分隔、重排、截断链表时，末尾节点的 `Next` 未显式置为 `nil` | 如 LC 86 中大链表尾部必须显式 `large.Next = nil` |
| 3 | **断链丢失引用** | 改变 `cur.Next` 前未先保存原来的 `cur.Next` | 必须在反转或插入前定义 `nxt := cur.Next` |
| 4 | **头节点变动丢失** | 头节点可能被删除或被换到后面，但直接返回了原 `head` | 必须统一使用 `dummy := &ListNode{Next: head}` 并返回 `dummy.Next` |
| 5 | **取模越界 / 冗余循环** | LC 61 中 $k$ 可能远大于链表长度 $n$ | 先统计长度 $n$，执行 $k = k \pmod n$，避免无效多次循环 |

---

## 本地练习入口与仓库实现

- **基础数据结构定义**：`leetcode/node.go`
- **链表反转实现**：`leetcode/reverse-list.go`、`joy-of-coding/list/reverse.go`
- **两数相加**：`leetcode/2.add-two-numbers/`、`leetcode/add-two-number-ii.go`
- **K个一组反转**：`leetcode/25.reverse-nodes-in-k-group/`
- **分隔链表**：`leetcode/86.partition-list/`、`leetcode/partition-list.go`
- **删除排序链表重复元素**：`leetcode/remove-duplicates-from-sorted-list-ii.go`、`leetcode/remove-duplicates-from-sorted-list.go`
- **旋转链表**：`leetcode/rotate-list.go`
- **链表排序与插入**：`leetcode/147.insertion-sort-list/`、`leetcode/148.sort-list/`
- **相交与合并链表**：`leetcode/intersection-of-two-linked-lists.go`、`leetcode/23.merge-k-sorted-lists/`
