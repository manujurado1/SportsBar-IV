package modelos

import (
	"fmt"
)

type Nivel int

const (
	NivelMinimo Nivel = 0
	NivelMaximo Nivel = 10
	NivelMedio  Nivel = 5
)

var (
	ErrorNivelFueraDeRango = fmt.Errorf("El nivel de un jugador debe estar en el rango %d - %d", NivelMinimo, NivelMaximo)
)

func NuevoNivel(valor Nivel) (*Nivel, error) {
	if valor < NivelMinimo || valor > NivelMaximo {
		return nil, ErrorNivelFueraDeRango
	}

	nivel := Nivel(valor)
	return &nivel, nil
}

func (n *Nivel) AumentarNivel(cantidad Nivel) {
	nuevoNivel := *n + cantidad
	if nuevoNivel > NivelMaximo {
		*n = NivelMaximo
	} else {
		*n = nuevoNivel
	}
}

func (n *Nivel) DisminuirNivel(cantidad Nivel) {
	nuevoNivel := *n - cantidad
	if nuevoNivel < NivelMinimo {
		*n = NivelMinimo
	} else {
		*n = nuevoNivel
	}
}
