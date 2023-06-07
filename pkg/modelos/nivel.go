package modelos

type Nivel int

const (
	NivelMinimo                    Nivel = 0
	NivelMaximo                    Nivel = 10
	NivelPorOmision                Nivel = 5
	CantidadAumentarDisminuirNivel Nivel = 1
)

func NewNivel() Nivel {
	return NivelPorOmision
}

func (n Nivel) AumentarNivel() Nivel {
	if n == NivelMaximo {
		return n
	}

	return n + CantidadAumentarDisminuirNivel
}

func (n Nivel) DisminuirNivel() Nivel {
	if n == NivelMinimo {
		return n
	}

	return n - CantidadAumentarDisminuirNivel
}
