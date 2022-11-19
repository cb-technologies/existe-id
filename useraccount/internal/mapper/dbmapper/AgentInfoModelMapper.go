package dbmapper

import (
	"log"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/core"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/entity/db"
)

func ProtoAgentSignInInfoToAgentSignInInfoModel(agentSignInInfo *pb.AgentSignInInfo) *db.AgentSignInModel {
	return &db.AgentSignInModel{
		Email:    agentSignInInfo.Email,
		Password: agentSignInInfo.Password,
	}
}

func AgentSignInInfoModelToAgentSignInInfoProto(agentSignInInfoModel *db.AgentSignInModel) *pb.AgentSignInInfo {
	return &pb.AgentSignInInfo{
		Email:    agentSignInInfoModel.Email,
		Password: agentSignInInfoModel.Password,
	}
}

func ProtoAgentInfoToAgentInfoModel(agentInfo *pb.AgentInfo) *db.AgentInfoModel {
	hashedPassword, err := core.HashPassword(agentInfo.Password)
	if err != nil {
		log.Printf("Failed to hash the password %v", err)
	}

	return &db.AgentInfoModel{
		Nom:      agentInfo.Nom,
		Prenom:   agentInfo.Prenom,
		Email:    agentInfo.Email,
		Password: hashedPassword,
	}
}
