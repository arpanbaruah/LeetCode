/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalOrder(root *TreeNode) [][]int {
   
   
    
    if root == nil {
        return nil
    }
    
    
     //queue := make([]*TreeNode , 0)
    
    //level := make([]int,  0)
    
  //  queue = append(queue, root)
//    level = append(level, 0)
    
    queue := []*TreeNode{root}
    level := []int{0}
    
    var min , max int
    
    hMap := make(map[int][]int)
    
    for len(queue)>0 {
        
        node := queue[0]
        val := level[0]
        
        lst := hMap[val]
        if lst == nil {
            lst = []int{node.Val}
        } else {
            lst = append(lst, node.Val)
        }
        hMap[val] = lst
        
        queue = queue[1:]
        level = level[1:]
        
        if min > val {
            min = val
        }
        
        if max < val {
            max = val
        }
        
        if node.Left != nil {
            queue = append(queue, node.Left)
            level = append(level, val - 1)
        }
        
        if node.Right != nil {
            queue = append(queue, node.Right)
            level = append(level, val + 1)
        }
    }
    
    fList := make([][]int , 0)
    
    for i := min; i <= max; i++ {
        
        fList = append(fList, hMap[i])
    }
    
    return fList
    
    
}