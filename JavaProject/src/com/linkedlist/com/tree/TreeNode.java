package com.linkedlist.com.tree;

public class TreeNode {

    private TreeNode leftNode;
    private TreeNode rightNode;

    public TreeNode(int data){
        this.data = data;
    }


    public TreeNode findNode(int num){

        if(num == this.data) {
            return this;
        }
        if(num > this.data && (rightNode != null)){
            return rightNode.findNode(num);
        }
        if(leftNode != null) {
            return leftNode.findNode(num);
        }

        return null;
    }



    public TreeNode getLeftNode() {
        return leftNode;
    }

    public void setLightNode(TreeNode leftNode) {
        this.leftNode = rightNode;
    }

    public TreeNode getRightNode() {
        return rightNode;
    }

    public void setRightNode(TreeNode rightNode) {
        this.rightNode = rightNode;
    }

    public int getData() {
        return data;
    }


    private int data;


}
