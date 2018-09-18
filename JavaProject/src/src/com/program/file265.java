package src.com.program;

public class file265 {

        public int minCost(int[][] costs) {

            int red=  costs[0][0];
            int blue = costs[0][1];
            int green = costs[0][2];
            int tempRed = 0;
            int tempBlue = 0;
            int tempGreen = 0;

            int m = 3;
            for (int i =0 ; i < 3; i++){

                red = Math.min(blue, green) + costs[i][0];
                blue = Math.min(red, green) + costs[i][1];

        }

        return(1);
    }
}
