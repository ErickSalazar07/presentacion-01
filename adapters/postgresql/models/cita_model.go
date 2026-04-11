package models

import (
	"time"

	"appointments/domain/entities"
)

type CitaModel struct {
	ID        int                  `gorm:"primaryKey;autoIncrement"`
	Estado    entities.EstadoCita  `gorm:"not null;default:'pendiente'"`
	Motivo    string               `gorm:"not null"`
	FechaHora time.Time            `gorm:"not null"`
	CreatedAt time.Time            `gorm:"autoCreateTime"`

	PacienteID int `gorm:"not null;index"`
	DoctorID   int `gorm:"not null;index"`

	Paciente PacienteModel `gorm:"foreignKey:PacienteID"`
	Doctor   DoctorModel   `gorm:"foreignKey:DoctorID"`
}

func (CitaModel) TableName() string {
	return "citas"
}

// ToEntity convierte un CitaModel completo (Paciente y Doctor ya cargados por GORM)
// a la entidad de dominio.
func (m CitaModel) ToEntity() entities.Cita {
	return entities.Cita{
		ID:               m.ID,
		Estado:           m.Estado,
		Motivo:           m.Motivo,
		FechaHora:        m.FechaHora,
		CreatedAt:        m.CreatedAt,
		PacienteAsignado: m.Paciente.ToEntity(),
		DoctorAsignado:   m.Doctor.ToEntity(),
	}
}

// ToEntitySinRelaciones convierte solo los campos escalares.
// Se usa al mapear Paciente.Citas o Doctor.Citas para evitar
// referencias circulares.
func (m CitaModel) ToEntitySinRelaciones() entities.Cita {
	return entities.Cita{
		ID:        m.ID,
		Estado:    m.Estado,
		Motivo:    m.Motivo,
		FechaHora: m.FechaHora,
		CreatedAt: m.CreatedAt,
	}
}

func CitaModelFromEntity(e entities.Cita) CitaModel {
	return CitaModel{
		ID:         e.ID,
		Estado:     e.Estado,
		Motivo:     e.Motivo,
		FechaHora:  e.FechaHora,
		CreatedAt:  e.CreatedAt,
		PacienteID: e.PacienteAsignado.ID,
		DoctorID:   e.DoctorAsignado.ID,
	}
}
