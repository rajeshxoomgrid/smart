package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rajeshbond/smart/internal/auth/permission"
)

func IsSuper(role string) bool {
	return role == permission.RoleXoomAdmin || role == permission.RoleSuperAdmin
}

func ValidateTenantAccess(role, claimsEmpID, reqEmpID string) error {
	fmt.Println("================================================================================================")
	fmt.Println("Inside ValidateTenantAccess with role:", role, "claimsEmpID:", claimsEmpID, "reqEmpID:", reqEmpID)

	role = strings.ToLower(strings.TrimSpace(role))

	// ✅ Superadmin & Admin → full access
	if role == "superadmin" || role == "admin" {
		return nil
	}

	// ✅ Tenant Admin → restricted
	if role == "tenantadmin" || role == "tenantowner" {

		claimsTcode, err := Tcode(claimsEmpID)
		if err != nil {
			return errors.New("invalid claims employee id")
		}

		reqTcode, err := Tcode(reqEmpID)
		if err != nil {
			return errors.New("invalid request employee id")
		}

		if claimsTcode != reqTcode {
			return errors.New("tenant mismatch: not allowed for other Tenant")
		}

		return nil
	}

	// ❌ Other roles
	return errors.New("insufficient permissions")
}

func Tcode(employee_id string) (string, error) {
	parts := strings.SplitN(employee_id, "@", 2)

	if len(parts) < 2 || parts[1] == "" {
		return "", errors.New("invalid employee id format")
	}

	return strings.ToLower(strings.TrimSpace(parts[1])), nil
}

// Function Validate Tenant ID with Tenant code

func ValidateTenantAccesswithTenantCode(role string, claimsTenantID, reqTenantID int64) error {

	switch role {

	case permission.RoleSuperAdmin, permission.RoleXoomAdmin:
		// Full Access
		return nil

	case permission.RoleTenantAdmin, permission.RoleTenantOwner:
		// Restricted to own tenant
		if claimsTenantID != reqTenantID {
			fmt.Println("Rajesh failed ")
			return ErrTenantMismatch
		}

		return nil // ✅ IMPORTANT FIX

	default:
		return ErrUnauthorized
	}
}

func TenantRoleCheck(role string) error {
	switch role {
	case "superadmin", "admin", "tenantadmin", "tenantowner":
		return fmt.Errorf("not allowed to create admin role")
	default:
		return nil
	}
}

// Define Roles

type Role string

// var superRoles = map[Role]struct{}{
// 	permission.RoleSuperAdmin: {},
// 	permission.RoleXoomAdmin:  {},
// }

// func IsSuper(role string) bool {
// 	_, exists := superRoles[Role(strings.ToLower(role))]
// 	return exists
// }

func IsTenatAdminRole(reqRole string) bool {
	fmt.Print("Inside isTenanat Admin ", reqRole)
	allowedRoles := map[string]struct{}{
		// "tenantadmin": {},
		"tenantadmin": {},
		"tenantowner": {},
	}

	_, ok := allowedRoles[reqRole]
	fmt.Println("Bool value", ok)
	return ok
}

// Check is Xoomgrid superAdmin or Admin

func IsXoodGridAdmin(reqRole string) bool {

	allowedRoles := map[string]struct{}{
		"superadmin": {},
		"xoomadmin":  {},
	}

	_, ok := allowedRoles[reqRole]

	return ok
}

func IsXoomGridUser(reqRole string) bool {
	allowedRoles := map[string]struct{}{
		"superadmin": {},
		"xoomadmin":  {},
		"xoomuser":   {},
	}

	_, ok := allowedRoles[reqRole]

	return ok
}

func IsDistributorAdmin(reqRole string) bool {
	allowedRoles := map[string]struct{}{
		"distributor":      {},
		"distributoradmin": {},
	}

	_, ok := allowedRoles[reqRole]

	return ok
}

func IsDistributorRole(reqRole string) bool {
	allowedRoles := map[string]struct{}{
		"distributor":        {},
		"distributoradmin":   {},
		"distributorservice": {},
		"distributoruser":    {},
	}

	_, ok := allowedRoles[reqRole]

	return ok
}
