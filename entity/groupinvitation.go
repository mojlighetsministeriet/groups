package entity

import uuid "github.com/satori/go.uuid"

type GroupInvitation struct {
	ID        string `json:"id" gorm:"not null;unique;size:36" validate:"uuid4,required"`
	Group     Group  `gorm:"not null;size:36"`
	AccountID string `json:"accountId" gorm:"not null;size:36" validate:"uuid4,required"`
}

// BeforeSave will run before the struct is persisted with gorm
func (entity *GroupInvitation) BeforeSave() {
	if entity.ID == "" {
		entity.ID = uuid.NewV4().String()
	}
}
