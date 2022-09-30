package dbmapper

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
)

func protoAddressToDBAddress(address *pb.Address) *db.AddressModel {
	return &db.AddressModel{
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

func addressModelToProtoAddress(addressModel *db.AddressModel) *pb.Address {
	return &pb.Address{
		Number:    addressModel.Number,
		Avenue:    addressModel.Avenue,
		Quartier:  addressModel.Quartier,
		Commune:   addressModel.Commune,
		ZipCode:   addressModel.ZipCode,
		Province:  addressModel.Province,
		Ville:     addressModel.Ville,
		Reference: addressModel.Reference,
	}
}

func protoNamesToDBNames(names *pb.Names) *db.NamesModel {
	return &db.NamesModel{
		Nom:         names.Nom,
		Prenom:      names.Prenom, // TODO: fix to prenom
		MiddleNames: names.MiddleNames,
	}
}

func namesModelToProtoNames(namesModel *db.NamesModel) *pb.Names {
	return &pb.Names{
		Nom:         namesModel.Nom,
		Prenom:      namesModel.Prenom,
		MiddleNames: namesModel.MiddleNames,
	}
}

func protoOriginToDBOrigin(origin *pb.Origin) *db.OriginModel {
	return &db.OriginModel{
		Province: origin.Province,
		ChefLieu: origin.ChefLieu,
	}
}

func originModelToProtoOrigin(originModel *db.OriginModel) *pb.Origin {
	return &pb.Origin{
		Province: originModel.Province,
		ChefLieu: originModel.ChefLieu,
	}
}

func protoPhenotypeToDBPhenotype(phenotype *pb.Phenotype) *db.PhenotypeModel {
	return &db.PhenotypeModel{
		Height:   phenotype.Height,
		Weight:   phenotype.Weight,
		EyeColor: phenotype.EyeColor,
	}
}

func phenotypeModelToProtoPhenotype(phenotypeModel *db.PhenotypeModel) *pb.Phenotype {
	return &pb.Phenotype{
		Height:   phenotypeModel.Height,
		Weight:   phenotypeModel.Weight,
		EyeColor: phenotypeModel.EyeColor,
	}
}

func protoBiometricToDBBiometric(biometric *pb.Biometric) *db.BiometricModel {
	return &db.BiometricModel{
		Photos:      biometric.Photos,
		FingerPrint: biometric.FingerPrint,
	}
}

func biometricModelToProtoBiometric(biometricModel *db.BiometricModel) *pb.Biometric {
	return &pb.Biometric{
		Photos:      biometricModel.Photos,
		FingerPrint: biometricModel.FingerPrint,
	}
}

func protoDateOfBirthToDBDateOfBirth(dateofBirth *pb.DateOfBirth) *db.DateOfBirthModel {
	return &db.DateOfBirthModel{
		Day:   dateofBirth.Day,
		Month: dateofBirth.Month,
		Year:  dateofBirth.Year,
	}
}

func dateOfBirthModelToProtoDateOfBirth(dateOfBirthModel *db.DateOfBirthModel) *pb.DateOfBirth {
	return &pb.DateOfBirth{
		Day:   dateOfBirthModel.Day,
		Month: dateOfBirthModel.Month,
		Year:  dateOfBirthModel.Year,
	}
}

func cardValidityModelToProtoCardValidity(cardValidityModel *db.CardValidityModel) *pb.CardValidity {
	return &pb.CardValidity{
		IssueDate: &pb.Date{
			Year:  cardValidityModel.CardIssueDate.Year,
			Month: cardValidityModel.CardIssueDate.Month,
			Day:   cardValidityModel.CardIssueDate.Day,
		},
		ExpirationDate: &pb.Date{
			Year:  cardValidityModel.CardExpirationDate.Year,
			Month: cardValidityModel.CardExpirationDate.Month,
			Day:   cardValidityModel.CardExpirationDate.Day,
		},
	}
}

func PersonInfoRequestToPersonInfoModel(personInfoModel *pb.PersonInfoRequest) *db.PersonInfoModel {
	return &db.PersonInfoModel{
		Names:       *protoNamesToDBNames(personInfoModel.Names),
		Biometrics:  *protoBiometricToDBBiometric(personInfoModel.Biometrics),
		Address:     *protoAddressToDBAddress(personInfoModel.Address),
		Origins:     *protoOriginToDBOrigin(personInfoModel.Origins),
		Phenotypes:  *protoPhenotypeToDBPhenotype(personInfoModel.Phenotypes),
		DateOfBirth: *protoDateOfBirthToDBDateOfBirth(personInfoModel.DateOfBirth),
	}
}

func nationaIDNumberModelToProtoNationalID(nationalID *db.NationalIDNumberModel) *pb.NationalIDNumber {
	return &pb.NationalIDNumber{
		Id: nationalID.Id,
	}
}

func PersonInfoModelToPersonInfoResponse(personInfoModel *db.PersonInfoModel) *pb.PersonInfoResponse {
	return &pb.PersonInfoResponse{
		Names:        namesModelToProtoNames(&personInfoModel.Names),
		Biometrics:   biometricModelToProtoBiometric(&personInfoModel.Biometrics),
		Address:      addressModelToProtoAddress(&personInfoModel.Address),
		Origins:      originModelToProtoOrigin(&personInfoModel.Origins),
		Phenotypes:   phenotypeModelToProtoPhenotype(&personInfoModel.Phenotypes),
		DateOfBirth:  dateOfBirthModelToProtoDateOfBirth(&personInfoModel.DateOfBirth),
		CardValidity: cardValidityModelToProtoCardValidity(&personInfoModel.CardValidity),
		Id:           nationaIDNumberModelToProtoNationalID(&personInfoModel.NationalID),
	}
}
