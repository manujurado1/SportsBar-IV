package modelos

import (
	"fmt"
	"time"
)

var (
	ErrorNombreGrupoVacio = fmt.Errorf("Un grupo no puede tener como nombre un string vacío")
)

type GrupoAmigos struct {
	NombreDelGrupo string
	ListaAmigos    []Amigo
}

func NewGrupoAmigos(nombre string) (*GrupoAmigos, error) {
	if nombre == "" {
		return nil, ErrorNombreGrupoVacio
	}

	return &GrupoAmigos{
		NombreDelGrupo: nombre,
		ListaAmigos:    []Amigo{},
	}, nil
}

func (g *GrupoAmigos) CrearAmigoYAniadirAlGrupo(nickAmigo string, fechaNacimiento time.Time) (bool, error) {
	success := false
	Amigo, errorAmigo := NewAmigo(nickAmigo, fechaNacimiento)

	if errorAmigo != nil {
		return success, errorAmigo
	}

	g.ListaAmigos = append(g.ListaAmigos, Amigo)
	success = true

	return success, nil
}
