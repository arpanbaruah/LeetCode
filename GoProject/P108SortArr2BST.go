/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    
    return sortedArray(nums, 0, len(nums)-1);
}

func sortedArray(nums []int, start , end int) *TreeNode {
    
    if start > end {
        return nil
    }
    
    mid := (start + end)/2
    
    node := &TreeNode{nums[mid], nil, nil}
    
    node.Left = sortedArray(nums, start, mid - 1)
    node.Right = sortedArray(nums, mid + 1, end)
    
    return node
    
}