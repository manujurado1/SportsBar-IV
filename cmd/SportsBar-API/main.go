package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manujurado1/SportsBar-IV/pkg/modelos"
)

var (
	GruposAmigos = map[string]modelos.GrupoAmigos{}
)

func inicializarAPI() {
	r := gin.Default()
	r.GET("/grupos-amigos", obtenerGruposAmigos)
	r.GET("/grupos-amigos/:nombre-grupo", obtenerGrupoAmigos)
	r.POST("/grupos-amigos", crearGrupoAmigos)
	r.POST("/aniadir-amigo/:nombre-grupo", aniadirAmigoAlGrupo)
	r.POST("/cambiar-disponibilidad-amigo/:nombre-grupo", cambiarDisponibilidadAmigo)
	r.GET("/obtener-equipos-igualados/:nombre-grupo", obtenerEquiposIgualados)
	r.POST("/actualizar-niveles-tras-partido/:nombre-grupo", modificarNivelesTrasPartido)
	_ = r.Run()
}
func obtenerGruposAmigos(c *gin.Context) {
	c.JSON(http.StatusAccepted, GruposAmigos)
}

func obtenerGrupoAmigos(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	// Verificar si el grupo de amigos  existe
	grupo, ok := GruposAmigos[nombreGrupo]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": "No existe un grupo de amigos con ese nombre"})
		return
	}

	c.JSON(http.StatusAccepted, grupo)
}

func crearGrupoAmigos(c *gin.Context) {
	var nuevoGrupo struct {
		NombreDelGrupo string
		ListaAmigos    []struct {
			Nick            string
			FechaNacimiento time.Time
		}
	}
	if err := c.ShouldBindJSON(&nuevoGrupo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el grupo de amigos ya existe
	if _, ok := GruposAmigos[nuevoGrupo.NombreDelGrupo]; ok {
		c.JSON(http.StatusConflict, gin.H{"error": "Ya existe un grupo de amigos con ese nombre"})
		return
	}

	// Crear un nuevo grupo de amigos utilizando el constructor NewGrupoAmigos
	ListaAmigos := []modelos.Amigo{}

	for _, amigo := range nuevoGrupo.ListaAmigos {
		amigoNuevo, err := modelos.NewAmigo(amigo.Nick, amigo.FechaNacimiento)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ListaAmigos = append(ListaAmigos, amigoNuevo)
	}

	grupo, err := modelos.NewGrupoAmigos(nuevoGrupo.NombreDelGrupo, ListaAmigos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Agregar el nuevo grupo de amigos a la lista
	GruposAmigos[nuevoGrupo.NombreDelGrupo] = *grupo

	c.JSON(http.StatusCreated, grupo)
}

func aniadirAmigoAlGrupo(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	// Verificar si el grupo de amigos  existe
	grupo, ok := GruposAmigos[nombreGrupo]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": "No existe un grupo de amigos con ese nombre"})
		return
	}

	var amigoAAniadir struct {
		Nick            string
		FechaNacimiento time.Time
	}

	if err := c.ShouldBindJSON(&amigoAAniadir); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := grupo.CrearAmigoYAniadirAlGrupo(amigoAAniadir.Nick, amigoAAniadir.FechaNacimiento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	GruposAmigos[grupo.NombreDelGrupo] = grupo

	c.JSON(http.StatusCreated, grupo)
}

func cambiarDisponibilidadAmigo(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	// Verificar si el grupo de amigos  existe
	grupo, ok := GruposAmigos[nombreGrupo]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": "No existe un grupo de amigos con ese nombre"})
		return
	}

	var amigoAModificar struct {
		Identificador string
		Disponible    bool
	}
	if err := c.ShouldBindJSON(&amigoAModificar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := grupo.CambiarDisponibilidadAmigo(amigoAModificar.Identificador, amigoAModificar.Disponible)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, grupo.NivelYDisponibilidadAmigos)
}

func obtenerEquiposIgualados(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	var EquiposIgualados struct {
		Equipo1 []string
		Equipo2 []string
	}

	// Verificar si el grupo de amigos  existe
	grupo, ok := GruposAmigos[nombreGrupo]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": "No existe un grupo de amigos con ese nombre"})
		return
	}

	Equipo1, Equipo2, err := grupo.CrearDosEquiposIgualados(grupo.NivelYDisponibilidadAmigos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	EquiposIgualados.Equipo1 = Equipo1.ListaNombreAmigoDentoDelGrupo
	EquiposIgualados.Equipo2 = Equipo2.ListaNombreAmigoDentoDelGrupo

	c.JSON(http.StatusAccepted, EquiposIgualados)
}

func modificarNivelesTrasPartido(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	var partido struct {
		Equipo1          []string
		ResultadoEquipo1 uint
		Equipo2          []string
		ResultadoEquipo2 uint
	}

	if err := c.ShouldBindJSON(&partido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el grupo de amigos  existe
	grupo, ok := GruposAmigos[nombreGrupo]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": "No existe un grupo de amigos con ese nombre"})
		return
	}

	Equipo := modelos.NewEquipo()
	Equipo1 := Equipo.RellenarEquipo(partido.Equipo1)
	Equipo2 := Equipo.RellenarEquipo(partido.Equipo2)

	err := grupo.ModificarNivelesTrasPartido(Equipo1, partido.ResultadoEquipo1, Equipo2, partido.ResultadoEquipo2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, grupo.NivelYDisponibilidadAmigos)
}

func main() {
	inicializarAPI()
}
