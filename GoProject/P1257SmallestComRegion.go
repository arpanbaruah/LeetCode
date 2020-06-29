func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
    
    parent := make(map[string]string)
    
    for _, region := range regions {
        
        for _, reg := range region[1:] {
            parent[reg] = region[0]
        }
    }
    
    
    ancestor := make(map[string]struct{})
    var exists = struct{}{}
    
    for {
        ancestor[region1] = exists
        reg, ok := parent[region1]
        if ok {
            region1 = reg
        } else {
            break
        }
    }
    
    for {
        
        _, ok := ancestor[region2]
        if ok {
            return region2
        } else {
            region2 = parent[region2]
        }
        
    }
    
    return ""
    
}