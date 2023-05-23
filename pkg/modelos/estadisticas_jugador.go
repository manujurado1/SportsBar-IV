package modelos

import "fmt"

type EstadisticasJugador struct {
	Nivel          uint
	Disponibilidad bool
}

// Constructor de la estructura de EstadísticasJugador que crea dicha estructura comprobando que el nivel introducido está dentro
// del rango de niveles permitidos que puede tener un jugador (Nivel mínimo = 1, Nivel máximo = 100).
func crearEstadisticasJugador(Nivel uint, Disponibilidad bool) (*EstadisticasJugador, error) {
	if Nivel >= 1 && Nivel <= 100 {
		Estadistica := EstadisticasJugador{Nivel, Disponibilidad}
		return &Estadistica, nil
	} else {
		return nil, fmt.Errorf("El nivel del jugador debe estar entre 1 y 100")
	}
}

func (this *EstadisticasJugador) setDisponibilidad(Disponibilidad bool) {
	this.Disponibilidad = Disponibilidad
}

// Función que modifica el atributo nivel dentro de la estructura de datos. Se pasa por parámetro la cantidad de nivel a modificar.
// Si este entero es un número positivo serán los niveles a sumar y si es negativo, los que se deben restar.
// Si esta alteración del nivel sobrepasa los rangos de nivel permitidos (De 1 a 100) se truncará el nivel al máximo/mínimo permitido.
func (this *EstadisticasJugador) setNivel(Cantidad int) {
	NuevoNivel := int(this.Nivel) + Cantidad

	if NuevoNivel <= 0 {
		NuevoNivel = 1
	} else if NuevoNivel > 100 {
		NuevoNivel = 100
	}
	this.Nivel = uint(NuevoNivel)

}
