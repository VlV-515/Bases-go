package interfaces

import (
	"fmt"
	"math"
)

//! Las interfaces no se implementan, al menos no manualmente
//! Las interfaces se cumplen

type Forma interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func imprimirArea(f Forma) {
	fmt.Printf("El area es: %.2f", f.Area())
}

func App() {
	circulo := Circulo{Radio: 5}
	imprimirArea(circulo)
}
