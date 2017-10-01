package entity

import uuid "github.com/satori/go.uuid"

// GroupInvitation is used to keep track of one account that invited another to a group
type GroupInvitation struct {
	ID            string `json:"id" gorm:"not null;unique;size:36" validate:"uuid4,required"`
	GroupID       string `json:"groupId" gorm:"not null;size:36" validate:"uuid4,required"`
	ToAccountID   string `json:"toAccountId" gorm:"not null;size:36" validate:"uuid4,required"`
	FromAccountID string `json:"fromAccountId" gorm:"not null;size:36" validate:"uuid4,required"`
}

// BeforeSave will run before the struct is persisted with gorm
func (entity *GroupInvitation) BeforeSave() {
	if entity.ID == "" {
		entity.ID = uuid.NewV4().String()
	}
}
