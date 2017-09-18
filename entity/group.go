package entity

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Group represents an entity that can be used to access the system
type Group struct {
	ID                 string   `json:"id" gorm:"not null;unique;size:36" validate:"uuid4,required"`
	Name               string   `json:"name" gorm:"not null;unique validate:"required"`
	Description        string   `json:"description"`
	Projects           []string `json:"projects" gorm:"-"`
	ProjectsSerialized string   `gorm:"projects"`
	Members            []string `json:"members" gorm:"-"`
	MembersSerialized  string   `gorm:"members"`
}

// BeforeSave will run before the struct is persisted with gorm
func (entity *Group) BeforeSave() {
	if entity.ID == "" {
		entity.ID = uuid.NewV4().String()
	}
	fmt.Println(entity.ID)
	entity.ProjectsSerialized = strings.Join(entity.Projects, ",")
	entity.MembersSerialized = strings.Join(entity.Members, ",")
}

// AfterFind will run after the struct has been read from persistence
func (entity *Group) AfterFind() {
	entity.Projects = strings.Split(entity.ProjectsSerialized, ",")
	entity.Members = strings.Split(entity.MembersSerialized, ",")
}

// LoadGroupFromID will fetch the entity from the persistence
func LoadGroupFromID(databaseConnection *gorm.DB, id string) (entity Group, err error) {
	err = databaseConnection.Where("id = ?", id).First(&entity).Error
	return
}
