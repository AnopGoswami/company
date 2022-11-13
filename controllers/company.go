package controllers

import (
	"company/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service *Controller) RegisterCompany(context *gin.Context) {

	var input models.CompanyRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	company := models.Company{
		Name:        input.Name,
		Description: input.Description,
		Employees:   input.Employees,
		Type:        input.Type,
		Registered:  true,
	}
	record := service.storage.Create(&company)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, company)
}

func (service *Controller) GetCompany(context *gin.Context) {
	var company models.Company
	if err := service.storage.Where("id = ?", context.Param("id")).First(&company).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No record found"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, company)
}

func (service *Controller) UpdateCompany(context *gin.Context) {

	var input models.CompanyUpdateRequest
	var company models.Company

	if err := service.storage.Where("id = ?", context.Param("id")).First(&company).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No record found"})
		context.Abort()
		return
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	service.storage.Model(&company).Updates(input)
	context.JSON(http.StatusOK, company)
}

func (service *Controller) DeleteCompany(context *gin.Context) {

	var company models.Company
	if err := service.storage.Where("id = ?", context.Param("id")).First(&company).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "No record found"})
		context.Abort()
		return
	}

	service.storage.Delete(company)
	context.JSON(http.StatusOK, gin.H{"success": true})
}
