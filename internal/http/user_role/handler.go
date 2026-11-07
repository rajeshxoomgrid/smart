package userrole

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	tenantIDstr := chi.URLParam(r, "tenantID")
	// Convert tenantID string to int64
	tenantID, err := strconv.ParseInt(tenantIDstr, 10, 64)
	roleId, err := h.service.GetRoleIDByName(ctx, roleIDStr, tenantID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, roleId)

}

func (h *Handler) GetTheUserNameByID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Authentication
	claims, err := auth.MustUserClaims(ctx)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	// Get URL parameters
	roleIDStr := chi.URLParam(r, "roleID")
	tenantIDStr := chi.URLParam(r, "tenantID")

	// Convert roleID to int64
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil || roleID <= 0 {
		response.Error(w, http.StatusBadRequest, "Invalid role ID")
		return
	}

	// Convert tenantID to int64
	tenantID, err := strconv.ParseInt(tenantIDStr, 10, 64)
	if err != nil || tenantID <= 0 {
		response.Error(w, http.StatusBadRequest, "Invalid tenant ID")
		return
	}

	// Fetch role name
	roleName, err := h.service.store.GetTheUserNameByID(ctx, roleID, claims.TenantID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, roleName)

}

func (h *Handler) DeleteUserRole(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Authentication
	claims, ok := auth.GetUserClaimsFromContext(ctx)
	if !ok {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	// Check admin role
	role := claims.Role

	if !auth.IsXoodGridAdmin(role) {
		response.Error(w, http.StatusForbidden, response.NotAuthorized)
		return
	}

	// Decode request body
	var req DeleteRole

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate IDs
	if req.RoleID <= 0 {
		response.Error(w, http.StatusBadRequest, "Invalid role ID")
		return
	}

	if req.TenantID <= 0 {
		response.Error(w, http.StatusBadRequest, "Invalid tenant ID")
		return
	}

	// Service call
	deleted, err := h.service.DeleteUserrole(
		ctx,
		req.RoleID,
		req.TenantID,
	)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !deleted {
		response.Error(w, http.StatusNotFound, "User role not found")
		return
	}

	response.JSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "User role deleted successfully",
	})
}
