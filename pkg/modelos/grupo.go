package modelos

import (
	"fmt"
	"time"
)

var (
	ErrorNombreGrupoVacio = fmt.Errorf("Un grupo no puede tener como nombre un string vacío")
	ErrorAmigoDuplicado   = fmt.Errorf("Ya existe un amigo con ese identificador")
	ErrorAmigoInexistente = fmt.Errorf("No existe el amigo indicado dentro del grupo de amigos")
)

const (
	DisponibilidadPorDefecto bool = true
)

type GrupoAmigos struct {
	NombreDelGrupo       string
	ListaAmigos          []Amigo
	NivelAmigos          map[string]Nivel
	DisponibilidadAmigos map[string]bool
}

func NewGrupoAmigos(nombre string) (*GrupoAmigos, error) {
	if nombre == "" {
		return nil, ErrorNombreGrupoVacio
	}

	return &GrupoAmigos{
		NombreDelGrupo:       nombre,
		ListaAmigos:          []Amigo{},
		NivelAmigos:          make(map[string]Nivel),
		DisponibilidadAmigos: make(map[string]bool),
	}, nil
}

func (g *GrupoAmigos) CrearAmigoYAniadirAlGrupo(nickAmigo string, fechaNacimiento time.Time) (bool, error) {
	success := false
	Amigo, errorAmigo := NewAmigo(nickAmigo, fechaNacimiento)

	if errorAmigo != nil {
		return success, errorAmigo
	}

	if g.ExisteAmigo(Amigo) {
		return success, ErrorAmigoDuplicado
	}

	g.ListaAmigos = append(g.ListaAmigos, Amigo)
	g.NivelAmigos[Amigo.ObtenerId()] = NewNivel()
	g.DisponibilidadAmigos[Amigo.ObtenerId()] = DisponibilidadPorDefecto
	success = true

	return success, nil
}

func (g *GrupoAmigos) ExisteAmigo(amigo Amigo) bool {
	existe := false
	idAmigo := amigo.ObtenerId()

	for _, amigo := range g.ListaAmigos {
		if amigo.ObtenerId() == idAmigo {
			existe = true
		}
	}
	return existe

}

func (g *GrupoAmigos) CambiarDisponibilidadAmigo(amigo Amigo, disponibilidad bool) (bool, error) {
	success := false

	if !g.ExisteAmigo(amigo) {
		return success, ErrorAmigoInexistente
	}

	g.DisponibilidadAmigos[amigo.ObtenerId()] = disponibilidad
	success = true

	return success, nil
}

func (g *GrupoAmigos) AumentarNivelAmigo(amigo Amigo) (bool, error) {
	success := false

	if !g.ExisteAmigo(amigo) {
		return success, ErrorAmigoInexistente
	}

	g.NivelAmigos[amigo.ObtenerId()] = g.NivelAmigos[amigo.ObtenerId()].AumentarNivel()
	success = true

	return success, nil
}

func (g *GrupoAmigos) DisminuirNivelAmigo(amigo Amigo) (bool, error) {
	success := false

	if !g.ExisteAmigo(amigo) {
		return success, ErrorAmigoInexistente
	}

	g.NivelAmigos[amigo.ObtenerId()] = g.NivelAmigos[amigo.ObtenerId()].DisminuirNivel()
	success = true

	return success, nil
}
