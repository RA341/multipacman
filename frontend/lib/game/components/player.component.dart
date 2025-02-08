import 'dart:async';
import 'dart:collection';

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame/extensions.dart';
import 'package:flame/text.dart';
import 'package:flutter/material.dart';
import 'package:multipacman/game/components/ghost.component.dart';
import 'package:multipacman/game/components/pellet.component.dart';
import 'package:multipacman/game/components/powerup.component.dart';
import 'package:multipacman/game/components/utils.dart';
import 'package:multipacman/game/connection_manager/game.manager.dart';

// base class to contain all players will use
class PlayerComponent extends SpriteAnimationComponent with CollisionCallbacks {
  late Map<Direction, SpriteAnimation> animations;
  Direction currentDirection = Direction.right;
  int moveSpeed = 1;
  double movePoint = 1;
  String text;
  Color textColor;
  late TextComponent playerNameText;

  final int textOffsetY = 10;
  final int textOffsetX = 10;
  final String spriteId;
  final GameManager manager;

  // store last 10 position moves
  final ListQueue<Vector2> _previousPosition = ListQueue(5);
  var isCollided = false;
  Direction? collidedDir;

  PlayerComponent({
    required this.manager,
    required this.spriteId,
    required this.animations,
    required this.textColor,
    required this.text,
    required Vector2 position,
  }) : super(
          size: Vector2(blockSize, blockSize),
          position: position,
        ) {
    // Initialize the TextComponent
    // position: Vector2(
    //   position.x - textOffsetX,
    //   position.y - textOffsetY,
    // ),

    playerNameText = TextComponent(
      text: text,
      // You can customize this to any text you like
      // anchor: Anchor.topCenter,
      // Position above the player
      textRenderer: TextPaint(
        style: TextStyle(fontSize: 16, color: textColor),
      ),
    );

    // debugMode = true;
    add(
      CircleHitbox(
        radius: 20,
        position: Vector2(5, 5),
      ),
    );



    animation = animations[currentDirection];
  }

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);
    if (other is PelletComponent || other is PowerUpComponent) {
      // print('collision with non block elements');
      // handled by the pacman class
      return;
    }

    if (spriteId != "pacman" && other is GhostComponent) {
      // ghost don't collide with each other
      return;
    }

    if (intersectionPoints.isEmpty) return;

    Vector2 center = position + Vector2(20, 20); // Circle's center
    Vector2 averagePoint = intersectionPoints.reduce((a, b) => a + b) /
        intersectionPoints.length.toDouble();

    Vector2 diff = averagePoint - center;

    if (diff.x.abs() > diff.y.abs()) {
      if (diff.x < 0) {
        // print("Collided from LEFT");
        collidedDir = Direction.left;
      } else {
        // print("Collided from RIGHT");
        collidedDir = Direction.right;
      }
    } else {
      if (diff.y < 0) {
        // print("Collided from TOP");
        collidedDir = Direction.up;
      } else {
        // print("Collided from BOTTOM");
        collidedDir = Direction.down;
      }
    }
  }

  @override
  void onCollisionEnd(PositionComponent other) {
    super.onCollisionEnd(other);
    isCollided = false;
    collidedDir = null;
  }

  bool isAllowedToMove(Direction dir) {
    return !(collidedDir == dir);
  }

  @override
  FutureOr<void> onLoad() async {
    super.onLoad();

    playerNameText.position = Vector2(
      size.x / 2 - playerNameText.size.x / 2, // Center horizontally
      playerNameText.size.y - 20, // Place above sprite with 5px padding
    );
  }

  @override
  void update(double dt) {
    super.update(dt);

    checkIfMoved();

    // Update the text position to stay above the player
    playerNameText.position = Vector2(
      x,
      y + textOffsetY,
    );
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);
    playerNameText.render(canvas); // Render the text above the player
  }

  void changeDirection(Direction direction) {
    if (currentDirection != direction) {
      currentDirection = direction;
      animation = animations[direction];
    }
  }

  void checkIfMoved() {
    // Check if the position has changed
    if (_previousPosition.contains(position)) {
      neutral();
    } else {
      _previousPosition.add(position.clone());
    }
  }

  void neutral() {
    changeDirection(Direction.right);
  }

  void up() {
    final dirMov = Direction.up;

    if (isAllowedToMove(dirMov)) {
      y += -movePoint * moveSpeed;
      changeDirection(dirMov);
      manager.updatePosControllingSprite();
    }
  }

  void down() {
    final dirMov = Direction.down;

    if (isAllowedToMove(dirMov)) {
      y += movePoint * moveSpeed;
      changeDirection(dirMov);
      manager.updatePosControllingSprite();
    }
  }

  void left() {
    final dirMov = Direction.left;

    if (isAllowedToMove(dirMov)) {
      x += -movePoint * moveSpeed;
      changeDirection(dirMov);
      manager.updatePosControllingSprite();
    }
  }

  void right() {
    final dirMov = Direction.right;

    if (isAllowedToMove(dirMov)) {
      x += movePoint * moveSpeed;
      changeDirection(dirMov);
      manager.updatePosControllingSprite();
    }
  }
}
