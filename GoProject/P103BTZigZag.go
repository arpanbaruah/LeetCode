/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    
    if(root == nil) {
        return nil
    }
    
    que := make([]*TreeNode, 0)
    fList:= make([][]int , 0)
    que = append(que, root);
    dir := 1
    
    for len(que) > 0 {
        
        tList := make([]int, 0)
        
        for size := len(que); size > 0; size -- {
            
            temp := que[0]
            que = que[1:]
            if dir > 0 {
                tList = append(tList, temp.Val)
            } else {
                tList = append([]int{temp.Val}, tList...)
            }
            
            
            if temp.Left != nil {
                que = append(que, temp.Left)
            }
            
             if temp.Right != nil {
                que = append(que, temp.Right)
            }
            
            
        }
        
        dir *= -1
        
        fList = append(fList, tList)
        
        
    }
    
    return fList
    
}