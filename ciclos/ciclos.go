package ciclos

import "fmt"

func App() {
	//* Defer (Para ejecutar al terminar la funcion)
	// Varios defer se ejecutan en orden inverso (el último en registrarse corre primero)
	defer fmt.Println("FIN DE CICLOS 2")
	defer fmt.Println("FIN DE CICLOS 1")

	//* Condicional
	edad := 20
	if edad < 18 {
		fmt.Printf("Eres menor de edad: %d\n", edad)
	} else {
		fmt.Printf("Eres mayor de edad: %d\n", edad)
	}

	//* Bucles for (clasico)
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteración clásica: %d\n", i)
	}

	//* Bucle tipo while
	n := 0
	for n < 3 {
		fmt.Printf("Iteración while: %d\n", n)
		n++
	}

	//* Bucle infinito
	for {
		if n == 5 {
			fmt.Println("5 lo saltamos")
			n++
			continue
		}

		fmt.Printf("Iteración infinito: %d\n", n)
		n++

		if n > 7 {
			fmt.Println("Chaito")
			break
		}
	}

	//* Range (para recorrer slices, arrays, mapas y strings)
	varRange := []string{"uno", "dos", "tres"}

	for index, value := range varRange {
		fmt.Printf("Rango, índice: %d, valor: %s\n", index, value)
	}

	// Con _ se ignora el índice
	for _, value := range varRange {
		fmt.Printf("Rango, solo valor: %s\n", value)
	}

	// Desde Go 1.22 se puede recorrer un número directamente
	for i := range 3 {
		fmt.Printf("Rango numérico: %d\n", i)
	}

	//* Switch (no necesita break; un case puede tener varios valores)
	valorSwitch := 2
	switch valorSwitch {
	case 1:
		fmt.Println("Caso 1")
	case 2, 3:
		fmt.Println("Caso 2 o 3")
	default:
		fmt.Println("Caso default")
	}

}
