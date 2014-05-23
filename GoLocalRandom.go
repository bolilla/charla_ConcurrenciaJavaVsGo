package main

import "fmt"
import "math/rand"

func main() {
    s1 := rand.NewSource(42) // HL
    r1 := rand.New(s1)

    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()
}
