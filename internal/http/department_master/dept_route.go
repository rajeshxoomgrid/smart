package departmentmaster

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rajeshbond/smart/internal/auth"
)

func (m *DeptModeule) Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/test1", func(w http.ResponseWriter, req *http.Request) {
		log.Println("========== Department TEST HIT ==========")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Department route Test Ok"))
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.Verifier(m.tokenAuth))
		r.Use(auth.Authenticator(m.tokenAuth))
		r.Use(auth.UserContextInjector)
		r.Post("/createdept", m.DeptHandler.Create)

	})

	if err := chi.Walk(r, func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler,
	) error {
		log.Printf("Department route ROUTE: %-6s %s", method, route)
		return nil
	}); err != nil {
		log.Printf("Department route error: %v", err)
	}

	return r
}
