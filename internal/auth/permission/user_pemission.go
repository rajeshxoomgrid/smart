package permission

import (
	"strings"
)

func IsXoomUser(roler string) bool {
	switch strings.ToLower(roler) {
	case
		RoleSuperAdmin,
		RoleXoomAdmin,
		RoleXoomUser:
		return true
	}
	return false
}

func CanCreateDevice(role string) bool {
	switch strings.ToLower(role) {
	case
		RoleSuperAdmin,
		RoleXoomAdmin,
		RoleXoomUser:

		return true
	}

	return false
}

func CanCreateDeprtment(role string) bool {
	switch strings.ToLower(role) {
	case
		RoleSuperAdmin,
		RoleAdminTenant:

		return true
	}

	return false
}

func CanDeleteDevice(role string) bool {

	switch strings.ToLower(role) {

	case
		RoleSuperAdmin,
		RoleXoomAdmin:
		return true
	}

	return false
}

func CanViewDevice(role string) bool {

	switch strings.ToLower(role) {

	case
		RoleSuperAdmin,
		RoleXoomAdmin,
		RoleXoomUser,
		RoleDistributorAdmin,
		RoleTenantAdmin:

		return true
	}

	return false
}

func CanUpdateDevice(role string) bool {

	switch strings.ToLower(role) {

	case
		RoleSuperAdmin,
		RoleXoomAdmin,
		RoleXoomUser,
		RoleDistributorAdmin,
		RoleTenantAdmin:

		return true
	}

	return false
}

func HasPermission(
	role string,
	requiredPermission string,
) bool {

	role = strings.ToLower(strings.TrimSpace(role))

	permissions, ok := rolePermissions[role]
	if !ok {
		return false
	}

	_, ok = permissions[requiredPermission]

	return ok
}

func CanListDevice(role string) bool {

	switch strings.ToLower(role) {

	case
		RoleSuperAdmin,
		RoleXoomAdmin,
		RoleXoomUser:
		return true
	}

	return false
}

func ProductionLogViewwer(role string) bool {
	switch strings.ToLower(role) {

	case
		RoleSuperAdmin,
		RoleXoomAdmin,
		RoleXoomUser,
		RoleAdminTenant,
		RoleTenantOperator,
		RoleTenantUser:
		return true
	}

	return false
}
