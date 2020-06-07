func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    head := ListNode{0, nil}
    prev := &head
    
    total := 0
    rem := 0
    
    l1 = reverse(l1)
    l2 = reverse(l2)
    
    for l1 != nil && l2 != nil {
    
        total = l1.Val + l2.Val + rem
        prev.Next = &ListNode{total%10, nil}
        prev = prev.Next
        l1 = l1.Next
        l2 = l2.Next
        rem = total/10
        
    }
    
    for l1 != nil {
        
        total = l1.Val + rem
        prev.Next = &ListNode{total%10, nil}
        prev = prev.Next
        l1 = l1.Next
        rem = total/10
        
    }
    
    
    for l2 != nil {
        
        total = l2.Val + rem
        prev.Next = &ListNode{total%10, nil}
        prev = prev.Next
        l2 = l2.Next
        rem = total/10
        
    }
    
    
    if rem > 0 {
        prev.Next = &ListNode{rem, nil}
        prev = prev.Next
    }
    
    
    return reverse(head.Next)
    
}


func reverse(curr *ListNode) *ListNode {
    
    var prev *ListNode
    
    for curr != nil {
        
        pt := curr.Next
        curr.Next = prev
        prev = curr
        curr = pt
        
    }
    
    return prev
    
}