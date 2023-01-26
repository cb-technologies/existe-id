package dbmapper

import (
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
)

func protoAddressToDBAddress(address *pb.Address) *db.AddressModel {
	if address == nil {
		return &db.AddressModel{}
	}
	return &db.AddressModel{
		Number:          address.Number,
		Avenue:          address.Avenue,
		Quartier:        address.Quartier,
		Commune:         address.Commune,
		ZipCode:         address.ZipCode,
		Ville:           address.Ville,
		Reference:       address.Reference,
		ProvinceAddress: address.Province,
	}
}

func addressModelToProtoAddress(addressModel *db.AddressModel) *pb.Address {
	if addressModel == nil {
		return &pb.Address{}
	}
	return &pb.Address{
		Number:    addressModel.Number,
		Avenue:    addressModel.Avenue,
		Quartier:  addressModel.Quartier,
		Commune:   addressModel.Commune,
		ZipCode:   addressModel.ZipCode,
		Ville:     addressModel.Ville,
		Reference: addressModel.Reference,
		Province:  addressModel.ProvinceAddress,
	}
}

func ProtoNamesToDBNames(names *pb.Names) *db.NamesModel {
	if names == nil {
		return &db.NamesModel{}
	}
	return &db.NamesModel{
		Nom:         names.Nom,
		Prenom:      names.Prenom,
		MiddleNames: names.MiddleNames,
	}
}

func namesModelToProtoNames(namesModel *db.NamesModel) *pb.Names {
	if namesModel == nil {
		return &pb.Names{}
	}
	return &pb.Names{
		Nom:         namesModel.Nom,
		Prenom:      namesModel.Prenom,
		MiddleNames: namesModel.MiddleNames,
	}
}

func protoOriginToDBOrigin(origin *pb.Origin) *db.OriginModel {
	if origin == nil {
		return &db.OriginModel{}
	}
	return &db.OriginModel{
		Province:        origin.Province,
		ChefLieu:        origin.ChefLieu,
		LieuDeNaissance: origin.LieuDeNaissance,
	}
}

func originModelToProtoOrigin(originModel *db.OriginModel) *pb.Origin {
	if originModel == nil {
		return &pb.Origin{}
	}
	return &pb.Origin{
		Province:        originModel.Province,
		ChefLieu:        originModel.ChefLieu,
		LieuDeNaissance: originModel.LieuDeNaissance,
	}
}

func protoPhenotypeToDBPhenotype(phenotype *pb.Phenotype) *db.PhenotypeModel {
	if phenotype == nil {
		return &db.PhenotypeModel{}
	}
	return &db.PhenotypeModel{
		Height:   phenotype.Height,
		Weight:   phenotype.Weight,
		EyeColor: phenotype.EyeColor,
	}
}

func phenotypeModelToProtoPhenotype(phenotypeModel *db.PhenotypeModel) *pb.Phenotype {
	if phenotypeModel == nil {
		return &pb.Phenotype{}
	}
	return &pb.Phenotype{
		Height:   phenotypeModel.Height,
		Weight:   phenotypeModel.Weight,
		EyeColor: phenotypeModel.EyeColor,
	}
}

func protoBiometricToDBBiometric(biometric *pb.Biometric) *db.BiometricModel {
	if biometric == nil {
		return &db.BiometricModel{}
	}
	return &db.BiometricModel{
		Photos:      biometric.Photos,
		FingerPrint: biometric.FingerPrint,
		PhotoType:   biometric.PhotoType,
	}
}

func biometricModelToProtoBiometric(biometricModel *db.BiometricModel) *pb.Biometric {
	if biometricModel == nil {
		return &pb.Biometric{}
	}
	return &pb.Biometric{
		Photos:      biometricModel.Photos,
		FingerPrint: biometricModel.FingerPrint,
		PhotoType:   biometricModel.PhotoType,
	}
}

func protoSexToDBSex(sex *pb.Sex) *db.Sex {
	if sex == nil {
		return &db.Sex{}
	}
	return &db.Sex{
		Sex: sex.Sex.String(),
	}
}

func sexModelToProtoSex(sexModel *db.Sex) *pb.Sex {
	if sexModel == nil {
		return &pb.Sex{}
	}
	return &pb.Sex{
		Sex: pb.SexEnum(pb.SexEnum_value[sexModel.Sex]),
	}
}

