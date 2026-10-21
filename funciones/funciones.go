package funciones

import (
	"errors"
	"fmt"
	"math"
)

// * Funcion clasica
func suma(a, b int) int {
	return a + b
}

// * Funcion que devuelve mas de un valor
// math.Mod calcula el resto con decimales y no hace panic con divisor 0 (devuelve NaN)
func dividir(a, b float64) (float64, float64) {
	cociente := a / b
	resto := math.Mod(a, b)
	return cociente, resto
}

// * Retornos con nombre
// Los valores nombrados ya existen dentro de la funcion; un return sin argumentos los devuelve
func dividirEnteros(a, b int) (cociente, resto int) {
	cociente = a / b
	resto = a % b
	return
}

// * Error centinela (se puede comparar con errors.Is)
var ErrDivisionPorCero = errors.New("no puedes dividir por 0")

// * Funcion que devuelve un error (por regla, el error va al final)
func dividirError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("dividir %v entre %v: %w", a, b, ErrDivisionPorCero)
	}
	cociente := a / b
	return cociente, nil
}

// * Funcion con numero variable de argumentos
func imprimirNombres(nombres ...string) {
	for _, name := range nombres {
		fmt.Printf("Hola %s\n", name)
	}
}

// * Funciones como valores
// Una funcion puede recibirse como parametro y definirse de forma anonima
func aplicar(n int, f func(int) int) int {
	return f(n)
}

// * Closure
// La funcion contador retorna una funcion que podemos ejecutar en otro lado cuando queramos
func contador() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}

func closure() {
	//Obtenemos la funcion, la guardamos en una variable para ejecutarla cuando nos convenga
	cont := contador()
	fmt.Printf("Contador: %d\n", cont())
	fmt.Printf("Contador: %d\n", cont())
	fmt.Printf("Contador: %d\n", cont())
}

// * Añadir funcion a una estructura
// Go no tiene clases: un metodo es una funcion con un receptor (r Rectangulo)
type Rectangulo struct {
	Ancho, Alto float64
}

// Area usa receptor por valor: recibe una copia, no modifica el original
func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

// Escalar usa receptor por puntero: modifica el Rectangulo original
func (r *Rectangulo) Escalar(factor float64) {
	r.Ancho *= factor
	r.Alto *= factor
}

// * Punteros
// *int en un tipo = "puntero a int"
// *p en una expresion = el valor al que apunta p (desreferenciar)
// &x = la direccion de memoria de x
// Desreferenciar un puntero nil causa panic
func incrementarSinPuntero(valor int) {
	valor++
}
func incrementarConPuntero(valor *int) {
	*valor++
}

// * Punteros con new
func punteroReferencia() {
	p := new(int) //Inicializa un puntero con valor 0
	fmt.Printf("Direccion de memoria de puntero: %p\n", p)
	fmt.Printf("Valor de puntero: %d\n", *p)
	*p = 3
	fmt.Printf("Modificando el valor del puntero: %d\n", *p)
}

func puntero() {
	//El valor no se incrementa porque siempre se crea una copia de argumentos,
	//no una referencia (osea, no muta el original), a menos que se mande el puntero (&),
	//ahi go entiende que se debe sobreescribir el valor inicial
	val := 1
	fmt.Printf("Valor antes de incrementar: %d\n", val)
	incrementarSinPuntero(val)
	fmt.Printf("Valor despues de incrementar sin puntero: %d\n", val)
	incrementarConPuntero(&val)
	fmt.Printf("Valor despues de incrementar con puntero: %d\n", val)
}

func App() {
	fmt.Printf("Suma: %d\n", suma(1, 2))
	//---------------------------------------------
	cociente, resto := dividir(10.5, 2)
	fmt.Printf("Cociente: %.2f - Resto: %.2f\n", cociente, resto)
	//---------------------------------------------
	cocienteEntero, _ := dividirEnteros(10, 3) // _ ignora el resto
	fmt.Printf("Cociente entero: %d\n", cocienteEntero)
	//---------------------------------------------
	resultado, err := dividirError(10, 0)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Es division por cero:", errors.Is(err, ErrDivisionPorCero))
	} else {
		fmt.Printf("El resultado es: %.2f\n", resultado)
	}
	//---------------------------------------------
	imprimirNombres("Juan", "Pedro", "Tito")
	//---------------------------------------------
	fmt.Printf("Aplicar doble: %d\n", aplicar(3, func(x int) int { return x * 2 }))
	//---------------------------------------------
	closure()
	// ---------------------------------------------
	r := Rectangulo{Ancho: 10, Alto: 20}
	fmt.Printf("El area es: %.2f\n", r.Area())
	r.Escalar(2)
	fmt.Printf("El area escalada x2 es: %.2f\n", r.Area())
	// ---------------------------------------------
	puntero()
	punteroReferencia()
	// ---------------------------------------------
}
