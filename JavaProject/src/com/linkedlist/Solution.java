package com.linkedlist;


class Solution {
    public static ListNode addTwoNumbers(ListNode l1, ListNode l2) {

        ListNode finalNumber = new ListNode(0);
        ListNode d = finalNumber;

        int number = 0;
        int remainder = 0;
        while (l1 != null || l2 != null) {


            if (l1 != null) {
                number += l1.val;
                System.out.println("1. The number is " + number);
                l1 = l1.next;
            }

            if (l2 != null) {
                number += l2.val;
                System.out.println("2. The number is " + number);
                l2 = l2.next;
            }

            d.next = new ListNode(number % 10);
            //remainder = number%10;
            //finalNumber.setData(remainder);
            d = d.next;
            number = number / 10;


        }

        if (number % 10 == 1) {
            d.next = new ListNode(1);
            d = d.next;
        }

        return finalNumber.next;
    }


    public static void main(String args[]) {

      /*  MyLinkedList newList = new MyLinkedList();
        newList.addAtHead(2);
        newList.addAtHead(4);
        newList.addAtHead(5);
        newList.addAtIndex(2,1);
        newList.addAtHead(3);
      //  newList.addAtTail(4);
        newList.display();
        newList.deleteAtIndex(0);
        newList.display(); */
//(2 -> 4 -> 3) + (5 -> 6 -> 4)

        ListNode l1 = new ListNode(2);
        ListNode f1 = l1;
        l1.next = new ListNode(4);
        l1 = l1.next;
        l1.next = new ListNode(3);

        ListNode l2 = new ListNode(5);
        ListNode f2 = l2;
        l2.next = new ListNode(6);
        l2 = l2.next;
        l2.next = new ListNode(4);

        ListNode finalSol = addTwoNumbers(f1, f2);

        while (finalSol != null) {
            System.out.println("The values are " + finalSol.val);
            finalSol = finalSol.next;
        }


    }

}