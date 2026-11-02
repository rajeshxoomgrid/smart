package userrole

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/common/response"
)

type Handler struct {
	service   *Service
	tokenAuth jwtauth.JWTAuth
}

func NewHandler(service *Service, tokenAuth *jwtauth.JWTAuth) *Handler {
	return &Handler{
		service:   service,
		tokenAuth: *tokenAuth,
	}
}

// Create Role by super user

func (h *Handler) CreateUserRoleBySuper(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// 1. Authentication
	claims, ok := auth.GetUserClaimsFromContext(ctx)
	if !ok {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	// 2. Authorization
	if !auth.IsXoodGridAdmin(claims.Role) {
		response.Error(w, http.StatusForbidden, response.NotAuthorized)
		return
	}

	// 3. Decode JSON
	var req CreateRoleBySuper

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// 4. Service call
	createRole, err := h.service.CreateRoleBySuper(ctx, req, claims)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// 5. Success response
	response.JSON(w, http.StatusCreated, createRole)
}
func (h *Handler) CreateUserRole(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Auth
	// Get JWT claims from middleware

	claims, ok := auth.GetUserClaimsFromContext(ctx)

	if !ok {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	if !auth.IsXoodGridAdmin(claims.Role) {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	// Creater User ID

	var dto CreateRole

	// Decode JSON safely

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&dto); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Validate required fierd

	if dto.UserRole == "" {
		response.Error(w, http.StatusBadRequest, "user_role is required")
		return
	}
	userID := claims.UserID

	createRole, err := h.service.Create(ctx, dto, claims.TenantID, userID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, createRole)

}

// Get Role String by Id (passed from params as string)

func (h *Handler) GetUserRoleIDByName(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Auth
	// Get JWT claims from middleware

	_, ok := auth.GetUserClaimsFromContext(ctx)

	if !ok {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	roleIDStr := chi.URLParam(r, "roleIDstr")
	roleId, err := h.service.GetRoleIDByName(ctx, roleIDStr)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, roleId)

}
