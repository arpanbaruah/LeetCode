/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    
    if root == nil {
        return nil
    }
    
    fList := make([][]int, 0)
    list := make([]int , 0)
    
    var dfs func(node *TreeNode, sum int, res *[]int, path *[][]int)
    
    dfs = func(node *TreeNode, sum int, res *[]int, path *[][]int) {
        
        if node == nil {
            return
        }
        
        *res = append(*res, node.Val)
        sum = sum - node.Val
        
        if sum == 0 && node.Left == nil && node.Right == nil {
            result := append([]int{}, *res...)
            *path = append(*path, result)
             *res = (*res)[:len(*res)-1]
            return
        }
        
        dfs(node.Left, sum, res, path)
        dfs(node.Right, sum, res, path)
        
        *res = (*res)[:len(*res)-1]
    }
    
   
    dfs(root, sum, &list, &fList)
    
    
    return fList
 
    }
