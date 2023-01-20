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

// Función que añade un jugador en el map de JugadoresNiveles si ese nombre no tiene una entrada existente y el nivel es correcto
// y lo inserta en la lista de disponibles en caso de que así se indique
func (this *Grupo) crearJugador(NombreJugador string, Nivel int, Disponibilidad bool) (bool, error) {
	var success bool = false
	_, existe := this.JugadoresNiveles[NombreJugador]
	if existe == false {
		if Nivel >= 0 && Nivel <= 100 {

			this.JugadoresNiveles[NombreJugador] = Nivel
			this.JugadoresDisponibilidad[NombreJugador] = Disponibilidad
			success = true

		} else {
			return success, fmt.Errorf("El nivel del jugador debe estar entre 0 y 100")
		}
	} else {
		return success, fmt.Errorf("Ya existe un jugador con ese nombre")
	}

	return success, nil

}

func (this *Grupo) cambiarDisponibilidadJugador(NombreJugador string, Disponibilidad bool) (bool, error) {
	var success bool = false
	_, existe := this.JugadoresNiveles[NombreJugador]
	if existe == true {
		this.JugadoresDisponibilidad[NombreJugador] = Disponibilidad
		success = true
	} else {
		return success, fmt.Errorf("No existe un jugador con ese nombre")
	}

	return success, nil

}
