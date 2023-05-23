package modelos

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearJugador(t *testing.T) {

	grupo := CrearGrupo("GrupoTest")

	success, error := grupo.crearJugador("Manuel", 120, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "El nivel del jugador debe estar entre 0 y 100")

	success, error = grupo.crearJugador("Manuel", 50, false)
	estadisticaEsperada := EstadisticasJugador{50, false}
	assert.Equal(t, true, success)
	assert.Equal(t, 1, len(grupo.Jugadores))
	assert.Equal(t, estadisticaEsperada, grupo.Jugadores["Manuel"])

	success, error = grupo.crearJugador("Jorge", 70, true)
	estadisticaEsperada = EstadisticasJugador{70, true}
	assert.Equal(t, true, success)
	assert.Equal(t, 2, len(grupo.Jugadores))
	assert.Equal(t, estadisticaEsperada, grupo.Jugadores["Jorge"])

	success, error = grupo.crearJugador("Manuel", 75, false)
	assert.Equal(t, false, success)
	assert.Errorf(t, error, "Ya existe un jugador con ese nombre")

}

func TestCambiarDisponibilidadJugador(t *testing.T) {

	Estadistica := EstadisticasJugador{20, true}
	Jugadores := map[string]EstadisticasJugador{"Manuel": Estadistica}
	grupo := Grupo{Nombre: "GrupoTest", Jugadores: Jugadores}

	success, error := grupo.cambiarDisponibilidadJugador("Manuel", false)
	EstadisticaEsperada := EstadisticasJugador{20, false}
	assert.True(t, success)
	assert.Nil(t, error)
	assert.Equal(t, EstadisticaEsperada, grupo.Jugadores["Manuel"])

	success, error = grupo.cambiarDisponibilidadJugador("Jorge", false)
	assert.False(t, success)
	assert.Errorf(t, error, "No existe un jugador con ese nombre")

}

func TestCrearEquiposIgualadosParaPartido(t *testing.T) {

	// Comprobamos que si no hay los suficientes jugadores disponibles no se pueden crear equipos igualados
	Jugadores := map[string]EstadisticasJugador{"Manuel": EstadisticasJugador{30, false}, "Jorge": EstadisticasJugador{50, false},
		"Edu": EstadisticasJugador{10, false}, "Clara": EstadisticasJugador{90, true}, "Migue": EstadisticasJugador{100, true},
		"Alberto": EstadisticasJugador{70, false}, "Javi": EstadisticasJugador{20, true}, "Lorena": EstadisticasJugador{80, true},
		"Maria": EstadisticasJugador{60, true}, "Sergio": EstadisticasJugador{40, false}}

	grupo := Grupo{Nombre: "GrupoTest", Jugadores: Jugadores}

	Equipo1, Equipo2, error := grupo.crearEquiposIgualadosParaPartido(grupo.Jugadores)
	assert.Errorf(t, error, "La lista de jugadores disponibles debe ser de almenos 10 personas y ser un número par")

	// Comprobamos la posibilidad de que aunque todo haya ido bien, sea imposible conseguir 2 equipos igualados

	Jugadores = map[string]EstadisticasJugador{"Manuel": EstadisticasJugador{3, true}, "Jorge": EstadisticasJugador{5, true},
		"Edu": EstadisticasJugador{1, true}, "Clara": EstadisticasJugador{9, true}, "Migue": EstadisticasJugador{100, true},
		"Alberto": EstadisticasJugador{7, true}, "Javi": EstadisticasJugador{2, true}, "Lorena": EstadisticasJugador{8, true},
		"Maria": EstadisticasJugador{6, true}, "Sergio": EstadisticasJugador{4, true}}
	grupo.Jugadores = Jugadores

	Equipo1, Equipo2, error = grupo.crearEquiposIgualadosParaPartido(grupo.Jugadores)
	assert.Errorf(t, error, "No se ha conseguido crear 2 equipos igualados")
	igualados, error := grupo.estanIgualados(Equipo1, Equipo2)
	assert.False(t, igualados)

	// Comprobamos caso correcto

	Jugadores = map[string]EstadisticasJugador{"Manuel": EstadisticasJugador{30, true}, "Jorge": EstadisticasJugador{50, true},
		"Edu": EstadisticasJugador{10, true}, "Clara": EstadisticasJugador{90, true}, "Migue": EstadisticasJugador{100, true},
		"Alberto": EstadisticasJugador{70, true}, "Javi": EstadisticasJugador{20, true}, "Lorena": EstadisticasJugador{80, true},
		"Maria": EstadisticasJugador{60, true}, "Sergio": EstadisticasJugador{40, true}}

	grupo.Jugadores = Jugadores

	Equipo1Esperado := []string{"Migue", "Alberto", "Jorge", "Sergio", "Edu"}
	Equipo2Esperado := []string{"Clara", "Lorena", "Maria", "Manuel", "Javi"}

	Equipo1, Equipo2, error = grupo.crearEquiposIgualadosParaPartido(grupo.Jugadores)
	assert.Equal(t, Equipo1Esperado, Equipo1)
	assert.Equal(t, Equipo2Esperado, Equipo2)
	assert.Nil(t, error)
	igualados, error = grupo.estanIgualados(Equipo1, Equipo2)
	assert.True(t, igualados)

	// Comprobamos caso aleatorio para comprobar si el algoritmo es mejor que la aleatoriedad (issue)
	var NivelEquipo1 uint = 0
	var NivelEquipo2 uint = 0
	var NivelTotalEquipo1 uint = 0
	var NivelTotalEquipo2 uint = 0

	Jugadores = map[string]EstadisticasJugador{"Manuel": EstadisticasJugador{73, true}, "Jorge": EstadisticasJugador{70, true},
		"Edu": EstadisticasJugador{42, true}, "Clara": EstadisticasJugador{66, true}, "Migue": EstadisticasJugador{28, true},
		"Alberto": EstadisticasJugador{16, true}, "Javi": EstadisticasJugador{29, true}, "Lorena": EstadisticasJugador{50, true},
		"Maria": EstadisticasJugador{10, true}, "Sergio": EstadisticasJugador{18, true}}

	grupo.Jugadores = Jugadores
	Equipo1, Equipo2, error = grupo.crearEquiposIgualadosParaPartido(grupo.Jugadores)
	igualados, error = grupo.estanIgualados(Equipo1, Equipo2)
	assert.True(t, igualados)

	for _, jugador := range Equipo1 {
		NivelTotalEquipo1 += grupo.Jugadores[jugador].Nivel
	}

	for _, jugador := range Equipo2 {
		NivelTotalEquipo2 += grupo.Jugadores[jugador].Nivel
	}

	NivelEquipo1 = NivelTotalEquipo1 / uint(len(Equipo1))
	NivelEquipo2 = NivelTotalEquipo2 / uint(len(Equipo2))

	DiferenciaNiveles := math.Abs(float64(NivelEquipo1) - float64(NivelEquipo2))

	assert.True(t, DiferenciaNiveles < 15)

}

