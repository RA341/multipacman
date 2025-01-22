import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/game/game.connection.dart';

class GameContainer extends ConsumerWidget {
  const GameContainer({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final gameSate = ref.watch(gameStreamProvider);

    return gameSate.when(
        data: (data) {
          return Text(data.toString());
        },
        error: (error, stackTrace) {
          return Center(
            child: Column(
              children: [
                Text('Uanble to connect to game websocket'),
                Text(error.toString()),
              ],
            ),
          );
        },
        loading: () => Center(child: CircularProgressIndicator()));
  }
}
