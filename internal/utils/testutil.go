package utils

import (
	"company/controllers"
	"company/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func SetUpTest() (*gin.Engine, *controllers.Controller) {

	r := setUpRouter()
	db := setupStorage()
	ctrl := controllers.NewController(db)

	return r, ctrl
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func setupStorage() *gorm.DB {

	//sqlite db for testing
	db, err := createDb()
	if err != nil {
		panic(err)
	}

	return db
}

func createDb() (*gorm.DB, error) {

	//Open db connection and create database
	cxn := "file:memdb1?mode=memory&cache=shared"
	db, err := gorm.Open("sqlite3", cxn)
	if err != nil {
		return nil, err
	}

	//fill db with dummy record
	fillDb(db)

	return db, nil
}

func fillDb(db *gorm.DB) {

	//Migrate tables
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Company{})

	//Add user
	user := models.User{Username: "test user a", Password: "$2a$14$rNcCJCuCtPoCEjfEZ8341eUqLSiel/350gQX/jicwzqF.dUoodYmm"}
	db.Create(&user)

	//Add company
	company := models.Company{Name: "Test Company a", Description: "This is test company description", Employees: 10, Registered: true, Type: "NonProfit"}
	db.Create(&company)
}
