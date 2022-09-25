package db

import (
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/date"
	"gorm.io/gorm"
)

type AddressModel struct {
	Number    int32
	Avenue    string
	Quartier  string
	Commune   string
	ZipCode   string
	Province  string
	Ville     string
	Reference string
}

type NamesModel struct {
	Nom         string
	Prenom      string
	MiddleNames []string
}

type OriginModel struct {
	Province []string
	ChefLieu string
}

type PhenotypeModel struct {
	Height   float32
	Weight   float32
	EyeColor string
}

type BiometricModel struct {
	Photos      []byte
	FingerPrint []byte
}

type DateOfBirthModel struct {
	Day   string
	Month string
	Year  string
}

type CardValidityModel struct {
	CardIssueDate      date.Date
	CardExpirationDate date.Date
}

type PersonInfoModel struct {
	gorm.Model

	Names        NamesModel        `gorm:"embedded"`
	Biometrics   BiometricModel    `gorm:"embedded"`
	Address      AddressModel      `gorm:"embedded"`
	Origins      OriginModel       `gorm:"embedded"`
	Phenotypes   PhenotypeModel    `gorm:"embedded"`
	DateOfBirth  DateOfBirthModel  // TODO we have to change this into a more compact format. we can use the data type provided by go
	CardValidity CardValidityModel `gorm:"embedded"`
	UUID         uuid.UUID
}
