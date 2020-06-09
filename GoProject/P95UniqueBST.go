/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    
    if n==0 {
        return nil
    }
    return generateDFS(1, n)
}

func generateDFS(start int, end int) []*TreeNode {
    
    var allTrees []*TreeNode = make([]*TreeNode, 0)
    
    if start > end {
        allTrees = append(allTrees, nil)
        return allTrees
    }
    
    for i := start; i <= end; i++ {
        
        leftT := generateDFS(start, i-1)
        rightT := generateDFS(i+1, end)
        
        for _, v := range leftT {
            
            for _, V := range rightT {
                
                curr := &TreeNode{i, nil, nil}
                curr.Left = v
                curr.Right = V
                allTrees = append(allTrees, curr)
                
            }
        }
        
    }
    
    return allTrees
}