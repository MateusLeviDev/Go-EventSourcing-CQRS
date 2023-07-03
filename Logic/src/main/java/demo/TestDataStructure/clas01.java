//Author: Levi
//        jun/2023

package demo.TestDataStructure;

import demo.DataStructure.ArrayObject;
import demo.DataStructure.Lista;

import java.util.ArrayList;

public class clas01 {

    public static void main(String[] args) {

        ArrayList<String> arrayList = new ArrayList<>();

        arrayList.add("A");
        arrayList.add("B");
        arrayList.add("C");

        System.out.println(arrayList);

        arrayList.add(0, "As");

        System.out.println(arrayList);

        boolean existe = arrayList.contains("A");
        if (existe) {
            System.out.println("Existe");
        } else {
            System.out.println("Não existe");
        }

        int pos = arrayList.indexOf("As");
        if (pos > -1) {
            System.out.println("Existe no array");
        } else {
            System.out.println("Não existe no array");
        }

    }
}
