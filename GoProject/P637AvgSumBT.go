/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    
    if root == nil {
        return nil
    }
    
    que := make([]*TreeNode, 0)
    
    que = append(que, root)
    
    var avg []float64
    var sum float64
    
    
    for len(que) > 0 {
        
        sum = 0.0
        lenQ := float64(len(que))
        for size := lenQ; size > 0; size-- {
            
            temp := que[0]
            que = que[1:]
            sum += float64(temp.Val)
            if(temp.Left != nil) {
                que = append(que, temp.Left)
            }
            
            if(temp.Right != nil) {
                que = append(que, temp.Right)
            }
            
        }
        
        avg = append(avg,sum/lenQ)
        
    }
    
    return avg
    
    
}