package deliveries

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	mockUserType      = "*models.Users"
	mockUserTypeSlice = "*[]models.Users"
	userEmail         = "khomkrid@twinsynergy.co.th"
	userPassword      = "password"
	userFullName      = "Arashi L."
	userTelephone     = "0925905444"
	userRole          = 1
)

func getUser(email, password, name, telephone string, role int) paramsUser {
	return paramsUser{
		Email:     email,
		Password:  password,
		Name:      name,
		Telephone: telephone,
		RoleID:    role,
	}
}
func UserMockSetup(t *testing.T) (mockUseCase *mocks.UserUseCase, handler *userHandler) {
	SetupMock(t)
	var mockUser models.Users
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUseCase = new(mocks.UserUseCase)
	handler = NewUserHandelr(mockUseCase, "expect1", "expect2", "/user")
	return
}
func TestNewUserHandlerSuccess(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)

	app := fiber.New()
	app.Post("/user", handler.NewUser)
	req, err := http.NewRequest("POST", "/user", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
	//fail validate
	params := getUser(userEmail, "", userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	app.Post("/user", handler.NewUser)
	req, err = http.NewRequest("POST", "/user", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
	//fail validate
	params = getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	data, _ = json.Marshal(&params)
	payload = bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	app.Post("/user", handler.NewUser)
	req, err = http.NewRequest("POST", "/user", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestNewUserHandlerFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)

	app := fiber.New()
	params := getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(errors.New("error TestNewUserHandlerFail"))
	app.Post("/user", handler.NewUser)
	req, err := http.NewRequest("POST", "/user", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}