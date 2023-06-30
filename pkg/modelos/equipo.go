package modelos

type Equipo struct {
	ListaNombreAmigoDentoDelGrupo []string
}

func NewEquipo() Equipo {
	return Equipo{
		ListaNombreAmigoDentoDelGrupo: []string{},
	}
}

func (e Equipo) ObtenerEquipo() []string {
	return e.ListaNombreAmigoDentoDelGrupo
}

func (e Equipo) RellenarEquipo(nombreAmigoDentroDelGrupo []string) Equipo {
	return Equipo{
		ListaNombreAmigoDentoDelGrupo: nombreAmigoDentroDelGrupo,
	}
}
