import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/game/connection_manager/game.manager.dart';
import 'package:multipacman/game/game.dart';
import 'package:multipacman/providers.dart';

class GameContainer extends ConsumerWidget {
  const GameContainer({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final gameManager = ref.watch(gameManagerProvider);

    return gameManager.when(
      data: (manager) => Container(
        padding: EdgeInsets.all(20),
        child: Column(
          children: [
            Row(
              children: [
                ElevatedButton(
                  onPressed: () {
                    ref.read(lobbyIDProvider.notifier).state = 0;
                  },
                  child: Text('Back'),
                ),
                Text(
                  '',
                  style: TextStyle(),
                ),
              ],
            ),
            SizedBox(height: 10),
            Expanded(
              child: GameWidget.controlled(
                gameFactory: () => GameWorld(manager),
              ),
            ),
          ],
        ),
      ),
      error: (error, stackTrace) => Column(
        children: [
          Text('Unable to connect to lobby'),
          Text(error.toString()),
        ],
      ),
      loading: () => Center(
        child: Text(
          'Connecting to lobby... standby',
          style: TextStyle(fontSize: 30),
        ),
      ),
    );
  }
}
