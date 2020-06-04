public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        
        int rem = 0;
        int total = 0;
        ListNode head = new ListNode(0);
        ListNode prev = head;
        ListNode nl1 = reverse(l1);
        ListNode nl2 = reverse(l2);
        while(nl1 != null && nl2 != null){
            
            total = nl1.val + nl2.val + rem;
            prev.next = new ListNode(total%10);
            prev = prev.next;
            rem = total/10;
            nl1 = nl1.next;
            nl2 = nl2.next;
            
        }
        
        while(nl1 != null){
            total = nl1.val + rem;
            prev.next = new ListNode(total%10);
            prev = prev.next;
            rem = total/10;
            nl1 = nl1.next;
        }
        
        while(nl2 != null){
            total = nl2.val + rem;
            prev.next = new ListNode(total%10);
            prev = prev.next;
            rem = total/10;
            nl2 = nl2.next;
        }
        
        if(rem > 0){
           // ListNode curr = new ListNode(rem);
            prev.next = new ListNode(rem);
            prev = prev.next;
        }
        
        return reverse(head.next);
        
        
        
    }
    
    public ListNode reverse(ListNode curr){
        
        ListNode prev = null;
    
        
        while(curr != null){
            
            ListNode pt = curr.next;
            curr.next = prev;
            prev = curr;
            curr = pt;
            
        }
        
        return prev;
        
    }