package demo;

public class Main {
    public static void main(String[] args) {

        try {
            throw new Exc1();
        } catch (Exc0 e0) {
            System.out.println("Ex0 caugth");
        } catch (Exception e) {
            System.out.println("Exception caught");
        }
    }
}