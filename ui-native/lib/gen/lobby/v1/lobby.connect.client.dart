//
//  Generated code. Do not modify.
//  source: lobby/v1/lobby.proto
//

import "package:connectrpc/connect.dart" as connect;
import "lobby.pb.dart" as lobbyv1lobby;
import "lobby.connect.spec.dart" as specs;

extension type LobbyServiceClient (connect.Transport _transport) {
  /// todo figure out lobby streaming
  Future<lobbyv1lobby.ListLobbiesResponse> listLobbies(
    lobbyv1lobby.ListLobbiesRequest input, {
    connect.Headers? headers,
    connect.AbortSignal? signal,
    Function(connect.Headers)? onHeader,
    Function(connect.Headers)? onTrailer,
  }) {
    return connect.Client(_transport).unary(
      specs.LobbyService.listLobbies,
      input,
      signal: signal,
      headers: headers,
      onHeader: onHeader,
      onTrailer: onTrailer,
    );
  }

  Future<lobbyv1lobby.AddLobbiesResponse> addLobby(
    lobbyv1lobby.AddLobbiesRequest input, {
    connect.Headers? headers,
    connect.AbortSignal? signal,
    Function(connect.Headers)? onHeader,
    Function(connect.Headers)? onTrailer,
  }) {
    return connect.Client(_transport).unary(
      specs.LobbyService.addLobby,
      input,
      signal: signal,
      headers: headers,
      onHeader: onHeader,
      onTrailer: onTrailer,
    );
  }

  Future<lobbyv1lobby.DelLobbiesResponse> deleteLobby(
    lobbyv1lobby.DelLobbiesRequest input, {
    connect.Headers? headers,
    connect.AbortSignal? signal,
    Function(connect.Headers)? onHeader,
    Function(connect.Headers)? onTrailer,
  }) {
    return connect.Client(_transport).unary(
      specs.LobbyService.deleteLobby,
      input,
      signal: signal,
      headers: headers,
      onHeader: onHeader,
      onTrailer: onTrailer,
    );
  }
}
