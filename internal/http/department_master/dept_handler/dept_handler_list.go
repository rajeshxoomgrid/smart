package depthandler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/common/response"
	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

func (h *DeptHandler) ListOfDepartHanlderTenantWise(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Extract JWt claims from context

	claims, err := auth.MustUserClaims(ctx)

	if err != nil {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	// check user role

	if !auth.IsSuper(claims.Role) {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	var req deptdto.ListDepartmentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil && err != io.EOF {
		response.Error(w, http.StatusBadRequest, response.InvalidRequest)
		return
	}

	var tenantID int64

	if req.TenantID != nil && *req.TenantID > 0 {
		tenantID = *req.TenantID
	} else {
		tenantID = claims.TenantID
	}

	if tenantID <= 0 {
		response.Error(w, http.StatusBadRequest, "Invaild Tenant ID")
		return
	}

	resp, err := h.DeptService.ListAllDepartemntTenantService(ctx, tenantID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, response.InternalServerError)
		return
	}

	// Success response
	response.JSON(w, http.StatusOK, resp)

}
