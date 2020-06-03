func removeVowels(S string) string {
    
    vowel := "aeiou"
    for _,x := range vowel {
        S = strings.ReplaceAll(S, string(x), "")
    }
    
    return S
   
}