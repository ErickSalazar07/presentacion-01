package postgresql

import (
	"fmt"

	"gorm.io/gorm"

	"appointments/adapters/postgresql/models"
	"appointments/domain/entities"
	"appointments/domain/ports"
)

// Garantía de implementación en tiempo de compilación.
var _ ports.DoctorRepository = (*DoctorRepo)(nil)

type DoctorRepo struct {
	db *gorm.DB
}

func NuevoDoctorRepo(db *gorm.DB) *DoctorRepo {
	return &DoctorRepo{db: db}
}

func (r *DoctorRepo) Crear(doctor entities.Doctor) (entities.Doctor, error) {
	model := models.DoctorModelFromEntity(doctor)

	if err := r.db.Create(&model).Error; err != nil {
		return entities.Doctor{}, fmt.Errorf("error creando doctor: %w", err)
	}

	return model.ToEntity(), nil
}

func (r *DoctorRepo) BuscarPorID(id int) (entities.Doctor, error) {
	var model models.DoctorModel

	err := r.db.Preload("Citas").First(&model, id).Error
	if err != nil {
		return entities.Doctor{}, fmt.Errorf("doctor %d no encontrado: %w", id, err)
	}

	return model.ToEntity(), nil
}

func (r *DoctorRepo) BuscarTodos() ([]entities.Doctor, error) {
	var modelList []models.DoctorModel

	if err := r.db.Find(&modelList).Error; err != nil {
		return nil, fmt.Errorf("error consultando doctores: %w", err)
	}

	doctores := make([]entities.Doctor, 0, len(modelList))
	for _, m := range modelList {
		doctores = append(doctores, m.ToEntity())
	}

	return doctores, nil
}
