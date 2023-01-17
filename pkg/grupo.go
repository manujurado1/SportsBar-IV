package pkg

import "fmt"

type Grupo struct {
	Nombre               string
	JugadoresDisponibles []string
	JugadoresNiveles     map[string]int
}

// Constructor que crea un nuevo Grupo con el nombre pasado por parámetro, inicializando la lista de JugadoresDiponibles y el map de JugadoresNiveles vacío.
func CrearGrupo(Nombre string) (*Grupo, error) {

	JugadoresDisponibles := []string{}
	JugadoresNiveles := make(map[string]int)

	if Nombre == "" {
		return nil, fmt.Errorf("El nombre del grupo no puede ser un string vacío")
	}

	grupo := Grupo{Nombre, JugadoresDisponibles, JugadoresNiveles}
	return &grupo, nil
}
