package application

import (
	"fmt"

	"appointments/domain/entities"
	"appointments/domain/ports"
)

type PacienteService struct {
	repo ports.PacienteRepository
}

func NuevoPacienteService(repo ports.PacienteRepository) *PacienteService {
	return &PacienteService{repo: repo}
}

// Caso de uso 1: Registrar paciente.
func (s *PacienteService) Registrar(paciente entities.Paciente) (entities.Paciente, error) {
	if paciente.Nombre == "" {
		return entities.Paciente{}, fmt.Errorf("el nombre es requerido")
	}
	if paciente.Email == "" {
		return entities.Paciente{}, fmt.Errorf("el email es requerido")
	}
	if paciente.Telefono == "" {
		return entities.Paciente{}, fmt.Errorf("el teléfono es requerido")
	}
	if paciente.FechaNacimiento.IsZero() {
		return entities.Paciente{}, fmt.Errorf("la fecha de nacimiento es requerida")
	}

	return s.repo.Crear(paciente)
}

func (s *PacienteService) BuscarPorID(id int) (entities.Paciente, error) {
	if id <= 0 {
		return entities.Paciente{}, fmt.Errorf("id inválido: %d", id)
	}
	return s.repo.BuscarPorID(id)
}

func (s *PacienteService) BuscarTodos() ([]entities.Paciente, error) {
	return s.repo.BuscarTodos()
}
