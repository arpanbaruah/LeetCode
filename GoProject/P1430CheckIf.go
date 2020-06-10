func isValidSequence(root *TreeNode, arr []int) bool {
    
    if root == nil {
        return false;
    }
    
    return dfs(root, arr, 0);
}

func dfs(root *TreeNode, arr []int, depth int) bool {
    
    if root == nil || depth >= len(arr) || arr[depth] != root.Val {
        return false;
    }
    
    if root.Left == nil && root.Right  == nil {
        return depth + 1 == len(arr)
    }
    
    return dfs(root.Left, arr, depth + 1) || dfs(root.Right, arr, depth + 1)
    
}