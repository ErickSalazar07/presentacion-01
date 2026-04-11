package models

import (
	"time"

	"appointments/domain/entities"
)

type PacienteModel struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	Nombre          string    `gorm:"not null"`
	Email           string    `gorm:"uniqueIndex;not null"`
	Telefono        string    `gorm:"uniqueIndex;not null"`
	FechaNacimiento time.Time `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`

	Citas []CitaModel `gorm:"foreignKey:PacienteID"`
}

func (PacienteModel) TableName() string {
	return "pacientes"
}

func (m PacienteModel) ToEntity() entities.Paciente {
	citas := make([]entities.Cita, 0, len(m.Citas))
	for _, c := range m.Citas {
		citas = append(citas, c.ToEntitySinRelaciones())
	}
	return entities.Paciente{
		ID:              m.ID,
		Nombre:          m.Nombre,
		Email:           m.Email,
		Telefono:        m.Telefono,
		FechaNacimiento: m.FechaNacimiento,
		CreatedAt:       m.CreatedAt,
		Citas:           citas,
	}
}

func PacienteModelFromEntity(e entities.Paciente) PacienteModel {
	return PacienteModel{
		ID:              e.ID,
		Nombre:          e.Nombre,
		Email:           e.Email,
		Telefono:        e.Telefono,
		FechaNacimiento: e.FechaNacimiento,
		CreatedAt:       e.CreatedAt,
	}
}
