var list []int

func inorderTraversal(root *TreeNode) []int {
    
    list = make([]int, 0)
    dfs(root)
    return list
}

func dfs(root *TreeNode) {
    
    if root != nil {
        
        dfs(root.Left)
        list = append(list, root.Val)
        dfs(root.Right)
        
    }
    
    
}