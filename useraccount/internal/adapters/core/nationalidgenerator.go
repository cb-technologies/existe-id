package core

import (
	"log"
	"strconv"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"math/rand"
)

func (adapter Adapter) GenerateNationalID(personInfo *pb.PersonInfoRequest) (*string) {
	return generateUniqueNationalID(personInfo)
}

func generateUniqueNationalID(personInfo *pb.PersonInfoRequest) (*string) {
	// Create the map which will allow us to get the correct abbreviation
	province_abb := ""
	congo_provinces := map[string]string{
        "Bas-Uélé": "BU",
		"Équateur": "ET", // remember to always use a comma, even in the last key-value pair
		"Haut-Katanga": "HK",
		"Haut-Lomami": "HL",
		"Haut-Uélé": "HU",
		"Ituri":"IT",
		"Kasaï": "KS",
		"Kasaï-Central":"LL",
		"Kasaï Oriental":"KO",
		"KINSHASA":"KN",
		"Kongo Central": "BC",
		"Kwango": "KG",
		"Kwilu":"KU",
		"Lomami":"LM",
		"Lualaba": "LB",
		"Mai-Ndombe":"MA",
		"Maniema":"MN",
		"Mongala":"MO",
		"Nord-Kivu":"NK",
		"Nord-Ubangi":"NU",
		"Sankuru":"SN",
		"Sud-Kivu":"SK",
		"Sud-Ubangi":"SU",
		"Tanganyika":"TG",
		"Tshopo":"TO",
		"Tshuapa":"TP",
	}
	
	for key, _ := range congo_provinces {
		if(key == personInfo.Address.Province){
			province_abb := congo_provinces[key]
			log.Printf("The abbreviation of %s is %s\n", personInfo.Address.Province, province_abb)
		}
	}
	log.Printf("The abbreviation of %s is %s\n", personInfo.Address.Province, province_abb)

	id_number := 100000000 + rand.Intn(999999999-100000000)
	nationalIdStr := strconv.FormatInt(int64(id_number), 16)
	final_id := province_abb + "-" + nationalIdStr
	return &final_id
}
