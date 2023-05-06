package modelos

type EstadisticasJugador struct {
	Nivel          uint
	Disponibilidad bool
}

func (this *EstadisticasJugador) setDisponibilidad(Disponibilidad bool) {
	this.Disponibilidad = Disponibilidad
}
