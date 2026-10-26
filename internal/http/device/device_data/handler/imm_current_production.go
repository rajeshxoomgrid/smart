package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (h *Handler) GetImmShiftProductionCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//----------------------------------------------------------
	// Authentication
	//----------------------------------------------------------

	claims, err := auth.MustUserClaims(ctx)
	if err != nil {

		response.Error(
			w,
			http.StatusUnauthorized,
			auth.UnAuthorised,
		)

		return
	}

	//----------------------------------------------------------
	// Authorization
	//----------------------------------------------------------
	if !permission.ProductionLogViewwer(claims.Role) {

		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)

		return
	}

	var req dto.ImmDataFetchRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {

		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	if req.Limit <= 0 {
		req.Limit = 1
	}

	//----------------------------------------------------------
	// Service
	//----------------------------------------------------------

	resp, err := h.Service.GetImmLatestProductionCount(ctx, req, claims.TenantID)

	if err != nil {

		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	//----------------------------------------------------------
	// Success
	//----------------------------------------------------------

	response.JSON(
		w,
		http.StatusOK,
		resp,
	)

	// item,err:= h.
}
