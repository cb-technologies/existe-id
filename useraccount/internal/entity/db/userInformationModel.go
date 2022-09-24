package db

import (
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/date"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model

	Number    int32
	Avenue    string
	Quartier  string
	Commune   string
	ZipCode   string
	Province  string
	Ville     string
	Reference string
}

type Names struct {
	gorm.Model

	Nom         string
	Prenom      string
	MiddleNames []string
}

type Origin struct {
	gorm.Model

	Province []string
	ChefLieu string
}

type Phenotype struct {
	Height   float32
	Weight   float32
	EyeColor string
}

type Biometric struct {
	Photos      []byte
	FingerPrint []byte
}

type DateOfBirth struct {
	Day   string
	Month string
	Year  string
}

type CardValidity struct {
	IssueDate      date.Date
	ExpirationDate date.Date
}

type PersonInfoResponse struct {
	gorm.Model

	Names        *Names
	Biometrics   *Biometric
	Address      *Address
	Origins      *Origin
	Phenotypes   *Phenotype
	DateOfBirth  *DateOfBirth
	CardValidity *CardValidity
	UUID         *uuid.UUID
}

type PersonInfoRequest struct {
	gorm.Model

	Names       *Names
	Biometrics  *Biometric
	Address     *Address
	Origins     *Origin
	Phenotypes  *Phenotype
	DateOfBirth *DateOfBirth
	UUID        *uuid.UUID
}
