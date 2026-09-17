package main

import (
	"fmt"
	"sync"
)

// Objetivo: Implementar una versión del problema de los Filósofos Comensales.
// Hay 5 filósofos y 5 tenedores (recursos). Cada filósofo necesita 2 tenedores para comer.
// Estrategia segura: imponer un **orden global** al tomar los tenedores (primero el menor ID, luego el mayor)
// para evitar deadlock. También puedes limitar concurrencia (ej. mayordomo).
// TODO: completa la lógica de toma/soltado de tenedores y bucle de pensar/comer.

type tenedor struct{ mu sync.Mutex }

func filosofo(id int, izq, der *tenedor, wg *sync.WaitGroup) {
	// TODO: desarrolla el código para el filósofo

	//se agrega wg.Done() para indicar a waitgroup que ya termino
	defer wg.Done()
	fmt.Printf("[filósofo %d] satisfecho\n", id)

	// se agregan variables para guarda los canales izq y der
	primero := izq
	segundo := der

	// se agrega condicional para determinar por que lado coger el tenedor
	if id%2 == 0 {
		primero = der
		segundo = izq
	}
	// se llama a la func pensar para pasar el id del filosofo
	pensar(id)

	// Tomar el primer tenedor.
	primero.mu.Lock()
	fmt.Printf("[filósofo %d] toma un tenedor\n", id)

	// Tomar el segundo tenedor.
	segundo.mu.Lock()
	fmt.Printf("[filósofo %d] toma el segundo tenedor\n", id)

	// se llama la func comer
	comer(id)

	// se libera para que otra goroutine la pueda usar
	segundo.mu.Unlock()
	primero.mu.Unlock()

	// se imprime que el filosofo suelta los tenedores y que esta satisfecho
	fmt.Printf("[filósofo %d] suelta los tenedores\n", id)
	fmt.Printf("[filósofo %d] satisfecho\n", id)
}

func pensar(id int) {
	fmt.Printf("[filósofo %d] pensando...\n", id)
	// TODO: simular tiempo de pensar

}

func comer(id int) {
	fmt.Printf("[filósofo %d] COMIENDO\n", id)
	// TODO: simular tiempo de pensar

}

func main() {
	const n = 5
	var wg sync.WaitGroup
	wg.Add(n)

	// crear tenedores
	forks := make([]*tenedor, n)
	for i := 0; i < n; i++ {
		// TODO: inicializar cada tenedor i

	}

	// lanzar filósofos
	for i := 0; i < n; i++ {
		izq := forks[i]
		der := forks[(i+1)%n]
		// TODO: lanzar goroutine para el filósofo i

	}

	wg.Wait()
	fmt.Println("Todos los filósofos han comido sin deadlock.")
}
