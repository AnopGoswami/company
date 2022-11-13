package models

type Company struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"unique;size:15;not null"`
	Description string `json:"description" gorm:"size:3000"`
	Employees   int    `json:"employees" gorm:"not null""`
	Registered  bool   `json:"registered" gorm:"not null"`
	Type        string `json:"type" gorm:"not null"`
}

type CompanyRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=15"`
	Description string `json:"description"`
	Employees   int    `json:"employees" binding:"required"`
	Type        string `json:"type" binding:"required,oneof=Corporations NonProfit Cooperative SoleProprietorship"`
}

type CompanyUpdateRequest struct {
	Name        string `json:"name" binding:"min=1,max=15"`
	Description string `json:"description"`
	Employees   int    `json:"employees"`
	Type        string `json:"type" binding:"oneof=Corporations NonProfit Cooperative SoleProprietorship"`
}
