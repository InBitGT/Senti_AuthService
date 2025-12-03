package migration

import (
	"fmt"

	"AuthService/db"
	"AuthService/internal/modules/auth"
)

func Migration() {
	database := db.Database()
	err := database.AutoMigrate(
		&auth.Tenant{},
		&auth.Role{},
		&auth.Permission{},
		&auth.RolePermission{},
		&auth.User{},
		&auth.RefreshToken{},
		&auth.UserOTP{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Migraciones de auth ejecutadas")
}
