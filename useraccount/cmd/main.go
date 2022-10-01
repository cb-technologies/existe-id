package main

import (
	"fmt"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/application/api"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/core/national_id_generator"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
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

	if err != nil {
		fmt.Println("The error is ", err)
	}

	fmt.Println("Connection succesful!")

	// Trying to mimic a request just to see if the GORM code is working
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

	err = application.AddNewPersonInfo(&personTest)

	if err != nil {
		fmt.Println("Test Failed! A demain faut dormir")
	} else {
		fmt.Println("Person created successfully")
	}

}
