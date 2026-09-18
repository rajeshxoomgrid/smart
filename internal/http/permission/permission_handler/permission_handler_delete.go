package permissionhandler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
)

func (h *PermissionHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	// --------------------------------------------------
	// Get ID
	// --------------------------------------------------

	idString := chi.URLParam(
		r,
		"id",
	)

	id, err := strconv.ParseInt(
		idString,
		10,
		64,
	)

	if err != nil || id <= 0 {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid permission id",
		)
		return
	}

	// --------------------------------------------------
	// Get JWT claims
	// --------------------------------------------------

	claims, err := auth.MustUserClaims(ctx)

	if err != nil {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	if claims == nil {
		response.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	//----------------------------------------------------------------------
	// Authorization
	//----------------------------------------------------------------------

	if !permission.IsXoomUser(claims.Role) {
		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)
		return
	}

	// --------------------------------------------------
	// deleted_by comes from JWT
	// --------------------------------------------------

	deletedBy := claims.UserID

	// --------------------------------------------------
	// Service
	// --------------------------------------------------

	err = h.PermissionService.Delete(
		ctx,
		id,
		&deletedBy,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			response.Error(
				w,
				http.StatusNotFound,
				"permission not found",
			)
			return
		}

		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	// --------------------------------------------------
	// 204 No Content
	// --------------------------------------------------

	w.WriteHeader(http.StatusNoContent)
}
