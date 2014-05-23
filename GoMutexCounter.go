package main

import "fmt"

//BEGIN OMIT
import "sync"
type contador struct {
	mu sync.Mutex
	x  int64
}
func (c *contador) incrementar() {
	c.mu.Lock() // HL
	defer c.mu.Unlock() // HL
	c.x += 1
}
func (c *contador) decrementar() {
	c.mu.Lock() // HL
	defer c.mu.Unlock() // HL
	c.x -= 1
}
func (c *contador) valor() (resultado int64) {
	c.mu.Lock() // HL
	defer c.mu.Unlock() // HL
	resultado = c.x
	return
}
//END OMIT
func main() {
	var cont contador
	fmt.Println("Pequeño Padawan, el contador es", cont.valor())
	cont.incrementar()
	fmt.Println("Pequeño Padawan, el contador es", cont.valor())
	cont.decrementar()
	fmt.Println("Pequeño Padawan, el contador es", cont.valor())
}
