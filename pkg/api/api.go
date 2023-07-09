package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manujurado1/SportsBar-IV/pkg/modelos"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	GruposAmigos = map[string]modelos.GrupoAmigos{}
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Bienvenido a SportsBar API REST, accede a /docs/index.html para consultar la documentación")
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/grupo-amigos", obtenerGruposAmigos)
	r.GET("/grupo-amigos/:nombre-grupo", obtenerGrupoAmigos)
	r.PUT("/grupo-amigos", crearGrupoAmigos)
	r.POST("/grupo-amigos/:nombre-grupo/amigo", aniadirAmigoAlGrupo)
	r.POST("/grupo-amigos/:nombre-grupo/disponibilidad-amigo", cambiarDisponibilidadAmigo)
	r.GET("/grupo-amigos/:nombre-grupo/equipos-igualados", obtenerEquiposIgualados)
	r.POST("/grupo-amigos/:nombre-grupo/resultado-partido", modificarNivelesTrasPartido)
	return r
}

// obtenerGruposAmigos	godoc
// @Summary 			Obtener los grupos de amigos
// @Description 		Obtener todos los grupos de amigos que se han creado
// @Produce 			application/json
// @Tags 				GrupoAmigos
// @Success				200 {object} map[string]modelos.GrupoAmigos{}
// @Router 				/grupo-amigos [get]
func obtenerGruposAmigos(c *gin.Context) {
	c.JSON(http.StatusOK, GruposAmigos)
}

// obtenerGrupoAmigo	godoc
// @Summary 			Obtener el grupo de amigos con el nombre indicado
// @Description 		Se obtiene el grupo de amigo con el nombre indicado
// @Produce 			application/json
// @Tags 				GrupoAmigos
// @Param				nombre-grupo path string true "Obtener grupo de amigos"
// @Success				200 {object} modelos.GrupoAmigos{}
// @Router 				/grupo-amigos/{nombre-grupo} [get]
func obtenerGrupoAmigos(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	// Verificar si el grupo de amigos  existe
	grupo, ok := GruposAmigos[nombreGrupo]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": "No existe un grupo de amigos con ese nombre"})
		return
	}

	c.JSON(http.StatusOK, grupo)
}

// crearGrupoAmigos		godoc
// @Summary 			Crear grupo amigos
// @Description 		Crear un nuevo grupo de amigos
// @Produce 			application/json
// @Tags 				GrupoAmigos
// @Param				NuevoGrupo body modelos.NuevoGrupo true "Nombre del grupo y lista de amigos"
// @Success				201 {object} modelos.GrupoAmigos{}
// @Router 				/grupo-amigos [put]
func crearGrupoAmigos(c *gin.Context) {

	nuevoGrupo := modelos.NuevoGrupo{}

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

// añadirAmigo			godoc
// @Summary 			Añadir amigo
// @Description 		Añadir un nuevo amigo al grupo
// @Produce 			application/json
// @Tags 				GrupoAmigos
// @Param				nombre-grupo path string true "Obtener grupo de amigos"
// @Param				AmigoAAñadir body modelos.AmigoAAniadir true "Nick y fecha de nacimiento del amigo a añadir"
// @Success				201 {object} modelos.GrupoAmigos{}
// @Router 				/grupo-amigos/{nombre-grupo}/amigo [post]
func aniadirAmigoAlGrupo(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	// Verificar si el grupo de amigos  existe
	grupo, ok := GruposAmigos[nombreGrupo]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": "No existe un grupo de amigos con ese nombre"})
		return
	}

	amigoAAniadir := modelos.AmigoAAniadir{}

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

// cambiarDisponibilidadAmigo	godoc
// @Summary 					Cambiar disponibilidad amigo
// @Description 				Cambiar la disponibilidad de un amigo del grupo
// @Produce 					application/json
// @Tags 						GrupoAmigos
// @Param						nombre-grupo path string true "Obtener grupo de amigos"
// @Param						AmigoAModificar body modelos.AmigoAModificar true "Identificador y disponibilidad del amigo al que se le quiere cambiar la disponibilidad"
// @Success						200 {object} modelos.RespuestaModificarAmigo
// @Router 						/grupo-amigos/{nombre-grupo}/disponibilidad-amigo/ [post]
func cambiarDisponibilidadAmigo(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	// Verificar si el grupo de amigos  existe
	grupo, ok := GruposAmigos[nombreGrupo]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": "No existe un grupo de amigos con ese nombre"})
		return
	}

	amigoAModificar := modelos.AmigoAModificar{}
	if err := c.ShouldBindJSON(&amigoAModificar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := grupo.CambiarDisponibilidadAmigo(amigoAModificar.Identificador, amigoAModificar.Disponible)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grupo.NivelYDisponibilidadAmigos)
}

// obtenerEquiposIgualados		godoc
// @Summary 					Obtener equipos igualados
// @Description 				Obtener equipos igualados en función de los jugadores disponibles de ese equipo
// @Produce 					application/json
// @Tags 						GrupoAmigos
// @Param						nombre-grupo path string true "Obtener grupo de amigos"
// @Success						200 {object} modelos.EquiposIgualados{}
// @Router 						/grupo-amigos/{nombre-grupo}/equipos-igualados/ [get]
func obtenerEquiposIgualados(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	EquiposIgualados := modelos.EquiposIgualados{}

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

	c.JSON(http.StatusOK, EquiposIgualados)
}

// modificarNiveles				godoc
// @Summary 					Introducir resultado del partido
// @Description 				Introducir el resultado del partido para que el sistema modifique los niveles necesarios.
// @Produce 					application/json
// @Tags 						GrupoAmigos
// @Param						nombre-grupo path string true "Obtener grupo de amigos"
// @Param						Partido body modelos.Partido true "Equipos y Resultado de ambos equipos en el partido"
// @Success						200 {object} modelos.RespuestaModificarAmigo
// @Router 						/grupo-amigos/{nombre-grupo}/resultado-partido [post]
func modificarNivelesTrasPartido(c *gin.Context) {
	nombreGrupo := c.Param("nombre-grupo")

	partido := modelos.Partido{}

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

	c.JSON(http.StatusOK, grupo.NivelYDisponibilidadAmigos)
}
