func lengthOfLongestSubstring(s string) int {
    var m = make(map[rune]int)    
    start := 0
    maxV := 0    
    for i, v := range s {        
        _, ok := m[v]        
        if ok {
            start = max(start, m[v]+1)
        }      
        m[v] = i
        maxV = max(maxV, i - start + 1)
    }    
    return maxV    
}