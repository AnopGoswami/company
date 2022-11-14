package tests

import (
	"bytes"
	"company/controllers"
	"company/middlewares"
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
	api := router.Group("/v1")
	{
		api.GET("/healthcheck", service.HealthCheck)
		api.POST("/token", service.GenerateToken)
		api.POST("/user", service.RegisterUser)
		api.POST("/company", service.RegisterCompany).Use(middlewares.Auth())
		api.GET("/company/:id", service.GetCompany).Use(middlewares.Auth())
		api.PATCH("/company/:id", service.UpdateCompany).Use(middlewares.Auth())
		api.DELETE("/company/:id", service.DeleteCompany).Use(middlewares.Auth())
	}
	gin.SetMode(gin.ReleaseMode)
}

func TestHealthcheck(t *testing.T) {

	mockResponse := `{"status":"ok"}`
	req, _ := http.NewRequest("GET", HelathCheckEndpoint, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddUser(t *testing.T) {

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
	company := models.Company{Name: "Test Company 1", Description: "This is test company description", Employees: 10, Registered: true, Type: "NonProfit"}
	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", CompanyEndpoint, bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetCompany(t *testing.T) {
	req, _ := http.NewRequest("GET", CompanyEndpoint+"/1", nil)
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateCompany(t *testing.T) {
	company := models.Company{Name: "Test Company D", Description: "This is test company description", Employees: 10, Registered: true, Type: "NonProfit"}
	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("PATCH", CompanyEndpoint+"/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCompany(t *testing.T) {
	req, _ := http.NewRequest("DELETE", CompanyEndpoint+"/1", nil)
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
