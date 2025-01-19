import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/config.dart';
import 'package:multipacman/gen/auth/v1/auth.pbgrpc.dart';
import 'package:multipacman/grpc/api.dart';
import 'package:multipacman/utils.dart';

final authApiProvider = Provider<AuthApi>((ref) {
  final channel = ref.watch(grpcChannelProvider);
  final client = AuthServiceClient(channel);

  return AuthApi(client);
});

class AuthApi {
  final AuthServiceClient apiClient;

  AuthApi(this.apiClient);

  Future<void> login({
    required String user,
    required String pass,
  }) async {
    final token = await apiClient.login(AuthRequest(
      username: user,
      password: pass,
    ));

    if (token.authToken.isEmpty) {
      throw Exception("token returned was empty");
    }

    await prefs.setString('apikey', token.authToken);
  }

  Future<void> register({
    required String user,
    required String pass,
    required String passVerify,
  }) async {
    await apiClient.register(RegisterUserRequest(
      username: user,
      password: pass,
      passwordVerify: passVerify,
    ));
  }

  Future<UserResponse?> testToken({required String token}) async {
    if (token == '') {
      logger.w("Token is empty");
      return null;
    }

    try {
      final resp = await apiClient.test(
        AuthResponse(authToken: token),
      );
      return resp;
    } catch (e) {
      logger.e('Incorrect token', error: e);
      return null;
    }
  }
}
