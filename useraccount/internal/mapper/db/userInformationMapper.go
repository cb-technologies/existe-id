package db

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
)

func protoAddressToDBAddress(address *pb.Address) *db.Address {
	return &db.Address{
		Number:    address.Number,
		Avenue:    address.Avenue,
		Quartier:  address.Quartier,
		Commune:   address.Commune,
		ZipCode:   address.ZipCode,
		Province:  address.Province,
		Ville:     address.Ville,
		Reference: address.Reference,
	}
}

func protoNamesToDBNames(names *pb.Names) *db.Names {
	return &db.Names{
		Nom:         names.Nom,
		Prenom:      names.Prenom, // TODO: fix to prenom
		MiddleNames: names.MiddleNames,
	}
}

func protoOriginToDBOrigin(origin *pb.Origin) *db.Origin {
	return &db.Origin{
		Province: origin.Province,
		ChefLieu: origin.ChefLieu,
	}
}

func protoPhenotypeToDBPhenotype(phenotype *pb.Phenotype) *db.Phenotype {
	return &db.Phenotype{
		Height:   phenotype.Height,
		Weight:   phenotype.Weight,
		EyeColor: phenotype.EyeColor,
	}
}

func protoBiometricToDBBiometric(biometric *pb.Biometric) *db.Biometric {
	return &db.Biometric{
		Photos:      biometric.Photos,
		FingerPrint: biometric.FingerPrint,
	}
}

func protoDateOfBirthToDBDateOfBirth(dateofBirth *pb.DateOfBirth) *db.DateOfBirth {
	return &db.DateOfBirth{
		Day:   dateofBirth.Day,
		Month: dateofBirth.Month,
		Year:  dateofBirth.Year,
	}
}

func protoPersonInfoRequestToDBPersonInfoRequest(personInfo *pb.PersonInfoRequest) *db.PersonInfoRequest {
	return &db.PersonInfoRequest{
		Names:       protoNamesToDBNames(personInfo.Names),
		Biometrics:  protoBiometricToDBBiometric(personInfo.Biometrics),
		Address:     protoAddressToDBAddress(personInfo.Address),
		Origins:     protoOriginToDBOrigin(personInfo.Origins),
		Phenotypes:  protoPhenotypeToDBPhenotype(personInfo.Phenotypes),
		DateOfBirth: protoDateOfBirthToDBDateOfBirth(personInfo.DateOfBirth),
	}
}

func dbCardValidityToProotCardValidity(cardValidity *db.CardValidity) *pb.CardValidity {
	return &pb.CardValidity{
		IssueDate: &pb.Date{
			Year:  cardValidity.IssueDate.Year,
			Month: cardValidity.IssueDate.Month,
			Day:   cardValidity.IssueDate.Day,
		},
		ExpirationDate: &pb.Date{
			Year:  cardValidity.ExpirationDate.Year,
			Month: cardValidity.ExpirationDate.Month,
			Day:   cardValidity.ExpirationDate.Day,
		},
	}
}
