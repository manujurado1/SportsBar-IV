package pkg

import (
	"fmt"
	"sort"
	"strings"
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

// Función que devuelve un diccionario clave-valor de los jugadores disponibles ordenados por nivel de manera descendente
// y su nivel asociado
func (this *Grupo) ordenarJugadoresDisponilesPorNivel(JugadoresNiveles map[string]int, JugadoresDisponibles []string) ([]string, error) {
	if (JugadoresNiveles == nil || len(JugadoresNiveles) == 0) || (JugadoresDisponibles == nil || len(JugadoresDisponibles) == 0) {
		return nil, fmt.Errorf("Los parámetros no son correctos, almenos un parámetro esta vacío")
	}

	JugadoresOrdenados := make([]string, 0, len(JugadoresNiveles))

	for Jugador := range JugadoresNiveles {
		JugadoresOrdenados = append(JugadoresOrdenados, Jugador)
	}

	sort.SliceStable(JugadoresOrdenados, func(i, j int) bool {
		return JugadoresNiveles[JugadoresOrdenados[i]] > JugadoresNiveles[JugadoresOrdenados[j]]
	})

	for _, Jugador := range JugadoresOrdenados {
		if existeEnArreglo(JugadoresDisponibles, Jugador) == false {
			JugadoresOrdenados = eliminarElementoArreglo(JugadoresOrdenados, Jugador)
		}
	}
	return JugadoresOrdenados, nil
}

// Función que reparte los 2 jugadores en 2 equipos con nivel similar siguiendo las siguientes reglas partiendo de
// una lista con los jugadores ordenados de mayor a menor nivel:
// 1. Colocar a los 2 mejores jugadores uno en cada equipo
// Para el resto de jugadores:
// 2. Ir añadiendo el siguiente jugador al equipo que sume menos nivel entre todos sus componentes si es que este no está lleno
// 2.1 Si está lleno, añadirlo al otro equipo
func (this *Grupo) repartirJugadoresEnEquiposPorNivel(JugadoresNiveles map[string]int, JugadoresDisponibles []string) ([]string, []string, error) {
	var Equipo1 []string
	var Equipo2 []string
	var SumaNivelEquipo1 int
	var SumaNivelEquipo2 int

	ListaJugadoresOrdenadosPorNivel, err := this.ordenarJugadoresDisponilesPorNivel(JugadoresNiveles, JugadoresDisponibles)

	if err != nil {
		return nil, nil, fmt.Errorf(err.Error())
	}

	for _, Jugador := range ListaJugadoresOrdenadosPorNivel {

		if len(Equipo1) == 0 && len(Equipo2) == 0 {
			Equipo1 = append(Equipo1, Jugador)
			SumaNivelEquipo1 += JugadoresNiveles[Jugador]
		} else if len(Equipo2) == 0 {
			Equipo2 = append(Equipo2, Jugador)
			SumaNivelEquipo2 += JugadoresNiveles[Jugador]
		} else if (SumaNivelEquipo1 < SumaNivelEquipo2) && (len(Equipo1) < (len(ListaJugadoresOrdenadosPorNivel) / 2)) {
			Equipo1 = append(Equipo1, Jugador)
			SumaNivelEquipo1 += JugadoresNiveles[Jugador]
		} else if len(Equipo2) < (len(ListaJugadoresOrdenadosPorNivel) / 2) {
			Equipo2 = append(Equipo2, Jugador)
			SumaNivelEquipo2 += JugadoresNiveles[Jugador]
		} else {
			Equipo1 = append(Equipo1, Jugador)
			SumaNivelEquipo1 += JugadoresNiveles[Jugador]
		}
	}

	return Equipo1, Equipo2, nil
}

func (this *Grupo) crearEquiposIgualadosParaPartido(JugadoresNiveles map[string]int, JugadoresDisponibles []string) (string, error) {
	ListaValida, err := this.validarListaJugadoresDisponiblesParaPartido(JugadoresDisponibles)
	var Equipos string
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	if ListaValida == true {
		Equipo1, Equipo2, err := this.repartirJugadoresEnEquiposPorNivel(JugadoresNiveles, JugadoresDisponibles)
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
		Equipos = fmt.Sprintf("Equipo 1 = %v Equipo 2 = %v", strings.Join(Equipo1, ","), strings.Join(Equipo2, ","))

	} else {
		return "", fmt.Errorf("La lista de jugadores disponibles no es válida ya que no es divisible entre 2 o es menor de 10")
	}

	return Equipos, nil

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
