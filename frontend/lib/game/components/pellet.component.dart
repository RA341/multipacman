import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame/flame.dart';
import 'package:multipacman/game/components/pacman.component.dart';
import 'package:multipacman/game/components/utils.dart';

class PelletComponent extends SpriteComponent with CollisionCallbacks {
  final Vector2 vectorId;

  PelletComponent({
    required super.position,
    required this.vectorId,
  });

  @override
  Future<void> onLoad() async {
    await super.onLoad();

    await add(CircleHitbox(
      radius: 5,
      position: Vector2(20, 20),
    ));

    sprite = Sprite(
      await Flame.images.load('centrepoint.png'),
      srcSize: Vector2(blockSize, blockSize),
    );
  }

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);
    if (other is PacmanComponent) {
      removeFromParent();
      other.eatPellet(vectorId);
    }
  }
}
