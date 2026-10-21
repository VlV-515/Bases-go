package jsondatos

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ! JSON en Go: el paquete encoding/json convierte entre structs y texto JSON
// ! Marshal = struct -> JSON (codificar) | Unmarshal = JSON -> struct (decodificar)
// ! Solo se convierten los campos EXPORTADOS (con mayuscula inicial)
// ! Las tags (`json:"..."`) controlan como se llama cada campo en el JSON

type Direccion struct {
	Calle  string `json:"calle"`
	Ciudad string `json:"ciudad"`
}

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`

	// omitempty: si el campo esta vacio (""), no aparece en el JSON
	Email string `json:"email,omitempty"`

	// Puntero: distingue "no vino" (nil -> null) de "vino con 0"
	Edad *int `json:"edad"`

	// "-" : el campo NUNCA se convierte (util para contrasenas)
	Clave string `json:"-"`

	Etiquetas []string  `json:"etiquetas"`
	Direccion Direccion `json:"direccion"`

	// Minuscula = no exportado = json lo ignora
	privado string
}

func edadTexto(u Usuario) string {
	if u.Edad == nil {
		return "sin edad"
	}
	return fmt.Sprintf("%d", *u.Edad)
}

func App() {
	//* Marshal: struct -> JSON
	edad := 30
	ana := Usuario{
		ID:        1,
		Nombre:    "Ana",
		Edad:      &edad,
		Clave:     "secreta",
		Etiquetas: []string{"admin", "dev"},
		Direccion: Direccion{Calle: "Calle 1", Ciudad: "CDMX"},
		privado:   "no sale en el JSON",
	}

	datos, err := json.Marshal(ana)
	if err != nil {
		fmt.Println("Error al codificar:", err)
		return
	}
	// Marshal devuelve []byte, por eso se convierte a string para imprimir
	// Ojo: no aparecen Email (omitempty), Clave ("-") ni privado (minuscula)
	fmt.Println("Marshal:", string(datos))

	// MarshalIndent -> JSON con sangria, mas facil de leer
	bonito, _ := json.MarshalIndent(ana, "", "  ")
	fmt.Println("MarshalIndent:")
	fmt.Println(string(bonito))

	// Campos nil: puntero nil -> null, slice nil -> null
	beto := Usuario{ID: 2, Nombre: "Beto", Email: "beto@correo.com"}
	datosBeto, _ := json.Marshal(beto)
	fmt.Println("Con nil:", string(datosBeto))
	//--------------------------------
	//* Unmarshal: JSON -> struct
	// Se pasa un PUNTERO (&) para que Unmarshal pueda llenar el struct
	texto := `{"id":3,"nombre":"Carla","edad":25,"etiquetas":["qa"],"direccion":{"calle":"Av 5","ciudad":"GDL"},"campo_extra":true}`
	var carla Usuario
	if err := json.Unmarshal([]byte(texto), &carla); err != nil {
		fmt.Println("Error al decodificar:", err)
		return
	}
	// Los campos que no existen en el struct (campo_extra) se ignoran en silencio
	fmt.Println("Unmarshal:", carla.ID, carla.Nombre, edadTexto(carla), carla.Etiquetas, carla.Direccion.Ciudad)

	// Si el JSON no trae un campo, queda en su valor cero
	var sinEdad Usuario
	json.Unmarshal([]byte(`{"id":4,"nombre":"Dani"}`), &sinEdad)
	fmt.Println("Sin edad:", edadTexto(sinEdad), "| Etiquetas nil:", sinEdad.Etiquetas == nil)
	//--------------------------------
	//* Errores al decodificar
	var malo Usuario
	err = json.Unmarshal([]byte(`{"id": }`), &malo)
	fmt.Println("JSON invalido:", err)

	// Tipo equivocado: "abc" no cabe en un int
	err = json.Unmarshal([]byte(`{"id":"abc"}`), &malo)
	fmt.Println("Tipo equivocado:", err)

	// errors.As permite sacar el detalle del error (tema de errores propios)
	var errTipo *json.UnmarshalTypeError
	if errors.As(err, &errTipo) {
		fmt.Printf("Campo: %s, llego: %s, esperaba: %s\n", errTipo.Field, errTipo.Value, errTipo.Type)
	}
	//--------------------------------
	//* Lista de objetos -> slice de structs
	lista := `[{"id":1,"nombre":"A"},{"id":2,"nombre":"B"}]`
	var usuarios []Usuario
	if err := json.Unmarshal([]byte(lista), &usuarios); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Lista:", len(usuarios), "usuarios, el segundo es", usuarios[1].Nombre)
	//--------------------------------
	//* JSON sin estructura conocida: map[string]any
	var generico map[string]any
	json.Unmarshal([]byte(`{"nombre":"Ana","edad":30,"activo":true,"notas":[1,2]}`), &generico)
	// Los numeros JSON siempre llegan como float64, aunque sean enteros
	fmt.Printf("Tipos: nombre=%T edad=%T activo=%T notas=%T\n",
		generico["nombre"], generico["edad"], generico["activo"], generico["notas"])
	if edadGen, ok := generico["edad"].(float64); ok {
		fmt.Println("Edad como float64:", edadGen)
	}
	//--------------------------------
	//* Encoder y Decoder: trabajan con flujos (io.Writer / io.Reader)
	// Asi se usan en una API: json.NewDecoder(r.Body) y json.NewEncoder(w)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(carla) // Encode ya agrega un salto de linea al final

	dec := json.NewDecoder(strings.NewReader(`{"id":9,"nombre":"Eli","campo_raro":1}`))
	dec.DisallowUnknownFields() // ahora los campos desconocidos SI dan error
	var estricto Usuario
	err = dec.Decode(&estricto)
	fmt.Println("Decoder estricto:", err)
}
