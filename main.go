package main

import (
	"bases-go/ciclos"
	"bases-go/concurrencia"
	"bases-go/funciones"
	"bases-go/interfaces"
	"bases-go/mutex"
	"bases-go/tiposdedatos"
)

func main() {
	tiposdedatos.App()
	ciclos.App()
	funciones.App()
	interfaces.App()
	concurrencia.App()
	mutex.App()
}
