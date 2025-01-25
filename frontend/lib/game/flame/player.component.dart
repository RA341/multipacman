import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame/extensions.dart';
import 'package:flame/text.dart';
import 'package:flutter/material.dart';
import 'package:multipacman/game/flame/utils.dart';

// base class to contain all players will use
class PlayerComponent extends SpriteAnimationComponent {
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

  PlayerComponent({
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
    playerNameText = TextComponent(
      text: text,
      // You can customize this to any text you like
      position: Vector2(
        position.x - textOffsetX,
        position.y - textOffsetY,
      ),
      anchor: Anchor.topCenter,
      // Position above the player
      textRenderer: TextPaint(
        style: TextStyle(
          fontSize: 16,
          color: textColor,
        ),
      ),
    );

    // add(playerNameText);
    // debugMode = true;
    add(
      CircleHitbox(
        radius: 20,
        position: Vector2(5, 5),
        isSolid: true
      )
        ..debugColor = Colors.green
        ..debugMode = true,
    );
    animation = animations[currentDirection];
  }

  @override
  void update(double dt) {
    super.update(dt);

    // Update the text position to stay above the player
    playerNameText.position = Vector2(
      x,
      y - textOffsetY,
    );
  }

  // @override
  // void render(Canvas canvas) {
  //   super.render(canvas);
  //   playerNameText.render(canvas); // Render the text above the player
  // }

  void changeDirection(Direction direction) {
    if (currentDirection != direction) {
      currentDirection = direction;
      animation = animations[direction];
    }
  }

  void neutral() {
    changeDirection(Direction.right);
  }

  void up() {
    y += -movePoint * moveSpeed;
    changeDirection(Direction.up);
  }

  void down() {
    y += movePoint * moveSpeed;
    changeDirection(Direction.down);
  }

  void left() {
    x += -movePoint * moveSpeed;
    changeDirection(Direction.left);
  }

  void right() {
    x += movePoint * moveSpeed;
    changeDirection(Direction.right);
  }
}
