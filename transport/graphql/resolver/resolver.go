package resolver

import "appointments/application"

// Resolver es el punto de entrada de todos los resolvers GraphQL.
// Recibe los application services por inyección de dependencias —
// nunca instancia repos ni conexiones directamente.
type Resolver struct {
	PacienteService *application.PacienteService
	DoctorService   *application.DoctorService
	CitaService     *application.CitaService
}
