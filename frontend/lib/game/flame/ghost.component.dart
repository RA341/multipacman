import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame/sprite.dart';
import 'package:flutter/material.dart';
import 'package:multipacman/game/flame/pacman.component.dart';
import 'package:multipacman/game/flame/player.component.dart';
import 'package:multipacman/game/flame/utils.dart';

class GhostComponent extends PlayerComponent with CollisionCallbacks {
  GhostComponent(
    String spriteId,
    SpriteSheet spriteSheet,
    int baseIndex,
    Vector2 pos,
  ) : super(
          spriteId: spriteId,
          animations: getGhostAnimMap(spriteSheet, baseIndex),
          position: pos,
          text: 'Ghost',
          textColor: Colors.yellow,
        );

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);

    if (other is PacmanComponent) {
      if (other.isPoweredUp) {
        print('Ghost: $spriteId messed with powered pacman');
        removeFromParent();
      } else {
        print('Ghost: $spriteId eating pacman');
        other.removeFromParent();
      }
    }
  }
}
