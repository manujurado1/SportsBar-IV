package modelos

import "errors"

var (
	ErrorFaltaNombre                  = errors.New("Se necesita introducir un nombre para crear el grupo")
	ErrorAmigoRepetido                = errors.New("Ya existe un amigo con ese nombre en el grupo de amigos")
	ErrorAmigosDisponiblesInadecuados = errors.New("Para crear 2 equipos debe haber un número par de amigos disponibles igual o superior a 10")
	ErrorImposibilidadEquipos         = errors.New("Ha sido imposible crear 2 equipos igualados con los amigos disponibles")
)

type GrupoAmigos struct {
	Nombre string
	Amigos []*Amigo
}
