package main

import (
	"fmt"
	"sync"
	"time"
)

type curreloType int

type funcionDeTrabajo func(int, curreloType)

func main() {
	var numTrabajadores int = 3
	canalTrabajo, puntoDeEncuentro := pool(numTrabajadores, andarCaminos) // HL
	for i := 0; i < 10; i += 1 {
		fmt.Println("Enviando tarea:", i)
		canalTrabajo <- curreloType(i) // HL
	}
	close(canalTrabajo) // HL
	fmt.Println("Tareas enviadas. Esperando a que terminen")
	puntoDeEncuentro.Wait() // HL
	fmt.Println("Fin")
}

//SLIDE CHANGE OMIT

func pool(tamaño int, funcion funcionDeTrabajo) (chan curreloType, *sync.WaitGroup) { //con 'ñ'
	resultChan := make(chan curreloType, tamaño) // HL
	var resultEncuentro = new(sync.WaitGroup)
	resultEncuentro.Add(tamaño)
	for i := 0; i < tamaño; i += 1 {
		go worker(resultChan, i, resultEncuentro, funcion) // HL
	}
	return resultChan, resultEncuentro
}

func worker(entradaDeCurro chan curreloType, id int, encuentro *sync.WaitGroup, funcion funcionDeTrabajo) {
	for curro := range entradaDeCurro { // HL
		funcion(id, curro) // HL
	} // HL
	fmt.Println("Fin Padawan", id)
	encuentro.Done() // HL
}

func andarCaminos(curranteId int, curro curreloType) { // HL
	aDormir := (100 + 200*time.Duration(int(curro)%4)) * time.Millisecond
	time.Sleep(aDormir)
	fmt.Println("Padawan", curranteId, "andados caminos de la fuerza:", curro)
}

//END OMIT
