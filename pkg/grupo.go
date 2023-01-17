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

// Función que añade un jugador en el map de JugadoresNiveles si ese nombre no tiene una entrada existente y el nivel es correcto
// y lo inserta en la lista de disponibles en caso de que así se indique
func (this *Grupo) crearJugador(NombreJugador string, Nivel int, Disponibilidad bool) (bool, error) {
	var success bool = false
	_, existe := this.JugadoresNiveles[NombreJugador]
	if existe == false {
		if Nivel >= 0 && Nivel <= 100 {

			this.JugadoresNiveles[NombreJugador] = Nivel
			success = true
			if Disponibilidad {
				this.JugadoresDisponibles = append(this.JugadoresDisponibles, NombreJugador)
			}
		} else {
			return success, fmt.Errorf("El nivel del jugador debe estar entre 0 y 100")
		}
	} else {
		return success, fmt.Errorf("Ya existe un jugador con ese nombre")
	}

	return success, nil

}

// Función que indica si la lista de strings pasada como primer parámetro contiene el string pasado como segundo parámetro
func existeEnArreglo(Arreglo []string, Busqueda string) (bool, error) {
	if len(Arreglo) > 0 {
		for _, Nombre := range Arreglo {
			if Nombre == Busqueda {
				return true, nil
			}
		}
		return false, nil
	} else {
		return false, fmt.Errorf("El arreglo está vacío")
	}

}
