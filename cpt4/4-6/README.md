# 灵活选择键值

## 例题

447 [回旋镖的数量](https://leetcode-cn.com/problems/number-of-boomerangs/)

利用i为轴，以到其的距离为key，value就是距离i点相同的个数

之后再用排列组合的思路求解

**注意**

4-6_477_number_of_boomerangs/number_of_boomerangs.go
golang函数 作为变量貌似比较消耗性能

## 习题

149 [直线上最多的点数](https://leetcode-cn.com/problems/max-points-on-a-line/)


# 关于golang的一些东西

## 除法

-1 % -2 == -1

1 % -2 == 1

## 斜率的计算

使用最大公约数

```go
package aa
func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
```
