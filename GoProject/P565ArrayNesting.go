func arrayNesting(nums []int) int {
    
    maxV := 0
    
    //visited array for memoization
    visited := make([]bool, len(nums))
    
    for i, v := range nums{
        if(!visited[i]){
            
            count := 0
            start := v
            
            for count == 0 || start != v {
                
                start  = nums[start]
                count += 1
                visited[start] = true
        
            } 
            
            maxV = max(maxV, count)
            
        }
    }
        
        return maxV
    
    }
    
    func max(x int , y int) int {
        if x > y {
            return x
        } else {
            return y
        }
    }