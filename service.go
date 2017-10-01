package main

import (
	"crypto/rsa"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/mojlighetsministeriet/groups/entity"
	"github.com/mojlighetsministeriet/utils/jwt"
)

// Service is the main service that holds web server and database connections and so on
type Service struct {
	DatabaseConnection *gorm.DB
	Router             *echo.Echo
	Log                echo.Logger
	PublicKey          *rsa.PublicKey
}

// Initialize will prepeare the service by connecting to database and Creating a web server instance (but it will not start listening until service.Listen() is run)
func (service *Service) Initialize(identityProviderURL string, databaseType string, databaseConnectionString string) (err error) {
	service.Router = echo.New()

	service.Log = service.Router.Logger
	service.Log.SetLevel(log.INFO)

	service.PublicKey, err = jwt.FetchPublicKey(identityProviderURL + "/public-key")
	if err != nil {
		return
	}

	service.DatabaseConnection, err = gorm.Open(databaseType, databaseConnectionString)
	if err != nil {
		return
	}

	service.DatabaseConnection.Debug()

	err = service.DatabaseConnection.AutoMigrate(&entity.Group{}).Error
	if err != nil {
		return
	}

	err = service.DatabaseConnection.AutoMigrate(&entity.GroupInvitation{}).Error
	if err != nil {
		return
	}

	err = service.DatabaseConnection.AutoMigrate(&entity.Project{}).Error
	if err != nil {
		return
	}

	err = service.DatabaseConnection.AutoMigrate(&entity.GroupInvitation{}).Error
	if err != nil {
		return
	}

	return
}

func (service *Service) CreateGroup(context echo.Context) error {
	group := entity.Group{}
	err := context.Bind(&group)
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Bad Request\"}"))
	}

	validate := validator.New()
	err = validate.Struct(group)
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Bad Request\"}"))
	}

	err = service.DatabaseConnection.Create(&group).Error
	if err != nil {
		service.Log.Error(err)
		return context.JSONBlob(http.StatusInternalServerError, []byte("{\"message\":\"Internal Server Error\"}"))
	}

	return context.JSON(http.StatusCreated, group)
}

func (service *Service) UpdateGroup(context echo.Context) error {
	group := entity.Group{}
	err := service.DatabaseConnection.Where("id = ?", context.Param("id")).Find(&group).Error
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Not Found\"}"))
	}

	err = context.Bind(&group)
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Bad Request\"}"))
	}

	validate := validator.New()
	err = validate.Struct(group)
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Bad Request\"}"))
	}

	err = service.DatabaseConnection.Save(&group).Error
	if err != nil {
		service.Log.Error(err)
		return context.JSONBlob(http.StatusInternalServerError, []byte("{\"message\":\"Internal Server Error\"}"))
	}

	return context.JSON(http.StatusOK, group)
}

func (service *Service) DeleteGroup(context echo.Context) error {
	group := entity.Group{}
	err := service.DatabaseConnection.Where("id = ?", context.Param("id")).Find(&group).Error
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Not Found\"}"))
	}

	err = service.DatabaseConnection.Save(&group).Error
	if err != nil {
		service.Log.Error(err)
		return context.JSONBlob(http.StatusInternalServerError, []byte("{\"message\":\"Internal Server Error\"}"))
	}

	return context.JSON(http.StatusOK, group)
}

func (service *Service) CreateProject(context echo.Context) error {
	project := entity.Project{}
	err := context.Bind(&project)
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Bad Request\"}"))
	}

	validate := validator.New()
	err = validate.Struct(project)
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Bad Request\"}"))
	}

	err = service.DatabaseConnection.Create(&project).Error
	if err != nil {
		service.Log.Error(err)
		return context.JSONBlob(http.StatusInternalServerError, []byte("{\"message\":\"Internal Server Error\"}"))
	}

	return context.JSON(http.StatusCreated, project)
}

func (service *Service) UpdateProject(context echo.Context) error {
	project := entity.Project{}
	err := service.DatabaseConnection.Where("id = ?", context.Param("id")).Find(&project).Error
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Not Found\"}"))
	}

	err = context.Bind(&project)
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Bad Request\"}"))
	}

	validate := validator.New()
	err = validate.Struct(project)
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Bad Request\"}"))
	}

	err = service.DatabaseConnection.Save(&project).Error
	if err != nil {
		service.Log.Error(err)
		return context.JSONBlob(http.StatusInternalServerError, []byte("{\"message\":\"Internal Server Error\"}"))
	}

	return context.JSON(http.StatusOK, project)
}

func (service *Service) DeleteProject(context echo.Context) error {
	project := entity.Project{}
	err := service.DatabaseConnection.Where("id = ?", context.Param("id")).Find(&project).Error
	if err != nil {
		return context.JSONBlob(http.StatusBadRequest, []byte("{\"message\":\"Not Found\"}"))
	}

	err = service.DatabaseConnection.Save(&project).Error
	if err != nil {
		service.Log.Error(err)
		return context.JSONBlob(http.StatusInternalServerError, []byte("{\"message\":\"Internal Server Error\"}"))
	}

	return context.JSON(http.StatusOK, project)
}

// Listen will make the service start listning for incoming requests
func (service *Service) Listen(address string) (err error) {
	err = service.Router.Start(address)
	return
}

// Close will shut down the service and any of it's related components
func (service *Service) Close() {
	service.DatabaseConnection.Close()
}
