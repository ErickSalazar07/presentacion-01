package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"appointments/adapters/postgresql"
	"appointments/application"
	"appointments/infrastructure/db"
	"appointments/transport/graphql/generated"
	"appointments/transport/graphql/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}

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
	http.Handle("/graphql", enableCORS(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("servidor corriendo en http://%s:8081\n", getLocalIP())
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
