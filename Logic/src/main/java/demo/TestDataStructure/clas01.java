//Author: Levi
//        jun/2023

package demo.TestDataStructure;

import demo.DataStructure.Array;
import demo.DataStructure.ArrayObject;

public class clas01 {

    public static void main(String[] args) {

        ArrayObject arrayObject = new ArrayObject(3);

        Contact contact1 = new Contact("Levi", "1123546", "mateus@gmail");
        Contact contact2 = new Contact("Levi", "1123546", "mateus@gmail");
        Contact contact3 = new Contact("Levi", "1123546", "mateus@gmail");

        arrayObject.addElement(contact1);
        arrayObject.addElement(contact2);
        arrayObject.addElement(contact3);

        System.out.println(arrayObject);

    }
}
