  public class Contador{
    private Lock cerrojo = new Lock();
    private long cont = 0;
    public long inc(){
      cerrojo.lock(); // HL
      cont += 1
      long resultado = cont
      cerrojo.unlock(); // HL
      return resultado;
    }
  }
