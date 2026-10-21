package concurrencia

import (
	"fmt"
	"sync"
	"time"
)

// ! Una goroutine NO es un hilo del sistema operativo: es una tarea ligera
// ! que el runtime de Go reparte entre pocos hilos (puede haber miles)
// ! Se utiliza la palabra reservada 'go' para crear una goroutine
// ! Channels se utilizan para comunicar cosas entre goroutines

func decirHola(canal chan<- string) {
	time.Sleep(1 * time.Second)
	canal <- "Hola desde goroutine"
}

func imprimirHola(canal <-chan string) {
	msg := <-canal
	fmt.Println(msg)
}

func App() {
	//* Canal sin bufer: enviar y recibir se esperan mutuamente (sincroniza)
	canal := make(chan string)
	go decirHola(canal)
	imprimirHola(canal)
	//--------------------------------
	//* Cerrar un canal y recorrerlo con range (siempre cierra quien envia)
	canal2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			canal2 <- i
		}
		close(canal2)
	}()

	for num := range canal2 {
		fmt.Println("Numero recibido:", num)
	}

	// Recibir con ok: ok es false cuando el canal esta cerrado y vacio
	v, ok := <-canal2
	fmt.Println("Canal cerrado -> valor:", v, "ok:", ok)
	//--------------------------------
	//* Canal con bufer: acepta hasta 3 envios sin que nadie reciba
	conBufer := make(chan int, 3)
	conBufer <- 1
	conBufer <- 2
	conBufer <- 3
	fmt.Println("Bufer -> len:", len(conBufer), "cap:", cap(conBufer))
	// Un cuarto envio bloquearia hasta que alguien reciba
	fmt.Println("Recibido del bufer:", <-conBufer)
	//--------------------------------
	//* WaitGroup: esperar a varias goroutines sin usar Sleep
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d termino\n", id)
		}(i)
	}
	wg.Wait()
	//--------------------------------
	//* Select: espera varios canales a la vez, con timeout
	lento := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		lento <- "respuesta tardia"
	}()

	select {
	case msg := <-lento:
		fmt.Println(msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout: la goroutine tardo demasiado")
	}
	//--------------------------------
	//* Deadlock (NO descomentar salvo para verlo una vez)
	// Enviar a un canal sin bufer sin que nadie reciba bloquea para siempre:
	// fatal error: all goroutines are asleep - deadlock!
	// bloqueado := make(chan int)
	// bloqueado <- 1
}
