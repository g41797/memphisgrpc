package plain

import "context"

type ManagerServiceClient interface {
	Manage(ctx context.Context, in *ManageRequest) (*Status, error)
}

type ManagerServiceServer interface {
	Manage(ctx context.Context, req *ManageRequest) (*Status, error)
}