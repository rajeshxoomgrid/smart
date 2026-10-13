package shift

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/jwtauth/v5"
	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/common/response"
)

type Handler struct {
	Service   *Service
	tokenAuth *jwtauth.JWTAuth
}

func NewHandler(service *Service, tokenAuth *jwtauth.JWTAuth) *Handler {
	return &Handler{
		tokenAuth: tokenAuth,
		Service:   service,
	}
}

// ==========================================
// BULK CREATE (Multi Shift API)
// ==========================================
// ==========================================
// BULK CREATE (Multi Shift API)
// ==========================================
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// 🔐 Authentication
	claims, ok := auth.GetUserClaimsFromContext(ctx)
	if !ok {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	fmt.Println("=============================>", claims)

	if !auth.IsTenatAdminRole(claims.Role) {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	// Decode Request
	var req BulkCreateShiftRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest,
			"Invalid JSON body: "+err.Error())
		return
	}

	if len(req) == 0 {
		response.Error(w,
			http.StatusBadRequest,
			"Request body cannot be empty")
		return
	}

	// Save
	if err := h.Service.CreateBulk(
		ctx,
		claims.UserID,
		req,
	); err != nil {

		switch {

		case strings.Contains(err.Error(), "overlap"):
			response.Error(w,
				http.StatusBadRequest,
				err.Error())

		case strings.Contains(err.Error(), "24 hours"):
			response.Error(w,
				http.StatusBadRequest,
				err.Error())

		case strings.Contains(err.Error(), "already exists"):
			response.Error(w,
				http.StatusBadRequest,
				err.Error())

		case strings.Contains(err.Error(), "invalid time"):
			response.Error(w,
				http.StatusBadRequest,
				err.Error())

		default:
			response.Error(w,
				http.StatusInternalServerError,
				err.Error())
		}

		return
	}

	response.JSON(
		w,
		http.StatusCreated,
		"Shift configuration saved successfully",
	)
}
