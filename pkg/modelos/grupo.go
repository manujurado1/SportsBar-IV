package modelos

import (
	"fmt"
	"time"
)

const (
	DisponibilidadPorDefecto bool = true
)

var (
	ErrorNombreGrupoVacio = fmt.Errorf("Un grupo no puede tener como nombre un string vacío")
	ErrorAmigoDuplicado   = fmt.Errorf("Ya existe un amigo con ese identificador")
	ErrorAmigoInexistente = fmt.Errorf("No existe el amigo indicado dentro del grupo de amigos")
	ErrorListaAmigosVacia = fmt.Errorf("No se puede crear un grupo de amigos con una lista de amigos vacía")
)

var EstadoAmigoPorDefecto EstadoAmigo = EstadoAmigo{
	Nivel:      NivelPorOmision,
	Disponible: DisponibilidadPorDefecto,
}

type EstadoAmigo struct {
	Nivel      Nivel
	Disponible bool
}
type GrupoAmigos struct {
	NombreDelGrupo             string
	ListaAmigos                []Amigo
	NivelYDisponibilidadAmigos map[string]EstadoAmigo
}

func NewGrupoAmigos(nombre string, listaAmigos []Amigo) (*GrupoAmigos, error) {
	if nombre == "" {
		return nil, ErrorNombreGrupoVacio
	}
	if len(listaAmigos) == 0 {
		return nil, ErrorListaAmigosVacia
	}

	estadosAmigos := map[string]EstadoAmigo{}

	for _, amigo := range listaAmigos {
		estadosAmigos[amigo.ObtenerId()] = EstadoAmigoPorDefecto
	}

	return &GrupoAmigos{
		NombreDelGrupo:             nombre,
		ListaAmigos:                listaAmigos,
		NivelYDisponibilidadAmigos: estadosAmigos,
	}, nil
}

func (g *GrupoAmigos) CrearAmigoYAniadirAlGrupo(nickAmigo string, fechaNacimiento time.Time) error {
	amigo, errorAmigo := NewAmigo(nickAmigo, fechaNacimiento)

	if errorAmigo != nil {
		return errorAmigo
	}

	_, existe := g.NivelYDisponibilidadAmigos[amigo.ObtenerId()]

	if existe {
		return ErrorAmigoDuplicado
	}

	g.ListaAmigos = append(g.ListaAmigos, amigo)
	g.NivelYDisponibilidadAmigos[amigo.ObtenerId()] = EstadoAmigoPorDefecto

	return nil
}

func (g *GrupoAmigos) CambiarDisponibilidadAmigo(amigo Amigo, disponibilidad bool) error {
	estado, existe := g.NivelYDisponibilidadAmigos[amigo.ObtenerId()]

	if !existe {
		return ErrorAmigoInexistente
	}

	estado.Disponible = disponibilidad
	g.NivelYDisponibilidadAmigos[amigo.ObtenerId()] = estado

	return nil
}

func (g *GrupoAmigos) AumentarNivelAmigo(amigo Amigo) error {
	estado, existe := g.NivelYDisponibilidadAmigos[amigo.ObtenerId()]

	if !existe {
		return ErrorAmigoInexistente
	}

	estado.Nivel = estado.Nivel.AumentarNivel()
	g.NivelYDisponibilidadAmigos[amigo.ObtenerId()] = estado

	return nil
}

func (g *GrupoAmigos) DisminuirNivelAmigo(amigo Amigo) error {
	estado, existe := g.NivelYDisponibilidadAmigos[amigo.ObtenerId()]

	if !existe {
		return ErrorAmigoInexistente
	}

	estado.Nivel = estado.Nivel.DisminuirNivel()
	g.NivelYDisponibilidadAmigos[amigo.ObtenerId()] = estado

	return nil
}
