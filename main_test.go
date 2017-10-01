package main

import (
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

	groupService.Router.
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

const publicKeyPEMFixture = `
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtn3txTNTthQH2ZN6RxSn
fYvbURQvzse5uq3mqcLjqaNIHPT9AtO8eSgLs23uwku8tkABBVihFmGHNhuf2tVa
mq3AU/bJcoRjqYOTr6AifiAsoZ46n8pFGl7zeNwZZSxHvPZ+uXZiTreU9kiomoAs
e6HdwHOXdcj3MMIGzO7zRZE3diMWS2sDmEfY3iApQ5DXqcqxxkih/YPSF3WiDyX0
y6gI5B43Bmrx174r7FkNwllZhjYrMErdMQ463g2axhJmHC96bFvXk6OX0sLZACqK
xMlFE61R5ZrvpFtLPZiEanqQXeM8fYejm2KoJCMr4vcglOraKpvO3+f5Vq67VMjf
1477/3iYTb2DwfYAexvBWTff5ei0EUZzKOkfVUhpC0aH1Nq6MCG570SF9I4bQ72d
oP/6D4JlX+div6ZY5jVcpKuC1soGDYRhNNsfucJ+ZOQ/ibKrrSOMQ5pAYS3ZU+2z
02wP5SSaD1SwIzGXDZumCvUDQ/T144tSBRm8A2bHJSwuK5Un/6jNFbHLg5S5p8Mv
Maajn1/A/z9UsA4nUFYxBixTQ8Yr8o9x4PPeiMjcJWEJs1MfaaiRvZpLh1TQ7OyV
F4gSVHbWTQuv2ZdqfDjz3JXiO4zM9LmGV/CL2lQ7OuBV/DB5A9SjVCgjO2MI0oL1
8BLWTs0bx/tRDK95JM8bTgcCAwEAAQ==
-----END PUBLIC KEY-----
`
