func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    lf := &ListNode{0, nil}
    curr := lf
    
    
    for l1 != nil && l2 != nil {
        
        if l1.Val > l2.Val {
            lf.Next = &ListNode{l2.Val, nil}
            lf = lf.Next
            l2 = l2.Next
        } else {
            lf.Next = &ListNode{l1.Val, nil}
            lf = lf.Next
            l1 = l1.Next
        }
    }
    
    for l1 != nil {
        
        lf.Next = &ListNode{l1.Val, nil}
        l1 = l1.Next
        lf = lf.Next
        
    }
    
    for l2 != nil {
        
        lf.Next = &ListNode{l2.Val, nil}
        l2 = l2.Next
        lf = lf.Next
        
    }
    
    return curr.Next
    
    
}