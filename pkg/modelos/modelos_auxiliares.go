package modelos

import "time"

type NuevoGrupo struct {
	NombreDelGrupo string
	ListaAmigos    []struct {
		Nick            string
		FechaNacimiento time.Time
	}
}

type AmigoAAniadir struct {
	Nick            string
	FechaNacimiento time.Time
}

type AmigoAModificar struct {
	Identificador string
	Disponible    bool
}

type RespuestaModificarAmigo map[string]EstadoAmigo

type EquiposIgualados struct {
	Equipo1 []string
	Equipo2 []string
}

type Partido struct {
	Equipo1          []string
	ResultadoEquipo1 uint
	Equipo2          []string
	ResultadoEquipo2 uint
}
