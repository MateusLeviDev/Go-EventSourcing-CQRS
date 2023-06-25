//Author: Levi
//        jun/2023

package demo;

import java.util.Scanner;

public class numberOne {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int firstValue = scanner.nextInt();
        int secondValue = scanner.nextInt();
        int thirdValue = scanner.nextInt();
        int sum = firstValue + secondValue + thirdValue;

        System.out.println(sum);


        scanner.close();
    }
}