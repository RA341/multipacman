import 'dart:convert';

import 'package:flame/components.dart';
import 'package:flame/flame.dart';
import 'package:flame/sprite.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/game/components/ghost.component.dart';
import 'package:multipacman/game/components/pacman.component.dart';
import 'package:multipacman/game/components/player.component.dart';
import 'package:multipacman/game/components/utils.dart';
import 'package:multipacman/game/connection_manager/game.connection.dart';
import 'package:multipacman/game/models/player.model.dart';
import 'package:multipacman/game/models/state.model.dart';
import 'package:multipacman/grpc/api.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/utils.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

final gameManagerProvider =
    FutureProvider.autoDispose<GameManager>((ref) async {
  final lobbyId = ref.watch(lobbyIDProvider);
  if (lobbyId == 0) {
    throw Exception("Invalid lobby ID");
  }

  final token = ref.watch(apiTokenProvider);
  final baseUrl = Uri.parse(ref.watch(basePathProvider));

  final wsUrl = getWsUrl(baseUrl, lobbyId);
  final channel = await createConnection(wsUrl, token);

  final man = GameManager(gameChannel: channel);
  await man.setupManager();

  await man.waitForGameState();

  ref.onDispose(() async => await man.dispose());
  return man;
});

class GameManager {
  final WebSocketChannel gameChannel;

  GameManager({required this.gameChannel});

  // sprites
  late final PacmanComponent pacman;
  final ghostList = <String, GhostComponent>{};

  get ghostIds => ghostList.values.toList();

  GameStateModel? gameState;
  var isClosed = false;

  Exception? streamError;

  final connectedPlayers = <String, PlayerModel>{};
  final spritePlayers = <String, PlayerComponent>{};

  String get controllingSpriteId => gameState!.controllingSpriteId;

  late final PlayerComponent controllingSprite;

  void assignControllingSprite() {
    if (controllingSpriteId == 'pacman') {
      controllingSprite = pacman;
    } else {
      final tmp = ghostList[controllingSpriteId];
      if (tmp == null) {
        throw Exception("Unidentified sprite ID $controllingSpriteId");
      }

      controllingSprite = tmp;
    }

    // initial position data
    sendPosData(
      x: controllingSprite.x,
      y: controllingSprite.y,
      dir: controllingSprite.currentDirection,
    );
  }

  Future<void> waitForGameState() {
    return Future<void>(
      () async {
        // wait for game state
        while (gameState == null) {
          if (streamError != null) {
            print('error detectd fuiture');
            throw streamError!;
          }
          await Future.delayed(Duration(milliseconds: 100));
        }
      },
    );
  }

  Future<void> dispose() async {
    isClosed = true;
    gameChannel.sink.close();
    logger.d('leaving game');
  }

  void updateCharacterPos(PlayerModel posData) {
    print(posData.spriteType);
    connectedPlayers[posData.playerid] = posData;
    if (posData.spriteType == controllingSpriteId) {
      // do not update self causes issue on web (fucking flutter web)
      print('self character');
      return;
    }

    final tmp = spritePlayers[posData.playerid];
    tmp?.position = Vector2(
      posData.x,
      posData.y,
    );

    final dd = Direction.values.byName(posData.dir);
    tmp?.changeDirection(dd);
  }

  Future<void> setupManager() async {
    // load sprites (not loaded in game yet)
    await _loadPacman();
    await _loadGhosts();
    listenForMessages();
  }

  Future<void> _loadPacman() async {
    final pacmanSpriteSheet = SpriteSheet(
      image: await Flame.images.load('pacmanSpriteSheet.png'),
      srcSize: Vector2(blockSize, blockSize),
    );

    pacman = PacmanComponent(
      pacmanSpriteSheet,
      0,
      Vector2(100, 200),
      this,
    );
    pacman.neutral();
  }

  Future<void> _loadGhosts() async {
    final image = await Flame.images.load('ghosts.png');
    final ghostSpriteSheet = SpriteSheet(
      image: image,
      srcSize: Vector2(blockSize, blockSize), // Size of each frame
    );

    for (int x = 0; x < 12; x += 4) {
      // Create animations for each character
      final iter = x ~/ 4;

      // Create components for each character
      final character = GhostComponent(
        "gh$iter",
        ghostSpriteSheet,
        x,
        // place ghosts in 50 pixel intervals
        Vector2(700 + (iter * 50), 400),
        this,
      );

      ghostList.addAll({"gh$iter": character});
    }
  }

  void listenForMessages() {
    logger.i('listening for game messages');
    gameChannel.stream.listen(
      (event) {
        try {
          handleMessages(event);
        } catch (e) {
          streamError = Exception(e);
        }
      },
      onDone: () {
        print('game stream done');
      },
      onError: (Object error, StackTrace st) {
        print('error ');
        print(error);
        print(st);
        streamError = Exception(error);
      },
      cancelOnError: true,
    );
  }

  void handleMessages(String input) {
    final message = jsonDecode(input) as Map<String, dynamic>;

    if (!message.containsKey('type')) {
      logger.w('Unknown message ${message.keys.toString()}');
      throw Exception('Error: $input');
    }

    print(message);
    final messageType = message['type'] as String;

    switch (messageType) {
      case 'active':
        addNewPlayer(message);
        logger.i("Adding new player: connected $connectedPlayers");
        logger.i("Adding new player: sprite $spritePlayers");

        return;
      case 'dis':
        final player = PlayerModel.fromJson(message);
        connectedPlayers.remove(player.playerid);
        spritePlayers[player.playerid]?..playerNameText.text = "";
        spritePlayers.remove(player.playerid);
        return;
      case 'mov':
        final player = PlayerModel.fromJson(message);
        updateCharacterPos(player);
      case 'state':
        gameState = GameStateModel.fromJson(message);
        return;
      default:
        logger.w("Unknown message type: $messageType");
        return;
    }
  }

  void addNewPlayer(Map<String, dynamic> message) {
    final player = PlayerModel.fromJson(message);
    if (player.spriteType == "pacman") {
      connectedPlayers.addAll({player.playerid: player});
      spritePlayers.addAll({
        player.playerid: pacman..playerNameText.text = player.user,
      });
      return;
    }

    var playerSprite = ghostList[player.spriteType];
    if (playerSprite == null) {
      throw Exception("Unknown sprite ID");
    }

    connectedPlayers.addAll({player.playerid: player});
    spritePlayers.addAll({
      player.playerid: playerSprite..playerNameText.text = player.user,
    });
  }

  void updatePosControllingSprite() {
    sendPosData(
      x: controllingSprite.x,
      y: controllingSprite.y,
      dir: controllingSprite.currentDirection,
    );
  }

  void sendPosData({
    required double x,
    required double y,
    required Direction dir,
  }) {
    final msg = <String, dynamic>{
      "type": "mov",
      "dir": dir.name,
      "x": x.toString(),
      "y": y.toString(),
    };
    sendMessage(msg);
  }

  void sendMessage(Map<String, dynamic> msg) {
    if (isClosed) return;
    msg.addAll({"secretToken": gameState!.playerSecretToken});
    gameChannel.sink.add(jsonEncode(msg));
  }

  void handlePelletAction() {}

  void handlePowerUpAction() {}
}
