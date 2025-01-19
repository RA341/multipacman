import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/clients/auth_api.dart';
import 'package:multipacman/grpc/api.dart';

final authStatusProvider = FutureProvider<bool>((ref) async {
  final token = ref.watch(apiTokenProvider);
  final authApi = ref.watch(authApiProvider);

  return await authApi.testToken(token: token);
});
