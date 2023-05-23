package modelos

import (
	"fmt"
	"math"
	"sort"
)

type Grupo struct {
	Nombre    string
	Jugadores map[string]EstadisticasJugador
	Historial []DatosPartido
}

// Constructor que crea un nuevo Grupo con el nombre pasado por parámetro, inicializando  el map de JugadoresNiveles y el de Jugadores Disponibilidad vacíos.
func CrearGrupo(Nombre string) *Grupo {

	Jugadores := make(map[string]EstadisticasJugador)
	NombreGrupo := Nombre
	if NombreGrupo == "" {
		NombreGrupo = "Grupo"
	}
	grupo := Grupo{NombreGrupo, Jugadores, []DatosPartido{}}
	return &grupo
}

// Función que añade un jugador en el map de JugadoresNiveles si ese nombre no tiene una entrada existente y el nivel es correcto
// y lo inserta en la lista de disponibles en caso de que así se indique
func (this *Grupo) crearJugador(NombreJugador string, Nivel uint, Disponibilidad bool) (bool, error) {
	var success bool = false
	_, existe := this.Jugadores[NombreJugador]
	if existe == false {
		estadisticas, error := crearEstadisticasJugador(Nivel, Disponibilidad)
		if error == nil {
			this.Jugadores[NombreJugador] = *estadisticas
			success = true
		} else {
			return success, error
		}
	} else {
		return success, fmt.Errorf("Ya existe un jugador con ese nombre")
	}

	return success, nil
}

func (this *Grupo) cambiarDisponibilidadJugador(NombreJugador string, Disponibilidad bool) (bool, error) {
	var success bool = false
	_, existe := this.Jugadores[NombreJugador]
	if existe == true {
		estadisticas := this.Jugadores[NombreJugador]
		estadisticas.setDisponibilidad(Disponibilidad)
		this.Jugadores[NombreJugador] = estadisticas
		success = true
	} else {
		return success, fmt.Errorf("No existe un jugador con ese nombre")
	}

	return success, nil

}

