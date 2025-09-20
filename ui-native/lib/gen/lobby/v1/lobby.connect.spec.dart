//
//  Generated code. Do not modify.
//  source: lobby/v1/lobby.proto
//

import "package:connectrpc/connect.dart" as connect;
import "lobby.pb.dart" as lobbyv1lobby;

abstract final class LobbyService {
  /// Fully-qualified name of the LobbyService service.
  static const name = 'lobby.v1.LobbyService';

  /// todo figure out lobby streaming
  static const listLobbies = connect.Spec(
    '/$name/ListLobbies',
    connect.StreamType.unary,
    lobbyv1lobby.ListLobbiesRequest.new,
    lobbyv1lobby.ListLobbiesResponse.new,
  );

  static const addLobby = connect.Spec(
    '/$name/AddLobby',
    connect.StreamType.unary,
    lobbyv1lobby.AddLobbiesRequest.new,
    lobbyv1lobby.AddLobbiesResponse.new,
  );

  static const deleteLobby = connect.Spec(
    '/$name/DeleteLobby',
    connect.StreamType.unary,
    lobbyv1lobby.DelLobbiesRequest.new,
    lobbyv1lobby.DelLobbiesResponse.new,
  );
}
