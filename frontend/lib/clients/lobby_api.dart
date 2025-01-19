import 'dart:async';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pbgrpc.dart';
import 'package:multipacman/grpc/api.dart';

final lobbyProvider = AsyncNotifierProvider<LobbyNotifier, List<Lobby>>(() {
  return LobbyNotifier();
});

class LobbyNotifier extends AsyncNotifier<List<Lobby>> {
  late final LobbyServiceClient apiClient;

  @override
  FutureOr<List<Lobby>> build() {
    apiClient = _getClientInst();
    return _listCategory();
  }

  LobbyServiceClient _getClientInst() {
    final channel = ref.watch(grpcChannelProvider);
    final authInterceptor = ref.watch(authInterceptorProvider);

    return LobbyServiceClient(
      channel,
      interceptors: [authInterceptor],
    );
  }

  Future<void> addCategory(String lname) async {
    state = const AsyncValue.loading();

    state = await AsyncValue.guard(() async {
      await apiClient.addLobby(AddLobbiesRequest(lobbyName: lname));
      return await _listCategory();
    });
  }

  Future<void> deleteCategory(Lobby cat) async {
    state = const AsyncValue.loading();

    state = await AsyncValue.guard(() async {
      await apiClient.deleteLobby(DelLobbiesRequest(lobby: cat));
      return await _listCategory();
    });
  }

  Future<List<Lobby>> _listCategory() async {
    final resp = await apiClient.listLobbies(ListLobbiesRequest());
    return resp.lobbies;
  }
}
