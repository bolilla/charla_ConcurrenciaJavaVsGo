  public class Prueba extends Thread {
  
    @Override
    public void run() {
      try {
        Thread.sleep(4000);
        System.out.println("¡Ya va!");
      } catch (InterruptedException e) {
        e.printStackTrace();
      }
    }
  
    public static void main(String[] args) throws InterruptedException {
      Prueba hilo = new Prueba();
      hilo.start();
      System.out.println("¡Vete viniendo!");
      hilo.join(); // HL
    }
  }
