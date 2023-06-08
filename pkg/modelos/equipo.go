package modelos

type Equipo struct {
	listaNombreAmigoDentoDelGrupo []string
}

func NewEquipo() Equipo {
	return Equipo{
		listaNombreAmigoDentoDelGrupo: []string{},
	}
}

func (e Equipo) ObtenerEquipo() []string {
	return e.listaNombreAmigoDentoDelGrupo
}

func (e Equipo) RellenarEquipo(nombreAmigoDentroDelGrupo []string) Equipo {
	return Equipo{
		listaNombreAmigoDentoDelGrupo: nombreAmigoDentroDelGrupo,
	}
}
