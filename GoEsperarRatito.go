  package main
  import (
    "fmt"
    "time"
  )
  func main() {
    frasecitas := []string{
      "El miedo es el camino hacia el lado oscuro",
      "El miedo lleva a la ira",
      "La ira lleva al odio",
      "El odio lleva al sufrimiento",
      "El sufrimiento al lado oscuro",
    }
    for _, frase := range frasecitas {
      time.Sleep(1000 * time.Millisecond) // HL
      fmt.Println(frase)
    }
  }
