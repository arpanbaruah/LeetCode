/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    
    return isValid(root, nil, nil)
    
}


func isValid(root , lower , upper *TreeNode) bool{
    
    if(root == nil) {
        return true
    }
    
    if(lower != nil && root.Val <= lower.Val) {
        return false
    }
    if(upper != nil && root.Val >= upper.Val) {
        return false
    }
    
    return isValid(root.Left, lower, root) && isValid(root.Right, root, upper)
}