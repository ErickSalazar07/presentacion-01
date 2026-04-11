package main

import (
	"log"
	"net/http"

	"appointments/application"
	"appointments/infrastructure/db"
	"appointments/adapters/postgresql"
	"appointments/transport/graphql/generated"
	"appointments/transport/graphql/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// 1. Conexión a DB
	database, err := db.Nueva()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Migraciones
	if err := db.Migrar(database); err != nil {
		log.Fatal(err)
	}

	// 3. Repositorios (Adapters)
	pacienteRepo := postgresql.NuevoPacienteRepo(database)
	doctorRepo := postgresql.NuevoDoctorRepo(database)
	citaRepo := postgresql.NuevoCitaRepo(database)

	// 4. Application services
	pacienteService := application.NuevoPacienteService(pacienteRepo)
	doctorService := application.NuevoDoctorService(doctorRepo)
	citaService := application.NuevoCitaService(citaRepo, pacienteRepo, doctorRepo)

	// 5. Resolver raíz (inyección de dependencias)
	resolverRoot := &resolver.Resolver{
		PacienteService: pacienteService,
		DoctorService:   doctorService,
		CitaService:     citaService,
	}

	// 6. GraphQL server
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolverRoot},
		),
	)

	// 7. Endpoints HTTP
	http.Handle("/graphql", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	log.Println("🚀 servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
