//Author: Levi
//        jun/2023

package demo.Logic;

import java.util.Scanner;

public class numberThree {

    public static void main(String[] args){
        Scanner scanner = new Scanner(System.in);

        System.out.println("Informe dois valores:");
        int firstValue = scanner.nextInt();
        int secondValue = scanner.nextInt();
        int result = smallerNumber(firstValue, secondValue);

        System.out.println(result);

//        simple way without smallerNumber
//        int smaller = Math.min(firstValue, secondValue);
//        System.out.println(smaller);

        scanner.close();
    }

    public static int smallerNumber(int a, int b) {
        if(a < b) {
            return a;
        } return b;
    }
}