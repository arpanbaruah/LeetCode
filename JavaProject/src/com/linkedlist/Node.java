package com.linkedlist;

public class Node {

    private int data;
    private Node next;

    public Node(int data){
        this.data=data;
        this.next = null;

    }

    public void setData(int newData){
        this.data = newData;
    }

    public void setNext(Node newNode){
        this.next= newNode;
    }

    public int getData(){

        return data;
    }


    public Node getNode(){
        return next;
    }
}
