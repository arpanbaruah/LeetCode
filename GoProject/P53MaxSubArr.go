func maxSubArray(nums []int) int {
    
    maxV := 0
    curMax := -2147483647
    
    for _,v := range nums {
        maxV = max(maxV + v, v)
        curMax = max(curMax, maxV)
        
    }
    return curMax
    
}

func max(x int, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}