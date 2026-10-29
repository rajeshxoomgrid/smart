package parthandler

import (
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/common/response"
)

func (h *PartHandler) PartCreateHanlder(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	//----------------------------------------------------------
	// Authentication
	//----------------------------------------------------------

	claims, err := auth.MustUserClaims(ctx)

	// if permission.HasPermission()

	if err != nil {

		response.Error(
			w,
			http.StatusUnauthorized,
			auth.UnAuthorised,
		)

		return
	}

	response.JSON(w, http.StatusCreated, claims)

}
