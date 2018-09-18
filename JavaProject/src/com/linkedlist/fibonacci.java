package com.linkedlist;

import java.util.*;

public class fibonacci {

    public static int findFibonacci(int val){

        int result  = 0;
        if(val==0 || val == 1){
            result = 1;

        }

        else {
             result = findFibonacci(val-1) + findFibonacci(val-2);

        }
        return result;
    }


    public static int dynamicFibonacci(int n){

        int[] memo = new int[n+1];
        return dynamicFiboMemo(n+1,memo);

    }

    public static int dynamicFiboMemo(int n, int[] memo){

        int result;
        //int[] Memo = memo;
        if(memo[n]!=0){
            return memo[n];

        }
        if(n==0 ||n==1){
            return 1;
        }
        else{
            result = dynamicFiboMemo(n-1, memo) + dynamicFiboMemo(n-2,memo);
            memo[n] = result;
        }
        return result;

    }



    public static void main(String args[]){
        System.out.println("The fibonacci value is " + dynamicFibonacci(10));
    }
}
