// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: auth/v1/auth.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/RA341/multipacman/generated/auth/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "auth.v1.AuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthServiceLoginProcedure is the fully-qualified name of the AuthService's Login RPC.
	AuthServiceLoginProcedure = "/auth.v1.AuthService/Login"
	// AuthServiceRegisterProcedure is the fully-qualified name of the AuthService's Register RPC.
	AuthServiceRegisterProcedure = "/auth.v1.AuthService/Register"
	// AuthServiceTestProcedure is the fully-qualified name of the AuthService's Test RPC.
	AuthServiceTestProcedure = "/auth.v1.AuthService/Test"
	// AuthServiceGuestLoginProcedure is the fully-qualified name of the AuthService's GuestLogin RPC.
	AuthServiceGuestLoginProcedure = "/auth.v1.AuthService/GuestLogin"
	// AuthServiceLogoutProcedure is the fully-qualified name of the AuthService's Logout RPC.
	AuthServiceLogoutProcedure = "/auth.v1.AuthService/Logout"
)

// AuthServiceClient is a client for the auth.v1.AuthService service.
type AuthServiceClient interface {
	Login(context.Context, *connect.Request[v1.AuthRequest]) (*connect.Response[v1.UserResponse], error)
	Register(context.Context, *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error)
	Test(context.Context, *connect.Request[v1.AuthResponse]) (*connect.Response[v1.UserResponse], error)
	GuestLogin(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.UserResponse], error)
	Logout(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error)
}

// NewAuthServiceClient constructs a client for the auth.v1.AuthService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	authServiceMethods := v1.File_auth_v1_auth_proto.Services().ByName("AuthService").Methods()
	return &authServiceClient{
		login: connect.NewClient[v1.AuthRequest, v1.UserResponse](
			httpClient,
			baseURL+AuthServiceLoginProcedure,
			connect.WithSchema(authServiceMethods.ByName("Login")),
			connect.WithClientOptions(opts...),
		),
		register: connect.NewClient[v1.RegisterUserRequest, v1.RegisterUserResponse](
			httpClient,
			baseURL+AuthServiceRegisterProcedure,
			connect.WithSchema(authServiceMethods.ByName("Register")),
			connect.WithClientOptions(opts...),
		),
		test: connect.NewClient[v1.AuthResponse, v1.UserResponse](
			httpClient,
			baseURL+AuthServiceTestProcedure,
			connect.WithSchema(authServiceMethods.ByName("Test")),
			connect.WithClientOptions(opts...),
		),
		guestLogin: connect.NewClient[v1.Empty, v1.UserResponse](
			httpClient,
			baseURL+AuthServiceGuestLoginProcedure,
			connect.WithSchema(authServiceMethods.ByName("GuestLogin")),
			connect.WithClientOptions(opts...),
		),
		logout: connect.NewClient[v1.Empty, v1.Empty](
			httpClient,
			baseURL+AuthServiceLogoutProcedure,
			connect.WithSchema(authServiceMethods.ByName("Logout")),
			connect.WithClientOptions(opts...),
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	login      *connect.Client[v1.AuthRequest, v1.UserResponse]
	register   *connect.Client[v1.RegisterUserRequest, v1.RegisterUserResponse]
	test       *connect.Client[v1.AuthResponse, v1.UserResponse]
	guestLogin *connect.Client[v1.Empty, v1.UserResponse]
	logout     *connect.Client[v1.Empty, v1.Empty]
}

// Login calls auth.v1.AuthService.Login.
func (c *authServiceClient) Login(ctx context.Context, req *connect.Request[v1.AuthRequest]) (*connect.Response[v1.UserResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// Register calls auth.v1.AuthService.Register.
func (c *authServiceClient) Register(ctx context.Context, req *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// Test calls auth.v1.AuthService.Test.
func (c *authServiceClient) Test(ctx context.Context, req *connect.Request[v1.AuthResponse]) (*connect.Response[v1.UserResponse], error) {
	return c.test.CallUnary(ctx, req)
}

// GuestLogin calls auth.v1.AuthService.GuestLogin.
func (c *authServiceClient) GuestLogin(ctx context.Context, req *connect.Request[v1.Empty]) (*connect.Response[v1.UserResponse], error) {
	return c.guestLogin.CallUnary(ctx, req)
}

// Logout calls auth.v1.AuthService.Logout.
func (c *authServiceClient) Logout(ctx context.Context, req *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error) {
	return c.logout.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the auth.v1.AuthService service.
type AuthServiceHandler interface {
	Login(context.Context, *connect.Request[v1.AuthRequest]) (*connect.Response[v1.UserResponse], error)
	Register(context.Context, *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error)
	Test(context.Context, *connect.Request[v1.AuthResponse]) (*connect.Response[v1.UserResponse], error)
	GuestLogin(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.UserResponse], error)
	Logout(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authServiceMethods := v1.File_auth_v1_auth_proto.Services().ByName("AuthService").Methods()
	authServiceLoginHandler := connect.NewUnaryHandler(
		AuthServiceLoginProcedure,
		svc.Login,
		connect.WithSchema(authServiceMethods.ByName("Login")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceRegisterHandler := connect.NewUnaryHandler(
		AuthServiceRegisterProcedure,
		svc.Register,
		connect.WithSchema(authServiceMethods.ByName("Register")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceTestHandler := connect.NewUnaryHandler(
		AuthServiceTestProcedure,
		svc.Test,
		connect.WithSchema(authServiceMethods.ByName("Test")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceGuestLoginHandler := connect.NewUnaryHandler(
		AuthServiceGuestLoginProcedure,
		svc.GuestLogin,
		connect.WithSchema(authServiceMethods.ByName("GuestLogin")),
		connect.WithHandlerOptions(opts...),
	)
	authServiceLogoutHandler := connect.NewUnaryHandler(
		AuthServiceLogoutProcedure,
		svc.Logout,
		connect.WithSchema(authServiceMethods.ByName("Logout")),
		connect.WithHandlerOptions(opts...),
	)
	return "/auth.v1.AuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthServiceLoginProcedure:
			authServiceLoginHandler.ServeHTTP(w, r)
		case AuthServiceRegisterProcedure:
			authServiceRegisterHandler.ServeHTTP(w, r)
		case AuthServiceTestProcedure:
			authServiceTestHandler.ServeHTTP(w, r)
		case AuthServiceGuestLoginProcedure:
			authServiceGuestLoginHandler.ServeHTTP(w, r)
		case AuthServiceLogoutProcedure:
			authServiceLogoutHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) Login(context.Context, *connect.Request[v1.AuthRequest]) (*connect.Response[v1.UserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.Login is not implemented"))
}

func (UnimplementedAuthServiceHandler) Register(context.Context, *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.Register is not implemented"))
}

func (UnimplementedAuthServiceHandler) Test(context.Context, *connect.Request[v1.AuthResponse]) (*connect.Response[v1.UserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.Test is not implemented"))
}

func (UnimplementedAuthServiceHandler) GuestLogin(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.UserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.GuestLogin is not implemented"))
}

func (UnimplementedAuthServiceHandler) Logout(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.Logout is not implemented"))
}
