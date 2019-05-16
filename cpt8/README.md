# 递归和回溯法

## 树形问题

### 例题

17 [电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)

- 字符串的合法性
- 空字符串
- 多个解的顺序

递归法 --> 回溯法 --> O(2^n) --> 复杂度为指数级 --> 暴力破解，效率比较低 

--> 升级：动态规划

--> 升级：剪枝法

---

## 什么是回溯

### 习题

93 [复原IP地址](https://leetcode-cn.com/problems/restore-ip-addresses/)

131 [分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/)

---

## 排列问题

### 例题

46 [全排列](https://leetcode-cn.com/problems/permutations/)

### 习题

47 [全排列 II](https://leetcode-cn.com/problems/permutations-ii/)

---

## 组合问题

### 例题

77 [组合](https://leetcode-cn.com/problems/combinations/)

回溯法剪枝

```
//for i := startIdx; i <= n; i++ {
for i := startIdx; i <= n-(k-len(arr))+1; i++ {
    combineHelper(n, k, i+1, append(arr, i), res)
}
```

### 习题

39 [组合总和](https://leetcode-cn.com/problems/combination-sum/)

40 [组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/  )

41 [组合总和 III](https://leetcode-cn.com/problems/combination-sum-iii/)

78 [子集](https://leetcode-cn.com/problems/subsets/)

90 [子集 II](https://leetcode-cn.com/problems/subsets-ii/)

401 [二进制手表](https://leetcode-cn.com/problems/binary-watch/)

---

## 二维平面上的回溯法

### 例题

79 [单词搜索](https://leetcode-cn.com/problems/word-search/)

搜索找不到进行下一个搜索时记得取消当前点的访问记录

TODO: https://leetcode-cn.com/problems/word-search-ii/
      https://leetcode-cn.com/problems/implement-trie-prefix-tree/

---

## floodfill算法

### 例题

200 [岛屿的个数](https://leetcode-cn.com/problems/number-of-islands/)

由于golang的测试代码是并发执行，所以放在函数外部的变量会被共享，所以不要把写类型数据放在函数外部

### 习题

130 [被围绕的区域](https://leetcode-cn.com/problems/surrounded-regions/)

417 [太平洋大西洋水流问题](https://leetcode-cn.com/problems/pacific-atlantic-water-flow/)

---

## 回溯法是经典人工智能的基础

### 例题

51 [N皇后](https://leetcode-cn.com/problems/n-queens/)

- 每一皇后占有独立的一行，一列，两斜线
- 一行一行逐行判断
- `col[i]` 表示i列已经被占用
- `dia1[i]` 正对角线 `2*n-1`个 i+j和相等
- `dia2[i]` 反对角线 `2*n-1`个 i-j+n-1相等

### 习题

37 [解数独](https://leetcode-cn.com/problems/sudoku-solver/)



