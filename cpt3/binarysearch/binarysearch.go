// 二分查找法
package binarysearch

var NotFound = -1

func binarySearch(arr []int, size int, target int) (idx int) {
    var (
        l = 0
        r = size - 1 // 在 [l..r] 的范围内寻找target
        mid int
    )
    for l <= r {
        // 这个会整形溢出
        //mid = (l + r ) /2
        mid = l + (r - l) / 2
        if arr[mid] == target {
            idx = mid
            return
        }
        if target > arr[mid] {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    idx = NotFound
    return
}

