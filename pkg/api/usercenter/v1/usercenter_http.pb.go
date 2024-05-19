// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.19.4
// source: usercenter/v1/usercenter.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserServiceCreateUser = "/usercenter.v1.UserService/CreateUser"
const OperationUserServiceGetUser = "/usercenter.v1.UserService/GetUser"
const OperationUserServiceListUser = "/usercenter.v1.UserService/ListUser"
const OperationUserServiceLogin = "/usercenter.v1.UserService/Login"
const OperationUserServiceLogout = "/usercenter.v1.UserService/Logout"

type UserServiceHTTPServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserReply, error)
	GetUser(context.Context, *GetUserRequest) (*UserReply, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/auth/login", _UserService_Login0_HTTP_Handler(srv))
	r.POST("/v1/auth/logout", _UserService_Logout0_HTTP_Handler(srv))
	r.POST("/v1/users", _UserService_CreateUser0_HTTP_Handler(srv))
	r.GET("/v1/users", _UserService_ListUser0_HTTP_Handler(srv))
	r.GET("/v1/users/{username}", _UserService_GetUser0_HTTP_Handler(srv))
}

func _UserService_Login0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _UserService_Logout0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _UserService_CreateUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/users/{username}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserResponse, error) {
	var out ListUserResponse
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/auth/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/auth/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
