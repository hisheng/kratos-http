package service

import (
	"context"
	"github.com/hisheng/kratos-http/api"
)

type UserService struct {
}

func (svc UserService) ListUser(context.Context, *api.ListUserRequest) (*api.ListUserReply, error) {
	return &api.ListUserReply{Id: 100}, nil
}

func (svc UserService) CreateUser(context.Context, *api.CreateUserRequest) (*api.CreateUserReply, error) {
	return nil, nil
}
