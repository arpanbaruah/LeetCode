
//Defined Custom type ByLength
type ByLength []string


//3 Functions are required to be used sort.Sort method. Less, Swap and Len
//Len gives the length of the string
func (s ByLength) Len() int {
    return len(s)
}


//Compares 2 string based on the length
func (s ByLength) Less (i, j int) bool {
    
    return len(s[i]) < len(s[j])
}


//Swap uses to swap the words
func(s ByLength) Swap(i, j int) {
    
    s[i], s[j] = s[j], s[i]
}


func longestStrChain(words []string) int {
    
    //This will sort the String array using the above methods
    sort.Sort(ByLength(words))
    
    total := 0
    
    mapVal  := make(map[string]int)
    
    for _, word := range words {
        
        best := 0
        
        for i := 0; i < len(word); i++ {
            
            tWord := word[:i] + word[i+1 : len(word)]
            best = max(best, mapVal[tWord] + 1)
            
        }
        
        mapVal[word] = best
        total = max(total, best)
    }
    
    return total
    
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}