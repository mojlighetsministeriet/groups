package entity_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mojlighetsministeriet/groups/entity"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGroupBeforeSave(test *testing.T) {
	group := entity.Group{
		Projects: []string{"b25eb146-6a97-405d-b1d9-5ea0ca5e9866", "8d9741da-5d8d-4269-b5f0-93624f154ae4"},
		Members:  []string{"7c935ca4-768b-4a76-ae21-bed3d218a5e0", "f68be5ab-1fca-4924-88d2-3a85042a5f42"},
	}
	assert.Equal(test, 0, len(group.ID))
	group.BeforeSave()
	assert.Equal(test, 36, len(group.ID))
	assert.Equal(test, "b25eb146-6a97-405d-b1d9-5ea0ca5e9866,8d9741da-5d8d-4269-b5f0-93624f154ae4", group.ProjectsSerialized)
	assert.Equal(test, "7c935ca4-768b-4a76-ae21-bed3d218a5e0,f68be5ab-1fca-4924-88d2-3a85042a5f42", group.MembersSerialized)
}

func TestGroupAfterFind(test *testing.T) {
	group := entity.Group{
		ProjectsSerialized: "b25eb146-6a97-405d-b1d9-5ea0ca5e9866,8d9741da-5d8d-4269-b5f0-93624f154ae4",
		MembersSerialized:  "7c935ca4-768b-4a76-ae21-bed3d218a5e0,f68be5ab-1fca-4924-88d2-3a85042a5f42",
	}
	group.AfterFind()
	assert.Equal(test, []string{"b25eb146-6a97-405d-b1d9-5ea0ca5e9866", "8d9741da-5d8d-4269-b5f0-93624f154ae4"}, group.Projects)
	assert.Equal(test, []string{"7c935ca4-768b-4a76-ae21-bed3d218a5e0", "f68be5ab-1fca-4924-88d2-3a85042a5f42"}, group.Members)
}

func TestGroupLoadGroupFromID(test *testing.T) {
	databaseConnection, err := gorm.Open("sqlite3", "/tmp/group-test-"+uuid.NewV4().String()+".db")
	assert.NoError(test, err)
	defer databaseConnection.Close()

	err = databaseConnection.AutoMigrate(&entity.Group{}).Error
	assert.NoError(test, err)

	group := entity.Group{ID: uuid.NewV4().String(), Name: "The Group"}
	err = databaseConnection.Create(&group).Error
	assert.NoError(test, err)

	loadedGroup, err := entity.LoadGroupFromID(databaseConnection, group.ID)
	assert.NoError(test, err)
	assert.Equal(test, "The Group", loadedGroup.Name)
	assert.Equal(test, group.ID, loadedGroup.ID)
}
