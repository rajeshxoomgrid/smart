package permissionhandler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (h *PermissionHandler) Update(
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
	// Decode request
	// --------------------------------------------------

	var req permissiondto.UpdatePermissionRequest

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	// --------------------------------------------------
	// updated_by comes from JWT
	// --------------------------------------------------

	req.UpdatedBy = &claims.UserID

	// --------------------------------------------------
	// Service
	// --------------------------------------------------

	permission, err := h.PermissionService.Update(
		ctx,
		id,
		&req,
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
	// Response
	// --------------------------------------------------

	response.JSON(
		w,
		http.StatusOK,
		permissiondto.PermissionResponse{
			Data: permission,
		},
	)
}
