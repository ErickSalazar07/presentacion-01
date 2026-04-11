package models

import (
	"time"

	"appointments/domain/entities"
)

type DoctorModel struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Nombre       string    `gorm:"not null"`
	Especialidad string    `gorm:"not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	Citas []CitaModel `gorm:"foreignKey:DoctorID"`
}

func (DoctorModel) TableName() string {
	return "doctores"
}

func (m DoctorModel) ToEntity() entities.Doctor {
	citas := make([]entities.Cita, 0, len(m.Citas))
	for _, c := range m.Citas {
		citas = append(citas, c.ToEntitySinRelaciones())
	}
	return entities.Doctor{
		ID:           m.ID,
		Nombre:       m.Nombre,
		Especialidad: m.Especialidad,
		Email:        m.Email,
		CreatedAt:    m.CreatedAt,
		Citas:        citas,
	}
}

func DoctorModelFromEntity(e entities.Doctor) DoctorModel {
	return DoctorModel{
		ID:           e.ID,
		Nombre:       e.Nombre,
		Especialidad: e.Especialidad,
		Email:        e.Email,
		CreatedAt:    e.CreatedAt,
	}
}
