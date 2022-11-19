package core

import (
	"log"
	"strconv"

	"github.com/sony/sonyflake"
)

func (adapter Adapter) GenerateNationalID() (*string, error) {
	return generateUniqueNationalID()
}

func generateUniqueNationalID() (*string, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Printf("flake.NextID() failed with %s\n", err)
	}
	nationalIdStr := strconv.FormatInt(int64(id), 16)

	return &nationalIdStr, err
}
