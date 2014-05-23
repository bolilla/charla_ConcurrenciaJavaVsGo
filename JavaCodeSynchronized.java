  public class SynchronizedCounter {
    private long  c  = 0;
    public synchronized void increment() { // HL
      c++;
    }
    public synchronized void decrement() { // HL
      c--;
    }
    public synchronized long value() { // HL
      return c;
    }
  }
