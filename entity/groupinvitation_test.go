package entity_test

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mojlighetsministeriet/groups/entity"
	"github.com/stretchr/testify/assert"
)

func TestGroupInvitationBeforeSave(test *testing.T) {
	invitation := entity.GroupInvitation{}
	assert.Equal(test, 0, len(invitation.ID))
	invitation.BeforeSave()
	assert.Equal(test, 36, len(invitation.ID))
}
