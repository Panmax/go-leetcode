package main

func getLastMoment_20220111(n int, left []int, right []int) int {
    var ans int
    for _, val := range left {
        ans = max(ans, val)
    }
    for _, val := range right {
        ans = max(ans, n-val)
    }
    return ans
}