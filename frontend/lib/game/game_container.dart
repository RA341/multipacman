import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/game/connection_manager/game.manager.dart';
import 'package:multipacman/game/game.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/utils.dart';

class GameContainer extends ConsumerWidget {
  const GameContainer({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final gameManager = ref.watch(gameManagerProvider);
    final gameStatus = ref.watch(gameStatusProvider);

    return gameStatus.isNotEmpty
        ? GameExitMessage()
        : gameManager.when(
            data: (manager) => GameCont(),
            error: (error, stackTrace) {
              // move to home page after 3 seconds
              Future.delayed(
                Duration(seconds: 3),
                () => ref.modState(lobbyIDProvider, 0),
              );

              return GameError(error: error);
            },
            loading: () => GameLoader(),
          );
  }
}

class GameCont extends ConsumerWidget {
  const GameCont({
    super.key,
  });

  @override
  Widget build(BuildContext context, ref) {
    final manager = ref.watch(gameManagerProvider).value!;

    return Container(
      padding: EdgeInsets.all(20),
      child: Column(
        children: [
          Row(
            children: [
              ElevatedButton(
                onPressed: () {
                  ref.modState(lobbyIDProvider, 0);
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
    );
  }
}

class GameExitMessage extends ConsumerWidget {
  const GameExitMessage({
    super.key,
  });

  @override
  Widget build(BuildContext context, ref) {
    final gameStatus = ref.watch(gameStatusProvider);

    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Text(
            gameStatus,
            style: TextStyle(fontSize: 50),
          ),
          SizedBox(height: 50),
          ElevatedButton(
            onPressed: () {
              ref.invalidate(gameStatusProvider);
              ref.invalidate(lobbyIDProvider);
            },
            child: Text('Back home'),
          )
        ],
      ),
    );
  }
}

class GameError extends StatelessWidget {
  const GameError({super.key, required this.error});

  final Object error;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text('Unable to connect to lobby'),
        Text(error.toString()),
      ],
    );
  }
}

class GameLoader extends StatelessWidget {
  const GameLoader({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Text(
        'Connecting to lobby... standby',
        style: TextStyle(fontSize: 30),
      ),
    );
  }
}
