//Author: Levi
//        jun/2023

package demo.Array;

import java.util.Scanner;

public class Array {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
//        double[] temperature = new double[365];
//        temperature[0] = 40;
//        temperature[1] = 41;
//        temperature[2] = 42;
//        temperature[3] = 43;
//        temperature[4] = 44;
//
//        for(double temp : temperature) {
//            System.out.println(temp);
//        }

        int[] firstArray = new int[5];
        int[] secondArray = new int[firstArray.length];

        for(int i=0; i<firstArray.length; i++) {
            System.out.println("Value: " + i); //value on array [0] = 1 example
            firstArray[i] = scanner.nextInt();
            secondArray[i] = firstArray[i];
        }

        System.out.println("First: ");
        for(int i = 0; i < firstArray.length; i++) {
            System.out.print(firstArray[i] + ", ");
        }
        System.out.println();

        System.out.println("Second: ");
        for(int i = 0; i < firstArray.length; i++) {
            System.out.print(secondArray[i] + ", ");
        }
            scanner.close();

    }
}
