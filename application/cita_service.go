package application

import (
	"fmt"

	"appointments/domain/entities"
	"appointments/domain/ports"
)

type CitaService struct {
	citaRepo     ports.CitaRepository
	pacienteRepo ports.PacienteRepository
	doctorRepo   ports.DoctorRepository
}

func NuevoCitaService(
	citaRepo ports.CitaRepository,
	pacienteRepo ports.PacienteRepository,
	doctorRepo ports.DoctorRepository,
) *CitaService {
	return &CitaService{
		citaRepo:     citaRepo,
		pacienteRepo: pacienteRepo,
		doctorRepo:   doctorRepo,
	}
}

// Caso de uso 3: Agendar cita médica.
// Verifica que el paciente y el doctor existan antes de crear la cita.
func (s *CitaService) Agendar(cita entities.Cita) (entities.Cita, error) {
	if cita.Motivo == "" {
		return entities.Cita{}, fmt.Errorf("el motivo es requerido")
	}
	if cita.FechaHora.IsZero() {
		return entities.Cita{}, fmt.Errorf("la fecha y hora es requerida")
	}

	paciente, err := s.pacienteRepo.BuscarPorID(cita.PacienteAsignado.ID)
	if err != nil {
		return entities.Cita{}, fmt.Errorf("paciente no encontrado: %w", err)
	}

	doctor, err := s.doctorRepo.BuscarPorID(cita.DoctorAsignado.ID)
	if err != nil {
		return entities.Cita{}, fmt.Errorf("doctor no encontrado: %w", err)
	}

	cita.PacienteAsignado = paciente
	cita.DoctorAsignado = doctor
	cita.Estado = entities.EstadoPendiente

	return s.citaRepo.Crear(cita)
}

// Caso de uso 4: Consultar citas por paciente.
func (s *CitaService) BuscarPorPaciente(pacienteID int) ([]entities.Cita, error) {
	if pacienteID <= 0 {
		return nil, fmt.Errorf("id de paciente inválido: %d", pacienteID)
	}
	return s.citaRepo.BuscarPorPaciente(pacienteID)
}

// Caso de uso 5: Consultar citas por doctor.
func (s *CitaService) BuscarPorDoctor(doctorID int) ([]entities.Cita, error) {
	if doctorID <= 0 {
		return nil, fmt.Errorf("id de doctor inválido: %d", doctorID)
	}
	return s.citaRepo.BuscarPorDoctor(doctorID)
}

// Caso de uso 6: Actualizar estado de la cita.
func (s *CitaService) ActualizarEstado(id int, estado entities.EstadoCita) (entities.Cita, error) {
	if id <= 0 {
		return entities.Cita{}, fmt.Errorf("id de cita inválido: %d", id)
	}
	if !estado.IsValid() {
		return entities.Cita{}, fmt.Errorf("estado inválido: %s", estado)
	}

	return s.citaRepo.ActualizarEstado(id, estado)
}

// Caso de uso 7: Cancelar cita.
// Es un caso especial de ActualizarEstado con EstadoCancelada.
func (s *CitaService) Cancelar(id int) (entities.Cita, error) {
	return s.ActualizarEstado(id, entities.EstadoCancelada)
}
