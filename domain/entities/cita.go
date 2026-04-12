package entities

import "time"

// EstadoCita representa el ciclo de vida de una cita medica.
// Usar un tipo dedicado (en lugar de string plano) permite al compilador
// detectar valores invalidos y hace el dominio mas expresivo.
type EstadoCita string

const (
	EstadoPendiente  EstadoCita = "pendiente"
	EstadoConfirmada EstadoCita = "confirmada"
	EstadoCancelada  EstadoCita = "cancelada"
	EstadoCompletada EstadoCita = "completada"
)

// IsValid verifica que el estado sea uno de los valores permitidos.
// Se usa en la capa de application antes de persistir.
func (e EstadoCita) IsValid() bool {
	switch e {
	case EstadoPendiente, EstadoConfirmada, EstadoCancelada, EstadoCompletada:
		return true
	}
	return false
}

// Cita es la entidad central del sistema: resuelve la relacion
// muchos-a-muchos entre Paciente y Doctor.
// Contiene IDs foráneos (PacienteID, DoctorID) para persistencia, y
// punteros opcionales a las entidades completas para cuando el caso
// de uso necesite los datos relacionados.
type Cita struct {
	ID         int
	Estado     EstadoCita
	Motivo     string
	FechaHora  time.Time
	CreatedAt  time.Time

	// Relaciones con Paciente y Doctor
	PacienteAsignado Paciente
	DoctorAsignado   Doctor
}
