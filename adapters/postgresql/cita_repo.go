package postgresql

import (
	"fmt"

	"gorm.io/gorm"

	"appointments/adapters/postgresql/models"
	"appointments/domain/entities"
	"appointments/domain/ports"
)

// Garantía de implementación en tiempo de compilación.
var _ ports.CitaRepository = (*CitaRepo)(nil)

type CitaRepo struct {
	db *gorm.DB
}

func NuevoCitaRepo(db *gorm.DB) *CitaRepo {
	return &CitaRepo{db: db}
}

func (r *CitaRepo) Crear(cita entities.Cita) (entities.Cita, error) {
	model := models.CitaModelFromEntity(cita)

	// Preload de Paciente y Doctor para retornar la cita completa.
	err := r.db.Create(&model).Error
	if err != nil {
		return entities.Cita{}, fmt.Errorf("error creando cita: %w", err)
	}

	return r.buscarModelPorID(model.ID)
}

func (r *CitaRepo) BuscarPorPaciente(pacienteID int) ([]entities.Cita, error) {
	var modelList []models.CitaModel

	err := r.db.
		Preload("Paciente").
		Preload("Doctor").
		Where("paciente_id = ?", pacienteID).
		Find(&modelList).Error
	if err != nil {
		return nil, fmt.Errorf("error consultando citas del paciente %d: %w", pacienteID, err)
	}

	return toCitaEntities(modelList), nil
}

func (r *CitaRepo) BuscarPorDoctor(doctorID int) ([]entities.Cita, error) {
	var modelList []models.CitaModel

	err := r.db.
		Preload("Paciente").
		Preload("Doctor").
		Where("doctor_id = ?", doctorID).
		Find(&modelList).Error
	if err != nil {
		return nil, fmt.Errorf("error consultando citas del doctor %d: %w", doctorID, err)
	}

	return toCitaEntities(modelList), nil
}

func (r *CitaRepo) ActualizarEstado(id int, estado entities.EstadoCita) (entities.Cita, error) {
	err := r.db.
		Model(&models.CitaModel{}).
		Where("id = ?", id).
		Update("estado", estado).Error
	if err != nil {
		return entities.Cita{}, fmt.Errorf("error actualizando estado de cita %d: %w", id, err)
	}

	return r.buscarModelPorID(id)
}

// buscarModelPorID es un helper interno que carga una cita completa con sus relaciones.
func (r *CitaRepo) buscarModelPorID(id int) (entities.Cita, error) {
	var model models.CitaModel

	err := r.db.
		Preload("Paciente").
		Preload("Doctor").
		First(&model, id).Error
	if err != nil {
		return entities.Cita{}, fmt.Errorf("cita %d no encontrada: %w", id, err)
	}

	return model.ToEntity(), nil
}

// toCitaEntities convierte una lista de CitaModel a entidades de dominio.
func toCitaEntities(modelList []models.CitaModel) []entities.Cita {
	citas := make([]entities.Cita, 0, len(modelList))
	for _, m := range modelList {
		citas = append(citas, m.ToEntity())
	}
	return citas
}
