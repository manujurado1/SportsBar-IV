package modelos

type EstadisticasJugador struct {
	Nivel          uint
	Disponibilidad bool
}

func (this *EstadisticasJugador) setDisponibilidad(Disponibilidad bool) {
	this.Disponibilidad = Disponibilidad
}

func (this *EstadisticasJugador) setNivel(Nivel int) {
	if Nivel <= 0 {
		this.Nivel = 1
	} else if Nivel > 100 {
		this.Nivel = 100
	} else {
		this.Nivel = uint(Nivel)
	}
}
