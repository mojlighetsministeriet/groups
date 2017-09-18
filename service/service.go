package service

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/mojlighetsministeriet/groups/entity"
)

// Service is the main service that holds web server and database connections and so on
type Service struct {
	DatabaseConnection *gorm.DB
	Router             *echo.Echo
	Log                echo.Logger
}

// Initialize will prepeare the service by connecting to database and creating a web server instance (but it will not start listening until service.Listen() is run)
func (service *Service) Initialize(databaseType string, databaseConnectionString string) (err error) {
	service.Router = echo.New()

	service.Log = service.Router.Logger
	service.Log.SetLevel(log.INFO)

	service.DatabaseConnection, err = gorm.Open(databaseType, databaseConnectionString)
	if err != nil {
		return
	}

	service.DatabaseConnection.Debug()

	err = service.DatabaseConnection.AutoMigrate(&entity.Group{}).Error
	if err != nil {
		return
	}

	err = service.DatabaseConnection.AutoMigrate(&entity.Project{}).Error
	if err != nil {
		return
	}

	return
}

// Listen will make the service start listning for incoming requests
func (service *Service) Listen(address string) (err error) {
	service.Router.Logger.Error(service.Router.Start(address))
	return
}

// Close will shut down the service and any of it's related components
func (service *Service) Close() {
	service.DatabaseConnection.Close()
}
