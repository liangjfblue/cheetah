// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api.proto

//protoc --proto_path=. --micro_out=. --go_out=. ./api.proto

package micro_srv_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterRespond, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginRespond, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetRespond, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListRespond, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "micro.srv.auth"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterRespond, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(RegisterRespond)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginRespond, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(LoginRespond)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetRespond, error) {
	req := c.c.NewRequest(c.name, "User.Get", in)
	out := new(GetRespond)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListRespond, error) {
	req := c.c.NewRequest(c.name, "User.List", in)
	out := new(ListRespond)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Register(context.Context, *RegisterRequest, *RegisterRespond) error
	Login(context.Context, *LoginRequest, *LoginRespond) error
	Get(context.Context, *GetRequest, *GetRespond) error
	List(context.Context, *ListRequest, *ListRespond) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Register(ctx context.Context, in *RegisterRequest, out *RegisterRespond) error
		Login(ctx context.Context, in *LoginRequest, out *LoginRespond) error
		Get(ctx context.Context, in *GetRequest, out *GetRespond) error
		List(ctx context.Context, in *ListRequest, out *ListRespond) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterRespond) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *LoginRequest, out *LoginRespond) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) Get(ctx context.Context, in *GetRequest, out *GetRespond) error {
	return h.UserHandler.Get(ctx, in, out)
}

func (h *userHandler) List(ctx context.Context, in *ListRequest, out *ListRespond) error {
	return h.UserHandler.List(ctx, in, out)
}
