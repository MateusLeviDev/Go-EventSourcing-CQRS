//Author: Levi
//        jun/2023
// IRREGULAR ARRAY

package demo.Matriz;

import java.util.Scanner;

public class Matriz2 {

    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);
        System.out.println("Number of people interviewed: ");
        int numInterviewed = scanner.nextInt();

        String[][] names = new String[numInterviewed][];

        for (int i = 0; i<numInterviewed;i++) {
            System.out.println("How many children do you have? ");
            int numberOfChildren = scanner.nextInt();

            names[i] = new String[numberOfChildren];

            for (int j=0; j<numberOfChildren; j++) {
                System.out.println("What´s their name(s)?");
                names[i][j] = scanner.next();
            }
        }

        for (int i = 0; i<numInterviewed;i++) {
            System.out.println("Pessoa " + (i+1) + " tem " + names[i] + "filhos");
            for (int j=0; j<names[i].length; j++) {
                System.out.println(names[i][j]);
            }
        }

        scanner.close();
    }

}
