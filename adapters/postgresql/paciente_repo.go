package postgresql

import (
	"fmt"

	"gorm.io/gorm"

	"appointments/adapters/postgresql/models"
	"appointments/domain/entities"
	"appointments/domain/ports"
)

// Garantía de implementación en tiempo de compilación.
// Si PacienteRepo no implementa ports.PacienteRepository, el compilador falla.
var _ ports.PacienteRepository = (*PacienteRepo)(nil)

type PacienteRepo struct {
	db *gorm.DB
}

func NuevoPacienteRepo(db *gorm.DB) *PacienteRepo {
	return &PacienteRepo{db: db}
}

func (r *PacienteRepo) Crear(paciente entities.Paciente) (entities.Paciente, error) {
	model := models.PacienteModelFromEntity(paciente)

	if err := r.db.Create(&model).Error; err != nil {
		return entities.Paciente{}, fmt.Errorf("error creando paciente: %w", err)
	}

	return model.ToEntity(), nil
}

func (r *PacienteRepo) BuscarPorID(id int) (entities.Paciente, error) {
	var model models.PacienteModel

	err := r.db.Preload("Citas").First(&model, id).Error
	if err != nil {
		return entities.Paciente{}, fmt.Errorf("paciente %d no encontrado: %w", id, err)
	}

	return model.ToEntity(), nil
}

func (r *PacienteRepo) BuscarTodos() ([]entities.Paciente, error) {
	var modelList []models.PacienteModel

	if err := r.db.Find(&modelList).Error; err != nil {
		return nil, fmt.Errorf("error consultando pacientes: %w", err)
	}

	pacientes := make([]entities.Paciente, 0, len(modelList))
	for _, m := range modelList {
		pacientes = append(pacientes, m.ToEntity())
	}

	return pacientes, nil
}
