package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"appointments/adapters/postgresql/models"
)

// Crea y retorna una conexión a PostgreSQL usando GORM.
// Lee las variables de entorno para construir el DSN.
func Nueva() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "default"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USERNAME", "default"),
		getEnv("DB_PASSWORD", "default"),
		getEnv("DB_NAME", "default"),
		getEnv("DB_SSLMODE", "disable"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	log.Println("Conexión a PostgreSQL establecida")
	return db, nil
}

// Migrar ejecuta las migraciones automáticas para crear/actualizar
// las tablas según los modelos GORM definidos en adapters/postgresql/models.
func Migrar(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.PacienteModel{},
		&models.DoctorModel{},
		&models.CitaModel{},
	)
	if err != nil {
		return fmt.Errorf("error ejecutando migraciones: %w", err)
	}

	log.Println("Migraciones ejecutadas correctamente")
	return nil
}

// getEnv retorna el valor de una variable de entorno,
// o el valor por defecto si no está definida.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
