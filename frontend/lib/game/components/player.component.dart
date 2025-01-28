import 'dart:async';

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame/extensions.dart';
import 'package:flame/text.dart';
import 'package:flutter/material.dart';
import 'package:multipacman/game/components/utils.dart';
import 'package:multipacman/game/connection_manager/game.manager.dart';

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
  final GameManager manager;

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
  FutureOr<void> onLoad() async {
    super.onLoad();

    await add(playerNameText);
    // Position the text above the sprite
    playerNameText.position = Vector2(
      size.x / 2 - playerNameText.size.x / 2, // Center horizontally
      -playerNameText.size.y - 5, // Place above sprite with 5px padding
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

  get dirF => changeDirection;

  void changeDirection(Direction direction) {
    if (currentDirection != direction) {
      currentDirection = direction;
      animation = animations[direction];
    }
  }

  void moveFromNetwork(Vector2 pos, Direction dir) {
    position = pos;
  }

  void sendMove(Vector2 vector) {}

  Vector2 _previousPosition = Vector2(0, 0);

  void checkIfMoved() {
    // Check if the position has changed
    if (_previousPosition == position) {
      neutral();
    } else {
      _previousPosition = position.clone();
    }
  }

  void neutral() {
    changeDirection(Direction.right);
  }

  void up() {
    y += -movePoint * moveSpeed;
    changeDirection(Direction.up);
    manager.updatePosControllingSprite();
  }

  void down() {
    y += movePoint * moveSpeed;
    changeDirection(Direction.down);
    manager.updatePosControllingSprite();
  }

  void left() {
    x += -movePoint * moveSpeed;
    changeDirection(Direction.left);
    manager.updatePosControllingSprite();
  }

  void right() {
    x += movePoint * moveSpeed;
    changeDirection(Direction.right);
    manager.updatePosControllingSprite();
  }
}
