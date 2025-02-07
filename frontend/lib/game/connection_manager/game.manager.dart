import 'dart:async';
import 'dart:convert';

import 'package:flame/components.dart';
import 'package:flame/flame.dart';
import 'package:flame/palette.dart';
import 'package:flame/sprite.dart';
import 'package:flame/text.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/game/components/ghost.component.dart';
import 'package:multipacman/game/components/pacman.component.dart';
import 'package:multipacman/game/components/pellet.component.dart';
import 'package:multipacman/game/components/player.component.dart';
import 'package:multipacman/game/components/powerup.component.dart';
import 'package:multipacman/game/components/utils.dart';
import 'package:multipacman/game/connection_manager/game.connection.dart';
import 'package:multipacman/game/models/player.model.dart';
import 'package:multipacman/game/models/state.model.dart';
import 'package:multipacman/grpc/api.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/utils.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

final gameStatusProvider = StateProvider<String>((ref) {
  return '';
});

final gameManagerProvider =
    FutureProvider.autoDispose<GameManager>((ref) async {
  final lobbyId = ref.watch(lobbyIDProvider);
  if (lobbyId == 0) {
    throw Exception("Invalid lobby ID");
  }

  await landscapeModeOnly(true);

  ref.onDispose(() async => await landscapeModeOnly(false));

  final token = ref.watch(apiTokenProvider);
  final baseUrl = Uri.parse(ref.watch(basePathProvider));

  final wsUrl = getWsUrl(baseUrl, lobbyId);
  final channel = await createConnection(wsUrl, token);

  final gameStatusStream = StreamController<String>();

  gameStatusStream.stream.listen(
    (event) {
      Future.delayed(
        Duration(seconds: 5),
        () {
          ref.invalidate(lobbyIDProvider);
          ref.invalidate(gameStatusProvider);
        },
      );

      ref.read(gameStatusProvider.notifier).state = event;
    },
  );

  final man = GameManager(
    gameChannel: channel,
    gameStatusStream: gameStatusStream,
  );

  await man.setupManager();

  await man.waitForGameState();

  ref.onDispose(() async => await man.dispose());
  return man;
});

class GameManager {
  final WebSocketChannel gameChannel;

  final StreamController<String> gameStatusStream;

  GameManager({
    required this.gameChannel,
    required this.gameStatusStream,
  });

  // sprites
  late final PacmanComponent pacman;
  final ghostList = <String, GhostComponent>{};

  GameStateModel? gameState;
  final connectedPlayers = <String, PlayerModel>{};
  final spritePlayers = <String, PlayerComponent>{};

  String get controllingSpriteId => gameState!.controllingSpriteId;

  late final PlayerComponent controllingSprite;

  final pelletMap = <int, PelletComponent>{};
  final powerMap = <int, PowerUpComponent>{};

  bool streamComplete = false;

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
    sendPositionAction(
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
          await Future.delayed(Duration(milliseconds: 100));
        }
      },
    );
  }

  Future<void> dispose() async {
    gameStatusStream.close();
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
      (event) => handleMessages(event),
      onDone: () {
        streamComplete = true;
        sendGameStatusMessage("Game connection was closed");
      },
      onError: (Object error, StackTrace st) =>
          sendGameStatusMessage(error.toString()),
      cancelOnError: true,
    );
  }

  void sendGameStatusMessage(String mess) {
    try {
      gameStatusStream.add(mess);
    } catch (E) {
      logger.w('Unable to send status message');
    }
  }

  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // receive stuff
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
        handleAddNewPlayer(message);
        logger.i("Adding new player: connected $connectedPlayers");
        logger.i("Adding new player: sprite $spritePlayers");
        return;
      case 'dis':
        final player = PlayerModel.fromJson(message);
        connectedPlayers.remove(player.playerid);
        spritePlayers[player.playerid]?..playerNameText.text = "";
        spritePlayers.remove(player.playerid);
        return;
      case "nopow":
        pacman.endPowerUp();
        return;
      case "pel":
        handlePelletMessage(message);
        return;
      case "pow":
        handlePowerUpMessage(message);
        return;
      case "gho":
        // ghost eaten message
        final ghostId = message["ghostId"] as String?;
        if (ghostId == null) {
          logger.w("No ghost id detected");
          return;
        }
        ghostList[ghostId]?.removeFromParent();
        return;
      case "pacd":
        // pacman eaten message (game over)
        print('pacman was eaten');
        pacman.removeFromParent();
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

  String? checkGameOver() {
    if (pacman.isRemoved) {
      return "Pacman was eaten, ghosts win !!!";
    }

    // result in true if all ghosts are eaten
    final allGhostsEaten =
        ghostList.values.every((element) => element.isRemoved);

    if (allGhostsEaten) {
      return "All ghosts were eaten, Pacman wins !!!";
    }

    if (pelletMap.isEmpty && powerMap.isEmpty) {
      return "all pellets and power ups were eaten Pacman wins";
    }

    return null;
  }

  void handleAddNewPlayer(Map<String, dynamic> message) {
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

  void handlePelletMessage(Map<String, dynamic> msg) {
    final tileId = msg['id'] as int;
    pelletMap[tileId]?.removeFromParent();
    pelletMap.remove(tileId);
  }

  void handlePowerUpMessage(Map<String, dynamic> msg) {
    final tileId = msg['id'] as int;

    powerMap[tileId]?.removeFromParent();
    powerMap.remove(tileId);

    pacman.startPowerUp();
  }

  ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // send stuff
  void updatePosControllingSprite() {
    sendPositionAction(
      x: controllingSprite.x,
      y: controllingSprite.y,
      dir: controllingSprite.currentDirection,
    );
  }

  void sendPositionAction({
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

  void sendPelletAction(int id) {
    final msg = <String, dynamic>{
      "type": "pel",
      "id": id,
    };

    sendMessage(msg);
  }

  void sendPowerUpAction(int id) {
    final msg = <String, dynamic>{
      "type": "pow",
      "id": id,
    };

    sendMessage(msg);
  }

  void sendPacmanGhostCollisionAction(String ghostId) {
    final msg = <String, dynamic>{
      "type": "gho",
      "ghId": ghostId,
    };
    sendMessage(msg);
  }

  void sendEatPacmanAction() {
    final msg = <String, dynamic>{
      "type": "pacded",
    };
  }

  void sendMessage(Map<String, dynamic> msg) {
    try {
      msg.addAll({"secretToken": gameState!.playerSecretToken});
      gameChannel.sink.add(jsonEncode(msg));
    } catch (e) {
      logger.w('Unable to send message: ${e.toString()}');
    }
  }

  ///////////////////////////////////////////////////////////////////////////////
  // text

  final regular = TextPaint(
    style: TextStyle(
      fontSize: 60.0,
      color: BasicPalette.red.color,
      backgroundColor: Colors.black,
      fontWeight: FontWeight.bold,
    ),
  );
  late final gameOverText = TextComponent(
    textRenderer: regular,
    anchor: Anchor.topCenter,
    priority: 0,
  );

  Duration exitTimeOut = Duration(seconds: 5);

  void showGameOverText(Vector2 textPos, String message) {
    gameOverText.position = textPos;
    gameOverText.text = message;
    Future.delayed(
      exitTimeOut,
      () => sendGameStatusMessage('Game Over'),
    );
  }
}
