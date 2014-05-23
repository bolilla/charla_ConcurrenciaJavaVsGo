  public class SleepMessages {
      public static void main(String args[])
          throws InterruptedException {
          String importantInfo[] = {
              "Mares eat oats",
              "Does eat oats",
              "Little lambs eat ivy",
              "A kid will eat ivy too"
          };
  
          for (int i = 0;
               i < importantInfo.length;
               i++) {
              Thread.sleep(4000); // this.Wait(4000); // HL
              System.out.println(importantInfo[i]);
          }
      }
  }
