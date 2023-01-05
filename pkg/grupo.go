package pkg

import (
	"fmt"
)

type Grupo struct {
	Nombre               string
	JugadoresDisponibles []string
	JugadoresNiveles     map[string]int
}

// Función que crea un jugador y lo inserta en la lista de disponibles en caso de que así se indique
func (this *Grupo) crearJugador(NombreJugador string, Nivel int, Disponibilidad bool) error {
	_, existe := this.JugadoresNiveles[NombreJugador]
	if existe == false {
		if Nivel >= 0 && Nivel <= 10 {

			this.JugadoresNiveles[NombreJugador] = Nivel
			if Disponibilidad {
				this.JugadoresDisponibles = append(this.JugadoresDisponibles, NombreJugador)
			}
		} else {
			return fmt.Errorf("El nivel del jugador debe estar entre 0 y 10")
		}
	} else {
		return fmt.Errorf("Ya existe un jugador con ese nombre")
	}

	return nil

}

func (this *Grupo) cambiarDisponibilidadJugador(NombreJugador string) error {
	_, existe := this.JugadoresNiveles[NombreJugador]
	if existe == true {
		if existeEnArreglo(this.JugadoresDisponibles, NombreJugador) == true {
			this.JugadoresDisponibles = eliminarElementoArreglo(this.JugadoresDisponibles, NombreJugador)
		} else {
			this.JugadoresDisponibles = append(this.JugadoresDisponibles, NombreJugador)
		}
	} else {
		return fmt.Errorf("No existe un jugador con ese nombre")
	}

	return nil

}

// Función que comprueba que la lista de jugadores disponibles cumple las condiciones para jugar un partido
func (this *Grupo) validarListaJugadoresDisponiblesParaPartido(JugadoresDisponibles []string) (bool, error) {
	if JugadoresDisponibles != nil && len(JugadoresDisponibles) > 0 {
		if len(this.JugadoresDisponibles) >= 10 && len(this.JugadoresDisponibles)%2 == 0 {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, fmt.Errorf("La lista de jugadores disponibles está vacía")
	}

}

func existeEnArreglo(Arreglo []string, Busqueda string) bool {
	for _, Nombre := range Arreglo {
		if Nombre == Busqueda {
			return true
		}
	}
	return false
}

func eliminarElementoArreglo(Arreglo []string, ElementoAEliminar string) []string {
	var ArregloAux []string
	for _, Elemento := range Arreglo {
		if Elemento != ElementoAEliminar {
			ArregloAux = append(ArregloAux, Elemento)
		}
	}

	return ArregloAux
}
