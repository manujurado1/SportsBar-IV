package modelos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	amigo1, _              = NewAmigo("Javi", time.Date(1999, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo2, _              = NewAmigo("Migue", time.Date(2001, time.January, 22, 0, 0, 0, 0, time.Now().Location()))
	amigo3, _              = NewAmigo("Jose", time.Date(1999, time.December, 1, 0, 0, 0, 0, time.Now().Location()))
	amigo4, _              = NewAmigo("Javi", time.Date(1999, time.September, 7, 0, 0, 0, 0, time.Now().Location()))
	amigo5, _              = NewAmigo("Guille", time.Date(1999, time.February, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo6, _              = NewAmigo("Manu", time.Date(1999, time.June, 27, 0, 0, 0, 0, time.Now().Location()))
	amigo7, _              = NewAmigo("Alex", time.Date(1999, time.July, 13, 0, 0, 0, 0, time.Now().Location()))
	amigo8, _              = NewAmigo("Jorge", time.Date(1999, time.July, 3, 0, 0, 0, 0, time.Now().Location()))
	amigo9, _              = NewAmigo("Sergio", time.Date(1999, time.May, 4, 0, 0, 0, 0, time.Now().Location()))
	amigo10, _             = NewAmigo("Fran", time.Date(1999, time.March, 8, 0, 0, 0, 0, time.Now().Location()))
	ListaAmigosCompleta    = []Amigo{amigo1, amigo2, amigo3, amigo4, amigo5, amigo6, amigo7, amigo8, amigo9, amigo10}
	EstadosAmigosImposible = map[string]EstadoAmigo{
		"Javi17August":     {NivelMaximo, DisponibilidadPorDefecto},
		"Migue22January":   {NivelMinimo, DisponibilidadPorDefecto},
		"Jose1December":    {NivelMinimo, DisponibilidadPorDefecto},
		"Javi7September":   {NivelMinimo, DisponibilidadPorDefecto},
		"Guille17February": {NivelMinimo, DisponibilidadPorDefecto},
		"Manu27June":       {NivelMinimo, DisponibilidadPorDefecto},
		"Alex13July":       {NivelMinimo, DisponibilidadPorDefecto},
		"Jorge3July":       {NivelMinimo, DisponibilidadPorDefecto},
		"Sergio4May":       {NivelMinimo, DisponibilidadPorDefecto},
		"Fran8March":       {NivelMinimo, DisponibilidadPorDefecto},
	}
	EstadosAmigosPosible = map[string]EstadoAmigo{
		"Javi17August":     {Nivel(2), DisponibilidadPorDefecto},
		"Migue22January":   {Nivel(6), DisponibilidadPorDefecto},
		"Jose1December":    {Nivel(1), DisponibilidadPorDefecto},
		"Javi7September":   {Nivel(9), DisponibilidadPorDefecto},
		"Guille17February": {Nivel(0), DisponibilidadPorDefecto},
		"Manu27June":       {Nivel(5), DisponibilidadPorDefecto},
		"Alex13July":       {Nivel(4), DisponibilidadPorDefecto},
		"Jorge3July":       {Nivel(3), DisponibilidadPorDefecto},
		"Sergio4May":       {Nivel(8), DisponibilidadPorDefecto},
		"Fran8March":       {Nivel(7), DisponibilidadPorDefecto},
	}
)

func TestNewGrupoAmigos(t *testing.T) {

	amigo, _ := NewAmigo("Javi", time.Date(1999, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo2, _ := NewAmigo("Migue", time.Date(2001, time.January, 22, 0, 0, 0, 0, time.Now().Location()))
	listaAmigos := []Amigo{amigo, amigo2}
	estadoAmigos := map[string]EstadoAmigo{}
	estadoAmigos[amigo.ObtenerId()] = EstadoAmigoPorDefecto
	estadoAmigos[amigo2.ObtenerId()] = EstadoAmigoPorDefecto

	//Caso incorrecto: Nombre vacío
	grupo, errorGrupo := NewGrupoAmigos("", listaAmigos)

	assert.EqualError(t, errorGrupo, "Un grupo no puede tener como nombre un string vacío")
	assert.Nil(t, grupo)

	//Caso incorrecto: Lista de amigos vacía
	grupo2, errorGrupo2 := NewGrupoAmigos("Nombre", []Amigo{})

	assert.EqualError(t, errorGrupo2, "No se puede crear un grupo de amigos con una lista de amigos vacía")
	assert.Nil(t, grupo2)

	//Caso correcto
	grupo3, errorGrupo3 := NewGrupoAmigos("Prueba", listaAmigos)
	grupoEsperado := GrupoAmigos{"Prueba", listaAmigos, estadoAmigos}

	assert.Nil(t, errorGrupo3)
	assert.Equal(t, grupoEsperado, *grupo3)
}

func TestCrearAmigoYAniadirAlGrupo(t *testing.T) {

	amigo, _ := NewAmigo("Javi", time.Date(1999, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo2, _ := NewAmigo("Migue", time.Date(2001, time.January, 22, 0, 0, 0, 0, time.Now().Location()))
	listaAmigos := []Amigo{amigo, amigo2}

	//Caso incorrecto por nombre amigo vacío
	grupo, _ := NewGrupoAmigos("Prueba", listaAmigos)
	errorAmigo := grupo.CrearAmigoYAniadirAlGrupo("", time.Date(1990, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	assert.EqualError(t, errorAmigo, "El nick de un amigo no puede ser un string vacío")

	//Caso correcto añadiendo 2 amigos con el mismo nombre
	grupo2, _ := NewGrupoAmigos("Prueba2", listaAmigos)

	errorAmigo = grupo2.CrearAmigoYAniadirAlGrupo("Jose", time.Date(1990, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	assert.Nil(t, errorAmigo)

	errorAmigo = grupo2.CrearAmigoYAniadirAlGrupo("Jose", time.Date(1990, time.August, 19, 0, 0, 0, 0, time.Now().Location()))
	assert.Nil(t, errorAmigo)
	assert.Equal(t, 4, len(grupo2.ListaAmigos))
}

func TestModificarDisponibilidadAmigo(t *testing.T) {

	amigo, _ := NewAmigo("Javi", time.Date(1999, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo2, _ := NewAmigo("Migue", time.Date(2001, time.January, 22, 0, 0, 0, 0, time.Now().Location()))
	listaAmigos := []Amigo{amigo, amigo2}
	grupo, _ := NewGrupoAmigos("Prueba", listaAmigos)

	//Caso incorrecto intentando cambiar disponibilidad de un amigo que no está en el grupo

	assertModificarAmigoInexistente(t, grupo, "Disponibilidad")

	//Caso correcto
	_ = grupo.CrearAmigoYAniadirAlGrupo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	amigoNuevo := grupo.ListaAmigos[2]
	errorDisponibilidad := grupo.CambiarDisponibilidadAmigo(amigoNuevo, false)
	estado := grupo.NivelYDisponibilidadAmigos[amigoNuevo.ObtenerId()]

	assert.Nil(t, errorDisponibilidad)
	assert.False(t, estado.Disponible)

}

func TestAumentarNivelAmigo(t *testing.T) {
	amigo, _ := NewAmigo("Javi", time.Date(1999, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo2, _ := NewAmigo("Migue", time.Date(2001, time.January, 22, 0, 0, 0, 0, time.Now().Location()))
	listaAmigos := []Amigo{amigo, amigo2}
	grupo, _ := NewGrupoAmigos("Prueba", listaAmigos)

	//Caso incorrecto intentando aumentar nivel de un amigo que no está en el grupo

	assertModificarAmigoInexistente(t, grupo, "Aumentar")

	//Caso correcto sin alcanzar el límite
	_ = grupo.CrearAmigoYAniadirAlGrupo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	amigoAniadido := grupo.ListaAmigos[2]
	errorNivel := grupo.AumentarNivelAmigo(amigoAniadido)

	assert.Nil(t, errorNivel)
	assert.Greater(t, grupo.NivelYDisponibilidadAmigos[amigoAniadido.ObtenerId()].Nivel, NivelPorOmision)

	//Caso correcto alcanzando el límite

	estado := grupo.NivelYDisponibilidadAmigos[amigoAniadido.ObtenerId()]
	estado.Nivel = NivelMaximo
	grupo.NivelYDisponibilidadAmigos[amigoAniadido.ObtenerId()] = estado
	errorNivel = grupo.AumentarNivelAmigo(amigoAniadido)
	assert.Nil(t, errorNivel)
	assert.Equal(t, grupo.NivelYDisponibilidadAmigos[amigoAniadido.ObtenerId()].Nivel, NivelMaximo)
}

func TestDisminuirNivelAmigo(t *testing.T) {

	amigo, _ := NewAmigo("Javi", time.Date(1999, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo2, _ := NewAmigo("Migue", time.Date(2001, time.January, 22, 0, 0, 0, 0, time.Now().Location()))
	listaAmigos := []Amigo{amigo, amigo2}
	grupo, _ := NewGrupoAmigos("Prueba", listaAmigos)

	//Caso incorrecto intentando disminuir de un amigo que no está en el grupo
	assertModificarAmigoInexistente(t, grupo, "Disminuir")

	//Caso correcto sin alcanzar el límite
	_ = grupo.CrearAmigoYAniadirAlGrupo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	amigoAniadido := grupo.ListaAmigos[2]
	errorNivel := grupo.DisminuirNivelAmigo(amigoAniadido)

	assert.Nil(t, errorNivel)
	assert.Less(t, grupo.NivelYDisponibilidadAmigos[amigoAniadido.ObtenerId()].Nivel, NivelPorOmision)

	//Caso correcto alcanzando el límite

	estado := grupo.NivelYDisponibilidadAmigos[amigoAniadido.ObtenerId()]
	estado.Nivel = NivelMinimo
	grupo.NivelYDisponibilidadAmigos[amigoAniadido.ObtenerId()] = estado
	errorNivel = grupo.DisminuirNivelAmigo(amigoAniadido)
	assert.Nil(t, errorNivel)
	assert.Equal(t, grupo.NivelYDisponibilidadAmigos[amigoAniadido.ObtenerId()].Nivel, NivelMinimo)

}

func assertModificarAmigoInexistente(t *testing.T, grupo *GrupoAmigos, tipo string) {
	amigoNuevo, _ := NewAmigo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	errorEsperado := ErrorAmigoInexistente.Error() + ": " + amigoNuevo.ObtenerId()

	var errorInexistente error

	if tipo == "Aumentar" {
		errorInexistente = grupo.AumentarNivelAmigo(amigoNuevo)
	} else if tipo == "Disminuir" {
		errorInexistente = grupo.DisminuirNivelAmigo(amigoNuevo)
	} else if tipo == "Disponibilidad" {
		errorInexistente = grupo.CambiarDisponibilidadAmigo(amigoNuevo, false)
	}

	assert.EqualError(t, errorInexistente, errorEsperado)
}

func TestCrearEquiposIgualados(t *testing.T) {

	// Caso incorrecto no hay jugadores suficientes
	amigo, _ := NewAmigo("Javi", time.Date(1999, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo2, _ := NewAmigo("Migue", time.Date(2001, time.January, 22, 0, 0, 0, 0, time.Now().Location()))
	listaAmigos := []Amigo{amigo, amigo2}
	grupo, _ := NewGrupoAmigos("Prueba", listaAmigos)

	equipo1, equipo2, errorCrearEquipos := grupo.CrearDosEquiposIgualados(grupo.NivelYDisponibilidadAmigos)

	assert.Equal(t, Equipo{}, equipo1)
	assert.Equal(t, Equipo{}, equipo2)
	assert.EqualError(t, errorCrearEquipos, ErrorJugadoresDisponiblesNoAptosParaCrearEquipos.Error())

	//Caso incorrecto por imposibilidad de crear 2 equipos igualados
	grupo.ListaAmigos = ListaAmigosCompleta

	grupo.NivelYDisponibilidadAmigos = EstadosAmigosImposible

	equipo1, equipo2, errorCrearEquipos = grupo.CrearDosEquiposIgualados(grupo.NivelYDisponibilidadAmigos)

	assert.Equal(t, Equipo{}, equipo1)
	assert.Equal(t, Equipo{}, equipo2)
	assert.EqualError(t, errorCrearEquipos, ErrorImposibilidadCrearEquiposIgualados.Error())

	//Caso correcto
	grupo.NivelYDisponibilidadAmigos = EstadosAmigosPosible

	equipo1, equipo2, errorCrearEquipos = grupo.CrearDosEquiposIgualados(grupo.NivelYDisponibilidadAmigos)

	listaEquipo1Esperada := []string{"Javi7September", "Migue22January", "Alex13July", "Jorge3July", "Guille17February"}
	listaEquipo2Esperada := []string{"Sergio4May", "Fran8March", "Manu27June", "Javi17August", "Jose1December"}

	assert.Equal(t, listaEquipo1Esperada, equipo1.listaNombreAmigoDentoDelGrupo)
	assert.Equal(t, listaEquipo2Esperada, equipo2.listaNombreAmigoDentoDelGrupo)
	assert.Nil(t, errorCrearEquipos)

}
