  package main
  
  import (
    "fmt"
    "time"
  )
  
  func otraFuncion(sincronizador chan bool) {
    time.Sleep(4000 * time.Millisecond)
    fmt.Println("¡Ya va!")
    sincronizador <- true // HL
  }
  func main() {
    llamameCuandoTermines := make(chan bool)
    go otraFuncion(llamameCuandoTermines) // HL
    fmt.Println("¡Vete viniendo!")
    <-llamameCuandoTermines // HL
  }
