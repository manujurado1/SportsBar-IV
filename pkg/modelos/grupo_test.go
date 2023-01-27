package modelos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearJugador(t *testing.T) {

	grupo, error := CrearGrupo("GrupoTest")
	assert.Nil(t, error)

	success, error := grupo.crearJugador("Manuel", 120, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "El nivel del jugador debe estar entre 0 y 100")

	success, error = grupo.crearJugador("Manuel", 50, false)
	assert.Equal(t, true, success)
	assert.Equal(t, 1, len(grupo.JugadoresNiveles))
	assert.Equal(t, 1, len(grupo.JugadoresDisponibilidad))
	assert.Equal(t, 50, grupo.JugadoresNiveles["Manuel"])
	assert.False(t, grupo.JugadoresDisponibilidad["Manuel"])

	success, error = grupo.crearJugador("Jorge", 70, true)
	assert.Equal(t, true, success)
	assert.Equal(t, 2, len(grupo.JugadoresNiveles))
	assert.Equal(t, 2, len(grupo.JugadoresDisponibilidad))
	assert.Equal(t, 70, grupo.JugadoresNiveles["Jorge"])
	assert.True(t, grupo.JugadoresDisponibilidad["Jorge"])

	success, error = grupo.crearJugador("Manuel", 75, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "Ya existe un jugador con ese nombre")

}

func TestCambiarDisponibilidadJugador(t *testing.T) {

	JugadoresNiveles := map[string]int{"Manuel": 20}
	JugadoresDisponibilidad := map[string]bool{"Manuel": true}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibilidad: JugadoresDisponibilidad, JugadoresNiveles: JugadoresNiveles}

	success, error := grupo.cambiarDisponibilidadJugador("Manuel", false)
	assert.True(t, success)
	assert.Nil(t, error)
	assert.False(t, grupo.JugadoresDisponibilidad["Manuel"])

	success, error = grupo.cambiarDisponibilidadJugador("Jorge", false)
	assert.False(t, success)
	assert.Errorf(t, error, "No existe un jugador con ese nombre")

}

func TestConseguirListaJugadoresDisponibles(t *testing.T) {
	ListaEsperada := []string{"Javier"}

	JugadoresNiveles := map[string]int{"Manuel": 20, "Javier": 50, "Jorge": 80}
	JugadoresDisponibilidad := map[string]bool{"Manuel": false, "Javier": true, "Jorge": false}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibilidad: JugadoresDisponibilidad, JugadoresNiveles: JugadoresNiveles}

	lista, error := grupo.conseguirListaJugadoresDisponibles(grupo.JugadoresDisponibilidad)
	assert.Nil(t, error)
	assert.Equal(t, ListaEsperada, lista)

	grupo.JugadoresDisponibilidad = make(map[string]bool)
	grupo.JugadoresNiveles = make(map[string]int)
	lista, error = grupo.conseguirListaJugadoresDisponibles(grupo.JugadoresDisponibilidad)
	assert.Nil(t, lista)
	assert.Errorf(t, error, "El map de JugadoresDisponibilidad no puede estar vacío")

}

func TestValidarListaJugadoresDisponiblesParaPartido(t *testing.T) {
	ListaMala1 := []string{}
	ListaMala2 := []string{"Jugador1", "Jugador2"}
	ListaMala3 := []string{"Jugador1", "Jugador2", "Jugador3", "Jugador4", "Jugador5", "Jugador6", "Jugador7",
		"Jugador8", "Jugador9", "Jugador10", "Jugador11"}
	ListaBuena := []string{"Jugador1", "Jugador2", "Jugador3", "Jugador4", "Jugador5", "Jugador6", "Jugador7",
		"Jugador8", "Jugador9", "Jugador10", "Jugador11", "Jugador12"}

	grupo := Grupo{Nombre: "GrupoTest"}

	success, error := grupo.validarListaJugadoresDisponiblesParaPartido(ListaMala1)
	assert.False(t, success)
	assert.Errorf(t, error, "La lista de jugadores disponibles no puede estar vacía")

	success, error = grupo.validarListaJugadoresDisponiblesParaPartido(ListaMala2)
	assert.False(t, success)
	assert.Errorf(t, error, "Debe haber almenos 10 jugadores disponibles para jugar un partido")

	success, error = grupo.validarListaJugadoresDisponiblesParaPartido(ListaMala3)
	assert.False(t, success)
	assert.Errorf(t, error, "El número de jugadores debe ser un número par")

	success, error = grupo.validarListaJugadoresDisponiblesParaPartido(ListaBuena)
	assert.True(t, success)
	assert.Nil(t, error)
}

