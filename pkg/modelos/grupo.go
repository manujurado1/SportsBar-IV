package modelos

import "fmt"

type Grupo struct {
	Nombre                  string
	JugadoresDisponibilidad map[string]bool
	JugadoresNiveles        map[string]int
}

// Constructor que crea un nuevo Grupo con el nombre pasado por parámetro, inicializando  el map de JugadoresNiveles y el de Jugadores Disponibilidad vacíos.
func CrearGrupo(Nombre string) (*Grupo, error) {

	JugadoresDisponibilidad := make(map[string]bool)
	JugadoresNiveles := make(map[string]int)

	if Nombre == "" {
		return nil, fmt.Errorf("El nombre del grupo no puede ser un string vacío")
	}

	grupo := Grupo{Nombre, JugadoresDisponibilidad, JugadoresNiveles}
	return &grupo, nil
}
