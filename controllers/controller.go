package controllers

import (
	"github.com/jinzhu/gorm"
)

type Controller struct {
	storage *gorm.DB
}

//Controller initialises Controller instance
func NewController(storage *gorm.DB) *Controller {
	return &Controller{storage: storage}
}
