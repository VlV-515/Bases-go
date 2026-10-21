package mutex

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ! Mutex protege datos compartidos: solo una goroutine a la vez entra a la seccion critica
// ! RWMutex permite muchos lectores a la vez, o un solo escritor
// ! Para ver errores de concurrencia: go run -race .

// Contador guarda el mutex junto al dato que protege (forma idiomatica).
// El receptor siempre es puntero: copiar un struct con mutex es un bug (go vet lo detecta).
type Contador struct {
	mu sync.Mutex
	n  int
}

func (c *Contador) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock() // defer libera aunque haya un return temprano o un panic
	c.n++
}

func (c *Contador) Valor() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func App() {
	//* Sin mutex: data race
	// Descomenta para verlo: el resultado casi nunca es 1000 y 'go run -race .' lo reporta
	// var sinProteger int
	// var wgRace sync.WaitGroup
	// for i := 0; i < 1000; i++ {
	// 	wgRace.Add(1)
	// 	go func() {
	// 		defer wgRace.Done()
	// 		sinProteger++
	// 	}()
	// }
	// wgRace.Wait()
	// fmt.Println("Sin mutex:", sinProteger)

	//* Con Mutex: siempre da 1000
	var contador Contador
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.Incrementar()
		}()
	}
	wg.Wait()
	fmt.Println("Con mutex:", contador.Valor())
	//--------------------------------
	//* Con atomic: para un contador simple es mas ligero que un mutex
	var atomico atomic.Int64
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomico.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println("Con atomic:", atomico.Load())
	//--------------------------------
	//* RWMutex: un escritor y varios lectores que leen varias veces
	var valor int
	var mu sync.RWMutex //Read Write Mutex

	//* Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			mu.Lock()
			valor++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	//* Readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				mu.RLock()
				fmt.Printf("Lector %d ve: %d\n", id, valor)
				mu.RUnlock()
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Esperamos a que terminen todas las goroutines, sin Sleep arbitrario
	wg.Wait()
}
