package depthandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
)

// GetByID handles GET /dept/{id}
func (h *DeptHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Get ID from URL
	idStr := chi.URLParam(r, "id")

	// Get authenticated user claims
	claims, err := auth.MustUserClaims(ctx)
	if err != nil {
		response.Error(
			w,
			http.StatusUnauthorized,
			response.NotAuthorized,
		)
		return
	}

	// Permission check
	if !permission.CanCreateDeprtment(claims.Role) {
		response.Error(
			w,
			http.StatusUnauthorized,
			response.NotAuthorized,
		)
		return
	}

	// Convert ID to int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		response.Error(
			w,
			http.StatusBadRequest,
			"invalid department id",
		)
		return
	}

	// Service call
	dept, err := h.DeptService.GetByID(ctx, id)
	if err != nil {
		response.Error(
			w,
			http.StatusNotFound,
			err.Error(),
		)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(dept); err != nil {
		return
	}
}

// GetByName handles GET /dept/name/{name}
func (h *DeptHandler) GetByName(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Get department name from URL
	name := chi.URLParam(r, "name")

	// Get authenticated user claims
	claims, err := auth.MustUserClaims(ctx)
	if err != nil {
		response.Error(
			w,
			http.StatusUnauthorized,
			response.NotAuthorized,
		)
		return
	}

	// Permission check
	if !permission.CanCreateDeprtment(claims.Role) {
		response.Error(
			w,
			http.StatusUnauthorized,
			response.NotAuthorized,
		)
		return
	}

	// Validate name
	if name == "" {
		response.Error(
			w,
			http.StatusBadRequest,
			"department name is required",
		)
		return
	}

	// Service call
	dept, err := h.DeptService.GetByName(ctx, name)
	if err != nil {
		response.Error(
			w,
			http.StatusNotFound,
			err.Error(),
		)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(dept); err != nil {
		return
	}
}
