import 'dart:convert';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/gen/lobby/v1/lobby.connect.client.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pb.dart';
import 'package:multipacman/grpc/api.dart';
import 'package:multipacman/model/lobby.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

final lobbyApiProvider = Provider<LobbyServiceClient>((ref) {
  final channel = ref.watch(grpcChannelProvider);

  return LobbyServiceClient(channel);
});

final lobbyListProvider = StreamProvider<List<LobbyModel>>((ref) async* {
  final token = ref.watch(apiTokenProvider);
  final baseUrl = Uri.parse(ref.watch(basePathProvider));

  final wsUrl = getWsUrl(baseUrl);

  final channel = WebSocketChannel.connect(
    wsUrl,
    // https://github.com/orgs/deepgram/discussions/175
    // sends token via the Sec-WebSocket-Protocol header
    protocols: [token],
  );

  final lobbyStream = channel.stream.map((event) {
    final json = jsonDecode(event) as List<dynamic>;
    return json.map((e) => LobbyModel.fromJson(e)).toList();
  });

  await for (final lobbyList in lobbyStream) {
    print(lobbyList);
    yield lobbyList;
  }
});

final grpcLobbyListProvider = StreamProvider<List<Lobby>>((ref) async* {
  final grpcLobbies =
      ref.watch(lobbyApiProvider).listLobbies(ListLobbiesRequest());

  await for (final lobbyList in grpcLobbies) {
    yield lobbyList.lobbies;
  }
});

Uri getWsUrl(Uri baseUrl) {
  var host = '';
  if (baseUrl.scheme == "https") {
    host = 'wss://${baseUrl.host}';
  } else {
    host = 'ws://${baseUrl.host}:${baseUrl.port}';
  }

  return Uri.parse('$host/api/lobbies');
}
