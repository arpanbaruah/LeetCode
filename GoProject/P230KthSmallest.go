/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    
    list := make([]int , 0)
    
    dfs(root, &list);
    
    return list[k-1];
    
}

func dfs(root *TreeNode, list *[]int) {
    
    if root == nil {
        return 
    }
    
    dfs(root.Left, list)
    *list = append(*list, root.Val)
    dfs(root.Right, list)
    
}