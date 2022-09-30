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

	// // Trying to mimic a request just to see if the GORM code is working
	// namesTest := pb.Names{
	// 	Nom:         "Nicolas",
	// 	Prenom:      "Nkiere",
	// 	MiddleNames: []string{"Bamanissa"},
	// }
	// biometricsTest := pb.Biometric{
	// 	Photos:      []uint8{1, 2},
	// 	FingerPrint: []uint8{2, 3},
	// }

	// addressTest := pb.Address{
	// 	Number:   1,
	// 	Avenue:   "Nicolas",
	// 	Quartier: "Santa Clara",
	// }

	// originTest := pb.Origin{
	// 	Province: []string{"Bandundu"},
	// }

	// phenotypeTest := pb.Phenotype{
	// 	EyeColor: "blue",
	// }

	// dateOfBirthTest := pb.DateOfBirth{
	// 	Day:   "Monday",
	// 	Month: "February",
	// 	Year:  "2002",
	// }

	// personTest := pb.PersonInfoRequest{
	// 	Names:       &namesTest,
	// 	Biometrics:  &biometricsTest,
	// 	Address:     &addressTest,
	// 	Origins:     &originTest,
	// 	Phenotypes:  &phenotypeTest,
	// 	DateOfBirth: &dateOfBirthTest,
	// }

	// err = postgres.AddNewPersonInfo(&personTest)

	nationalID := pb.NationalIDNumber{Id: "2"}
	_, err = postgres.FindPersonInfo(&nationalID)

	if err != nil {
		fmt.Println("Test Failed! A demain faut dormir")
	} else {
		fmt.Println("---------------------------")

		fmt.Println("Person created successfully")
	}

	postgres.CloseDBConnection()

}
