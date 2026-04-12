package entities

import "time"

// Doctor representa un medico en el sistema de citas.
// Entidad de dominio pura, sin dependencias externas.
type Doctor struct {
	ID          int
	Nombre      string
	Especialidad string
	Email       string
	CreatedAt   time.Time

	// Relacion: un doctor puede atender multiples citas.
	Citas []Cita
}