func TestConsultarHistorial(t *testing.T) {
	listaPartidos := []DatosPartido{}
	partido1 := DatosPartido{ResultadoEquipo1: 3, ResultadoEquipo2: 2, Equipo1: []string{"j1", "j2", "j3", "j4", "j5"}, Equipo2: []string{"j6", "j7", "j8", "j9", "j10"}}
	partido2 := DatosPartido{ResultadoEquipo1: 0, ResultadoEquipo2: 5, Equipo1: []string{"j1", "j2", "j3", "j4", "j5"}, Equipo2: []string{"j6", "j7", "j8", "j9", "j10"}}
	listaPartidos = append(listaPartidos, partido1, partido2)
	grupo := CrearGrupo("GrupoTest")
	grupo.Historial = []DatosPartido{partido1, partido2}

	historial, err := grupo.ConsultarHistorial()
	assert.Nil(t, err)
	assert.Equal(t, listaPartidos, historial)
}

func TestListaMejoresJugadores(t *testing.T) {
	grupo := CrearGrupo("")
	Jugadores := map[string]EstadisticasJugador{"Manuel": EstadisticasJugador{30, true}, "Jorge": EstadisticasJugador{50, true},
		"Edu": EstadisticasJugador{10, true}, "Clara": EstadisticasJugador{90, true}, "Migue": EstadisticasJugador{100, true},
		"Alberto": EstadisticasJugador{70, true}, "Javi": EstadisticasJugador{20, true}, "Lorena": EstadisticasJugador{80, true},
		"Maria": EstadisticasJugador{60, true}, "Sergio": EstadisticasJugador{40, true}}

	grupo.Jugadores = Jugadores

	resultadoEsperado := []string{"Migue - 100", "Clara - 90", "Lorena - 80", "Alberto - 70", "Maria - 60", "Jorge - 50", "Sergio - 40", "Manuel - 30", "Javi - 20", "Edu - 10"}
	resultado, err := grupo.ObtenerMejoresJugadores()
	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado)
}
