/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    
    if(root == nil){
        return true
    }
    return isHelper(root, root)
    
}


func isHelper(left *TreeNode, right *TreeNode) bool{
    
    if(left == nil && right == nil) {
        return true
    }
    
    if(right == nil || left == nil) {
        return false
    }
    
    return left.Val == right.Val && isHelper(left.Left, right.Right) && isHelper(left.Right, right.Left)
}