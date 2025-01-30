import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:multipacman/game/components/player.component.dart';

class BlockComponent extends SpriteComponent with CollisionCallbacks {
  final int tileId;

  BlockComponent({
    required Vector2 position,
    required Sprite sprite,
    required this.tileId,
  }) : super(
          position: position,
          autoResize: true,
          sprite: sprite,
        ) {
    add(
      RectangleHitbox(priority: 1),
    );
  }

  // @override
  // void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
  //   super.onCollision(intersectionPoints, other);
  //   if (other is PlayerComponent) {
  //     final otherPos = other.position;
  //     other.position = otherPos;
  //   }
  // }
}