func ProtoDateOfBirthToDBDateOfBirth(dateofBirth *pb.DateOfBirth) *db.DateOfBirthModel {
	if dateofBirth == nil {
		return &db.DateOfBirthModel{}
	}
	return &db.DateOfBirthModel{
		Day:   dateofBirth.Day,
		Month: dateofBirth.Month,
		Year:  dateofBirth.Year,
	}
}

func dateOfBirthModelToProtoDateOfBirth(dateOfBirthModel *db.DateOfBirthModel) *pb.DateOfBirth {
	if dateOfBirthModel == nil {
		return &pb.DateOfBirth{}
	}
	return &pb.DateOfBirth{
		Day:   dateOfBirthModel.Day,
		Month: dateOfBirthModel.Month,
		Year:  dateOfBirthModel.Year,
	}
}

func cardValidityModelToProtoCardValidity(cardValidityModel *db.CardValidityModel) *pb.CardValidity {
	if cardValidityModel == nil {
		return &pb.CardValidity{}
	}
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

func protoQRCodeToQRCodeModel(qrcode *pb.QRCode) *db.QRCodeModel {
	if qrcode == nil {
		return &db.QRCodeModel{}
	}
	return &db.QRCodeModel{
		QRCode: qrcode.Qrcode,
	}
}

func QRCodeModelToProtoQRCode(qrcode *db.QRCodeModel) *pb.QRCode {
	if qrcode == nil {
		return &pb.QRCode{}
	}
	return &pb.QRCode{
		Qrcode: qrcode.QRCode,
	}
}

func PersonInfoRequestToPersonInfoModel(personInfoModel *pb.PersonInfoRequest) *db.PersonInfoModel {
	if personInfoModel == nil {
		return &db.PersonInfoModel{}
	}
	return &db.PersonInfoModel{
		Names:       *ProtoNamesToDBNames(personInfoModel.Names),
		Biometrics:  *protoBiometricToDBBiometric(personInfoModel.Biometrics),
		Address:     *protoAddressToDBAddress(personInfoModel.Address),
		Origins:     *protoOriginToDBOrigin(personInfoModel.Origins),
		Phenotypes:  *protoPhenotypeToDBPhenotype(personInfoModel.Phenotypes),
		DateOfBirth: *ProtoDateOfBirthToDBDateOfBirth(personInfoModel.DateOfBirth),
		Sex:         *protoSexToDBSex(personInfoModel.Sex),
		QRCode:      *protoQRCodeToQRCodeModel(personInfoModel.Qrcode),
	}
}

func nationaIDNumberModelToProtoNationalID(nationalID *db.NationalIDNumberModel) *pb.NationalIDNumber {
	if nationalID == nil {
		return &pb.NationalIDNumber{}
	}
	return &pb.NationalIDNumber{
		Id: nationalID.NationalID,
	}
}

func ProtoNationalIDNumberToNationalIDNumberModel(nationalID *pb.NationalIDNumber) *db.NationalIDNumberModel {
	if nationalID == nil {
		return &db.NationalIDNumberModel{}
	}
	return &db.NationalIDNumberModel{
		NationalID: nationalID.Id,
	}
}

func PersonInfoModelToPersonInfoResponse(personInfoModel *db.PersonInfoModel) *pb.PersonInfoResponse {
	if personInfoModel == nil {
		return &pb.PersonInfoResponse{}
	}
	return &pb.PersonInfoResponse{
		Names:        namesModelToProtoNames(&personInfoModel.Names),
		Biometrics:   biometricModelToProtoBiometric(&personInfoModel.Biometrics),
		Address:      addressModelToProtoAddress(&personInfoModel.Address),
		Origins:      originModelToProtoOrigin(&personInfoModel.Origins),
		Phenotypes:   phenotypeModelToProtoPhenotype(&personInfoModel.Phenotypes),
		DateOfBirth:  dateOfBirthModelToProtoDateOfBirth(&personInfoModel.DateOfBirth),
		CardValidity: cardValidityModelToProtoCardValidity(&personInfoModel.CardValidity),
		Id:           nationaIDNumberModelToProtoNationalID(&personInfoModel.NationalID),
		Sex:          sexModelToProtoSex(&personInfoModel.Sex),
		Qrcode:       QRCodeModelToProtoQRCode(&personInfoModel.QRCode),
	}
}
