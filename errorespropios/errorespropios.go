package errorespropios

import (
	"errors"
	"fmt"
)

// ! error es una interfaz: type error interface { Error() string }
// ! Cualquier tipo con un metodo Error() string es un error
// ! Tres herramientas:
// !   1) Error centinela (var ErrX = errors.New) -> se compara con errors.Is
// !   2) Tipo de error propio (struct)           -> se extrae con errors.As
// !   3) Envolver con %w                         -> agrega contexto sin perder el error original

// * 1) Errores centinela: valores fijos que se pueden comparar
var (
	ErrNoEncontrado = errors.New("no encontrado")
	ErrNoAutorizado = errors.New("no autorizado")
)

// * 2) Tipo de error propio: puede cargar datos extra
type ErrorValidacion struct {
	Campo   string
	Mensaje string
}

// Receptor por puntero: se usa y se devuelve como &ErrorValidacion{...}
func (e *ErrorValidacion) Error() string {
	return fmt.Sprintf("validacion: el campo %s %s", e.Campo, e.Mensaje)
}

// ErrorConsulta envuelve otro error. Unwrap deja que errors.Is y errors.As
// sigan buscando dentro de Err.
type ErrorConsulta struct {
	Tabla string
	Err   error
}

func (e *ErrorConsulta) Error() string {
	return fmt.Sprintf("consulta en %s: %v", e.Tabla, e.Err)
}

func (e *ErrorConsulta) Unwrap() error {
	return e.Err
}

// Capa "base de datos"
func buscarUsuario(id int) (string, error) {
	if id <= 0 {
		return "", &ErrorValidacion{Campo: "id", Mensaje: "debe ser mayor a 0"}
	}
	if id != 1 {
		return "", &ErrorConsulta{Tabla: "usuarios", Err: ErrNoEncontrado}
	}
	return "Ana", nil
}

// Capa "servicio": agrega contexto con %w, NO pierde el error original
func servicio(id int) (string, error) {
	nombre, err := buscarUsuario(id)
	if err != nil {
		return "", fmt.Errorf("servicio buscando usuario %d: %w", id, err)
	}
	return nombre, nil
}

// Capa "handler": traduce el error a un codigo HTTP (asi se usa en una API REST)
func codigoHTTP(err error) int {
	var errValidacion *ErrorValidacion
	switch {
	case err == nil:
		return 200
	case errors.As(err, &errValidacion):
		return 400
	case errors.Is(err, ErrNoAutorizado):
		return 401
	case errors.Is(err, ErrNoEncontrado):
		return 404
	default:
		return 500
	}
}

// ! TRAMPA CLASICA: devolver un puntero nil dentro de un error
// ! La interfaz guarda (tipo, valor); con el tipo presente ya NO es nil
func conTrampa() error {
	var e *ErrorValidacion // puntero nil
	return e
}

func sinTrampa() error {
	return nil // devuelve nil directo
}

func App() {
	//* Recorrer los tres casos por todas las capas
	for _, id := range []int{1, 0, 7} {
		nombre, err := servicio(id)
		if err != nil {
			fmt.Printf("id %d -> error: %v\n", id, err)
			fmt.Printf("         codigo HTTP: %d\n", codigoHTTP(err))
			continue
		}
		fmt.Printf("id %d -> %s (codigo HTTP: %d)\n", id, nombre, codigoHTTP(err))
	}
	//--------------------------------
	//* errors.Is: "¿este error, o alguno que envuelve, es ErrNoEncontrado?"
	_, err := servicio(7)
	fmt.Println("errors.Is NoEncontrado:", errors.Is(err, ErrNoEncontrado))
	fmt.Println("errors.Is NoAutorizado:", errors.Is(err, ErrNoAutorizado))
	// Con == NO funciona: el error esta envuelto
	fmt.Println("err == ErrNoEncontrado:", err == ErrNoEncontrado)
	//--------------------------------
	//* errors.As: "¿hay un *ErrorConsulta en la cadena? sacalo"
	var errConsulta *ErrorConsulta
	if errors.As(err, &errConsulta) {
		fmt.Println("errors.As -> tabla:", errConsulta.Tabla)
	}

	_, err = servicio(0)
	var errValidacion *ErrorValidacion
	if errors.As(err, &errValidacion) {
		fmt.Println("errors.As -> campo:", errValidacion.Campo, "| mensaje:", errValidacion.Mensaje)
	}
	//--------------------------------
	//* errors.Unwrap: abre UNA capa
	_, err = servicio(7)
	fmt.Println("Unwrap 1:", errors.Unwrap(err))
	fmt.Println("Unwrap 2:", errors.Unwrap(errors.Unwrap(err)))
	//--------------------------------
	//* errors.Join: junta varios errores en uno (util al validar varios campos)
	unidos := errors.Join(
		&ErrorValidacion{Campo: "nombre", Mensaje: "esta vacio"},
		&ErrorValidacion{Campo: "email", Mensaje: "no es valido"},
	)
	fmt.Println("Join:")
	fmt.Println(unidos)
	fmt.Println("Join contiene ErrorValidacion:", errors.As(unidos, &errValidacion))
	//--------------------------------
	//* La trampa del puntero nil
	fmt.Println("conTrampa() == nil:", conTrampa() == nil, "(deberia ser true, pero es false)")
	fmt.Println("sinTrampa() == nil:", sinTrampa() == nil)
}
