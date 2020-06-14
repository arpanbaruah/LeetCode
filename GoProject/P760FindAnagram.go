func anagramMappings(A []int, B []int) []int {
    
    mp := make(map[int][]int)
    for i, v := range B {
        mp[v] = append(mp[v], i)
    }
    
    res := make([]int , len(B))
    
    for i, v := range A {
        temp := mp[v]
        res[i] = temp[0]
        temp = temp[1:]
    }
    
    return res
}}+