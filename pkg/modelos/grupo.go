package modelos

import (
	"fmt"
	"sort"
)

type Grupo struct {
	Nombre                  string
	JugadoresDisponibilidad map[string]bool
	JugadoresNiveles        map[string]uint
}

// Constructor que crea un nuevo Grupo con el nombre pasado por parámetro, inicializando  el map de JugadoresNiveles y el de Jugadores Disponibilidad vacíos.
func CrearGrupo(Nombre string) *Grupo {

	JugadoresDisponibilidad := make(map[string]bool)
	JugadoresNiveles := make(map[string]uint)
	NombreGrupo := Nombre
	if NombreGrupo == "" {
		NombreGrupo = "Grupo"
	}
	grupo := Grupo{NombreGrupo, JugadoresDisponibilidad, JugadoresNiveles}
	return &grupo
}

// Función que añade un jugador en el map de JugadoresNiveles si ese nombre no tiene una entrada existente y el nivel es correcto
// y lo inserta en la lista de disponibles en caso de que así se indique
func (this *Grupo) crearJugador(NombreJugador string, Nivel uint, Disponibilidad bool) (bool, error) {
	var success bool = false
	_, existe := this.JugadoresNiveles[NombreJugador]
	if existe == false {
		if Nivel >= 1 && Nivel <= 100 {

			this.JugadoresNiveles[NombreJugador] = Nivel
			this.JugadoresDisponibilidad[NombreJugador] = Disponibilidad
			success = true

		} else {
			return success, fmt.Errorf("El nivel del jugador debe estar entre 1 y 100")
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

	if len(ListaJugadoresDisponibles) >= 10 && len(ListaJugadoresDisponibles)%2 == 0 {
		success = true

	} else {
		return success, fmt.Errorf("La lista de jugadores disponibles debe ser de almenos 10 personas y ser un número par")
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
func (this *Grupo) repartirJugadoresDisponiblesEn2Equipos(ListaJugadoresDisponiblesOrdenados []string) ([]string, []string, error) {

	var EquipoA []string
	var EquipoB []string
	var NivelTotalEquipoA uint
	var NivelTotalEquipoB uint

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

		return EquipoA, EquipoB, nil

	} else {
		return nil, nil, fmt.Errorf("La lista de jugadores disponibles no puede estar vacía")
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

	Equipo1, Equipo2, error := this.repartirJugadoresDisponiblesEn2Equipos(ListaJugadoresDisponiblesOrdenados)

	if error != nil {
		return nil, nil, fmt.Errorf(error.Error())
	}

	Igualados, error := this.estanIgualados(Equipo1, Equipo2)
	if !Igualados {
		return Equipo1, Equipo2, fmt.Errorf("No se ha conseguido crear 2 equipos igualados")
	}

	return Equipo1, Equipo2, nil
}

func (this *Grupo) estanIgualados(Equipo1 []string, Equipo2 []string) (bool, error) {

	Igualados := false
	var NivelTotalEquipo1 uint = 0
	var NivelTotalEquipo2 uint = 0

	if len(Equipo1) == 0 || len(Equipo2) == 0 {
		return Igualados, fmt.Errorf("Los equipos no pueden estar vacíos")
	}

	if len(Equipo1) == len(Equipo2) && len(Equipo1) >= 5 {

		for _, jugador := range Equipo1 {
			NivelTotalEquipo1 += this.JugadoresNiveles[jugador]
		}

		for _, jugador := range Equipo2 {
			NivelTotalEquipo2 += this.JugadoresNiveles[jugador]
		}

		NivelEquipo1 := NivelTotalEquipo1 / uint(len(Equipo1))
		NivelEquipo2 := NivelTotalEquipo2 / uint(len(Equipo2))

		if (NivelEquipo1+10 > NivelEquipo2) && (NivelEquipo2+10 > NivelEquipo1) {
			Igualados = true
		}

	}

	return Igualados, nil

}
