package core

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
	golangDate "github.com/rickb777/date"
	"google.golang.org/genproto/googleapis/type/date"
)

func (adapter Adapter) GetCardValidity() *db.CardValidityModel {
	return &db.CardValidityModel{
		CardExpirationDate: date.Date{Day: int32(golangDate.Today().Day()), Month: int32(golangDate.Today().Month()), Year: int32(golangDate.Today().Year() + 5)},
		CardIssueDate:      date.Date{Day: int32(golangDate.Today().Day()), Month: int32(golangDate.Today().Month()), Year: int32(golangDate.Today().Year())},
	}
}
