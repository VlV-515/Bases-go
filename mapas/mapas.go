package mapas

import (
	"fmt"
	"maps"
	"slices"
	"sync"
)

// ! Un mapa es una tabla clave -> valor
// ! Las claves deben ser comparables (string, int, bool, structs de comparables...)
// ! NO pueden ser claves: slices, mapas ni funciones
// ! Un mapa es una referencia: pasarlo a una funcion NO lo copia

type Persona struct {
	Nombre string
	Edad   int
}

type Punto struct {
	X, Y int
}

// Almacen es un mapa protegido con RWMutex.
// Los mapas NO son seguros entre goroutines: escribir a la vez da
// 'fatal error: concurrent map writes'. Este patron se usa mucho en APIs.
type Almacen struct {
	mu    sync.RWMutex
	datos map[int]Persona
}

func NuevoAlmacen() *Almacen {
	return &Almacen{datos: make(map[int]Persona)}
}

func (a *Almacen) Guardar(id int, p Persona) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.datos[id] = p
}

func (a *Almacen) Obtener(id int) (Persona, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	p, ok := a.datos[id]
	return p, ok
}

func (a *Almacen) Eliminar(id int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.datos, id)
}

func escribirEnMapaNil() {
	// recover atrapa el panic para poder mostrarlo; no se usa como flujo normal
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recuperado:", r)
		}
	}()
	var nulo map[string]int
	nulo["a"] = 1
}

func agregarClave(m map[string]int) {
	m["nueva"] = 99 // modifica el mapa original, no una copia
}

func App() {
	//* Crear mapas
	literal := map[string]int{"a": 1, "b": 2}
	conMake := make(map[string]int) // mapa vacio listo para usar
	conMake["x"] = 10
	fmt.Println("Literal:", literal, "| make:", conMake)
	//--------------------------------
	//* Mapa nil: se puede LEER, pero escribir hace panic
	var nulo map[string]int
	fmt.Println("Mapa nil -> leer:", nulo["a"], "| len:", len(nulo), "| es nil:", nulo == nil)
	escribirEnMapaNil()
	//--------------------------------
	//* Leer una clave: la forma "coma ok"
	edades := map[string]int{"Ana": 30, "Beto": 0}

	// Si la clave no existe, devuelve el valor cero (0), igual que si valiera 0
	fmt.Println("Ana:", edades["Ana"], "| Beto:", edades["Beto"], "| Zoe (no existe):", edades["Zoe"])

	// ok distingue "existe con valor 0" de "no existe"
	if edad, ok := edades["Beto"]; ok {
		fmt.Println("Beto existe, edad:", edad)
	}
	if _, ok := edades["Zoe"]; !ok {
		fmt.Println("Zoe no existe")
	}
	//--------------------------------
	//* Agregar, actualizar y borrar
	edades["Carla"] = 25 // agregar
	edades["Ana"] = 31   // actualizar
	delete(edades, "Beto")
	delete(edades, "NoExiste") // borrar una clave inexistente no hace nada ni da error
	fmt.Println("Despues de cambios:", edades, "| len:", len(edades))
	//--------------------------------
	//* Recorrer: el orden es ALEATORIO, no dependas de el
	for nombre, edad := range edades {
		_ = nombre
		_ = edad
	}
	// Para un orden estable: ordena las claves primero
	claves := slices.Sorted(maps.Keys(edades))
	for _, nombre := range claves {
		fmt.Printf("%s tiene %d\n", nombre, edades[nombre])
	}
	//--------------------------------
	//* Los mapas son referencias
	original := map[string]int{"a": 1}
	agregarClave(original)
	otraVariable := original // NO copia: apuntan al mismo mapa
	otraVariable["b"] = 2
	fmt.Println("Original tras modificar:", original)

	// Copia real (superficial)
	copia := maps.Clone(original)
	copia["solo-en-copia"] = 0
	fmt.Println("Original:", len(original), "claves | copia:", len(copia), "claves")

	// Vaciar un mapa
	clear(copia)
	fmt.Println("Copia tras clear:", len(copia))
	//--------------------------------
	//* Struct como clave (mientras todos sus campos sean comparables)
	nombresPuntos := map[Punto]string{
		{X: 0, Y: 0}: "origen",
		{X: 1, Y: 2}: "punto A",
	}
	fmt.Println("Punto (1,2):", nombresPuntos[Punto{X: 1, Y: 2}])
	//--------------------------------
	//* Mapa de structs: NO se puede modificar un campo directamente
	personas := map[string]Persona{"ana": {Nombre: "Ana", Edad: 30}}
	// personas["ana"].Edad = 31 // ERROR de compilacion: cannot assign to struct field in map
	p := personas["ana"] // se saca una copia...
	p.Edad = 31
	personas["ana"] = p // ...y se vuelve a guardar
	fmt.Println("Reasignando el struct:", personas["ana"])

	// Alternativa: guardar punteros, asi el cambio es directo
	punteros := map[string]*Persona{"beto": {Nombre: "Beto", Edad: 20}}
	punteros["beto"].Edad = 21
	fmt.Println("Con punteros:", *punteros["beto"])
	//--------------------------------
	//* Patrones muy comunes
	// Conjunto (set): el valor struct{} no ocupa memoria
	conjunto := map[string]struct{}{}
	conjunto["go"] = struct{}{}
	conjunto["rust"] = struct{}{}
	if _, existe := conjunto["go"]; existe {
		fmt.Println("go esta en el conjunto, tamano:", len(conjunto))
	}

	// Contar repeticiones
	cuenta := map[rune]int{}
	for _, letra := range "hola mundo" {
		cuenta[letra]++ // si no existe, parte de 0
	}
	for _, letra := range slices.Sorted(maps.Keys(cuenta)) {
		if cuenta[letra] > 1 {
			fmt.Printf("La letra %q se repite %d veces\n", letra, cuenta[letra])
		}
	}

	// Agrupar: append sobre un slice nil funciona
	grupos := map[string][]string{}
	grupos["frutas"] = append(grupos["frutas"], "manzana")
	grupos["frutas"] = append(grupos["frutas"], "pera")
	grupos["verduras"] = append(grupos["verduras"], "zanahoria")
	fmt.Println("Grupos:", grupos)
	//--------------------------------
	//* Almacen seguro para goroutines (base de una API en memoria)
	almacen := NuevoAlmacen()
	almacen.Guardar(1, Persona{Nombre: "Ana", Edad: 30})
	if persona, ok := almacen.Obtener(1); ok {
		fmt.Println("Almacen id 1:", persona)
	}
	almacen.Eliminar(1)
	if _, ok := almacen.Obtener(1); !ok {
		fmt.Println("Almacen id 1 eliminado")
	}
}
