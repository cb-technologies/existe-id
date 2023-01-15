package serverGrpc

import (
	"context"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/constants/serviceresponse"
)

func (adapter Adapter) SignInAgent(ctx context.Context, agentSignInInfo *pb.AgentSignInInfo) (*pb.Response, error) {
	err := adapter.api.SignInAgent(agentSignInInfo)

	if err != nil {
		return &pb.Response{Status: serviceresponse.FAILURE}, err
	}

	return &pb.Response{Status: serviceresponse.SUCCESS}, err
}

func (adapter Adapter) SignUpAgent(ctx context.Context, agentInfo *pb.AgentInfo) (*pb.Response, error) {
	print("Backend ")
	print(agentInfo.Role)
	err := adapter.api.SignUpAgent(agentInfo)

	if err != nil {
		return &pb.Response{Status: serviceresponse.FAILURE}, err
	}

	return &pb.Response{Status: serviceresponse.SUCCESS}, err
}
