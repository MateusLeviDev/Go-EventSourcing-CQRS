//Author: Levi
//        jun/2023

package demo;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

    public class numberFive {

        public static void main(String[] args){
            Scanner scanner = new Scanner(System.in);

            int numberOfValues = scanner.nextInt();
            List<Integer> values = new ArrayList<>();

            for (int i = 0; i < numberOfValues; i++) {
                int value = scanner.nextInt();
                values.add(value);
            }

//            java 8+ methods
//                stream(): https://docs.oracle.com/javase/8/docs/api/java/util/stream/Stream.html
            int sum = values.stream().mapToInt(Integer::intValue).sum();
            int avg = sum/numberOfValues;

            System.out.println("Result: " + avg);

            scanner.close();
        }
    }