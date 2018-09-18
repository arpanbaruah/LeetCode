package com.linkedlist;

import java.util.*;

public class testLinkedList {

    public static void main(String args[]){

      /*  MyLinkedList newList = new MyLinkedList();
        newList.addAtHead(2);
        newList.addAtHead(4);
        newList.addAtHead(5);
        newList.addAtIndex(2,1);
        newList.addAtHead(3);
      //  newList.addAtTail(4);
        newList.display();
        newList.deleteAtIndex(0);
        newList.display();
//(2 -> 4 -> 3) + (5 -> 6 -> 4)

      ListNode l1 = new ListNode(2);
      ListNode f1 = l1;
      l1.next = new ListNode(4);
      l1 = l1.next;
      l1.next = new ListNode(3);

        ListNode l2 = new ListNode(5);
        ListNode f2 = l1;
        l2.next = new ListNode(6);
        l2 = l2.next;
        l2.next = new ListNode(4);

        ListNode finalSol = Solution.addTwoNumbers(f1,f2);

      while(finalSol!=null){
          System.out.println("The values are " + finalSol.val);
          finalSol = finalSol.next;
      } */

      Deque<String> deque = new LinkedList<String>();

      deque.add("Hello");
      deque.addLast("bolo");
      deque.addFirst("Sabko");


      for(String x : deque){
          System.out.println(x);
      }



    }
}
