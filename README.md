# Bases-go

Proyecto de aprendizaje de Go. Cada tema vive en su propio paquete, con una función `App()` llena de ejemplos y notas para estudiar. `main.go` las ejecuta en orden.

## Requisitos

- Go `1.27.3` (la versión declarada en `go.mod`).
- Los ejemplos usan funciones recientes: `for i := range 3` (Go 1.22+), y los paquetes `slices` y `maps` (Go 1.23+).

## Cómo correrlo

```bash
go run .
```

Para detectar data races en los ejemplos de concurrencia:

```bash
go run -race .
```

Para estudiar un solo tema, comenta las demás llamadas en [main.go](main.go).

## Estructura

| Paquete | Qué cubre |
|---|---|
| [tiposdedatos](tiposdedatos/tiposdedatos.go) | Tipos básicos, arrays, slices, mapas y structs |
| [ciclos](ciclos/ciclos.go) | `if`, `for`, `range`, `switch` y `defer` |
| [funciones](funciones/funciones.go) | Retornos múltiples, errores, variádicas, closures, métodos y punteros |
| [interfaces](interfaces/interfaces.go) | Interfaces, type switch, type assertion, `Stringer`, `any` y `error` |
| [concurrencia](concurrencia/concurrencia.go) | Goroutines, canales, `WaitGroup` y `select` |
| [mutex](mutex/mutex.go) | `Mutex`, `RWMutex`, `atomic` y data races |
| [jsondatos](jsondatos/jsondatos.go) | `encoding/json`: tags, `Marshal`, `Unmarshal`, `Encoder` y `Decoder` |
| [mapas](mapas/mapas.go) | Mapas a fondo y un almacén seguro para goroutines |
| [slicesavanzados](slicesavanzados/slicesavanzados.go) | `len`/`cap`, memoria compartida, `copy` y el paquete `slices` |
| [errorespropios](errorespropios/errorespropios.go) | Errores centinela, tipos propios, `errors.Is`, `errors.As` y `errors.Join` |

## Cheatsheet

Un resumen de todo lo anterior en una sola referencia: [cheatsheet-go.pdf](cheatsheet-go.pdf).

## Orden sugerido de estudio

1. `tiposdedatos`, `ciclos`, `funciones`
2. `slicesavanzados`, `mapas`
3. `interfaces`, `errorespropios`
4. `jsondatos`
5. `concurrencia`, `mutex`

## Convenciones del proyecto

- Un tema por paquete, con nombre en minúscula.
- Cada paquete expone `App()` y `main.go` las llama.
- Comentarios `//!` para notas importantes y `//*` para separar secciones.
- Se revisa con `gofmt -l .` y `go vet ./...` antes de cada commit.

## Siguiente paso

Construir una API REST con `net/http`: una API CRUD en memoria, con JSON, errores con código HTTP y tests con `httptest`.
