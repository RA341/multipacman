import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/clients/auth_api.dart';
import 'package:multipacman/gen/auth/v1/auth.pb.dart';
import 'package:multipacman/grpc/api.dart';

final userDataProvider = FutureProvider<UserResponse?>((ref) async {
  final token = ref.watch(apiTokenProvider);
  final authApi = ref.watch(authApiProvider);

  return await authApi.testToken(token: token);
});

final lobbyIDProvider = StateProvider<int>((ref) {
  return 0;
});

final deviceWidthProvider = StateProvider<double>((ref) {
  return 0.0;
});