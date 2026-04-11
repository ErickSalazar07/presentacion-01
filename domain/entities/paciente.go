package entities

import "time"

// Paciente representa un paciente en el sistema de citas medicas.
// no contiene tags de ORM ni dependencias de infraestructura.
type Paciente struct {
	ID              int
	Nombre          string
	Email           string
	Telefono        string
	FechaNacimiento time.Time
	CreatedAt       time.Time

	// Relacion: un paciente puede tener múltiples citas.
	Citas []Cita
}
