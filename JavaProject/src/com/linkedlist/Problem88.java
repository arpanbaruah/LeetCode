package com.linkedlist;

import java.util.Arrays;

public class Problem88 {

        public static void merge(int[] nums1, int m, int[] nums2, int n) {


            System.out.println("HELLO1");
            while(m > 0 && n > 0) {
                if (nums1[m - 1] >= nums2[n - 1]) {
                    System.out.println("The nums1 is " + nums1[m - 1]);
                    nums1[m + n - 1] = nums1[m - 1];
                    m--;
                } else {

                    System.out.println("The nums2 is " + nums2[n - 1]);
                    nums1[m + n - 1] = nums2[n - 1];
                    System.out.println("The final nums1 is " + Arrays.toString(nums1));
                    n--;
                    System.out.println("n " + n);
                }
            }
                while(n > 0){
                    nums1[n-1] = nums2[n-1];
                    n--;
                }

            System.out.println(Arrays.toString(nums1));
        }

        public static void main(String[] args){

            int[] nums1 = {0};
            int m = 0;
            int[] nums2 = {1};
            int n = 1;
            merge(nums1,m, nums2,n);

        }
    }

