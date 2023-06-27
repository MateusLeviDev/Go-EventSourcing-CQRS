//Author: Levi
//        jun/2023

package demo.Matriz;

public class Matriz1 {

    public static void main(String[] args) {

        int[][][] tridimensionalArray = new int[3][3][3];

        for (int i = 0; i < tridimensionalArray.length; i++) {
            for (int j = 0; j < tridimensionalArray[i].length; j++) {
                for (int k = 0; k < tridimensionalArray[i][j].length; k++) {
                    tridimensionalArray[i][j][k] = i + j + k;
                }
            }
        }
        int soma = 0;
        int somaPares = 0;
        int somaImpares = 0;
        for (int i = 0; i < tridimensionalArray.length; i++) {
            for (int j = 0; j < tridimensionalArray[i].length; j++) {
                for (int k = 0; k < tridimensionalArray[i][j].length; k++) {
                    soma += tridimensionalArray[i][j][k];

                    if (tridimensionalArray[i][j][k] % 2 == 0) {
                        somaPares += tridimensionalArray[i][j][k];
                    } else {
                        somaImpares += tridimensionalArray[i][j][k];
                    }
                }
            }
        }

        System.out.println("Sum: " + soma);
        System.out.println("SumEven: " + somaPares);
        System.out.println("SumOdd: " + somaImpares);
    }
}