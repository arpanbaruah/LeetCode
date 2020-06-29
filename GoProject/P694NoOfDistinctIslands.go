func numDistinctIslands(grid [][]int) int {
    
    
    set := make(map[string]struct{})
    
    var exist = struct{}{}
    
    var dfs func(grid [][]int, x, y,i,j int, str *string)
    
    dfs = func(grid [][]int, x, y,i,j int, str *string) {
        
         if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
        return
    }
    
    
    *str = *str + strconv.Itoa(i) + strconv.Itoa(j)
    grid[x][y] = 0
    
    
    dfs(grid, x + 1, y, i + 1, j , str)
    dfs(grid, x - 1, y, i - 1, j , str)
    dfs(grid, x, y - 1, i, j - 1 , str)
    dfs(grid, x, y + 1, i, j + 1 , str)
        
        
        
    }
    
    for i := 0 ; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            
            if(grid[i][j]==1){
                var str string
                dfs(grid, i , j, 0 , 0, &str)
                set[str]=exist
            }
            
            
        }
    }
        
    return len(set)
    
}
