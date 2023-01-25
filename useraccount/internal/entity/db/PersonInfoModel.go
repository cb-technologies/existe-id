package db

import (
	"google.golang.org/genproto/googleapis/type/date"
	"gorm.io/gorm"
)

type AddressModel struct {
	Number          int32
	Avenue          string
	Quartier        string
	Commune         string
	ZipCode         string
	Ville           string
	Reference       string
	ProvinceAddress string
}

type NamesModel struct {
	Nom         string
	Prenom      string
	MiddleNames []string `gorm:"serializer:json"`
}

type OriginModel struct {
	Province        []string `gorm:"serializer:json"`
	ChefLieu        string
	LieuDeNaissance string
}

type PhenotypeModel struct {
	Height   float32
	Weight   float32
	EyeColor string
}

type BiometricModel struct {
	Photos      []byte `gorm:"serializer:json"`
	FingerPrint []byte `gorm:"serializer:json"`
	PhotoType   string
}

type DateOfBirthModel struct {
	Day   string
	Month string
	Year  string
}

type Sex struct {
	Sex string
}

type CardValidityModel struct {
	CardIssueDate      date.Date `gorm:"serializer:json"`
	CardExpirationDate date.Date `gorm:"serializer:json"`
}

type NationalIDNumberModel struct {
	NationalID string
}

type QRCodeModel struct {
	QRCode string
}

type PersonInfoModel struct {
	gorm.Model

	Names        NamesModel            `gorm:"embedded"`
	Biometrics   BiometricModel        `gorm:"embedded"`
	Address      AddressModel          `gorm:"embedded"`
	Origins      OriginModel           `gorm:"embedded"`
	Phenotypes   PhenotypeModel        `gorm:"embedded"`
	DateOfBirth  DateOfBirthModel      `gorm:"serializer:json"` // TODO we have to change this into a more compact format. we can use the data type provided by go
	CardValidity CardValidityModel     `gorm:"embedded"`
	NationalID   NationalIDNumberModel `gorm:"embedded"`
	Sex          Sex                   `gorm:"embedded"`
	QRCode       QRCodeModel           `gorm:"embedded"`
}
