import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/game/game.dart';
import 'package:multipacman/game/connection_manager/game.connection.dart';

class GameContainer extends ConsumerWidget {
  const GameContainer({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    // final gameSate = ref.watch(gameStreamProvider);

    return Container(
      padding: EdgeInsets.all(20),
      child: GameWidget.controlled(gameFactory: GameWorld.new),
    );
  }
}
