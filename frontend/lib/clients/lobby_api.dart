import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pbgrpc.dart';
import 'package:multipacman/grpc/api.dart';

final lobbyApiProvider = Provider<LobbyServiceClient>((ref) {
  final channel = ref.watch(grpcChannelProvider);
  final authInterceptor = ref.watch(authInterceptorProvider);

  return LobbyServiceClient(
    channel,
    interceptors: [authInterceptor],
  );
});

final lobbyListProvider = FutureProvider<List<Lobby>>((ref) async {
  final client = ref.watch(lobbyApiProvider);
  return (await client.listLobbies(ListLobbiesRequest())).lobbies;
});
