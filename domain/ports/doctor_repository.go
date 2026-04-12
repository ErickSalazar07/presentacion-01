package ports

import "appointments/domain/entities"

// DoctorRepository define las operaciones de persistencia
// que el dominio necesita para la entidad Doctor.
// La implementación concreta vive en adapters/postgres.
type DoctorRepository interface {
	// Crear registra un nuevo doctor y retorna la entidad con su ID asignado.
	Crear(doctor entities.Doctor) (entities.Doctor, error)

	// BuscarPorID retorna un doctor por su identificador.
	BuscarPorID(id int) (entities.Doctor, error)

	// BuscarTodos retorna todos los doctores registrados.
	BuscarTodos() ([]entities.Doctor, error)
}
