/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    
    if root == nil {
        return nil
    }
    
    list := make([]int, 0)
    dfs(root, &list, 0)
    return list
    
}


func dfs(node *TreeNode, list *[]int, level int ) {
    
    if node == nil {
        return
    }
    
    if len(*list) == level {
        *list = append(*list, node.Val)
    }
    
    dfs(node.Right, list, level + 1)
    dfs(node.Left, list, level + 1)
    
    
}
