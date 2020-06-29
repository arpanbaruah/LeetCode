/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    
     h := height(root)
    
/*    nodes := 0
    
  //  h := height(root)
  //  fmt.Println(h)
    
    for root != nil {
        
        if height(root.Right) == h - 1 {
            nodes += (1 << h)
            root = root.Right
        } else {
            nodes += (1 << (h - 1))
            root = root.Left
        }
        h = h-1
    }
    
    return nodes
        
        
    } */
   
    if h < 0 {
        return 0
    }
    
    if height(root.Right) == (h - 1) {
        
        return  ( 1 << h ) + countNodes(root.Right)
    } else {
        
        return (1 << (h - 1)) + countNodes(root.Left)
    }
}
    


func height(root *TreeNode) int {
    
    if root == nil {
        return -1;
    }
    
    return 1 + height(root.Left)
    
}