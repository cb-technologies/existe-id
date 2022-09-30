package ports

import (
	"context"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
)

type GRPCPort interface {
	Run()
	AddNewPersonInfo(ctx context.Context, req *pb.PersonInfoRequest) (*pb.Response, error)
	UpdatePersonInfo(ctx context.Context, req *pb.EditPersonInfoParameters) (*pb.Response, error)
	FindPersonInfo(ctx context.Context, req *pb.NationalIDNumber) (*pb.PersonInfoResponse, error)
}
