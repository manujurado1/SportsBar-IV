package modelos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	//Caso incorrecto intentando cambiar disponibilidad de un amigo que no está en el grupo
	grupo, _ := NewGrupoAmigos("Prueba", listaAmigos)
	amigoNuevo, _ := NewAmigo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))

	errorDisponibilidad := grupo.CambiarDisponibilidadAmigo(amigoNuevo, false)
	errorEsperado := ErrorAmigoInexistente.Error() + ": " + amigoNuevo.ObtenerId()

	assert.EqualError(t, errorDisponibilidad, errorEsperado)

	//Caso correcto
	_ = grupo.CrearAmigoYAniadirAlGrupo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	amigoNuevo = grupo.ListaAmigos[2]
	errorDisponibilidad = grupo.CambiarDisponibilidadAmigo(amigoNuevo, false)
	estado := grupo.NivelYDisponibilidadAmigos[amigoNuevo.ObtenerId()]

	assert.Nil(t, errorDisponibilidad)
	assert.False(t, estado.Disponible)

}

func TestAumentarNivelAmigo(t *testing.T) {
	amigo, _ := NewAmigo("Javi", time.Date(1999, time.August, 17, 0, 0, 0, 0, time.Now().Location()))
	amigo2, _ := NewAmigo("Migue", time.Date(2001, time.January, 22, 0, 0, 0, 0, time.Now().Location()))
	listaAmigos := []Amigo{amigo, amigo2}

	//Caso incorrecto intentando cambiar disponibilidad de un amigo que no está en el grupo
	grupo, _ := NewGrupoAmigos("Prueba", listaAmigos)
	amigoNuevo, _ := NewAmigo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))

	errorNivel := grupo.AumentarNivelAmigo(amigoNuevo)
	errorEsperado := ErrorAmigoInexistente.Error() + ": " + amigoNuevo.ObtenerId()

	assert.EqualError(t, errorNivel, errorEsperado)

	//Caso correcto sin alcanzar el límite
	_ = grupo.CrearAmigoYAniadirAlGrupo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	amigoAniadido := grupo.ListaAmigos[2]
	errorNivel = grupo.AumentarNivelAmigo(amigoAniadido)

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

	//Caso incorrecto intentando cambiar disponibilidad de un amigo que no está en el grupo
	grupo, _ := NewGrupoAmigos("Prueba", listaAmigos)
	amigoNuevo, _ := NewAmigo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))

	errorNivel := grupo.AumentarNivelAmigo(amigoNuevo)
	errorEsperado := ErrorAmigoInexistente.Error() + ": " + amigoNuevo.ObtenerId()

	assert.EqualError(t, errorNivel, errorEsperado)

	//Caso correcto sin alcanzar el límite
	_ = grupo.CrearAmigoYAniadirAlGrupo("Migue", time.Date(1999, time.April, 27, 0, 0, 0, 0, time.Now().Location()))
	amigoAniadido := grupo.ListaAmigos[2]
	errorNivel = grupo.DisminuirNivelAmigo(amigoAniadido)

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
	_ = grupo.CrearAmigoYAniadirAlGrupo("Jose", time.Date(1999, time.December, 1, 0, 0, 0, 0, time.Now().Location()))
	_ = grupo.CrearAmigoYAniadirAlGrupo("Javi", time.Date(1999, time.September, 7, 0, 0, 0, 0, time.Now().Location()))
	_ = grupo.CrearAmigoYAniadirAlGrupo("Guille", time.Date(1999, time.February, 17, 0, 0, 0, 0, time.Now().Location()))
	_ = grupo.CrearAmigoYAniadirAlGrupo("Manu", time.Date(1999, time.June, 27, 0, 0, 0, 0, time.Now().Location()))
	_ = grupo.CrearAmigoYAniadirAlGrupo("Alex", time.Date(1999, time.July, 13, 0, 0, 0, 0, time.Now().Location()))
	_ = grupo.CrearAmigoYAniadirAlGrupo("Jorge", time.Date(1999, time.July, 3, 0, 0, 0, 0, time.Now().Location()))
	_ = grupo.CrearAmigoYAniadirAlGrupo("Sergio", time.Date(1999, time.May, 4, 0, 0, 0, 0, time.Now().Location()))
	_ = grupo.CrearAmigoYAniadirAlGrupo("Fran", time.Date(1999, time.March, 8, 0, 0, 0, 0, time.Now().Location()))

	for nombre := range grupo.NivelYDisponibilidadAmigos {
		grupo.NivelYDisponibilidadAmigos[nombre] = EstadoAmigo{Nivel: 1, Disponible: true}
	}

	grupo.NivelYDisponibilidadAmigos["Javi17August"] = EstadoAmigo{Nivel: NivelMaximo, Disponible: true}

	equipo1, equipo2, errorCrearEquipos = grupo.CrearDosEquiposIgualados(grupo.NivelYDisponibilidadAmigos)

	assert.Equal(t, Equipo{}, equipo1)
	assert.Equal(t, Equipo{}, equipo2)
	assert.EqualError(t, errorCrearEquipos, ErrorImposibilidadCrearEquiposIgualados.Error())

	//Caso correcto
	grupo.NivelYDisponibilidadAmigos["Guille17February"] = EstadoAmigo{Nivel: Nivel(0), Disponible: true}
	grupo.NivelYDisponibilidadAmigos["Jose1December"] = EstadoAmigo{Nivel: Nivel(1), Disponible: true}
	grupo.NivelYDisponibilidadAmigos["Javi17August"] = EstadoAmigo{Nivel: Nivel(2), Disponible: true}
	grupo.NivelYDisponibilidadAmigos["Jorge3July"] = EstadoAmigo{Nivel: Nivel(3), Disponible: true}
	grupo.NivelYDisponibilidadAmigos["Alex13July"] = EstadoAmigo{Nivel: Nivel(4), Disponible: true}
	grupo.NivelYDisponibilidadAmigos["Manu27June"] = EstadoAmigo{Nivel: Nivel(5), Disponible: true}
	grupo.NivelYDisponibilidadAmigos["Migue22January"] = EstadoAmigo{Nivel: Nivel(6), Disponible: true}
	grupo.NivelYDisponibilidadAmigos["Fran8March"] = EstadoAmigo{Nivel: Nivel(7), Disponible: true}
	grupo.NivelYDisponibilidadAmigos["Sergio4May"] = EstadoAmigo{Nivel: Nivel(8), Disponible: true}
	grupo.NivelYDisponibilidadAmigos["Javi7September"] = EstadoAmigo{Nivel: Nivel(9), Disponible: true}

	equipo1, equipo2, errorCrearEquipos = grupo.CrearDosEquiposIgualados(grupo.NivelYDisponibilidadAmigos)

	listaEquipo1Esperada := []string{"Javi7September", "Migue22January", "Alex13July", "Jorge3July", "Guille17February"}
	listaEquipo2Esperada := []string{"Sergio4May", "Fran8March", "Manu27June", "Javi17August", "Jose1December"}

	assert.Equal(t, listaEquipo1Esperada, equipo1.listaNombreAmigoDentoDelGrupo)
	assert.Equal(t, listaEquipo2Esperada, equipo2.listaNombreAmigoDentoDelGrupo)
	assert.Nil(t, errorCrearEquipos)

}
