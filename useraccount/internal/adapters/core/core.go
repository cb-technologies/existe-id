package core

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
	"github.com/sony/sonyflake"
	"log"
	"strconv"
)

func GenerateNationalID(model *db.PersonInfoModel) {
	model.NationalID = db.NationalIDNumberModel{NationalID: *generateUniqueNationalID()}
}

func generateUniqueNationalID() *string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	nationalIdStr := strconv.FormatInt(int64(id), 16)

	return &nationalIdStr
}
