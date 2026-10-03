/******************************************************************************
 *
 * MODULE      : Permission
 * FILE        : device.go
 *
 * DESCRIPTION :
 * Device Permissions
 *
 ******************************************************************************/

package permission

const (
	RoleSuperAdmin  = "superadmin"
	RoleXoomAdmin   = "xoomadmin"
	RoleXoomUser    = "xoomuser"
	RoleAdminTenant = "admintenant"
	RoleTenantOwner = "tenantowner"
	RoleTenantUser  = "tenantuser"
	RolePlantHeader = "plantheader"

	RoleDistributorAdmin   = "distributoradmin"
	RoleDistributorService = "distributorservice"
	RoleDistributorUser    = "distributoruser"

	RoleTenantAdmin       = "tenantadmin"
	RoleTenantSupervisor  = "tenantsupervisor"
	RoleTenantMaintenance = "tenantmaintenance"
	RoleTenantOperator    = "tenantoperator"
	RoleTenantViewer      = "tenantviewer"
)

const (

	//----------------------------------------------------------------------
	// Device Master
	//----------------------------------------------------------------------

	DeviceCreate = "DEVICE_CREATE"

	DeviceUpdate = "DEVICE_UPDATE"

	DeviceDelete = "DEVICE_DELETE"

	DeviceView = "DEVICE_VIEW"

	DeviceList = "DEVICE_LIST"

	//----------------------------------------------------------------------
	// Device Operations
	//----------------------------------------------------------------------

	DeviceUpdateStatus = "DEVICE_UPDATE_STATUS"

	DeviceUpdateFirmware = "DEVICE_UPDATE_FIRMWARE"

	DeviceUpdateLastSeen = "DEVICE_UPDATE_LAST_SEEN"

	DeviceRestart = "DEVICE_RESTART"

	DeviceResetCounter = "DEVICE_RESET_COUNTER"

	DeviceSyncTime = "DEVICE_SYNC_TIME"

	DeviceSendCommand = "DEVICE_SEND_COMMAND"

	DeviceOTAUpdate = "DEVICE_OTA_UPDATE"

	DeviceAssignMachine = "DEVICE_ASSIGN_MACHINE"

	DeviceAssignTenant = "DEVICE_ASSIGN_TENANT"

	DeviceGenerateSecret = "DEVICE_GENERATE_SECRET"

	DeviceRotateSecret = "DEVICE_ROTATE_SECRET"

	DeviceViewSecret = "DEVICE_VIEW_SECRET"

	DeviceUpdateMQTT = "DEVICE_UPDATE_MQTT"

	DeviceUpdateSoftAP = "DEVICE_UPDATE_SOFTAP"

	DeviceActivate = "DEVICE_ACTIVATE"

	DeviceDeactivate = "DEVICE_DEACTIVATE"

	DeviceExport = "DEVICE_EXPORT"

	DeviceImport = "DEVICE_IMPORT"
)
