package permission

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rajeshbond/smart/internal/auth"
)

func (m *PermissionModule) Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/test1", func(w http.ResponseWriter, req *http.Request) {
		log.Println("========== Permission TEST HIT ==========")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Permission master Test Ok"))
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.Verifier(m.tokenAuth))
		r.Use(auth.Authenticator(m.tokenAuth))
		r.Use(auth.UserContextInjector)
		// All Post Request
		r.Post("/createpermission", m.PermissionHandler.Create) //To Create the Permission
		// All the Put
		r.Put("/updatepermission/{id}", m.PermissionHandler.UpdatePermissionRequest)
		// All Get Request
		r.Get("/{id}", m.PermissionHandler.GetByID)

	})

	if err := chi.Walk(r, func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler,
	) error {
		log.Printf("Mold MODULE ROUTE: %-6s %s", method, route)
		return nil
	}); err != nil {
		log.Printf("Mold module route error: %v", err)
	}

	return r
}
