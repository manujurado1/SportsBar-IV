package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manujurado1/SportsBar-IV/pkg/modelos"
	"github.com/stretchr/testify/assert"
)

var (
	Amigo = `{
		"Nick": "Lorena",
		"FechaNacimiento": "2001-10-12T16:30:40Z"
	}`
	AmigoAModificar = `{
		"Identificador": "Alberto2January",
		"Disponible": false
	}`

	CuerpoTipo1 = `{
		"NombreDelGrupo": "Grupo1",
		"ListaAmigos": [
		  {
			"Nick": "Alberto",
			"FechaNacimiento": "2006-01-02T15:04:05Z"
		  },
		  {
			"Nick": "Migue",
			"FechaNacimiento": "2007-03-04T10:12:30Z"
		  }
		]
	}`
	CuerpoTipo2 = `{
		"NombreDelGrupo": "Grupo2",
		"ListaAmigos": [
		  {
			"Nick": "Juan",
			"FechaNacimiento": "2006-01-02T15:04:05Z"
		  },
		  {
			"Nick": "Jose",
			"FechaNacimiento": "2007-03-04T10:12:30Z"
		  }
		]
	}`
	CuerpoTipo3 = `{
		"NombreDelGrupo": "Grupo1",
		"ListaAmigos": [
		  {
			"Nick": "Amigo1",
			"FechaNacimiento": "2006-01-02T15:04:05Z"
		  },
		  {
			"Nick": "Amigo2",
			"FechaNacimiento": "2007-03-04T10:12:30Z"
		  },
		  {
			"Nick": "Amigo3",
			"FechaNacimiento": "2005-08-20T18:45:00Z"
		  },
		  {
			"Nick": "Amigo4",
			"FechaNacimiento": "2008-12-10T09:30:15Z"
		  },
		  {
			"Nick": "Amigo5",
			"FechaNacimiento": "2003-06-18T12:20:00Z"
		  },
		  {
			"Nick": "Amigo6",
			"FechaNacimiento": "2009-07-25T17:55:30Z"
		  },
		  {
			"Nick": "Amigo7",
			"FechaNacimiento": "2004-09-05T08:40:45Z"
		  },
		  {
			"Nick": "Amigo8",
			"FechaNacimiento": "2010-11-15T14:25:10Z"
		  },
		  {
			"Nick": "Amigo9",
			"FechaNacimiento": "2002-02-28T20:15:20Z"
		  },
		  {
			"Nick": "Amigo10",
			"FechaNacimiento": "2011-04-08T06:05:35Z"
		  },
		  {
			"Nick": "Amigo11",
			"FechaNacimiento": "2001-10-12T16:30:40Z"
		  },
		  {
			"Nick": "Amigo12",
			"FechaNacimiento": "2012-09-22T22:50:55Z"
		  }
		]
	}`
	CuerpoTipo4 = `{
		"Equipo1": [
		  "Amigo24March",
		  "Amigo815November",
		  "Amigo1112October",
		  "Amigo518June",
		  "Amigo928February",
		  "Amigo410December"
		],
		"ResultadoEquipo1": 0,
		"Equipo2": [
		  "Amigo320August",
		  "Amigo75September",
		  "Amigo108April",
		  "Amigo12January",
		  "Amigo625July",
		  "Amigo1222September"
		],
		"ResultadoEquipo2": 5
	  }`

	SalidaEsperadaTipo1 = `{"NombreDelGrupo":"Grupo1","ListaAmigos":[{"Identificador":"Alberto2January"},{"Identificador":"Migue4March"}],"NivelYDisponibilidadAmigos":{"Alberto2January":{"Nivel":5,"Disponible":true},"Migue4March":{"Nivel":5,"Disponible":true}}}`
	SalidaEsperadaTipo2 = `{"Grupo1":{"NombreDelGrupo":"Grupo1","ListaAmigos":[{"Identificador":"Alberto2January"},{"Identificador":"Migue4March"}],"NivelYDisponibilidadAmigos":{"Alberto2January":{"Nivel":5,"Disponible":true},"Migue4March":{"Nivel":5,"Disponible":true}}},"Grupo2":{"NombreDelGrupo":"Grupo2","ListaAmigos":[{"Identificador":"Juan2January"},{"Identificador":"Jose4March"}],"NivelYDisponibilidadAmigos":{"Jose4March":{"Nivel":5,"Disponible":true},"Juan2January":{"Nivel":5,"Disponible":true}}}}`
	SalidaEsperadaTipo3 = `{"NombreDelGrupo":"Grupo1","ListaAmigos":[{"Identificador":"Alberto2January"},{"Identificador":"Migue4March"},{"Identificador":"Lorena12October"}],"NivelYDisponibilidadAmigos":{"Alberto2January":{"Nivel":5,"Disponible":true},"Lorena12October":{"Nivel":5,"Disponible":true},"Migue4March":{"Nivel":5,"Disponible":true}}}`
	SalidaEsperadaTipo4 = `{"Alberto2January":{"Nivel":5,"Disponible":false},"Migue4March":{"Nivel":5,"Disponible":true}}`
	SalidaEsperadaTipo5 = `{"Equipo1":["Amigo518June","Amigo12January","Amigo410December","Amigo75September","Amigo1112October","Amigo24March"],"Equipo2":["Amigo815November","Amigo928February","Amigo320August","Amigo625July","Amigo108April","Amigo1222September"]}`
)

