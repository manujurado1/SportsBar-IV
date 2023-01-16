package pkg

import "fmt"

type Grupo struct {
	Nombre               string
	JugadoresDisponibles []string
	JugadoresNiveles     map[string]int
}

// Función que crea un jugador y lo inserta en la lista de disponibles en caso de que así se indique
func (this *Grupo) crearJugador(NombreJugador string, Nivel int, Disponibilidad bool) (bool, error) {
	var success bool = false
	_, existe := this.JugadoresNiveles[NombreJugador]
	if existe == false {
		if Nivel >= 0 && Nivel <= 10 {

			this.JugadoresNiveles[NombreJugador] = Nivel
			success = true
			if Disponibilidad {
				this.JugadoresDisponibles = append(this.JugadoresDisponibles, NombreJugador)
			}
		} else {
			return success, fmt.Errorf("El nivel del jugador debe estar entre 0 y 10")
		}
	} else {
		return success, fmt.Errorf("Ya existe un jugador con ese nombre")
	}

	return success, nil

}
