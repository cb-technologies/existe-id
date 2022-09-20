package ports

import (
	"context"

	"github.com/cb-technologies/existe-id/useraccount/useraccount/internal/adapters/framework/driver/grpc/pb"
)

type GRPCPort interface {
	Run()
	AddNewPersonInfo(ctx context.Context, req *pb.PersonInfo) (*pb.Response, error)
	UpdatePersonInfo(ctx context.Context, req *pb.EditPersonInfoParameters) (*pb.Response, error)
	FindPersonInfo(ctx context.Context, req *pb.UUID) (*pb.PersonInfo, error)
}
