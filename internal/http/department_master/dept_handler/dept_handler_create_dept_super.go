package depthandler

import (
	"encoding/json"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	deptdto "github.com/rajeshbond/smart/internal/http/department_master/dept_dto"
)

func (h *DeptHandler) CreateDeptHandlerSuper(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract jwt claims from context

	claims, err := auth.MustUserClaims(ctx)

	if err != nil {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}
	if !permission.CanCreateDeprtment(claims.Role) {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	// Parse Request

	var req deptdto.CreateDepartmentRequestSuper

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, response.InvalidRequest)
		return
	}

	resp, err := h.DeptService.CreateDepartmentBySuper(ctx, claims.UserID, req.TenantID, req.Department)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, resp)

}
