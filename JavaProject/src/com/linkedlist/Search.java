package com.linkedlist;

public class Search {

    public static int BinarySearch(int[] arr, int num, int start, int end){

        if(end >=1 ) {
            int mid = start + (end - 1) / 2;
            System.out.println("mid is " + mid);
         //   int result = 0;

            if (num == arr[mid]) {
                return mid;
            }

            if (num > arr[mid]) {
                //     start = mid + 1;
                return BinarySearch(arr, num, mid + 1, end);
            }
                System.out.println("Hello");
                return BinarySearch(arr, num, start, mid - 1);

        }

    return -1;

    }


    public static void main(String arg[]){
        int[] arr = {1,3, 5, 5,7,8};

        System.out.println("The result is " + BinarySearch(arr, 5,0, arr.length-1));

    }
}
