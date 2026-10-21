package slicesavanzados

import (
	"encoding/json"
	"fmt"
	"slices"
)

// ! Un slice NO guarda datos: es una "ventana" con 3 partes
// !   1) puntero a un array  2) len (cuantos ve)  3) cap (cuantos caben desde donde empieza)
// ! Dos slices pueden ver el MISMO array, por eso modificar uno puede afectar al otro

func modificarElementos(s []int) {
	s[0] = -1 // comparte el array: el cambio SI se ve afuera
}

func agregarSinEfecto(s []int) {
	s = append(s, 100) // s es una copia de la ventana: el len de afuera no cambia
}

func agregarConRetorno(s []int) []int {
	return append(s, 100) // por eso append siempre se reasigna: s = append(s, x)
}

func App() {
	//* Anatomia: len y cap
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("s: %v len=%d cap=%d\n", s, len(s), cap(s))
	//--------------------------------
	//* Rebanar: s[inicio:fin] (el fin NO se incluye)
	sub := s[1:3]
	fmt.Printf("sub: %v len=%d cap=%d\n", sub, len(sub), cap(sub)) // cap llega hasta el final de s
	fmt.Println("s[:2]:", s[:2], "| s[2:]:", s[2:], "| s[:]:", s[:])
	//--------------------------------
	//* Rebanar COMPARTE memoria
	original := []int{1, 2, 3, 4, 5}
	ventana := original[1:3] // [2 3]
	ventana[0] = 99
	fmt.Println("Tras ventana[0]=99, original:", original)

	// append dentro de la capacidad escribe EN el array original
	ventana = append(ventana, 100) // pisa original[3]
	fmt.Println("Tras append en la ventana, original:", original)

	// Expresion completa s[inicio:fin:limite]: limita la capacidad
	// asi el siguiente append reserva un array nuevo y no pisa nada
	seguro := []int{1, 2, 3, 4, 5}
	recortado := seguro[1:3:3]
	recortado = append(recortado, 100)
	fmt.Println("Con limite de cap, seguro queda:", seguro, "| recortado:", recortado)
	//--------------------------------
	//* Como crece un slice con append
	// Los valores exactos de cap dependen de la version de Go
	var crece []int
	capAnterior := cap(crece)
	for i := 0; i < 10; i++ {
		crece = append(crece, i)
		if cap(crece) != capAnterior {
			fmt.Printf("append #%d -> len=%d cap=%d (reservo array nuevo)\n", i+1, len(crece), cap(crece))
			capAnterior = cap(crece)
		}
	}
	//--------------------------------
	//* make: crear con len y cap
	ceros := make([]int, 3) // len 3, ya tiene 3 ceros
	ceros = append(ceros, 7)
	fmt.Println("make([]int, 3) + append:", ceros) // error comun: quedan los 3 ceros

	reservado := make([]int, 0, 10) // len 0, cap 10: ideal si sabes cuantos vas a agregar
	reservado = append(reservado, 1, 2)
	fmt.Printf("make([]int, 0, 10): %v len=%d cap=%d\n", reservado, len(reservado), cap(reservado))
	//--------------------------------
	//* nil vs vacio (importa en JSON!)
	var nulo []int
	vacio := []int{}
	fmt.Println("nulo == nil:", nulo == nil, "| vacio == nil:", vacio == nil, "| ambos len 0:", len(nulo), len(vacio))
	nulo = append(nulo, 1) // append funciona con slice nil
	jsonNulo, _ := json.Marshal([]int(nil))
	jsonVacio, _ := json.Marshal([]int{})
	fmt.Println("JSON de nil:", string(jsonNulo), "| JSON de vacio:", string(jsonVacio))
	// En una API, devuelve []int{} si quieres que el cliente reciba [] y no null
	//--------------------------------
	//* copy: copia real, sin compartir memoria
	fuente := []int{1, 2, 3}
	destino := make([]int, len(fuente)) // destino necesita tener espacio (len)
	n := copy(destino, fuente)
	destino[0] = 100
	fmt.Println("copy copio", n, "elementos | fuente:", fuente, "| destino:", destino)

	// Asignar NO copia, comparte
	alias := fuente
	alias[0] = 555
	fmt.Println("Tras alias[0]=555, fuente:", fuente)
	//--------------------------------
	//* Pasar slices a funciones
	datos := []int{1, 2, 3}
	modificarElementos(datos)
	fmt.Println("Tras modificarElementos:", datos)
	agregarSinEfecto(datos)
	fmt.Println("Tras agregarSinEfecto (sin cambio):", datos)
	datos = agregarConRetorno(datos)
	fmt.Println("Tras agregarConRetorno:", datos)
	//--------------------------------
	//* Borrar, insertar y filtrar
	base := []int{10, 20, 30, 40, 50}

	// Forma manual: borrar el indice 2 (cuidado: modifica el array de base)
	manual := slices.Clone(base)
	i := 2
	manual = append(manual[:i], manual[i+1:]...)
	fmt.Println("Borrar indice 2 (manual):", manual)

	// Con el paquete slices
	conSlices := slices.Clone(base)
	conSlices = slices.Delete(conSlices, 1, 2) // borra el rango [1:2]
	fmt.Println("slices.Delete(1,2):", conSlices)
	conSlices = slices.Insert(conSlices, 1, 99) // inserta 99 en el indice 1
	fmt.Println("slices.Insert(1, 99):", conSlices)

	// Filtrar: se arma un slice nuevo
	var pares []int
	for _, n := range []int{1, 2, 3, 4, 5, 6} {
		if n%2 == 0 {
			pares = append(pares, n)
		}
	}
	fmt.Println("Pares:", pares)
	//--------------------------------
	//* Utilidades del paquete slices
	numeros := []int{5, 2, 8, 1}
	fmt.Println("Contains 8:", slices.Contains(numeros, 8), "| Index de 8:", slices.Index(numeros, 8))
	slices.Sort(numeros)
	fmt.Println("Sort:", numeros, "| Min:", slices.Min(numeros), "| Max:", slices.Max(numeros))
	slices.Reverse(numeros)
	fmt.Println("Reverse:", numeros)
	// Los slices NO se comparan con ==, se usa slices.Equal
	fmt.Println("Equal:", slices.Equal([]int{1, 2}, []int{1, 2}))
	//--------------------------------
	//* Slice de slices (matriz)
	matriz := make([][]int, 3)
	for i := range matriz {
		matriz[i] = make([]int, 3) // cada fila se crea por separado
	}
	matriz[1][1] = 5
	fmt.Println("Matriz:", matriz)
}
