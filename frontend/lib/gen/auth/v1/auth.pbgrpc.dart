//
//  Generated code. Do not modify.
//  source: auth/v1/auth.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'auth.pb.dart' as $0;

export 'auth.pb.dart';

@$pb.GrpcServiceName('auth.v1.AuthService')
class AuthServiceClient extends $grpc.Client {
  static final _$login = $grpc.ClientMethod<$0.AuthRequest, $0.UserResponse>(
      '/auth.v1.AuthService/Login',
      ($0.AuthRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.UserResponse.fromBuffer(value));
  static final _$register = $grpc.ClientMethod<$0.RegisterUserRequest, $0.RegisterUserResponse>(
      '/auth.v1.AuthService/Register',
      ($0.RegisterUserRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.RegisterUserResponse.fromBuffer(value));
  static final _$test = $grpc.ClientMethod<$0.AuthResponse, $0.UserResponse>(
      '/auth.v1.AuthService/Test',
      ($0.AuthResponse value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.UserResponse.fromBuffer(value));

  AuthServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.UserResponse> login($0.AuthRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$login, request, options: options);
  }

  $grpc.ResponseFuture<$0.RegisterUserResponse> register($0.RegisterUserRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$register, request, options: options);
  }

  $grpc.ResponseFuture<$0.UserResponse> test($0.AuthResponse request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$test, request, options: options);
  }
}

@$pb.GrpcServiceName('auth.v1.AuthService')
abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'auth.v1.AuthService';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.AuthRequest, $0.UserResponse>(
        'Login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.AuthRequest.fromBuffer(value),
        ($0.UserResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.RegisterUserRequest, $0.RegisterUserResponse>(
        'Register',
        register_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RegisterUserRequest.fromBuffer(value),
        ($0.RegisterUserResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.AuthResponse, $0.UserResponse>(
        'Test',
        test_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.AuthResponse.fromBuffer(value),
        ($0.UserResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.UserResponse> login_Pre($grpc.ServiceCall call, $async.Future<$0.AuthRequest> request) async {
    return login(call, await request);
  }

  $async.Future<$0.RegisterUserResponse> register_Pre($grpc.ServiceCall call, $async.Future<$0.RegisterUserRequest> request) async {
    return register(call, await request);
  }

  $async.Future<$0.UserResponse> test_Pre($grpc.ServiceCall call, $async.Future<$0.AuthResponse> request) async {
    return test(call, await request);
  }

  $async.Future<$0.UserResponse> login($grpc.ServiceCall call, $0.AuthRequest request);
  $async.Future<$0.RegisterUserResponse> register($grpc.ServiceCall call, $0.RegisterUserRequest request);
  $async.Future<$0.UserResponse> test($grpc.ServiceCall call, $0.AuthResponse request);
}
