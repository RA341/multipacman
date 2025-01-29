import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/clients/lobby_api.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pb.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/ui/components/lobby_list.dart';
import 'package:multipacman/ui/components/utils.dart';
import 'package:multipacman/utils.dart';

class LobbyGridView extends ConsumerWidget {
  const LobbyGridView({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final lobbyList = ref.watch(lobbyListProvider);

    return lobbyList.when(
      data: (data) => LobbyList(data: data),
      error: (error, stackTrace) => Center(
        child: Column(
          children: [
            Text('Error fetching lobbies'),
            Text(error.toString()),
          ],
        ),
      ),
      loading: () {
        final val = ref.read(lobbyListProvider).valueOrNull;
        if (val != null) {
          return LobbyList(data: val);
        }

        return Center(child: CircularProgressIndicator());
      },
    );
  }
}
