import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/clients/lobby_api.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pb.dart';
import 'package:multipacman/ui/components/lobby_grid_view_if.dart';
import 'package:multipacman/ui/components/utils.dart';

class LobbyGridView extends ConsumerWidget {
  const LobbyGridView({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final lobbyList = ref.watch(lobbyListProvider);

    return lobbyList.when(
      data: (data) => LobbyGridViewIf(data: data),
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
          return LobbyGridViewIf(data: val);
        }

        return Center(child: CircularProgressIndicator());
      },
    );
  }
}

class LobbyActionBar extends ConsumerWidget {
  const LobbyActionBar({
    super.key,
    required this.item,
  });

  final Lobby item;

  @override
  Widget build(BuildContext context, ref) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: [
        IconButton(
          onPressed: () async {
            await runGrpcRequest(
              context,
              () async {
                await ref
                    .read(lobbyApiProvider)
                    .deleteLobby(DelLobbiesRequest(lobby: item));

                ref.invalidate(lobbyListProvider);
              },
            );
          },
          icon: Icon(Icons.delete),
        ),
        ElevatedButton(
          onPressed: () {},
          child: Text('Join'),
        )
      ],
    );
  }
}
