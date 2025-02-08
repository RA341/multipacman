import 'package:flame/components.dart';
import 'package:flame/sprite.dart';
import 'package:flutter/material.dart';
import 'package:multipacman/game/components/player.component.dart';
import 'package:multipacman/game/components/utils.dart';
import 'package:multipacman/game/connection_manager/game.manager.dart';
import 'package:multipacman/utils.dart';

class PacmanComponent extends PlayerComponent {
  @override
  int get moveSpeed => 1;

  final pelletsEaten = <Vector2>{};
  final powerUpEaten = <Vector2>{};


  PacmanComponent(
    SpriteSheet spriteSheet,
    int baseIndex,
    Vector2 pos,
    GameManager manager,
  ) : super(
          manager: manager,
          spriteId: 'pacman',
          animations: getPacmanAnimMap(spriteSheet, baseIndex),
          position: pos,
          text: manager.connectedPlayers['pacman']?.user ?? 'Pacman',
          textColor: Colors.yellow,
        );

  bool get isAllowedToEat => manager.controllingSpriteId == "pacman";

  void eatPellet(int pelletId) {
    if (!isAllowedToEat) {
      // no collision for non pacman controlling players
      return;
    }

    manager.sendPelletAction(pelletId);
  }

  void eatPowerUp(int powerUpId) {
    if (!isAllowedToEat) {
      // no collision for non pacman controlling players
      return;
    }

    manager.sendPowerUpAction(powerUpId);
  }

  void endPowerUp() {
    paint.colorFilter = null;
    logger.d('power up over');
  }

  void startPowerUp() {
    paint.colorFilter = ColorFilter.mode(Colors.red, BlendMode.srcATop);
  }

  @override
  void neutral() {
    // no anim when player has stopped
    animationTicker?.currentIndex = 0;
    playing = false;
  }

  @override
  void changeDirection(Direction direction) {
    // continue playing
    playing = true;
    super.changeDirection(direction);
  }
}
