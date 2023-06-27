//Author: Levi
//        jun/2023

package demo.Array;

import java.util.Scanner;

public class Array3 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int[] array1 = new int[15];
        int[] array2 = new int[array1.length];

        for(int i=0; i< array1.length; i++) {
            System.out.println("Value: " + i);
            array1[i] = scanner.nextInt();
            array2[i] = array1[i] * array1[i];
        }

        System.out.println("First: ");
        for(int i = 0; i < array1.length; i++) {
            System.out.print(array1[i] + ", ");
        }
        System.out.println();

        System.out.println("Second: ");
        for(int i = 0; i < array2.length; i++) {
            System.out.print(array2[i] + ", ");
        }

        scanner.close();
    }
}
