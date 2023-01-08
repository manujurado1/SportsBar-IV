package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearJugador(t *testing.T) {
	grupo := Grupo{}
	grupo.JugadoresNiveles = make(map[string]int)
	assert.Errorf(t, grupo.crearJugador("Manuel", 20, false), "El nivel del jugador debe estar entre 0 y 10")
	assert.Equal(t, nil, grupo.crearJugador("Manuel", 5, false))
	assert.Errorf(t, grupo.crearJugador("Manuel", 6, false), "Ya existe un jugador con ese nombre")

}

func TestCambiarDisponibilidadJugador(t *testing.T) {
	grupo := Grupo{}
	grupo.JugadoresNiveles = make(map[string]int)
	assert.Errorf(t, grupo.cambiarDisponibilidadJugador("Manuel"), "No existe un jugador con ese nombre")
	grupo.JugadoresNiveles["Manuel"] = 6
	assert.Equal(t, 0, len(grupo.JugadoresDisponibles))
	grupo.cambiarDisponibilidadJugador("Manuel")
	assert.Equal(t, 1, len(grupo.JugadoresDisponibles))
	assert.Equal(t, "Manuel", grupo.JugadoresDisponibles[0])
	grupo.cambiarDisponibilidadJugador("Manuel")
	assert.Equal(t, 0, len(grupo.JugadoresDisponibles))

}

func TestValidarListaJugadoresDisponibles(t *testing.T) {
	grupo := Grupo{}
	ListaValida, err := grupo.validarListaJugadoresDisponiblesParaPartido(grupo.JugadoresDisponibles)
	assert.Equal(t, ListaValida, false)
	assert.Errorf(t, err, "La lista de jugadores disponibles está vacía")
	grupo.JugadoresDisponibles = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	ListaValida, err = grupo.validarListaJugadoresDisponiblesParaPartido(grupo.JugadoresDisponibles)
	assert.Equal(t, ListaValida, false)
	assert.Nil(t, err)
	grupo.JugadoresDisponibles = append(grupo.JugadoresDisponibles, "10", "11")
	ListaValida, err = grupo.validarListaJugadoresDisponiblesParaPartido(grupo.JugadoresDisponibles)
	assert.Equal(t, ListaValida, false)
	assert.Nil(t, err)
	grupo.JugadoresDisponibles = append(grupo.JugadoresDisponibles, "12")
	ListaValida, err = grupo.validarListaJugadoresDisponiblesParaPartido(grupo.JugadoresDisponibles)
	assert.Equal(t, ListaValida, true)
	assert.Nil(t, err)
}

func TestRepartirJugadoresEnEquipoPorNivel(t *testing.T) {
	Entrada := map[string]int{"Manuel": 3, "Jorge": 5, "Edu": 1, "Clara": 9, "Migue": 10}
	JugadoresDisponibles := []string{"Manuel", "Jorge", "Edu", "Migue"}
	Equipo1Esperado := []string{"Migue", "Edu"}
	Equipo2Esperado := []string{"Jorge", "Manuel"}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibles: JugadoresDisponibles, JugadoresNiveles: Entrada}

	Equipo1, Equipo2, err := grupo.repartirJugadoresEnEquiposPorNivel(grupo.JugadoresNiveles, grupo.JugadoresDisponibles)

	assert.Equal(t, Equipo1, Equipo1Esperado)
	assert.Equal(t, Equipo2, Equipo2Esperado)
	assert.Nil(t, err)

}

func TestOrdenarJugadoresDisponiblesPorNivel(t *testing.T) {
	Entrada := map[string]int{"Manuel": 3, "Jorge": 5, "Edu": 1, "Clara": 9, "Migue": 10}
	JugadoresDisponibles := []string{"Manuel", "Jorge", "Edu", "Migue"}
	SalidaEsperada := []string{"Migue", "Jorge", "Manuel", "Edu"}
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibles: JugadoresDisponibles, JugadoresNiveles: Entrada}
	solucion, err := grupo.ordenarJugadoresDisponilesPorNivel(grupo.JugadoresNiveles, grupo.JugadoresDisponibles)

	assert.Equal(t, SalidaEsperada, solucion)
	assert.Nil(t, err)
}

func TestCrearEquiposIgualadosParaPartido(t *testing.T) {
	Entrada := map[string]int{"Manuel": 3, "Jorge": 5, "Edu": 1, "Clara": 9, "Migue": 10, "Alberto": 7, "Javi": 2, "Lorena": 8, "Maria": 6, "Sergio": 4}
	JugadoresDisponibles := []string{"Manuel", "Jorge", "Edu", "Migue", "Clara", "Alberto", "Javi", "Lorena", "Maria", "Sergio"}
	SalidaEsperada := "Equipo 1 = Migue,Alberto,Jorge,Sergio,Edu Equipo 2 = Clara,Lorena,Maria,Manuel,Javi"
	grupo := Grupo{Nombre: "GrupoTest", JugadoresDisponibles: JugadoresDisponibles, JugadoresNiveles: Entrada}
	solucion, err := grupo.crearEquiposIgualadosParaPartido(grupo.JugadoresNiveles, grupo.JugadoresDisponibles)

	assert.Equal(t, SalidaEsperada, solucion)
	assert.Nil(t, err)

}
