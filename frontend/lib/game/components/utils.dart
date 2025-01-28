import 'dart:convert';

import 'package:flame/sprite.dart';
import 'package:flutter/services.dart';

enum Direction { down, left, right, up }

const blockSize = 50.0;

Future<(int, int, List, List, List)> getMapMetadata() async {
  final mapJson =
      jsonDecode(await rootBundle.loadString('assets/tiles/map.json'))
          as Map<String, dynamic>;

  final height = mapJson['height'] as int;
  final width = mapJson['width'] as int;

  final mapLayer = mapJson['layers'];

  final mapElements = mapLayer[0]['data'] as List<dynamic>;
  final pelletElements = mapLayer[1]['data'] as List<dynamic>;
  final powerUpElements = mapLayer[2]['data'] as List<dynamic>;

  return (
    height,
    width,
    mapElements,
    pelletElements,
    powerUpElements,
  );
}

Map<Direction, SpriteAnimation> getGhostAnimMap(
  SpriteSheet spriteSheet,
  int baseIndex,
) {
  final sprites = {
    Direction.up: [spriteSheet.getSprite(0, baseIndex + 1)],
    Direction.down: [spriteSheet.getSprite(0, baseIndex + 2)],
    Direction.left: [spriteSheet.getSprite(0, baseIndex + 3)],
    Direction.right: [spriteSheet.getSprite(0, baseIndex)],
  };

// Create animations from the sprites
  return {
    Direction.down: SpriteAnimation.spriteList(
      sprites[Direction.down]!,
      stepTime: 0.2,
    ),
    Direction.left: SpriteAnimation.spriteList(
      sprites[Direction.left]!,
      stepTime: 0.2,
    ),
    Direction.right: SpriteAnimation.spriteList(
      sprites[Direction.right]!,
      stepTime: 0.2,
    ),
    Direction.up: SpriteAnimation.spriteList(
      sprites[Direction.up]!,
      stepTime: 0.2,
    ),
  };
}

Map<Direction, SpriteAnimation> getPacmanAnimMap(
  SpriteSheet spriteSheet,
  int baseIndex,
) =>
    {
      Direction.right: getPacmanDirAnim(spriteSheet, baseIndex),
      Direction.up: getPacmanDirAnim(spriteSheet, baseIndex + 3),
      Direction.down: getPacmanDirAnim(spriteSheet, baseIndex + 6),
      Direction.left: getPacmanDirAnim(spriteSheet, baseIndex + 9),
    };

SpriteAnimation getPacmanDirAnim(
  SpriteSheet spriteSheet,
  int baseIndex,
) =>
    spriteSheet.createAnimation(
      row: 0,
      stepTime: 0.095,
      from: baseIndex,
      to: baseIndex + 3,
      loop: true,
    );
