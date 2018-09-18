package com.linkedlist;

import java.util.*;

public class Container11 {


        public static int maxArea(int[] height) {


            int total = 0;
            int sumFront = 0;
            int sumBack = 0;

            Deque<Integer> newHeight = new LinkedList<Integer>();

            for(int i = 0; i<height.length-1; i++){
                for(int j = i+1; j<height.length; j++){
                    System.out.println("print element 1 " + height[i]);
                    System.out.println("print next element " + height[j]);
                    if(height[i] >= height[j]) {
                     //   int temp = height[i];
                        sumFront+=height[i];
                        System.out.println("the total sum is " + sumFront);
                    }
                }

                for(int j = i-1; j>=0; j--){
                    System.out.println("print element 2 " + height[i]);
                    System.out.println("print next element 2 " + height[j]);
                    if(height[i] <= height[j]) {
                        //   int temp = height[i];
                        sumBack+=height[i];
                        System.out.println("the total sum is " + sumBack);
                    }
                }

                total = Math.max(Math.max(sumFront, sumBack),total);

                sumFront = 0;
                sumBack = 0;
            }


            return total;
        }

        public static void main (String args[]){

            int[] val = {1,7,6,2,5,4,8,3,8};

            System.out.println("The total is " +  maxArea(val));
        }
}


