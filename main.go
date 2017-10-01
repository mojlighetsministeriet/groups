package main // import "github.com/mojlighetsministeriet/groups"

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mojlighetsministeriet/utils"
	"github.com/mojlighetsministeriet/utils/jwt"
)

func createService() (groupService *Service, err error) {
	groupService = &Service{}

	err = groupService.Initialize(
		utils.GetEnv("IDENTITY_PROVIDER_URL", "http://identity-provider"),
		utils.GetEnv("DATABASE_TYPE", "mysql"),
		utils.GetEnv(
			"DATABASE_CONNECTION",
			utils.GetFileAsString("/run/secrets/database-connection", "user:password@/dbname?charset=utf8mb4,utf8&parseTime=True&loc=Europe/Stockholm"),
		),
	)

	return
}

func createGroupResource(groupService *Service) {
	groupResource := groupService.Router.Group("/group")
	groupResource.Use(jwt.RequiredRoleMiddleware(groupService.PublicKey, "user"))

	groupResource.POST("", groupService.createGroup)
}

func main() {
	groupService, err := createService()
	if err != nil {
		groupService.Log.Error("Failed to initialize the service, make sure that you provided the correct identity-provider URL and database credentials.")
		groupService.Log.Error(err)
		panic("Cannot continue due to previous errors.")
	}

	defer groupService.Close()

	createGroupResource(groupService)

	err = groupService.Listen(":1323")
	if err != nil {
		panic(err)
	}
}
