package main

import (
	"bases-go/ciclos"
	"bases-go/concurrencia"
	"bases-go/errorespropios"
	"bases-go/funciones"
	"bases-go/interfaces"
	"bases-go/jsondatos"
	"bases-go/mapas"
	"bases-go/mutex"
	"bases-go/slicesavanzados"
	"bases-go/tiposdedatos"
)

func main() {
	tiposdedatos.App()
	ciclos.App()
	funciones.App()
	interfaces.App()
	concurrencia.App()
	mutex.App()
	jsondatos.App()
	mapas.App()
	slicesavanzados.App()
	errorespropios.App()
}
