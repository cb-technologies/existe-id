package main

import (
	"fmt"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"
)

func main() {

	fmt.Println("Existe ID")

	postgres, err := postgresSQL.NewAdapter()

	if err != nil {
		fmt.Println("The error is ", err)
	}

	fmt.Println("Connection succesful!")

	// Trying to mimic a request just to see if the GORM code is working
	namesTest := pb.Names{
		Nom:         "Bingoto",
		Prenom:      "Areely",
		MiddleNames: []string{"Bamanissa"},
	}
	biometricsTest := pb.Biometric{
		Photos:      []uint8{1, 2},
		FingerPrint: []uint8{2, 3},
	}

	addressTest := pb.Address{
		Number:   1,
		Avenue:   "Nicolas",
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

	my_id := "2"
	natinalId := pb.NationalIDNumber{
		Id: my_id,
	}
	err = postgres.AddNewPersonInfo(&personTest)
	postgres.UpdatePersonInfo(&natinalId)
	if err != nil {
		fmt.Println("Test Failed! A demain faut dormir")
	} else {
		fmt.Println("Person created successfully")
	}

	postgres.CloseDBConnection()

}
