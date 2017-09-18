package main // import "github.com/mojlighetsministeriet/groups"

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mojlighetsministeriet/groups/service"
	"github.com/mojlighetsministeriet/utils"
)

func main() {
	groupService := service.Service{}
	err := groupService.Initialize(
		utils.GetEnv("DATABASE_TYPE", "mysql"),
		utils.GetFileAsString("/run/secrets/database-connection", "user:password@/dbname?charset=utf8mb4,utf8&parseTime=True&loc=Europe/Stockholm"),
	)
	if err != nil {
		groupService.Log.Error("Failed to initialize the service, make sure that you provided the correct database credentials.")
		groupService.Log.Error(err)
		panic("Cannot continue due to previous errors.")
	}
	defer groupService.Close()

	groupService.Listen(":1323")
}
