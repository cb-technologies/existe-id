package national_id_generator

import (
	"github.com/sony/sonyflake"
	"log"
	"strconv"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (adapter Adapter) GenerateNationalID() (*string, error) {
	return generateUniqueNationalID()
}

func generateUniqueNationalID() (*string, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	nationalIdStr := strconv.FormatInt(int64(id), 16)

	return &nationalIdStr, err
}
