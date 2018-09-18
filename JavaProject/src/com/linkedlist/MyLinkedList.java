package com.linkedlist;

class MyLinkedList {

    public Node first;
    public Node last;
    public int size;
    /** Initialize your data structure here. */
    public MyLinkedList() {
        this.size = 0;
        first = null;
        last = null;

    }

    /** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
    public int get(int index) {

        Node ptr = first;
        for(int i =0;i < index;i++){

            ptr = ptr.getNode();
        }

        return ptr.getData();

    }

    /** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
    public void addAtHead(int val) {

        Node newNode = new Node(val);
        size ++;
        if(first == null) {
            first = newNode;
            newNode = last;
        }
        else {
            newNode.setNext(first);
            first = newNode;
        }

    }

    /** Append a node of value val to the last element of the linked list. */
    public void addAtTail(int val) {

        Node newNode = new Node(val);
        size++;
        if(first==null){
            first = newNode;
            newNode = last;
        }
        else {
            System.out.println("The value is " + newNode.getData());

            last.setNext(newNode);
            System.out.println("The value is 1  " + newNode.getData());
            last = newNode;
            System.out.println("The value is 2  " + newNode.getData());
        }

    }

    /** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
    public void addAtIndex(int index, int val) {

        Node newNode = new Node(val);
        if (index < size) {

            Node ptr = first;
            for (int i = 0; i < index-1; i++) {
                ptr = ptr.getNode();
            }

            Node temp = ptr.getNode();
            ptr.setNext(newNode);
            newNode.setNext(temp);
            size++;

        }
        else {
            System.out.println("Hello");
        }
    }

    /** Delete the index-th node in the linked list, if the index is valid. */
    public void deleteAtIndex(int index) {


       if(index < size) {
           Node ptr = first;
           if (index == 0) {
               Node temp = first;
               first = first.getNode();
               temp.setNext(null);
           } else {
               for (int i = 0; i < index - 1; i++) {
                   ptr = ptr.getNode();

               }
               Node prev = ptr;
               System.out.println(prev.getData());
               Node temp = ptr.getNode();
               prev.setNext(temp.getNode());
               temp.setNext(null);

           }
       }
    }


    public void display()
    {
        System.out.print("\nSingly Linked List = ");
        if (size == 0)
        {
            System.out.print("empty\n");
            return;
        }
        if (first.getNode() == null)
        {
            System.out.println(first.getData() );
            return;
        }
        Node ptr = first;
        System.out.print(first.getData()+ "->");
        ptr = first.getNode();
        while (ptr.getNode() != null)
        {
            System.out.print(ptr.getData()+ "->");
            ptr = ptr.getNode();
        }
        System.out.print(ptr.getData()+ "\n");
    }
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * MyLinkedList obj = new MyLinkedList();
 * int param_1 = obj.get(index);
 * obj.addAtHead(val);
 * obj.addAtTail(val);
 * obj.addAtIndex(index,val);
 * obj.deleteAtIndex(index);
 */