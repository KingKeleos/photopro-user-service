package userservice

import (
	"context"
	"time"

	"github.com/KingKeleos/photopro-user-service/user"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/KingKeleos/photopro-user-service/grpc"
)

type UserServer struct {
}

// mustEmbedUnimplementedUserServiceServer implements photopro_user_service.UserServiceServer.
func (s UserServer) mustEmbedUnimplementedUserServiceServer() {
	panic("unimplemented")
}

func NewUserServer() UserServer {
	return UserServer{}
}

func (s UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	resp := pb.CreateUserResponse{}
	location := user.Location{
		Name:         *req.User.Location.Name,
		Street:       *req.User.Location.Street,
		StreetNumber: *req.User.Location.StreetNumber,
		Country:      *req.User.Location.Country,
		FederalState: *req.User.Location.FederalState,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	user := user.User{
		Username:  *req.User.Username,
		Password:  *req.User.Password,
		EMail:     *req.User.Email,
		Location:  location,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user.Create(ctx)

	return &resp, nil
}

// mustEmbedUnimplementedUserServiceServer implements photopro_user_service.UserServiceServer.
func (s UserServer) mustEmbedUnimplementedUserServiceServer() {
	panic("unimplemented")
}

// CreateRole implements photopro_user_service.UserServiceServer.
func (s UserServer) CreateRole(context.Context, *pb.CreateRolesRequest) (*pb.CreateRolesResponse, error) {
	panic("unimplemented")
}

// DeleteUser implements photopro_user_service.UserServiceServer.
func (s UserServer) DeleteUser(context.Context, *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	panic("unimplemented")
}

// GetUser implements photopro_user_service.UserServiceServer.
func (s UserServer) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	panic("unimplemented")
}

// ListPermissions implements photopro_user_service.UserServiceServer.
func (s UserServer) ListPermissions(context.Context, *pb.ListPermissionsRequest) (*pb.ListPermissionsResponse, error) {
	panic("unimplemented")
}

// ListRoles implements photopro_user_service.UserServiceServer.
func (s UserServer) ListRoles(context.Context, *emptypb.Empty) (*pb.ListRolesResponse, error) {
	panic("unimplemented")
}

// UpdatePermissions implements photopro_user_service.UserServiceServer.
func (s UserServer) UpdatePermissions(context.Context, *pb.UpdatePermissionsRequest) (*pb.UpdatePermissionsResponse, error) {
	panic("unimplemented")
}

// UpdateRoles implements photopro_user_service.UserServiceServer.
func (s UserServer) UpdateRoles(context.Context, *pb.UpdateRolesRequest) (*pb.UpdateRolesResponse, error) {
	panic("unimplemented")
}

// UpdateUser implements photopro_user_service.UserServiceServer.
func (s UserServer) UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	panic("unimplemented")
}
