/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    

    queue := make([]*TreeNode , 0)
    
    fList := make([][]int , 0)
    
    if root == nil {
        return fList
    }
    
    queue = append(queue, root)
    
    for len(queue) > 0{
        
        temp := make([]int ,0)
        for size := len(queue); size > 0 ; size-- {
            
            node := queue[0]
            queue = queue[1:]
            temp = append(temp, node.Val)
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            
        }
        
        fList = append(fList,  temp)
        
        
    }
    return fList
    
}

