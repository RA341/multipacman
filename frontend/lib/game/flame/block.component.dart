import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:multipacman/game/flame/player.component.dart';

class BlockComponent extends SpriteComponent with CollisionCallbacks {
  final Vector2 vectorId;

  BlockComponent({
    required Vector2 position,
    required Sprite sprite,
    required this.vectorId,
  }) : super(
          position: position,
          autoResize: true,
          sprite: sprite,
        ) {
    add(
      RectangleHitbox(priority: 1),
    );
  }

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);
    if (other is PlayerComponent) {
      print('player colliding with wall');
      debugMode = true;
      // print(intersectionPoints);
      // other.position = intersectionPoints.last;
    }
  }
}
