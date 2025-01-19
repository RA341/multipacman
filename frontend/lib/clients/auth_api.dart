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
    final token = await apiClient.authenticate(AuthRequest(
      username: user,
      password: pass,
    ));
    await prefs.setString('apikey', token.authToken);
  }

  Future<void> register({
    required String user,
    required String pass,
    required String passVerify,
  }) async {
    await apiClient.registerUser(RegisterUserRequest(
      username: user,
      password: pass,
      passwordVerify: passVerify,
    ));
  }

  Future<bool> testToken({required String token}) async {
    if (token == '') {
      logger.w("Token is empty");
      return false;
    }

    try {
      await apiClient.test(
        AuthResponse(authToken: token),
      );
    } catch (e) {
      logger.e('Incorrect token', error: e);
      return false;
    }

    return true;
  }
}