func (this *Grupo) conseguirListaJugadoresDisponibles(Jugadores map[string]EstadisticasJugador) ([]string, error) {

	var ListaJugadoresDisponibles []string

	if len(Jugadores) > 0 {

		for key, value := range Jugadores {
			if value.Disponibilidad == true {
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

func (this *Grupo) ordenarListaJugadoresPorNivelDescendiente(ListaJugadoresDisponibles []string) ([]string, error) {

	if len(ListaJugadoresDisponibles) > 0 {

		sort.SliceStable(ListaJugadoresDisponibles, func(i, j int) bool {
			return this.Jugadores[ListaJugadoresDisponibles[i]].Nivel > this.Jugadores[ListaJugadoresDisponibles[j]].Nivel
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
				NivelTotalEquipoA += this.Jugadores[Jugador].Nivel
			} else if len(EquipoB) == 0 {
				EquipoB = append(EquipoB, Jugador)
				NivelTotalEquipoB += this.Jugadores[Jugador].Nivel
			} else if (NivelTotalEquipoA < NivelTotalEquipoB) && (len(EquipoA) < (len(ListaJugadoresDisponiblesOrdenados) / 2)) {
				EquipoA = append(EquipoA, Jugador)
				NivelTotalEquipoA += this.Jugadores[Jugador].Nivel
			} else if len(EquipoB) < (len(ListaJugadoresDisponiblesOrdenados) / 2) {
				EquipoB = append(EquipoB, Jugador)
				NivelTotalEquipoB += this.Jugadores[Jugador].Nivel
			} else {
				EquipoA = append(EquipoA, Jugador)
				NivelTotalEquipoA += this.Jugadores[Jugador].Nivel
			}
		}

		return EquipoA, EquipoB, nil

	} else {
		return nil, nil, fmt.Errorf("La lista de jugadores disponibles no puede estar vacía")
	}
}

func (this *Grupo) crearEquiposIgualadosParaPartido(Jugadores map[string]EstadisticasJugador) ([]string, []string, error) {

	ListaJugadoresDisponibles, error := this.conseguirListaJugadoresDisponibles(Jugadores)

	if error != nil {
		return nil, nil, fmt.Errorf(error.Error())
	}

	success, error := this.validarListaJugadoresDisponiblesParaPartido(ListaJugadoresDisponibles)

	if success == false && error != nil {
		return nil, nil, fmt.Errorf(error.Error())
	}

	ListaJugadoresDisponiblesOrdenados, error := this.ordenarListaJugadoresPorNivelDescendiente(ListaJugadoresDisponibles)

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

// Función que calcula el nivel de cada uno de los 2 equipos pasados por parámetro, siendo este nivel la media del nivel de los jugadores que lo forman.
// A continuación, comprueba si ambos equipos están igualados, lo que se traduce a que su diferencia de nivel sea menor a 10 niveles.
func (this *Grupo) estanIgualados(Equipo1 []string, Equipo2 []string) (bool, error) {

	Igualados := false
	var NivelTotalEquipo1 uint = 0
	var NivelTotalEquipo2 uint = 0

	if len(Equipo1) == 0 || len(Equipo2) == 0 {
		return Igualados, fmt.Errorf("Los equipos no pueden estar vacíos")
	}

	if len(Equipo1) == len(Equipo2) && len(Equipo1) >= 5 {

		for _, jugador := range Equipo1 {
			NivelTotalEquipo1 += this.Jugadores[jugador].Nivel
		}

		for _, jugador := range Equipo2 {
			NivelTotalEquipo2 += this.Jugadores[jugador].Nivel
		}

		NivelEquipo1 := NivelTotalEquipo1 / uint(len(Equipo1))
		NivelEquipo2 := NivelTotalEquipo2 / uint(len(Equipo2))

		if (NivelEquipo1+10 > NivelEquipo2) && (NivelEquipo2+10 > NivelEquipo1) {
			Igualados = true
		}

	}

	return Igualados, nil

}

func (this *Grupo) TerminarPartido(Equipo1 []string, ResultadoEquipo1 uint8, ResultadoEquipo2 uint8, Equipo2 []string) error {
	Partido := DatosPartido{Equipo1: Equipo1, ResultadoEquipo1: ResultadoEquipo1, Equipo2: Equipo2, ResultadoEquipo2: ResultadoEquipo2}
	this.Historial = append(this.Historial, Partido)

	if this.esResultadoIgualado(Partido.ResultadoEquipo1, Partido.ResultadoEquipo2) == false {
		this.AjustarNivelesPorResultadoAbultado(Partido)
	}

	return nil
}

func (this *Grupo) ConsultarHistorial() ([]DatosPartido, error) {

	return this.Historial, nil
}

func (this *Grupo) ObtenerMejoresJugadores() ([]string, error) {
	ListaTotalJugadores, err := this.ObtenerListaTotalJugadores()
	if err != nil {
		return []string{}, err
	}
	ListaMejoresJugadores, err := this.ordenarListaJugadoresPorNivelDescendiente(ListaTotalJugadores)
	if err != nil {
		return []string{}, err
	}

	ListaMejoresJugadoresConNivel := []string{}

	for _, jugador := range ListaMejoresJugadores {
		nuevoJugador := jugador + " - " + fmt.Sprint(this.Jugadores[jugador].Nivel)
		ListaMejoresJugadoresConNivel = append(ListaMejoresJugadoresConNivel, nuevoJugador)
	}

	return ListaMejoresJugadoresConNivel, nil

}

func (this *Grupo) ObtenerListaTotalJugadores() ([]string, error) {
	var ListaJugadoresDisponibles []string

	if len(this.Jugadores) > 0 {

		for key := range this.Jugadores {
			ListaJugadoresDisponibles = append(ListaJugadoresDisponibles, key)
		}
		return ListaJugadoresDisponibles, nil

	} else {
		return nil, fmt.Errorf("El map de Jugadores no puede estar vacío")
	}
}

// Función que comprueba si el resultado de un partido ha sido igualado
func (this *Grupo) esResultadoIgualado(ResultadoEquipo1 uint8, ResultadoEquipo2 uint8) bool {
	diferenciaGoles := math.Abs(float64(ResultadoEquipo1) - float64(ResultadoEquipo2))
	diferenciaGolesPartidoIgualado := 3

	if diferenciaGoles <= float64(diferenciaGolesPartidoIgualado) {
		return true
	} else {
		return false
	}
}

// Función que ajusta los niveles de los jugadores de un partido cuando el resultado de este no ha sido igualado
func (this *Grupo) AjustarNivelesPorResultadoAbultado(Partido DatosPartido) error {
	ModificacionNivelGanador := 5
	ModificacionNivelPerdedor := -5

	if Partido.ResultadoEquipo1 > Partido.ResultadoEquipo2 {
		this.AjustarNivelEquipo(Partido.Equipo1, ModificacionNivelGanador)
		this.AjustarNivelEquipo(Partido.Equipo2, ModificacionNivelPerdedor)
	} else {
		this.AjustarNivelEquipo(Partido.Equipo1, ModificacionNivelPerdedor)
		this.AjustarNivelEquipo(Partido.Equipo2, ModificacionNivelGanador)
	}

	return nil
}

// Función que modifica el nivel del equipo pasado como primer parámetro con la cantidad pasada como segundo parámetro.
func (this *Grupo) AjustarNivelEquipo(Equipo []string, Cantidad int) error {
	for _, jugador := range Equipo {
		estadisticas := this.Jugadores[jugador]
		estadisticas.modificarNivel(Cantidad)
		this.Jugadores[jugador] = estadisticas
	}
	return nil
}
