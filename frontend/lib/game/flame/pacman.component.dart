import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame/sprite.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:multipacman/game/flame/player.component.dart';
import 'package:multipacman/game/flame/utils.dart';

class PacmanComponent extends PlayerComponent with CollisionCallbacks {
  @override
  int get moveSpeed => 2;

  final pelletsEaten = <Vector2>{};
  final powerUpEaten = <Vector2>{};

  var isPoweredUp = false;

  PacmanComponent(SpriteSheet spriteSheet, int baseIndex, Vector2 pos)
      : super(
          spriteId: 'pacman',
          animations: getPacmanAnimMap(spriteSheet, baseIndex),
          position: pos,
          text: 'Pacman',
          textColor: Colors.yellow,
        );

  void eatPellet(Vector2 pelletId) {
    pelletsEaten.add(pelletId);
  }

  void eatPowerUp(Vector2 powerUpId) {
    if (isPoweredUp) {
      print('already powered');
      return;
    }

    // todo handle via the backend
    Future.delayed(
      Duration(seconds: kDebugMode ? 3 : 10),
      endPowerUp,
    );

    startPowerUp();
    powerUpEaten.add(powerUpId);
  }

  void endPowerUp() {
    isPoweredUp = false;
    paint.colorFilter = null;
    print('power up over');
  }

  void startPowerUp() {
    isPoweredUp = true;
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
