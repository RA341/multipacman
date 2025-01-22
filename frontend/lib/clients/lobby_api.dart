import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/gen/lobby/v1/lobby.connect.client.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pb.dart';
import 'package:multipacman/grpc/api.dart';

final lobbyApiProvider = Provider<LobbyServiceClient>((ref) {
  final channel = ref.watch(grpcChannelProvider);
  return LobbyServiceClient(channel);
});

final lobbyListProvider = StreamProvider<List<Lobby>>((ref) async* {
  final grpcLobbies =
      ref.watch(lobbyApiProvider).listLobbies(ListLobbiesRequest());

  await for (final lobbyList in grpcLobbies) {
    yield lobbyList.lobbies;
  }
});
