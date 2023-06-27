//Author: Levi
//        jun/2023

package demo;

public class FizzBuzz {
    public static void main(String[] args) {

        for (int i = 1; i <= 100 ; i++) {
            String result = switch (i % 15) {
                case 0 -> "FizzBuzz";
                case 3 -> "Fizz";
                case 5 -> "Buzz";
                default -> Integer.toString(i);
            };
            System.out.println(result);
        }
    }
}
