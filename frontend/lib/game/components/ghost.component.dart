import 'package:flame/components.dart';
import 'package:flame/sprite.dart';
import 'package:flutter/material.dart';
import 'package:multipacman/game/components/pacman.component.dart';
import 'package:multipacman/game/components/player.component.dart';
import 'package:multipacman/game/components/utils.dart';
import 'package:multipacman/game/connection_manager/game.manager.dart';

class GhostComponent extends PlayerComponent {
  GhostComponent(
    String spriteId,
    SpriteSheet spriteSheet,
    int baseIndex,
    Vector2 pos,
    GameManager manager,
  ) : super(
          manager: manager,
          spriteId: spriteId,
          animations: getGhostAnimMap(spriteSheet, baseIndex),
          position: pos,
          text: manager.connectedPlayers[spriteId]?.user ?? spriteId,
          textColor: decideGhostTextColor(spriteId),
        );

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);

    if (other is PacmanComponent) {
      manager.sendPacmanGhostCollisionAction(spriteId);
    }
  }

  static decideGhostTextColor(String spriteId) {
    if (spriteId == "gh0") {
      return Colors.red;
    }
    if (spriteId == "gh1") {
      return Colors.cyan;
    }
    if (spriteId == "gh2") {
      return Colors.pink;
    }
  }
}
