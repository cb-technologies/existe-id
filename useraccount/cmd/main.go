package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"
	//"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
)

func main() {

	fmt.Println("Existe ID")

	postgres, err := postgresSQL.NewAdapter()

	if err != nil {
		fmt.Println("The error is ", err)
	}

	fmt.Println("Connection succesful!")

	// Trying to mimic a request just to see if the GORM code is working
	// namesTest := db.NamesModel{
	// 	Nom:         "Bingoto",
	// 	Prenom:      "Areely",
	// 	MiddleNames: []string{"Bamanissa"},
	// }
	// biometricsTest := db.BiometricModel{
	// 	Photos:      []uint8{1, 2},
	// 	FingerPrint: []uint8{2, 3},
	// }

	// addressTest := db.AddressModel{
	// 	Number:   1,
	// 	Avenue:   "Nicolas",
	// 	Quartier: "Santa Clara",
	// }

	// originTest := db.OriginModel{
	// 	Province: []string{"Bandundu"},
	// }

	// phenotypeTest := db.PhenotypeModel{
	// 	EyeColor: "blue",
	// }

	// dateOfBirthTest := db.DateOfBirthModel{
	// 	Day:   "Monday",
	// 	Month: "February",
	// 	Year:  "2002",
	// }
	// my_UUID := uuid.New()

	// personTest := db.PersonInfoModel{
	// 	Names:       namesTest,
	// 	Biometrics:  biometricsTest,
	// 	Address:     addressTest,
	// 	Origins:     originTest,
	// 	Phenotypes:  phenotypeTest,
	// 	DateOfBirth: dateOfBirthTest,
	// 	UUID:        my_UUID,
	// }

	//err = postgres.AddNewPersonInfo(&personTest)
	my_id uuid.UUID := "915aa63c-68e8-4c0f-b884-033b2673ab14";
	postgres.UpdatePersonInfo()
	if err != nil {
		fmt.Println("Test Failed! A demain faut dormir")
	} else {
		fmt.Println("Person created successfully with UUID")
	}

	// 915aa63c-68e8-4c0f-b884-033b2673ab14
	postgres.CloseDBConnection()

}
