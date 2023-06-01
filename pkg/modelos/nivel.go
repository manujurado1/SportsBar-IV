package modelos

type Nivel int

const (
	NivelMinimo Nivel = 1
	NivelMaximo Nivel = 10
)

func (n *Nivel) Set(valor Nivel) {
	if valor < NivelMinimo {
		*n = NivelMinimo
	} else if valor > NivelMaximo {
		*n = NivelMaximo
	} else {
		*n = valor
	}
}
