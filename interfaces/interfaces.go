package interfaces

import (
	"bases-go/funciones"
	"fmt"
	"math"
)

//! Las interfaces no se implementan, se satisfacen de forma implicita
//! Si un tipo tiene los metodos que pide la interfaz, ya la cumple (no hay palabra 'implements')
//! Buenas practicas: interfaces chicas (1 o 2 metodos) y definidas donde se usan

// Forma representa cualquier figura con area calculable.
type Forma interface {
	Area() float64
}

// Perimetro representa cualquier figura con perimetro calculable.
type Perimetro interface {
	Perimetro() float64
}

// FormaCompleta compone dos interfaces: exige los metodos de ambas.
type FormaCompleta interface {
	Forma
	Perimetro
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

// String hace que Circulo cumpla fmt.Stringer: Println lo usa solo
func (c Circulo) String() string {
	return fmt.Sprintf("Circulo(radio %.1f)", c.Radio)
}

type Cuadrado struct {
	Lado float64
}

func (c Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

func (c Cuadrado) Perimetro() float64 {
	return 4 * c.Lado
}

// Verificacion en compilacion: si Circulo deja de cumplir Forma, no compila
var _ Forma = Circulo{}

func imprimirArea(f Forma) {
	fmt.Printf("El area es: %.2f\n", f.Area())
}

// describir usa un type switch para saber el tipo concreto detras de la interfaz
func describir(f Forma) {
	switch v := f.(type) {
	case Circulo:
		fmt.Println("Es un circulo de radio", v.Radio)
	case Cuadrado:
		fmt.Println("Es un cuadrado de lado", v.Lado)
	default:
		fmt.Printf("Tipo desconocido: %T\n", v)
	}
}

func App() {
	circulo := Circulo{Radio: 5}
	imprimirArea(circulo)
	//--------------------------------
	// Varios tipos tratados igual gracias a la interfaz
	formas := []Forma{
		circulo,
		Cuadrado{Lado: 4},
		// Rectangulo vive en otro paquete y no sabe que existe Forma, pero tiene Area() y ya la cumple
		funciones.Rectangulo{Ancho: 2, Alto: 3},
	}
	for _, f := range formas {
		imprimirArea(f)
		describir(f)
	}
	//--------------------------------
	// Type assertion: con ok no hace panic si el tipo no coincide
	if c, ok := formas[0].(Circulo); ok {
		fmt.Println("Assertion correcta, radio:", c.Radio)
	}
	if _, ok := formas[1].(Circulo); !ok {
		fmt.Println("formas[1] no es un Circulo")
	}
	//--------------------------------
	// Composicion: FormaCompleta exige Area() y Perimetro()
	var completa FormaCompleta = Cuadrado{Lado: 3}
	fmt.Printf("Area: %.2f - Perimetro: %.2f\n", completa.Area(), completa.Perimetro())
	//--------------------------------
	// Stringer: Println llama a String() automaticamente
	fmt.Println(circulo)
	//--------------------------------
	// any (interfaz vacia) acepta cualquier valor
	var cualquiera any = 42
	cualquiera = "ahora soy un string"
	fmt.Printf("any: %v (%T)\n", cualquiera, cualquiera)
	//--------------------------------
	// error tambien es una interfaz: type error interface { Error() string }
	var err error = fmt.Errorf("soy un error")
	fmt.Println(err.Error())
}
