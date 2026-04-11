package ports

import "appointments/domain/entities"

// PacienteRepository define las operaciones de persistencia
// que el dominio necesita para la entidad Paciente.
// La implementación concreta vive en adapters/postgres.
type PacienteRepository interface {
	// Crear registra un nuevo paciente y retorna la entidad con su ID asignado.
	Crear(paciente entities.Paciente) (entities.Paciente, error)

	// BuscarPorID retorna un paciente por su identificador.
	BuscarPorID(id int) (entities.Paciente, error)

	// BuscarTodos retorna todos los pacientes registrados.
	BuscarTodos() ([]entities.Paciente, error)
}
