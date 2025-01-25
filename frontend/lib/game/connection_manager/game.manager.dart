import 'package:multipacman/game/models/state.model.dart';
import 'package:multipacman/game/models/player.model.dart';
import 'package:multipacman/utils.dart';

class GameManager {
  void handleMessages(Map<String, dynamic> input) {
    if (!input.containsKey('type')) {
      logger.w('Unknown message type ${input.keys.toString()}');
      return;
    }

    final messageType = input['type'] as String;

    switch (messageType) {
      case 'active':
      case 'dis':
        final player = PlayerModel.fromJson(input);
        return;
      case 'state':
        final gameState = GameStateModel.fromJson(input);
        return;
    }
  }

  void setGameState(GameStateModel state) {
    // set game stuff
  }

  void removePlayerFromGame() {
    // remove player from game
  }

  void handlePelletAction() {}

  void handlePowerUpAction() {}
}
