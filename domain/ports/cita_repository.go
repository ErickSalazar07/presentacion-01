package ports

import "appointments/domain/entities"

// CitaRepository define las operaciones de persistencia
// que el dominio necesita para la entidad Cita.
// La implementación concreta vive en adapters/postgres.
type CitaRepository interface {
	// Crear agenda una nueva cita y retorna la entidad con su ID asignado.
	// Caso de uso 3: Agendar cita médica.
	Crear(cita entities.Cita) (entities.Cita, error)

	// BuscarPorPaciente retorna todas las citas de un paciente.
	// Caso de uso 4: Consultar citas por paciente.
	BuscarPorPaciente(pacienteID int) ([]entities.Cita, error)

	// BuscarPorDoctor retorna todas las citas de un doctor.
	// Caso de uso 5: Consultar citas por doctor.
	BuscarPorDoctor(doctorID int) ([]entities.Cita, error)

	// ActualizarEstado cambia el estado de una cita existente.
	// Caso de uso 6: Actualizar estado (pendiente, confirmada, completada).
	// Caso de uso 7: Cancelar cita (estado = cancelada).
	ActualizarEstado(id int, estado entities.EstadoCita) (entities.Cita, error)
}
