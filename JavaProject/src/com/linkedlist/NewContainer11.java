package com.linkedlist;

import java.util.*;

public class NewContainer11 {


    public static class setNum {

        private int key;
        private int value;

        public setNum(int key, int value){
            this.key=key;
            this.value = value;
        }

        public int getKey() {
            return key;
        }

        public void setKey(int key) {
            this.key = key;
        }

        public int getValue() {
            return value;
        }

        public void setValue(int value) {
            this.value = value;
        }


    }

    public static int maxArea(int[] height) {

        int first = 0;
        int last = height.length-1;
        int sum = 0;

      for(int i =height.length-1 ; i >=0; i--){

          sum = Math.max(sum, Math.min(height[first], height[last])*i);
          if(height[first] >= height[last]){
              last--;
          }
          else {
              first ++;
          }


      }
        return sum;

    }


    public static void main(String args[]){
        int[] val = {11,1,7,3,2};

       System.out.println(maxArea(val));
    }

}
