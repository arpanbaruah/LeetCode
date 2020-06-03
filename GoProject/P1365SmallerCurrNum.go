func smallerNumbersThanCurrent(nums []int) []int {
    
    numsN := make([]int, len(nums))
    
    for i,val := range nums {
        count := 0
        for _, val2 := range nums {
            
            if val2 < val {
                count += 1
            }
        }
        numsN[i] = count
    }
    
    return numsN
    
}