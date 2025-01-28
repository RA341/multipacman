import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flutter/material.dart';
import 'package:multipacman/game/components/pacman.component.dart';
import 'package:multipacman/game/components/utils.dart';

class PortalComponent extends PositionComponent with CollisionCallbacks {
  final Vector2 vectorId;
  final Direction portalId;

  PortalComponent({
    required this.portalId,
    required this.vectorId,
    required super.position,
  }) {
    add(RectangleHitbox());
    debugMode = true;
    debugColor = Colors.green;
  }

  final leftVector = Vector2(105, 455);
  final rightVector = Vector2(1461, 451);

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);
    if (other is PacmanComponent) {
      print('portal-ling to other tile');

      if (portalId == Direction.left) {
        other.position = leftVector;
      } else if (portalId == Direction.right) {
        other.position = rightVector;
      }
      // other.position =
    }
  }
}
