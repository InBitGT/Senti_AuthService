package migration

import (
	"fmt"

	"AuthService/db"
	"AuthService/internal/modules/otp"
	"AuthService/internal/modules/permission"
	"AuthService/internal/modules/permissionmodule"
	"AuthService/internal/modules/refreshtoken"
	"AuthService/internal/modules/role"
	"AuthService/internal/modules/tenant"
	"AuthService/internal/modules/user"
)

func Migration() {
	database := db.Database()

	err := database.AutoMigrate(
		&tenant.Tenant{},
		&role.Role{},
		&permission.Permission{},
		&permissionmodule.PermissionModule{},
		&user.User{},
		&refreshtoken.RefreshToken{},
		&otp.UserOTP{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migraciones de auth ejecutadas")
}
