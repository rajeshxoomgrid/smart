package devicedata

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rajeshbond/smart/internal/auth"
)

func (m *Module) Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/prodlog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Production Log Test Ok"))
	})

	// Provate router

	r.Group(func(r chi.Router) {
		r.Use(auth.Verifier(m.tokenAuth))
		r.Use(auth.Authenticator(m.tokenAuth))
		r.Use(auth.UserContextInjector)
		r.Post("/assembly", m.Handler.GetProductionLogByTenantIDAndDeviceID)
		r.Post("/assemblyshift", m.Handler.GetShiftProductionCount)
		r.Post("/assemblyhr", m.Handler.SavePreviousHour)
		r.Post("/yesterday", m.Handler.GetYesterdayProdcution)
		r.Post("/newassemblygr", m.Handler.NewPreviousHourly)

	})

	return r

}
