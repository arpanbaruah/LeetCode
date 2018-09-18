package com.linkedlist;

import java.util.*;

public class Short3 {

    public static int lengthOfLongestSubstring(String s) {



        char[] y = s.toCharArray();
        Queue<Character> newQue = new LinkedList<Character>();
        int total  = 0;
        int finalVal = 0;
        //int temp = 0;



            for (Character val : y) {
               // System.out.println("the value is " + val);
                if (!newQue.contains(val)) {
                    System.out.println("The value add " + val);
                    newQue.add(val);
                    total = newQue.size();
                    System.out.println("the total is 1 " + total);
                } else {
                    System.out.println("The total is 2 " + total);
                    if(finalVal < total) finalVal = total;
                    Character temp = newQue.peek();
                    System.out.println("The value to be removed is " + temp);
                    System.out.println("The value of val is " + val);
                    while (temp != val) {

                        System.out.println("The removed value " + temp);
                        newQue.remove();
                        temp = newQue.peek();
                    }
                   // System.out.println("The removed value 2 " + temp);
                   // newQue.remove();
                  //  total = 1;
                    newQue.remove();
                    newQue.add(val);


                    System.out.println("The new Start value is " + newQue.peek());
                }
            }
            if(total > finalVal) { finalVal = total; }

        return finalVal;
    }


    public static void main(String arg[]) {

        String x = "adsdaas";



        System.out.println("The size is " + lengthOfLongestSubstring(x));
    }


}
