package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

func (h *Handler) GetImmLimitDataDeviceWiseHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
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

	//----------------------------------------------------------
	// Decode POST request body
	//----------------------------------------------------------

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

	//----------------------------------------------------------
	// Validate request
	//----------------------------------------------------------

	if req.DeviceID == "" {
		response.Error(
			w,
			http.StatusBadRequest,
			"device_id is required",
		)
		return
	}

	if req.MachineID == "" {
		response.Error(
			w,
			http.StatusBadRequest,
			"machine_id is required",
		)
		return
	}

	//----------------------------------------------------------
	// Service
	//----------------------------------------------------------

	result, err := h.Service.GetDataLimitDeviceWise(
		ctx,
		&req,
		claims.TenantID,
	)
	if err != nil {
		log.Printf("Get IMM production data failed: %v", err)

		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	//----------------------------------------------------------
	// Response
	//----------------------------------------------------------

	response.JSON(w, http.StatusOK, result)
}

// func (h *Handler) GetImmLimitDataDeviceWiseHandler(w http.ResponseWriter, r *http.Request) {

// 	ctx := r.Context()

// 	//----------------------------------------------------------
// 	// Authentication
// 	//----------------------------------------------------------

// 	claims, err := auth.MustUserClaims(ctx)
// 	if err != nil {

// 		response.Error(
// 			w,
// 			http.StatusUnauthorized,
// 			auth.UnAuthorised,
// 		)

// 		return
// 	}

// 	//----------------------------------------------------------
// 	// Authorization
// 	//----------------------------------------------------------
// 	if !permission.ProductionLogViewwer(claims.Role) {

// 		response.Error(
// 			w,
// 			http.StatusForbidden,
// 			auth.PermissionDenied.Error(),
// 		)

// 		return
// 	}

// 	var req dto.ImmDataFetchRequest

// 	decoder := json.NewDecoder(r.Body)
// 	decoder.DisallowUnknownFields()

// 	if err := decoder.Decode(&req); err != nil {

// 		response.Error(
// 			w,
// 			http.StatusBadRequest,
// 			err.Error(),
// 		)

// 		return
// 	}

// 	result, err := h.Service.GetDataLimitDeviceWiseService(ctx, &req, claims.TenantID)

// 	if err != nil {
// 		log.Printf("Get IMM production data failed: %v", err)

// 		response.Error(
// 			w,
// 			http.StatusInternalServerError,
// 			err.Error(),
// 		)

// 		return
// 	}

// 	response.JSON(w, http.StatusOK, result)

// }
