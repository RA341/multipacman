//
//  Generated code. Do not modify.
//  source: auth/v1/auth.proto
//

import "package:connectrpc/connect.dart" as connect;
import "auth.pb.dart" as authv1auth;

abstract final class AuthService {
  /// Fully-qualified name of the AuthService service.
  static const name = 'auth.v1.AuthService';

  static const login = connect.Spec(
    '/$name/Login',
    connect.StreamType.unary,
    authv1auth.AuthRequest.new,
    authv1auth.UserResponse.new,
  );

  static const register = connect.Spec(
    '/$name/Register',
    connect.StreamType.unary,
    authv1auth.RegisterUserRequest.new,
    authv1auth.RegisterUserResponse.new,
  );

  static const test = connect.Spec(
    '/$name/Test',
    connect.StreamType.unary,
    authv1auth.AuthResponse.new,
    authv1auth.UserResponse.new,
  );

  static const guestLogin = connect.Spec(
    '/$name/GuestLogin',
    connect.StreamType.unary,
    authv1auth.Empty.new,
    authv1auth.UserResponse.new,
  );

  static const logout = connect.Spec(
    '/$name/Logout',
    connect.StreamType.unary,
    authv1auth.Empty.new,
    authv1auth.Empty.new,
  );
}
