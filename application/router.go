package application

//////////////////////////////////////////////
// Packages Mounted
//////////////////////////////////////////////
// 1.User Role
// 2.Tenant

// =======Command for MQTT Clients (esp 32)========
// MQTT Esp 32 reset count
// ================================================
//////////////////////////////////////////////
import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rajeshbond/smart/cmd/service"
	assemblymaster "github.com/rajeshbond/smart/internal/http/assembly_master"
	"github.com/rajeshbond/smart/internal/http/command"
	departmentmaster "github.com/rajeshbond/smart/internal/http/department_master"
	devicedata "github.com/rajeshbond/smart/internal/http/device/device_data"
	devicemaster "github.com/rajeshbond/smart/internal/http/device/device_master"
	internalsetup "github.com/rajeshbond/smart/internal/http/internal_setup"
	machine "github.com/rajeshbond/smart/internal/http/machine_imm"
	"github.com/rajeshbond/smart/internal/http/mold"
	"github.com/rajeshbond/smart/internal/http/permission"
	"github.com/rajeshbond/smart/internal/http/shift"
	shiftslot "github.com/rajeshbond/smart/internal/http/shift_slot"

	"github.com/rajeshbond/smart/internal/http/tenant"
	tenantshifts "github.com/rajeshbond/smart/internal/http/tenant_shifts"
	userrole "github.com/rajeshbond/smart/internal/http/user_role"
	"github.com/rajeshbond/smart/internal/http/users"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(app *App) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // React/Next.js frontend
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{"*"},
		// AllowedHeaders: []string{
		// 	"Accept",
		// 	"Authorization",
		// 	"Content-Type",
		// 	"X-CSRF-Token",
		// },
		// ExposedHeaders:   []string{"Link"},
		// AllowCredentials: true,
		// MaxAge:           300,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health Ok"))
	})

	// Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	tokenAuth := service.GetTokenAuth()

	// ================================================================
	// Mouldes mounting (Starts)
	// ================================================================

	// Internal setup
	internalModule := internalsetup.NewModule(app.DB.SQLDB)
	r.Mount("/internal", internalModule.Router())

	// 1.User Role---->

	userRoleModule := userrole.NewModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/user-role", userRoleModule.Router())

	// 2.Tenant---->

	tenantModule := tenant.NewModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/tenant", tenantModule.Router())

	// Permission router
	permissionModule := permission.NewPermissionModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/permission", permissionModule.Router())

	// Depertment

	departmentModule := departmentmaster.NewDeptModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/dept", departmentModule.Router())

	//3. Users
	usersModule := users.NewModule(app.DB.SQLDB, tokenAuth, userRoleModule.Service, tenantModule.Service, departmentModule.DeptService)
	r.Mount("/users", usersModule.Router())

	// Tenant shift
	tenantShiftMoulde := tenantshifts.NewModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/tenantshift", tenantShiftMoulde.Router())

	// Shift

	shiftModule := shift.NewModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/shift", shiftModule.Router())

	// shift slot

	shiftSlot := shiftslot.NewModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/shiftslot", shiftSlot.Router())

	// Production Log

	productionLogModule := devicedata.NewModule(app.DB.SQLDB, tokenAuth, shiftModule.Store)
	r.Mount("/proddata", productionLogModule.Router())

	// Mold Master

	moldModule := mold.NewMoldModeule(app.DB.SQLDB, tokenAuth)
	r.Mount("/mold", moldModule.Router())

	// Imm Machine

	machineImmModule := machine.NewMachineModule(app.DB.SQLDB, tokenAuth)
	r.Mount("/imachine", machineImmModule.Router())

	// Assebly master

	assemblyMasterModule := assemblymaster.NewModuleAssembly(app.DB.SQLDB, tokenAuth)
	r.Mount("/assemblymaster", assemblyMasterModule.Router())

	// Route Mounts Ends here <---------------

	// 4. Device master
	//==============================================================
	// Device Master
	//==============================================================

	deviceMasterModule := devicemaster.NewModule(app.DB.SQLDB, app.Config, tokenAuth)
	r.Mount("/device-master", deviceMasterModule.Routes())

	// MQTT Commands ----->(MQTT)

	commandModule := command.NewModule(app.MQTTClient)
	r.Mount("/api/v1/assembly-command", commandModule.Router())
	// ================================================================
	// Mouldes mounting (Ends)
	// ================================================================

	return r
}

// Testing Router

// assemblyMasterModule := assemblymaster.NewModuleAssembly(
// 		app.DB.SQLDB,
// 		tokenAuth,
// 	)

// 	assemblyMasterRouter := assemblyMasterModule.Router()

// 	log.Println("========== MOUNTING ASSEMBLY MASTER ==========")

// 	r.Mount("/assemblymaster", assemblyMasterRouter)

// 	log.Println("========== ASSEMBLY MASTER MOUNTED ==========")

// 	if err := chi.Walk(r, func(
// 		method string,
// 		route string,
// 		handler http.Handler,
// 		middlewares ...func(http.Handler) http.Handler,
// 	) error {
// 		log.Printf("ROOT ROUTE: %-6s %s", method, route)
// 		return nil
// 	}); err != nil {
// 		log.Printf("Root route walk error: %v", err)
// 	}
