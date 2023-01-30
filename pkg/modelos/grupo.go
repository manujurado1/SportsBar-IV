package modelos

import (
	"fmt"
	"sort"
)

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

func (this *Grupo) conseguirListaJugadoresDisponibles(JugadoresDisponibilidad map[string]bool) ([]string, error) {

	var ListaJugadoresDisponibles []string

	if len(JugadoresDisponibilidad) > 0 {

		for key, value := range JugadoresDisponibilidad {
			if value == true {
				ListaJugadoresDisponibles = append(ListaJugadoresDisponibles, key)
			}
		}
		return ListaJugadoresDisponibles, nil

	} else {
		return nil, fmt.Errorf("El map de JugadoresDisponibilidad no puede estar vacío")
	}

}

func (this *Grupo) validarListaJugadoresDisponiblesParaPartido(ListaJugadoresDisponibles []string) (bool, error) {
	success := false

	if len(ListaJugadoresDisponibles) > 0 {
		if len(ListaJugadoresDisponibles) >= 10 {
			if len(ListaJugadoresDisponibles)%2 == 0 {
				success = true

			} else {
				return success, fmt.Errorf("El número de jugadores debe ser un número par")
			}
		} else {
			return success, fmt.Errorf("Debe haber almenos 10 jugadores disponibles para jugar un partido")
		}

	} else {
		return success, fmt.Errorf("La lista de jugadores disponibles no puede estar vacía")
	}

	return success, nil
}

func (this *Grupo) ordenarListaJugadoresDisponiblesPorNivelDescendiente(ListaJugadoresDisponibles []string) ([]string, error) {

	if len(ListaJugadoresDisponibles) > 0 {

		sort.SliceStable(ListaJugadoresDisponibles, func(i, j int) bool {
			return this.JugadoresNiveles[ListaJugadoresDisponibles[i]] > this.JugadoresNiveles[ListaJugadoresDisponibles[j]]
		})

		return ListaJugadoresDisponibles, nil

	} else {
		return ListaJugadoresDisponibles, fmt.Errorf("La lista de jugadores disponibles no puede estar vacía")
	}
}

// Función que reparte los 2 jugadores en 2 equipos con nivel similar siguiendo las siguientes reglas partiendo de
// una lista con los jugadores ordenados de mayor a menor nivel:
// 1. Colocar a los 2 mejores jugadores uno en cada equipo
// Para el resto de jugadores:
// 2. Ir añadiendo el siguiente jugador al equipo que sume menos nivel entre todos sus componentes si es que este no está lleno
// 2.1 Si está lleno, añadirlo al otro equipo
func (this *Grupo) repartirJugadoresDisponiblesEn2Equipos(ListaJugadoresDisponiblesOrdenados []string) ([]string, int, []string, int, error) {

	var EquipoA []string
	var EquipoB []string
	var NivelTotalEquipoA int
	var NivelTotalEquipoB int

	if len(ListaJugadoresDisponiblesOrdenados) > 0 {

		for _, Jugador := range ListaJugadoresDisponiblesOrdenados {

			if len(EquipoA) == 0 && len(EquipoB) == 0 {
				EquipoA = append(EquipoA, Jugador)
				NivelTotalEquipoA += this.JugadoresNiveles[Jugador]
			} else if len(EquipoB) == 0 {
				EquipoB = append(EquipoB, Jugador)
				NivelTotalEquipoB += this.JugadoresNiveles[Jugador]
			} else if (NivelTotalEquipoA < NivelTotalEquipoB) && (len(EquipoA) < (len(ListaJugadoresDisponiblesOrdenados) / 2)) {
				EquipoA = append(EquipoA, Jugador)
				NivelTotalEquipoA += this.JugadoresNiveles[Jugador]
			} else if len(EquipoB) < (len(ListaJugadoresDisponiblesOrdenados) / 2) {
				EquipoB = append(EquipoB, Jugador)
				NivelTotalEquipoB += this.JugadoresNiveles[Jugador]
			} else {
				EquipoA = append(EquipoA, Jugador)
				NivelTotalEquipoA += this.JugadoresNiveles[Jugador]
			}
		}

		return EquipoA, NivelTotalEquipoA, EquipoB, NivelTotalEquipoB, nil

	} else {
		return nil, 0, nil, 0, fmt.Errorf("La lista de jugadores disponibles no puede estar vacía")
	}
}

func (this *Grupo) crearEquiposIgualadosParaPartido(JugadoresDisponibilidad map[string]bool) ([]string, []string, error) {

	ListaJugadoresDisponibles, error := this.conseguirListaJugadoresDisponibles(JugadoresDisponibilidad)

	if error != nil {
		return nil, nil, fmt.Errorf(error.Error())
	}

	success, error := this.validarListaJugadoresDisponiblesParaPartido(ListaJugadoresDisponibles)

	if success == false && error != nil {
		return nil, nil, fmt.Errorf(error.Error())
	}

	ListaJugadoresDisponiblesOrdenados, error := this.ordenarListaJugadoresDisponiblesPorNivelDescendiente(ListaJugadoresDisponibles)

	if error != nil {
		return nil, nil, fmt.Errorf(error.Error())
	}

	Equipo1, NivelTotalEquipo1, Equipo2, NivelTotalEquipo2, error := this.repartirJugadoresDisponiblesEn2Equipos(ListaJugadoresDisponiblesOrdenados)

	if error != nil {
		return nil, nil, fmt.Errorf(error.Error())
	}

	MargenNivelEquipo1 := float32(NivelTotalEquipo1) * 1.3
	MargenNivelEquipo2 := float32(NivelTotalEquipo2) * 1.3

	if MargenNivelEquipo1 < float32(NivelTotalEquipo2) || MargenNivelEquipo2 < float32(NivelTotalEquipo1) {
		return nil, nil, fmt.Errorf("No se han podido crear 2 equipos igualados (La diferencia entre ambos equipos es > al 30 por ciento)")
	}

	return Equipo1, Equipo2, nil
}
