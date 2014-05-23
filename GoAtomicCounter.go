package main

import "fmt"

//BEGIN OMIT
import "sync/atomic"

type contador int32

func (c *contador) incrementar() int32 {
	return atomic.AddInt32((*int32)(c), 1) // HL
}

func (c *contador) decrementar() int32 {
	return atomic.AddInt32((*int32)(c), -1) // HL
}

func (c *contador) get() int32 {
	return atomic.LoadInt32((*int32)(c)) // HL
}

//END OMIT
func main() {
	var cont contador
	fmt.Println("Pequeño Padawan, el contador es", cont)
	cont.incrementar()
	fmt.Println("Pequeño Padawan, el contador es", cont)
	cont.decrementar()
	fmt.Println("Pequeño Padawan, el contador es", cont)
}
