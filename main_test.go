package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateService(test *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"http://identity-provider/public-key",
		httpmock.NewStringResponder(http.StatusOK, publicKeyPEMFixture),
	)

	storage := "test-storage-" + uuid.NewV4().String() + ".db"
	defer os.Remove(storage)

	os.Setenv("DATABASE_TYPE", "sqlite3")
	os.Setenv("DATABASE_CONNECTION", storage)

	groupService, err := createService()
	assert.NoError(test, err)
	assert.NotEqual(test, nil, groupService)
}

func TestPOSTGroupWithInvalidToken(test *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"http://identity-provider/public-key",
		httpmock.NewStringResponder(http.StatusOK, publicKeyPEMFixture),
	)

	storage := "test-storage-" + uuid.NewV4().String() + ".db"
	defer os.Remove(storage)

	os.Setenv("DATABASE_TYPE", "sqlite3")
	os.Setenv("DATABASE_CONNECTION", storage)

	groupService, err := createService()
	assert.NoError(test, err)
	assert.NotEqual(test, nil, groupService)

	createGroupResource(groupService)

	groupJSON := `{"name":"The Amazing Circus Group","description":"We make truly amazing shows"}`

	request := httptest.NewRequest(echo.POST, "/group", strings.NewReader(groupJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	context := groupService.Router.NewContext(request, recorder)
	fmt.Println(context.IsTLS())
	//groupService.Router.
}

/*
func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
*/
