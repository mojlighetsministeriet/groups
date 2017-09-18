package entity_test

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mojlighetsministeriet/groups/entity"
	"github.com/stretchr/testify/assert"
)

func TestProjectBeforeSave(test *testing.T) {
	project := entity.Project{}
	assert.Equal(test, 0, len(project.ID))
	project.BeforeSave()
	assert.Equal(test, 36, len(project.ID))
}
