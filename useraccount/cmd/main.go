package main

import (
	"fmt"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/application/api"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/core/national_id_generator"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/ports"
)

func main() {

	fmt.Println("Existe ID")

	var postgres ports.PostgresSQLPort
	var application ports.APIPorts
	var core ports.IDCoreFunctionsPorts

	postgres, err := postgresSQL.NewAdapter()
	core = national_id_generator.NewAdapter()

	application = api.NewAdapter(postgres, core)
	my_postgres, err := postgresSQL.NewAdapter()
	if err != nil {
		fmt.Println("The error is ", err)
	}

	fmt.Println("Connection succesful!")

	//Trying to mimic a request just to see if the GORM code is working
	// namesTest := pb.NamesModel{
	// 	Nom:         "Nathan",
	// 	Prenom:      "Tuala",
	// 	MiddleNames: []string{"Mbuyi"},
	// }
	// biometricsTest := db.BiometricModel{
	// 	Photos:      []uint8{1, 2},
	// 	FingerPrint: []uint8{2, 3},
	// }

	// addressTest := db.AddressModel{
	// 	Number:   1,
	// 	Avenue:   "Ramba Place",
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

	namesTest := pb.Names{
		Nom:         "Andrea",
		Prenom:      "Mufuta",
		MiddleNames: []string{"Mbuyi"},
	}
	biometricsTest := pb.Biometric{
		Photos:      []uint8{1, 2},
		FingerPrint: []uint8{2, 3},
	}

	addressTest := pb.Address{
		Number:   1,
		Avenue:   "Ramba Place",
		Quartier: "Santa Clara",
	}

	originTest := pb.Origin{
		Province: []string{"Bandundu"},
	}

	phenotypeTest := pb.Phenotype{
		EyeColor: "blue",
	}

	dateOfBirthTest := pb.DateOfBirth{
		Day:   "Monday",
		Month: "February",
		Year:  "2002",
	}

	personTest := pb.PersonInfoRequest{
		Names:       &namesTest,
		Biometrics:  &biometricsTest,
		Address:     &addressTest,
		Origins:     &originTest,
		Phenotypes:  &phenotypeTest,
		DateOfBirth: &dateOfBirthTest,
	}

	nationalID := &db.NationalIDNumberModel{NationalID: "5f054e6280000de"}
	// personTest := db.PersonInfoModel{
	// 	Names:       namesTest,
	// 	Biometrics:  biometricsTest,
	// 	Address:     addressTest,
	// 	Origins:     originTest,
	// 	Phenotypes:  phenotypeTest,
	// 	DateOfBirth: dateOfBirthTest,
	// 	NationalID:  *nationalID,
	// }

	//err = application.AddNewPersonInfo(&personTest)
	//person, err := my_postgres.FindPersonInfo(nationalID)
	person, err := my_postgres.FindPersonInfo(nationalID)
	// if false {
	// 	my_postgres.UpdatePersonInfo(&personTest)
	// }
	if err != nil {
		fmt.Println("Test Failed! A demain faut dormir")
	} else {
		fmt.Println("---------------------------")
		fmt.Println(person.Names)
		fmt.Println("Person created successfully")
	}

}
