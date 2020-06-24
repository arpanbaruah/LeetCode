func postorderTraversal(root *TreeNode) []int {
    
    if root == nil {
        return nil
    }
    
    stack := make([]*TreeNode, 0)
    
    list := make([]int , 0)
    
    stack = append(stack, root)
    
    for len(stack) > 0 {
        
        temp := stack[len(stack)-1]
        list = append([]int{temp.Val}, list...)
        
        stack = stack[:len(stack)-1]
        
        if temp.Left != nil {
            stack = append(stack, temp.Left)
        }
        
         if temp.Right != nil {
            stack = append(stack, temp.Right)
        }
        
    }
    
    return list
    
}