//Author: Levi
//        jun/2023

package demo.TestDataStructure;

import demo.DataStructure.Array;

public class clas01 {

    public static void main(String[] args) {

        Array array = new Array(10);
//        array.addElement("Element 1");
//        array.addElement("Element 2");
//        array.addElement("Element 3");
//
//        System.out.println(array.getSize());
//        System.out.println(array);
//        System.out.println(array.findElement(0));

        array.addElement("B");
        array.addElement("C");
        array.addElement("E");
        array.addElement("F");
        array.addElement("G");

        System.out.println(array);

        array.addElement(0, "A");

        System.out.println(array);
    }
}
