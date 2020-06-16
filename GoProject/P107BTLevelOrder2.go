/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    
    que := make([]*TreeNode, 0)
    
    if(root == nil){
        return nil
    }
    
    list := make([][]int , 0)
    
    que = append(que, root)
    
    for len(que) > 0 {
        
        temp := make([]int , 0)
        
        for size := len(que); size > 0; size-- {
            
            node := que[0]
            temp = append(temp, node.Val)
            que = que[1:]
            if(node.Left != nil) {
                que = append(que, node.Left)
            }
            if(node.Right != nil){
                que = append(que, node.Right)
            }
            
        }
        
        list = append([][]int{temp}, list...)
        
    }
    
    return list
    
}