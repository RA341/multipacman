import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:multipacman/game/components/player.component.dart';

mixin KeyInputHandler on Game {
  final _pressedKeys = <LogicalKeyboardKey>{};

  KeyEventResult onKeyEventHandler(
    KeyEvent event,
    Set<LogicalKeyboardKey> keysPressed,
  ) {
    if (event is KeyDownEvent) {
      if (!_pressedKeys.contains(event.logicalKey)) {
        _pressedKeys.add(event.logicalKey); // Add key to the end (most recent)
      }
    } else if (event is KeyUpEvent) {
      _pressedKeys.remove(event.logicalKey); // Remove key when released
    }

    return KeyEventResult.handled;
  }

  void handleKeyInput(PlayerComponent sprite) {
    if (_pressedKeys.isEmpty) return; // No keys pressed
    final key = _pressedKeys.last;

    if (key == LogicalKeyboardKey.arrowUp) {
      sprite.up();
    } else if (key == LogicalKeyboardKey.arrowDown) {
      sprite.down();
    } else if (key == LogicalKeyboardKey.arrowRight) {
      sprite.right();
    } else if (key == LogicalKeyboardKey.arrowLeft) {
      sprite.left();
    }
  }
}
