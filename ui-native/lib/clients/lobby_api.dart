import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/gen/lobby/v1/lobby.connect.client.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pb.dart';
import 'package:multipacman/grpc/api.dart';

final lobbyApiProvider = Provider<LobbyServiceClient>((ref) {
  final channel = ref.watch(grpcChannelProvider);
  return LobbyServiceClient(channel);
});

final reconnectProvider = StateProvider<Duration>((ref) {
  return Duration.zero;
});

final lobbyListProvider = StreamProvider.autoDispose<List<Lobby>>((ref) async* {
  final lobbyApi = ref.watch(lobbyApiProvider);
  int attempt = 0;

  while (true) {
    try {
      final grpcLobbies = lobbyApi.listLobbies(ListLobbiesRequest());

      await for (final lobbyList in grpcLobbies) {
        ref.read(reconnectProvider.notifier).state = Duration.zero;
        attempt = 0; // Reset attempt on success
        yield lobbyList.lobbies;
      }
    } catch (e) {
      attempt++;
      final delay =
          Duration(milliseconds: 500 * (1 << attempt)); // Exponential backoff

      ref.read(reconnectProvider.notifier).state = delay;
      await Future.delayed(delay);

      if (attempt > 10) {
        // reset duration, on error, so that the message can be displayed
        ref.read(reconnectProvider.notifier).state = Duration.zero;
        throw Exception('Failed after multiple retries: $e');
      }
    }
  }
});
