//Author: Levi
//        jun/2023

package demo.Logic;

import java.util.Scanner;

public class numberTwo {

    public static void main(String[] args){
        Scanner scanner = new Scanner(System.in);

        int value = scanner.nextInt();
        double div = value/5;

        System.out.println(div);

        scanner.close();
    }
}
