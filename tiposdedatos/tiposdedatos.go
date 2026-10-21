package tiposdedatos

import (
	"fmt"
	"strings"
)

/*
*TIPOS DE DATOS
bool | true / false | Flag o condicionales
string | Cadena de caracteres | Para representar texto
int, int8, int16, int32, int64 | numero entero | Los bits es para controlar el tamaño del entero (int sin número depende de la arquitectura)
uint, uint8, uint16, uint32, uint64 | numero entero sin signo | Cuando no queremos negativos
float32, float64 | Para representar numeros reales | Numeros con punto => 32 o 64 es la cantidad de bits (precision)
byte == uint8 | Trabajar con datos binarios
rune == int32 | Para representar un caracter que ocupa mas de un byte
complex64 , complex128 | Para representar una parte real y una parte imaginaria | N + iN
*/

//var entero int = 10 //Forma inusual de crear variables

func App() {
	fmt.Println("Hola mundo") //PrintLine | Imprime una linea en consola

	//* Numeros
	entero := 10
	decimal := 3.14
	suma := entero + int(decimal) // int() trunca: 3.14 pasa a 3
	fmt.Println("Suma:", suma)

	//* Textos
	mensaje := "Mensaje "
	concatenado := mensaje + "concatenado"
	mayuscula := strings.ToUpper(concatenado)
	fmt.Println("Concatenado y mayuscula:", mayuscula)

	//* Booleanos
	esVerdadero := true
	fmt.Println("Boolean:", esVerdadero)

	//* Arrays
	// [X]int{} es un array: el tamaño es parte del tipo y no cambia.
	// [...]int{1, 2, 3} deja que el compilador cuente los elementos.
	arrayFijo := [3]int{1, 2, 3}
	//arrayFijo = append(arrayFijo, 1) //Esto da error
	fmt.Println("Array fijo:", arrayFijo)

	//* Slices
	// []int{} (sin tamaño) NO es un array, es un slice: una ventana sobre un array que crece con append.
	// En Go casi siempre se usan slices.
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("Slice:", slice)

	longitud := len(slice)  // Obtener la longitud (cuantos elementos tiene)
	capacidad := cap(slice) // Obtener la capacidad (cuantos elementos caben antes de que Go reserve memoria nueva)
	fmt.Println("Longitud:", longitud)
	fmt.Println("Capacidad:", capacidad)

	//* Mapas (diccionario)
	diccionario := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("Mapa:", diccionario)

	//* Structs
	type Persona struct {
		Nombre string
		Edad   int
	}

	persona := Persona{Nombre: "Hola mundo", Edad: 30}
	fmt.Println("Struct:", persona)
}
