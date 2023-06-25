package demo;

import java.util.Scanner;

public class numberFour {

    public static void main(String[] args){
        Scanner scanner = new Scanner(System.in);

        int value = scanner.nextInt();
        double result = Math.pow(value, 2);

        System.out.println(result);

        scanner.close();
    }
}