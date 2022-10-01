package main

import (
	"fmt"
<<<<<<< HEAD
=======

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/application/api"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/core/national_id_generator"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
>>>>>>> main

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

	if err != nil {
		fmt.Println("The error is ", err)
	}

	fmt.Println("Connection succesful!")

	// Trying to mimic a request just to see if the GORM code is working
<<<<<<< HEAD
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

	//err = postgres.AddNewPersonInfo(&personTest)
	postgres.UpdatePersonInfo("2")
	if err != nil {
		fmt.Println("Test Failed! A demain faut dormir")
	} else {
		fmt.Println("Person created successfully with UUID")
	}

	// 915aa63c-68e8-4c0f-b884-033b2673ab14
	postgres.CloseDBConnection()

=======
	//namesTest := pb.Names{
	//	Nom:         "Andrea",
	//	Prenom:      "Mufuta",
	//	MiddleNames: []string{"Mbuyi"},
	//}
	//biometricsTest := pb.Biometric{
	//	Photos:      []uint8{1, 2},
	//	FingerPrint: []uint8{2, 3},
	//}
	//
	//addressTest := pb.Address{
	//	Number:   1,
	//	Avenue:   "Ramba Place",
	//	Quartier: "Santa Clara",
	//}
	//
	//originTest := pb.Origin{
	//	Province: []string{"Bandundu"},
	//}
	//
	//phenotypeTest := pb.Phenotype{
	//	EyeColor: "blue",
	//}
	//
	//dateOfBirthTest := pb.DateOfBirth{
	//	Day:   "Monday",
	//	Month: "February",
	//	Year:  "2002",
	//}
	//
	//personTest := pb.PersonInfoRequest{
	//	Names:       &namesTest,
	//	Biometrics:  &biometricsTest,
	//	Address:     &addressTest,
	//	Origins:     &originTest,
	//	Phenotypes:  &phenotypeTest,
	//	DateOfBirth: &dateOfBirthTest,
	//}

	nationalID := &pb.NationalIDNumber{Id: "5f054e6280000de"}

	//err = application.AddNewPersonInfo(&personTest)
	person, err := application.FindPersonInfo(nationalID)

	if err != nil {
		fmt.Println("Test Failed! A demain faut dormir")
	} else {
		fmt.Println("---------------------------")
		fmt.Println(person.Names)
		fmt.Println("Person created successfully")
	}

>>>>>>> main
}
