package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	"github.com/rajeshbond/smart/internal/http/device/device_data/dto"
	"github.com/rajeshbond/smart/internal/http/device/device_data/helper"
)

func (h *Handler) GetYesterdayProdcution(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	// 1. Authentication
	claims, err := auth.MustUserClaims(ctx)
	if err != nil {
		response.Error(
			w,
			http.StatusUnauthorized,
			auth.UnAuthorised,
		)
		return
	}

	// 2. Authorization
	if !permission.ProductionLogViewwer(claims.Role) {
		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)
		return
	}

	// 3. Limit request body to 1 MiB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	// 4. Decode JSON
	var req dto.GetProductionYesterdayRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		var maxBytesErr *http.MaxBytesError

		if errors.As(err, &maxBytesErr) {
			response.Error(
				w,
				http.StatusRequestEntityTooLarge,
				"request body exceeds the 1 MiB limit",
			)
			return
		}

		response.Error(
			w,
			http.StatusBadRequest,
			"invalid JSON request body",
		)
		return
	}

	// Reject a second JSON value or trailing content.
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		response.Error(
			w,
			http.StatusBadRequest,
			"request body must contain exactly one JSON object",
		)
		return
	}

	// 5. Validate required fields.
	// Station and Variant are optional.
	if req.DeviceID == "" {
		response.Error(
			w,
			http.StatusBadRequest,
			"device_id is required",
		)
		return
	}

	if req.Date == "" {
		response.Error(
			w,
			http.StatusBadRequest,
			"date is required in YYYY-MM-DD format",
		)
		return
	}

	// 6. Use the authenticated tenant ID.
	// Never trust a tenant ID supplied by the request body.
	tenantID := strconv.FormatInt(claims.TenantID, 10)

	// 7. Call service.
	item, err := h.Service.GetYesterdatsProduction1(
		ctx,
		req,
		tenantID,
	)
	if err != nil {
		// Log the real error on the server, not in the API response.
		log.Printf(
			"GetYesterdayProdcution failed: tenant_id=%q device_id=%q station=%q variant=%q date=%q error=%v",
			tenantID,
			req.DeviceID,
			req.Station,
			req.Variant,
			req.Date,
			err,
		)

		switch {
		case errors.Is(err, helper.ErrInvalidRequest):
			response.Error(
				w,
				http.StatusBadRequest,
				err.Error(),
			)
			return

		case errors.Is(err, context.Canceled):
			response.Error(
				w,
				http.StatusRequestTimeout,
				"request was cancelled",
			)
			return

		case errors.Is(err, context.DeadlineExceeded):
			response.Error(
				w,
				http.StatusGatewayTimeout,
				"production query timed out",
			)
			return

		default:
			response.Error(
				w,
				http.StatusInternalServerError,
				"failed to retrieve yesterday's production data",
			)
			return
		}
	}

	// 8. Handle missing result.
	if item == nil {
		response.Error(
			w,
			http.StatusNotFound,
			"yesterday's production data not found",
		)
		return
	}

	// 9. Return successful response.
	response.JSON(
		w,
		http.StatusOK,
		item,
	)
}
