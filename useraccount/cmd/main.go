package main

import (
	"fmt"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/application/api"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/core/national_id_generator"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driven/postgresSQL"
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
	//my_postgres, err := postgresSQL.NewAdapter()
	if err != nil {
		fmt.Println("The error is ", err)
	}

	fmt.Println("Connection succesful!")

	// Add a new User

	namesTest := pb.Names{
		Nom:         "Elie",
		Prenom:      "Masanka",
		MiddleNames: []string{"Ntumba"},
	}
	biometricsTest := pb.Biometric{
		Photos:      []uint8{1, 2},
		FingerPrint: []uint8{2, 3},
	}

	addressTest := pb.Address{
		Number:   1,
		Avenue:   "Boma",
		Quartier: "OUA",
		Commune:  "Kintambo",
		ZipCode:  "1430",
	}

	phenotypeTest := pb.Phenotype{
		EyeColor: "brown",
	}

	originTest := pb.Origin{
		Province: []string{"BasCongo"},
		ChefLieu: "kibasa",
	}

	dateOfBirthTest := pb.DateOfBirth{
		Day:   "23",
		Month: "March",
		Year:  "1998",
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
		fmt.Println("Error Adding a User")
	} else {
		fmt.Println("---------------------------")
		fmt.Println("Person created successfully")
	}

	//Find a User
	//nationalID := &pb.NationalIDNumber{Id: "5f7ec5a5300002e"}
	//person, err := application.FindPersonInfo(nationalID)
	//if err != nil {
	//	fmt.Println("Error")
	//} else {
	//	fmt.Println("---------------------------")
	//	fmt.Println(person)
	//	fmt.Println("Person found successfully")
	//}

	// Update a User

	//namesTest := pb.Names{
	//	Nom:         "Elie",
	//	Prenom:      "Masanka",
	//	MiddleNames: []string{"Ntumba"},
	//}
	//biometricsTest := pb.Biometric{
	//	Photos:      []uint8{1, 2},
	//	FingerPrint: []uint8{2, 3},
	//}
	//
	//addressTest := pb.Address{
	//	Number:   1,
	//	Avenue:   "XXXXX",
	//	Quartier: "BBBB",
	//	Commune:  "Kintambo",
	//	ZipCode:  "1430",
	//}
	//
	//phenotypeTest := pb.Phenotype{
	//	EyeColor: "brown",
	//}
	//
	//dateOfBirthTest := pb.DateOfBirth{
	//	Day:   "23",
	//	Month: "March",
	//	Year:  "1998",
	//}
	//
	//personTest := pb.PersonInfoRequest{
	//	Names:       &namesTest,
	//	Biometrics:  &biometricsTest,
	//	Address:     &addressTest,
	//	Phenotypes:  &phenotypeTest,
	//	DateOfBirth: &dateOfBirthTest,
	//}
	//nationalID := &pb.NationalIDNumber{Id: "5f7ec5a5300002e"}
	//
	//newPersonTest := pb.EditPersonInfoParameters{
	//	PersonId:         nationalID,
	//	EditedPersonInfo: &personTest,
	//}
	//find_err := application.UpdatePersonInfo(&newPersonTest)
	//if find_err != nil {
	//	fmt.Println("Update Failing")
	//} else {
	//	fmt.Println("Success updating the person information")
	//}

}
