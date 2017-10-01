package entity

import (
	uuid "github.com/satori/go.uuid"
)

// Project represents an entity that can be used to access the system
type Project struct {
	ID          string `json:"id" gorm:"not null;unique;size:36" validate:"uuid4,required"`
	Name        string `json:"name" gorm:"not null" validate:"required"`
	Description string `json:"description"`
	GroupID     string `json:"groupId" gorm:"not null;size:36" validate:"uuid4,required"`
}

// BeforeSave will run before the struct is persisted with gorm
func (entity *Project) BeforeSave() {
	if entity.ID == "" {
		entity.ID = uuid.NewV4().String()
	}
}
