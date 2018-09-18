package com.linkedlist;

public class Problem74 {

        public static Boolean searchInsertBinary(int[] nums, int target, int start, int end   ){
            int mid = (start + end)/2;
            if(target == nums[mid]) return true;
            if(target > nums[mid]) return searchInsertBinary(nums,target, mid+1, end );
            if(target < nums[mid]) return searchInsertBinary(nums, target, start, mid-1);
            return false;

        }



    public static Boolean searchMatrix(int[][] matrix, int target) {

            int start = 0;
            int end = matrix.length-1;
        System.out.println("The end is "+ end);
            int size = matrix[0].length-1;

        System.out.println("The size is "+ size);
            if(target < matrix[start][start]) return false;
            if(target > matrix[end][size]) return false;
            return searchMatrixBinary(matrix, target, start, end, size);

    }

    public static Boolean searchMatrixBinary(int[][] matrix, int target, int start, int end, int size){

            int mid = (start + end)/2;
            if(target >= matrix[mid][0] && target <= matrix[mid][size]) return searchInsertBinary(matrix[mid], target, mid, size);
            if(target < matrix[mid][0]) { System.out.println("target < matrix " ); return searchMatrixBinary(matrix, target, start, mid-1,size); }

            System.out.println("target > matrix " );
            return searchMatrixBinary(matrix, target, mid+1, end,size);
    }


        public static void main (String arg[]){

            System.out.println("Hello");
            int[][] matrix = {{1, 3, 5, 7},
                    {10, 11, 16, 20},
                    {23, 30, 34, 50}};
            int target = 18;
            System.out.println("If the value exist " + searchMatrix(matrix, target));
        }
    }

    /*
    Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true


     */