func TestOrdenarListaJugadoresDisponiblesPorNivelDescenciente(t *testing.T) {

	JugadoresNiveles := map[string]int{"Manuel": 20, "Javier": 50, "Jorge": 80}
	JugadoresDisponibilidad := map[string]bool{"Manuel": true, "Javier": true, "Jorge": true}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibilidad: JugadoresDisponibilidad, JugadoresNiveles: JugadoresNiveles}

	lista, error := grupo.ordenarListaJugadoresDisponiblesPorNivelDescendiente([]string{"Manuel", "Javier", "Jorge"})
	assert.Nil(t, error)
	assert.Equal(t, []string{"Jorge", "Javier", "Manuel"}, lista)

	lista, error = grupo.ordenarListaJugadoresDisponiblesPorNivelDescendiente([]string{})
	assert.Errorf(t, error, "La lista de jugadores disponibles no puede estar vacía")

}

func TestRepartirJugadoresDisponiblesEn2Equipos(t *testing.T) {
	Entrada := map[string]int{"Manuel": 40, "Jorge": 50, "Edu": 10, "Clara": 9, "Migue": 100, "Lorena": 60, "Maria": 20}
	JugadoresDisponibilidad := map[string]bool{"Manuel": true, "Jorge": true, "Edu": true, "Clara": false, "Migue": true, "Lorena": true, "Maria": true}
	Equipo1Esperado := []string{"Migue", "Manuel", "Edu"}
	NivelTotalEquipo1Esperado := 150
	Equipo2Esperado := []string{"Lorena", "Jorge", "Maria"}
	NivelTotalEquipo2Esperado := 130
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibilidad: JugadoresDisponibilidad, JugadoresNiveles: Entrada}

	Equipo1, NivelTotalEquipo1, Equipo2, NivelTotalEquipo2, error := grupo.repartirJugadoresDisponiblesEn2Equipos([]string{"Migue", "Lorena", "Jorge", "Manuel", "Maria", "Edu"})

	assert.Equal(t, Equipo1Esperado, Equipo1)
	assert.Equal(t, NivelTotalEquipo1Esperado, NivelTotalEquipo1)
	assert.Equal(t, Equipo2Esperado, Equipo2)
	assert.Equal(t, NivelTotalEquipo2Esperado, NivelTotalEquipo2)
	assert.Nil(t, error)

	Equipo1, NivelTotalEquipo1, Equipo2, NivelTotalEquipo2, error = grupo.repartirJugadoresDisponiblesEn2Equipos([]string{})

	assert.Nil(t, Equipo1)
	assert.Zero(t, NivelTotalEquipo1)
	assert.Nil(t, Equipo1)
	assert.Zero(t, NivelTotalEquipo2)
	assert.Errorf(t, error, "La lista de jugadores disponibles no puede estar vacía")

}

func TestCrearEquiposIgualadosParaPartido(t *testing.T) {

	Entrada := map[string]int{"Manuel": 30, "Jorge": 50, "Edu": 10, "Clara": 90, "Migue": 100, "Alberto": 70,
		"Javi": 20, "Lorena": 80, "Maria": 60, "Sergio": 40}
	JugadoresDisponibilidad := map[string]bool{"Manuel": true, "Jorge": true, "Edu": true, "Migue": true, "Clara": true, "Alberto": true,
		"Javi": true, "Lorena": true, "Maria": true, "Sergio": true}

	// Siguiendo las indicaciones del issue #54, los equipos son equilibrados si los dos mejores jugadores están en equipos diferentes y siempre se añade el siguiente jugador
	// de más nivel al equipo que menos nivel sume, siempre que este equipo no esté completo (Se considera completo si su longitud es de la mitad de la lista de jugadores disponibles)
	//Por lo cuál, la salida correcta para este caso sería, paso por paso:
	// 1. Miguel al equipo 1 (100 puntos totales) y Clara al equipo 2 (90 puntos totales)
	// 2. Lorena iría al equipo 2 ya que suma menos nivel y este quedaría así : Clara,Lorena (170 puntos totales)
	// 3. El siguiente jugador sería Alberto que íria al equipo 1 junto a Miguel (170 puntos totales)
	// 4. La siguiente jugadora sería María, y ante el caso de que los equipos están igualados se prioriza el equipo 2 (230 puntos) (Esto se hace así para "contrarestar" la ventaja que tiene el equipo A teniendo al mejor jugador)
	// 5. El siguiente sería jorge que iría al equipo 1 (220 puntos)
	// 6. Sergio iría al equipo 1 de nuevo (260 puntos)
	// 7. Y por último Javi iría al equipo 2 (250 puntos) y Edu al equipo 1 (270 puntos) obteniendo 2 equipos con nivel total muy similares
	SalidaEsperada := "Equipo 1 = Migue,Alberto,Jorge,Sergio,Edu Equipo 2 = Clara,Lorena,Maria,Manuel,Javi"
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibilidad: JugadoresDisponibilidad, JugadoresNiveles: Entrada}

	solucion, error := grupo.crearEquiposIgualadosParaPartido(grupo.JugadoresDisponibilidad)

	assert.Equal(t, SalidaEsperada, solucion)
	assert.Nil(t, error)
}
