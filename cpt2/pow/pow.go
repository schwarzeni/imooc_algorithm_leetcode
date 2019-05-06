// 对某个数做n次方运算
package pow

// 使用递归，时间复杂度 log(n)
// 函数调用的次数
func betterPow(x float64, n int) (result float64){
    var (
        prevResult float64
    )

    if n == 0 {
        result = 1.0
        return
    }

    prevResult = betterPow(x, n/2)
    result = prevResult * prevResult

    if n % 2  != 0 {
        result = float64(x)  * result
    }
    return
}

func normalPow(x float64, n int) (result float64){
    result = 1
    for i := 0; i < n; i++ {
        result = result * x
    }
    return
}