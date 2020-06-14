/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    
    return buildTreeDFS(preorder, inorder, 0 , len(preorder)-1, 0, len(inorder)-1)
    
}


func buildTreeDFS(preorder []int, inorder []int, preStart , preEnd, inStart, inEnd int) *TreeNode{
    
    if preStart > preEnd || inStart > inEnd {
        return nil
    }
    
    node := &TreeNode{preorder[preStart], nil, nil}
    
    inPos := 0
    
    for i:= inStart; i <= inEnd; i++ {
        
        if preorder[preStart]==inorder[i] {
            inPos = i
        }
    }
    
    node.Left = buildTreeDFS(preorder, inorder, preStart + 1, preEnd, inStart, inPos-1)
    node.Right = buildTreeDFS(preorder, inorder, preStart + (inPos-inStart) + 1, preEnd, inPos+1, inEnd)
    
    return node
    
    
    
}