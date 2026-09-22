package internalsetup

import (
	"database/sql"

	deptservice "github.com/rajeshbond/smart/internal/http/department_master/dept_service"
	deptstore "github.com/rajeshbond/smart/internal/http/department_master/dept_store"
	"github.com/rajeshbond/smart/internal/http/tenant"
	userrole "github.com/rajeshbond/smart/internal/http/user_role"
	"github.com/rajeshbond/smart/internal/http/users"
)

type Module struct {
	Service *Service
}

func NewModule(db *sql.DB) *Module {

	// Initialize stores
	roleStore := userrole.NewStore(db)
	tenantStore := tenant.NewStore(db)
	deptStore := deptstore.NewDeptStore(db)
	userStore := users.NewStore(db)

	// Initialize services
	roleService := userrole.NewService(roleStore)
	tenantService := tenant.NewService(tenantStore)
	deptService := deptservice.NewDeptService(deptStore)
	userService := users.NewService(userStore, users.RoleProvider(roleService), tenantStore, deptService)

	// Initialize setup service
	setupService := NewService(db, tenantService, roleService, userService)

	return &Module{
		Service: setupService,
	}
}
