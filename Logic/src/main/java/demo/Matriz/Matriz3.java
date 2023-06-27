//Author: Levi
//        jun/2023

package demo.Matriz;

import java.util.Random;
import java.util.Scanner;

public class Matriz3 {

    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);

        int[][] array = new int[4][4];
        Random random = new Random();

        for (int i=0; i<array.length;i++) {
            for (int j=0; i<array[i].length; j++) {
                array[i][j] = random.nextInt(10);
            }
        }

        int high = 0;
        int line = 0;
        int col = 0;
        for (int i=0; i<array.length;i++) {
            for (int j=0; i<array[i].length; j++){
                if (array[i][j] > 0) {
                     high = array[i][j];
                     line = i;
                     col = j;
                }
            }
        }

        System.out.println("Hugh number: " + high);
        System.out.println("Line: " + line);
        System.out.println("Column: " + col);

        scanner.close();
    }
}
