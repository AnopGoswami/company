package tests

import (
	"bytes"
	"company/controllers"
	"company/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"company/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	HelathCheckEndpoint = "/v1/healthcheck"
	UserEndpoint        = "/v1/user"
	TokenEndpoint       = "/v1/token"
	CompanyEndpoint     = "/v1/company"
)

var (
	token   string
	router  *gin.Engine
	service *controllers.Controller
)

func init() {
	router, service = utils.SetUpTest()
}

func TestHealthcheck(t *testing.T) {

	mockResponse := `{"status":"ok"}`

	router.GET(HelathCheckEndpoint, service.HealthCheck)
	req, _ := http.NewRequest("GET", HelathCheckEndpoint, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddUser(t *testing.T) {

	router.POST(UserEndpoint, service.RegisterUser)

	user := models.User{
		Username: "new user",
		Password: "new passwd",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", UserEndpoint, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetToken(t *testing.T) {

	user := models.User{
		Username: "test user a",
		Password: "mypasswd",
	}

	jsonValue, _ := json.Marshal(user)

	//Get token
	router.POST(TokenEndpoint, service.GenerateToken)
	req, _ := http.NewRequest("POST", TokenEndpoint, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	//Set token
	type TokenReponse struct {
		Token string `json:"token"`
	}
	var response TokenReponse
	json.Unmarshal(w.Body.Bytes(), &response)
	token = response.Token
}

func TestAddCompany(t *testing.T) {

	router.POST(CompanyEndpoint, service.RegisterCompany)
	company := models.Company{Name: "Test Company 1", Description: "This is test company description", Employees: 10, Registered: true, Type: "NonProfit"}

	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", CompanyEndpoint, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetCompany(t *testing.T) {

	router.GET(CompanyEndpoint+"/:id", service.GetCompany)
	req, _ := http.NewRequest("GET", CompanyEndpoint+"/1", nil)
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateCompany(t *testing.T) {

	company := models.Company{Name: "Test Company D", Description: "This is test company description", Employees: 10, Registered: true, Type: "NonProfit"}
	jsonValue, _ := json.Marshal(company)
	router.PATCH("/v1/company/:id", service.UpdateCompany)
	req, _ := http.NewRequest("PATCH", "/v1/company/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteCompany(t *testing.T) {

	router.DELETE(CompanyEndpoint+"/:id", service.GetCompany)
	req, _ := http.NewRequest("DELETE", CompanyEndpoint+"/1", nil)
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
