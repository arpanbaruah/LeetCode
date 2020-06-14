/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    
    if(root == nil) {
        return nil
    }
    
    fList := make([][]int, 0)
    
    queue := make([]*Node, 0)
    
    queue = append(queue, root)
    
    for len(queue) > 0 {
        
        temp := make([]int, 0)
        for size := len(queue); size > 0 ; size-- {
            
            node := queue[0]
            
            temp = append(temp, node.Val)
            
            queue = queue[1:]
            
            for _,v := range node.Children {
                
                queue = append(queue, v)
            }
            
        }
        
        fList = append(fList, temp)
        
    }
    
    return fList
    
}