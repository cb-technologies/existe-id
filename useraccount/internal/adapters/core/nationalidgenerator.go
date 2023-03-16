package core

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func (adapter Adapter) GenerateNationalID(personInfo *pb.PersonInfoRequest) *string {
	return generateUniqueNationalID(personInfo)
}

func generateUniqueNationalID(personInfo *pb.PersonInfoRequest) *string {
	// Create the map which will allow us to get the correct abbreviation
	province_abb := ""
	congo_provinces := map[string]string{
		"BAS-UELE":       "BU",
		"EQUATEUR":       "ET", // remember to always use a comma, even in the last key-value pair
		"HAUT-KATANGA":   "HK",
		"HAUT-LOMAMI":    "HL",
		"HAUT-UELE":      "HU",
		"ITURI":          "IT",
		"KASAI":          "KS",
		"KASAI-CENTRAL":  "LL",
		"KASAI-ORIENTAL": "KO",
		"KINSHASA":       "KN",
		"KONGO-CENTRAL":  "BC",
		"KWANGO":         "KG",
		"KWILU":          "KU",
		"LOMAMI":         "LM",
		"LUALABA":        "LB",
		"MAI-NDOMBE":     "MA",
		"MANIEMA":        "MN",
		"MONGALA":        "MO",
		"NORD-KIVU":      "NK",
		"NORD-UBANGI":    "NU",
		"SANKURU":        "SN",
		"SUD-KIVU":       "SK",
		"SUD-UBANGI":     "SU",
		"TANGANYIKA":     "TG",
		"TSHOPO":         "TO",
		"TSHUAPA":        "TP",
	}

	for key, _ := range congo_provinces {
		if key == personInfo.Address.Province {
			province_abb = congo_provinces[key]
			log.Printf("The abbreviation of %s is %s\n", personInfo.Address.Province, province_abb)
		}
	}
	log.Printf("The abbreviation of %s is %s\n", personInfo.Address.Province, province_abb)
	rand.Seed(time.Now().UnixNano())
	id_number := 0 + rand.Intn(999999999-100000000)
	nationalIdStr := strconv.Itoa(id_number)
	log.Printf(nationalIdStr)
	final_id := province_abb + "-" + nationalIdStr
	return &final_id
}
