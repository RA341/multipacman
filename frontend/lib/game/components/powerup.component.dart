import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame/flame.dart';
import 'package:multipacman/game/components/pacman.component.dart';
import 'package:multipacman/game/components/utils.dart';

class PowerUpComponent extends SpriteComponent with CollisionCallbacks {
  final int tileId;

  PowerUpComponent({
    required super.position,
    required this.tileId,
  });

  @override
  Future<void> onLoad() async {
    await super.onLoad();

    sprite = Sprite(
      await Flame.images.load('powercent.png'),
      srcSize: Vector2(blockSize, blockSize),
    );

    await add(
      CircleHitbox(
        radius: 9,
        position: Vector2(16, 16),
      ),
    );
  }

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);

    if (other is PacmanComponent) {
      // print('Pacman collided with powerup');
      removeFromParent();
      other.eatPowerUp(tileId);
    }
  }
}
