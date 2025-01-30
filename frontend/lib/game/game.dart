import 'dart:async';
import 'dart:math';

import 'package:flame/camera.dart';
import 'package:flame/components.dart' hide Timer;
import 'package:flame/events.dart';
import 'package:flame/flame.dart';
import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:multipacman/game/components/block.component.dart';
import 'package:multipacman/game/components/pellet.component.dart';
import 'package:multipacman/game/components/powerup.component.dart';
import 'package:multipacman/game/components/utils.dart';
import 'package:multipacman/game/connection_manager/game.manager.dart';
import 'package:multipacman/game/mixins/keyboard_mixin.dart';

class GameWorld extends FlameGame
    with HasCollisionDetection, KeyboardEvents, KeyInputHandler {
  final GameManager manager;

  GameWorld(this.manager);

  late final CameraComponent cameraComponent;
  late final World gameWorld;

  final mapWidth = 1700.0;
  final mapHeight = 1000.0;

  @override
  Future<void> onLoad() async {
    await Flame.images.loadAllImages();

    // Initialize the world
    gameWorld = World();

    await buildAndLoadMap();
    // add sprites already loaded from manager
    await gameWorld.add(manager.pacman);
    await gameWorld.addAll(manager.ghostList.values);
    manager.assignControllingSprite();

    await lifecycleEventsProcessed;

    // Initialize the camera with adaptive viewport
    cameraComponent = CameraComponent(
      world: gameWorld,
      viewport: FixedResolutionViewport(
        resolution: Vector2(mapWidth, mapHeight),
      ),
    );

    // Add camera and world to the game
    addAll([cameraComponent, gameWorld]);

    // Center the camera on the map
    cameraComponent.viewfinder.position = Vector2(mapWidth / 2, mapHeight / 2);
    cameraComponent.viewfinder.anchor = Anchor.center;

    // start socket listener
    await super.onLoad();
  }

  @override
  KeyEventResult onKeyEvent(
    KeyEvent event,
    Set<LogicalKeyboardKey> keysPressed,
  ) =>
      onKeyEventHandler(event, keysPressed);

  @override
  void update(double dt) {
    super.update(dt);

    if (manager.gameOverText.text.isNotEmpty) return;
    final message = manager.checkGameOver();
    if (message != null) {
      manager.showGameOverText(
        Vector2(size.x / 2, size.y / 2),
        message,
      );
      gameWorld.add(manager.gameOverText);
      return;
    }

    // Track frame time
    // if (dt > 0.016) { // More than 16ms (less than 60 FPS)
    //   print('Frame drop detected: ${dt * 1000}ms');
    // }

    handleKeyInput(manager.controllingSprite);
  }

  @override
  void onGameResize(Vector2 size) {
    super.onGameResize(size);

    if (!isLoaded) return;

    final zoomX = size.x / mapWidth;
    final zoomY = size.y / mapHeight;
    final newZoom = min(zoomX, zoomY).clamp(0.1, 2.0);

    cameraComponent.viewfinder.zoom = newZoom;
  }

  Future<void> buildAndLoadMap() async {
    final (
      height,
      width,
      mapElements,
      pelletElements,
      powerUpElements,
    ) = await getMapMetadata();

    final blueTileSprite = await getTileSprite("secondTile.png");
    final redTileSprite = await getTileSprite("forthTile.png");

    var position = Vector2(0, 0);
    var globIndex = 0;

    for (int rowIndex = 0; rowIndex < height; rowIndex++) {
      for (int colIndex = 0; colIndex < width; colIndex++) {
        final mapGid = mapElements[globIndex] as int;
        final pelletGid = pelletElements[globIndex] as int;
        final powerUpGid = powerUpElements[globIndex] as int;
        final tileId = globIndex;

        globIndex++;
        position.x += blockSize;

        if (manager.gameState!.pelletsEaten.contains(tileId) ||
            manager.gameState!.powerUpsEaten.contains(tileId)) {
          continue;
        }

        if (mapGid == 5) {
          // todo portals
          continue;
        }

        if (mapGid == 0) {
          // decide pellet or power up
          if (pelletGid == 3) {
            final pellet = PelletComponent(
              tileId: tileId,
              position: position,
            );

            manager.pelletMap.addAll({tileId: pellet});
            await gameWorld.add(pellet);
          } else if (powerUpGid == 4) {
            final power = PowerUpComponent(
              tileId: tileId,
              position: position,
            );

            manager.powerMap.addAll({tileId: power});
            await gameWorld.add(power);
          }

          // empty tile
          continue;
        }

        await gameWorld.add(
          BlockComponent(
            tileId: tileId,
            sprite: mapGid != 2 ? blueTileSprite : redTileSprite,
            position: position,
          ),
        );
      }
      position.y += blockSize;
      position.x = 0;
    }
  }

  Future<Sprite> getTileSprite(String spriteImg) async {
    final blueTileSprite = Sprite(
      await images.load(spriteImg),
      srcSize: Vector2(blockSize, blockSize),
    );
    return blueTileSprite;
  }
}
