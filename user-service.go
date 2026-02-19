package userservice

import (
	"context"
	"log/slog"
	"time"

	"github.com/KingKeleos/photopro-user-service/user"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/KingKeleos/photopro-user-service/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServer() UserServer {
	return UserServer{}
}

func (s UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
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
	err := user.Create(ctx)

	resp := pb.CreateUserResponse{
		User: &pb.User{
			Username: &user.Username,
			Email:    &user.EMail,
			Password: &user.Password,
			Location: &pb.Location{
				Name:         &location.Name,
				Street:       &location.Street,
				StreetNumber: &location.StreetNumber,
				Country:      &location.Country,
				FederalState: &location.FederalState,
			},
		}}
	return &resp, err
}

// mustEmbedUnimplementedUserServiceServer implements photopro_user_service.UserServiceServer.
func (s UserServer) mustEmbedUnimplementedUserServiceServer() {
	panic("unimplemented")
}

// CreateRole implements photopro_user_service.UserServiceServer.
func (s UserServer) CreateRole(ctx context.Context, req *pb.CreateRolesRequest) (*pb.CreateRolesResponse, error) {
	panic("unimplemented")
}

// DeleteUser implements photopro_user_service.UserServiceServer.
func (s UserServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	user := user.User{
		ID: req.GetUserID(),
	}
	err := user.Delete(ctx)
	resp := pb.DeleteUserResponse{
		Username: &user.Username,
	}
	return &resp, err
}

// GetUser implements photopro_user_service.UserServiceServer.
func (s UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	panic("unimplemented")
}

// ListPermissions implements photopro_user_service.UserServiceServer.
func (s UserServer) ListPermissions(ctx context.Context, req *pb.ListPermissionsRequest) (*pb.ListPermissionsResponse, error) {
	panic("unimplemented")
}

// ListRoles implements photopro_user_service.UserServiceServer.
func (s UserServer) ListRoles(ctx context.Context, req *emptypb.Empty) (*pb.ListRolesResponse, error) {
	panic("unimplemented")
}

// UpdatePermissions implements photopro_user_service.UserServiceServer.
func (s UserServer) UpdatePermissions(ctx context.Context, req *pb.UpdatePermissionsRequest) (*pb.UpdatePermissionsResponse, error) {
	panic("unimplemented")
}

// UpdateRoles implements photopro_user_service.UserServiceServer.
func (s UserServer) UpdateRoles(context.Context, *pb.UpdateRolesRequest) (*pb.UpdateRolesResponse, error) {
	panic("unimplemented")
}

// UpdateUser implements photopro_user_service.UserServiceServer.
func (s UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	location := user.Location{
		Name:         *req.Location.Name,
		Street:       *req.Location.Street,
		StreetNumber: *req.Location.StreetNumber,
		Country:      *req.Location.Country,
		FederalState: *req.Location.FederalState,
		UpdatedAt:    time.Now(),
	}
	user := user.User{
		ID:        *req.UserID,
		Password:  *req.Password,
		EMail:     *req.Email,
		Location:  location,
		UpdatedAt: time.Now(),
	}

	err := user.Update(ctx)
	if err != nil {
		slog.Error("updating user", "error", err)
	}
	resp := pb.UpdateUserResponse{
		User: &pb.User{
			Email:    &user.EMail,
			Password: &user.Password,
			Location: &pb.Location{
				Name:         &location.Name,
				Street:       &location.Street,
				StreetNumber: &location.StreetNumber,
				Country:      &location.Country,
				FederalState: &location.FederalState,
			},
		},
	}
	return &resp, err
}
