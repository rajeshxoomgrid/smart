package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
)

// ============================================================
// SAVE PREVIOUS HOUR
// ============================================================
//
// API:
//
// POST /proddata/assemblyhr
//
// The handler does:
//
// 1. Authentication
// 2. Authorization
// 3. Decode request
// 4. Call service
// 5. Return response
//
// Hourly calculation itself is done in the service/store.
//
// ============================================================

func (h *Handler) SavePreviousHour(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	// ========================================================
	// AUTHENTICATION
	// ========================================================

	claims, err := auth.MustUserClaims(ctx)

	fmt.Println("<==========================================>")
	fmt.Println("User Tenant ID--->", claims.TenantID)
	fmt.Println("<==========================================>")

	if err != nil {

		response.Error(
			w,
			http.StatusUnauthorized,
			auth.UnAuthorised,
		)

		return
	}

	// ========================================================
	// AUTHORIZATION
	// ========================================================

	if !permission.ProductionLogViewwer(claims.Role) {

		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)

		return
	}

	// ========================================================
	// DECODE REQUEST
	// ========================================================

	var req dto.SaveHourlyProductionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		response.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)

		return
	}

	// ========================================================
	// BASIC REQUEST LOG
	// ========================================================

	log.Printf(
		"========== SAVE PREVIOUS HOUR REQUEST ==========\n"+
			"TenantID : %s\n"+
			"DeviceID : %s\n"+
			"MachineID: %s\n"+
			"Station  : %s\n"+
			"Variant  : %v\n"+
			"=================================================",
		req.TenantID,
		req.DeviceID,
		req.MachineID,
		req.Station,
		req.Variant,
	)

	// ========================================================
	// SERVICE
	// ========================================================

	item, err := h.Service.SavePreviousHour(
		ctx,
		req,
	)

	if err != nil {

		log.Printf(
			"❌ SavePreviousHour failed: %v",
			err,
		)

		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)

		return
	}

	// ========================================================
	// SUCCESS
	// ========================================================

	response.JSON(
		w,
		http.StatusOK,
		item,
	)
}
