package permissionhandler

import (
	"net/http"
	"strconv"

	"github.com/rajeshbond/smart/internal/auth"
	"github.com/rajeshbond/smart/internal/auth/permission"
	"github.com/rajeshbond/smart/internal/common/response"
	permissiondto "github.com/rajeshbond/smart/internal/http/permission/permission_dto"
)

func (h *PermissionHandler) List(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()

	claims, err := auth.MustUserClaims(ctx)

	if err != nil {
		response.Error(w, http.StatusUnauthorized, response.NotAuthorized)
		return
	}

	if claims == nil {
		response.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	//----------------------------------------------------------------------
	// Authorization
	//----------------------------------------------------------------------

	if !permission.IsXoomUser(claims.Role) {
		response.Error(
			w,
			http.StatusForbidden,
			auth.PermissionDenied.Error(),
		)
		return
	}

	query := r.URL.Query()

	// --------------------------------------------------
	// Default pagination
	// --------------------------------------------------

	filter := &permissiondto.PermissionFilter{
		Page:      1,
		PageSize:  20,
		Search:    query.Get("search"),
		SortBy:    query.Get("sort_by"),
		SortOrder: query.Get("sort_order"),
	}

	// --------------------------------------------------
	// Page
	// --------------------------------------------------

	if pageString := query.Get("page"); pageString != "" {

		page, err := strconv.Atoi(pageString)

		if err != nil || page <= 0 {
			response.Error(
				w,
				http.StatusBadRequest,
				"invalid page",
			)
			return
		}

		filter.Page = page
	}

	// --------------------------------------------------
	// Page size
	// --------------------------------------------------

	if pageSizeString := query.Get("page_size"); pageSizeString != "" {

		pageSize, err := strconv.Atoi(pageSizeString)

		if err != nil || pageSize <= 0 {
			response.Error(
				w,
				http.StatusBadRequest,
				"invalid page_size",
			)
			return
		}

		if pageSize > 100 {
			response.Error(
				w,
				http.StatusBadRequest,
				"page_size cannot be greater than 100",
			)
			return
		}

		filter.PageSize = pageSize
	}

	// --------------------------------------------------
	// Service
	// --------------------------------------------------

	permissions, total, err := h.PermissionService.List(
		ctx,
		filter,
	)

	if err != nil {
		response.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	// --------------------------------------------------
	// Response
	// --------------------------------------------------

	response.JSON(
		w,
		http.StatusOK,
		permissiondto.PermissionListResponse{
			Data:  permissions,
			Total: total,
		},
	)
}