func TestCrearGruposAmigos(t *testing.T) {
	GruposAmigos = make(map[string]modelos.GrupoAmigos)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/grupo-amigo", bytes.NewBuffer([]byte(CuerpoTipo1)))
	assert.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, SalidaEsperadaTipo1, w.Body.String())
}

func TestObtenerGrupoAmigo(t *testing.T) {
	GruposAmigos = make(map[string]modelos.GrupoAmigos)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/grupo-amigo", bytes.NewBuffer([]byte(CuerpoTipo1)))
	router.ServeHTTP(w, req)

	w2 := httptest.NewRecorder()
	req2, err := http.NewRequest("GET", "/grupo-amigo/Grupo1", nil)
	assert.Nil(t, err)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, SalidaEsperadaTipo1, w2.Body.String())
}

func TestObtenerGruposAmigos(t *testing.T) {
	GruposAmigos = make(map[string]modelos.GrupoAmigos)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/grupo-amigo", bytes.NewBuffer([]byte(CuerpoTipo1)))
	router.ServeHTTP(w, req)
	req, _ = http.NewRequest("POST", "/grupo-amigo", bytes.NewBuffer([]byte(CuerpoTipo2)))
	router.ServeHTTP(w, req)

	w2 := httptest.NewRecorder()
	req2, err := http.NewRequest("GET", "/grupos-amigos", nil)
	assert.Nil(t, err)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, SalidaEsperadaTipo2, w2.Body.String())
}

func TestAniadirAmigo(t *testing.T) {
	GruposAmigos = make(map[string]modelos.GrupoAmigos)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/grupo-amigo", bytes.NewBuffer([]byte(CuerpoTipo1)))
	router.ServeHTTP(w, req)

	w2 := httptest.NewRecorder()
	req2, err := http.NewRequest("POST", "/aniadir-amigo/Grupo1", bytes.NewBuffer([]byte(Amigo)))
	assert.Nil(t, err)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 201, w2.Code)
	assert.Equal(t, SalidaEsperadaTipo3, w2.Body.String())

}

func TestCambiarDisponibilidad(t *testing.T) {
	GruposAmigos = make(map[string]modelos.GrupoAmigos)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/grupo-amigo", bytes.NewBuffer([]byte(CuerpoTipo1)))
	router.ServeHTTP(w, req)

	w2 := httptest.NewRecorder()
	req2, err := http.NewRequest("POST", "/cambiar-disponibilidad-amigo/Grupo1", bytes.NewBuffer([]byte(AmigoAModificar)))
	assert.Nil(t, err)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, SalidaEsperadaTipo4, w2.Body.String())
}

func TestObtenerEquiposIgualados(t *testing.T) {
	GruposAmigos = make(map[string]modelos.GrupoAmigos)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/grupo-amigo", bytes.NewBuffer([]byte(CuerpoTipo3)))
	router.ServeHTTP(w, req)

	w2 := httptest.NewRecorder()
	req2, err := http.NewRequest("GET", "/obtener-equipos-igualados/Grupo1", bytes.NewBuffer([]byte(AmigoAModificar)))
	assert.Nil(t, err)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
}

func TestModificarNivelesTrasPartido(t *testing.T) {
	GruposAmigos = make(map[string]modelos.GrupoAmigos)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/grupo-amigo", bytes.NewBuffer([]byte(CuerpoTipo3)))
	router.ServeHTTP(w, req)

	w2 := httptest.NewRecorder()
	req2, err := http.NewRequest("POST", "/actualizar-niveles-tras-partido/Grupo1", bytes.NewBuffer([]byte(CuerpoTipo4)))
	assert.Nil(t, err)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
}
