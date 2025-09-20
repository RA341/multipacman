// This is a generated file - do not edit.
//
// Generated from auth/v1/auth.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'auth.pb.dart' as $0;
import 'auth.pbjson.dart';

export 'auth.pb.dart';

abstract class AuthServiceBase extends $pb.GeneratedService {
  $async.Future<$0.UserResponse> login(
      $pb.ServerContext ctx, $0.AuthRequest request);
  $async.Future<$0.RegisterUserResponse> register(
      $pb.ServerContext ctx, $0.RegisterUserRequest request);
  $async.Future<$0.UserResponse> test(
      $pb.ServerContext ctx, $0.AuthResponse request);
  $async.Future<$0.UserResponse> guestLogin(
      $pb.ServerContext ctx, $0.Empty request);
  $async.Future<$0.Empty> logout($pb.ServerContext ctx, $0.Empty request);

  $pb.GeneratedMessage createRequest($core.String methodName) {
    switch (methodName) {
      case 'Login':
        return $0.AuthRequest();
      case 'Register':
        return $0.RegisterUserRequest();
      case 'Test':
        return $0.AuthResponse();
      case 'GuestLogin':
        return $0.Empty();
      case 'Logout':
        return $0.Empty();
      default:
        throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx,
      $core.String methodName, $pb.GeneratedMessage request) {
    switch (methodName) {
      case 'Login':
        return login(ctx, request as $0.AuthRequest);
      case 'Register':
        return register(ctx, request as $0.RegisterUserRequest);
      case 'Test':
        return test(ctx, request as $0.AuthResponse);
      case 'GuestLogin':
        return guestLogin(ctx, request as $0.Empty);
      case 'Logout':
        return logout(ctx, request as $0.Empty);
      default:
        throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => AuthServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>>
      get $messageJson => AuthServiceBase$messageJson;
}
