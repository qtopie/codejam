# backtracking

搜索算法, 剪枝, state-space-tree

promising, non-promising 

1. DFS searching, and constructing state-space tree
2. promising? yes, add option to component; no, restore and back one more level up
3. find results

Pseudocode

```
backtrack(results, inputs, states, i) {
    if states is fesible {
        write results
        return
    }   

    try states for i 
    restore states
    try states for i+1
}
```

## Examples

n queens problem

```go
func nQueen(n int) int {
    
}

func backtrack(n int, count int, states []int, i int) {
    if i == n {
        // write results to count
        return
    }

    for j := 1; j <= n; j++ {
       // check conflicts 
       if not conflicts
       states = append(states, j)
       backtrack(n, count, states, i+1)
       states = states[:i+1]
    }
    
}
```


