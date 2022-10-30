package db

import "gorm.io/gorm"

type AgentInfoModel struct {
	gorm.Model

	Nom      string `gorm:"not null"`
	Prenom   string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"` // will make this unique and create an index on this field in the db
	Password string `gorm:"not null"`
}

type AgentSignInModel struct {
	Email    string
	Password string
}
