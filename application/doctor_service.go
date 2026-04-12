package application

import (
	"fmt"

	"appointments/domain/entities"
	"appointments/domain/ports"
)

type DoctorService struct {
	repo ports.DoctorRepository
}

func NuevoDoctorService(repo ports.DoctorRepository) *DoctorService {
	return &DoctorService{repo: repo}
}

// Caso de uso 2: Registrar doctor.
func (s *DoctorService) Registrar(doctor entities.Doctor) (entities.Doctor, error) {
	if doctor.Nombre == "" {
		return entities.Doctor{}, fmt.Errorf("el nombre es requerido")
	}
	if doctor.Email == "" {
		return entities.Doctor{}, fmt.Errorf("el email es requerido")
	}
	if doctor.Especialidad == "" {
		return entities.Doctor{}, fmt.Errorf("la especialidad es requerida")
	}

	return s.repo.Crear(doctor)
}

func (s *DoctorService) BuscarPorID(id int) (entities.Doctor, error) {
	if id <= 0 {
		return entities.Doctor{}, fmt.Errorf("id inválido: %d", id)
	}
	return s.repo.BuscarPorID(id)
}

func (s *DoctorService) BuscarTodos() ([]entities.Doctor, error) {
	return s.repo.BuscarTodos()
}
