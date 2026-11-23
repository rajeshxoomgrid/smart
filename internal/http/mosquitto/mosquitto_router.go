package mosquitto

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rajeshbond/smart/internal/auth"
)

func (m *MosquittoModule) Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/test1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mosquitto Route is Ok "))
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.Verifier(m.tokenAuth))
		r.Use(auth.Authenticator(m.tokenAuth))
		r.Use(auth.UserContextInjector)
		// r.Post("/createmold", m.Handler.BulkCreate)

	})

	return r
}
