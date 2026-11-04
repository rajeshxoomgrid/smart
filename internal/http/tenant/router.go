package tenant

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rajeshbond/smart/cmd/service"
	"github.com/rajeshbond/smart/internal/auth"
)

func (m *Module) Router() chi.Router {
	r := chi.NewRouter()
	tokenAuth := service.GetTokenAuth()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health Ok"))
	})
	// r.Post("/createtenant", m.Handler.CreateTenant)

	// private

	// Protected routes

	r.Group(func(r chi.Router) {
		r.Use(auth.Verifier(tokenAuth))
		r.Use(auth.Authenticator(tokenAuth))
		r.Use(auth.UserContextInjector)
		// Create User Tenant
		r.Post("/createtenant", m.Handler.CreateTenant)
		r.Post("/list", m.Handler.GetAllTenants)
		r.Post("/verifytenant", m.Handler.VerifyTenant)
		r.Patch("/updatetenant/{tenant_code}", m.Handler.UpdateTenant)
		r.Delete("/deletetenant/{tenant_code}", m.Handler.DeleteTenant)

	})
	return r
}
