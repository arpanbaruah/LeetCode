/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    
    if root == nil {
        return nil
    }
    
    stack := make([]*TreeNode , 0)
    
    list := make([]int, 0)
    
    stack = append(stack, root)
    
    for len(stack) > 0 {
        visit := stack[len(stack)-1]
        
        list = append(list, visit.Val)
        stack = stack[:len(stack)-1]
        
        if(visit.Right != nil) {
            stack = append(stack , visit.Right)
        }
        
        if(visit.Left != nil) {
            stack = append(stack , visit.Left)
        }
        
    }
    
    return list
}