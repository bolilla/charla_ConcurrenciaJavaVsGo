package main
import "fmt"
//slide1 BEGIN OMIT
func main() {
	c := newContador()
	fmt.Println("Contador", c.leer())
	c.sumar(42)
	fmt.Println("Contador", c.leer())
	c.matar()
	fmt.Println("Contador", c.leer())
}
type contador struct {
	l, s chan int  //leer y sumar
	v    int       //valor
}
func (c *contador) leer() int {
	return <-c.l // HL
}
func (c *contador) sumar(in int) {
	c.s <- in // HL
}
func (c *contador) matar() {
	close(c.l)
	close(c.s) // HL
}
//SLIDE CHANGE OMIT
func newContador() *contador {
	res := new(contador)
	res.l = make(chan int)
	res.s = make(chan int)
	go res.gestionar()
	return res
}

func (c *contador) gestionar() {
	for {
		select {
		case i, abierto := <-c.s: // HL
			if !abierto{
				break
			}
			c.v += i
		case c.l <- c.v: // HL
		}
	}
}
//END OMIT
