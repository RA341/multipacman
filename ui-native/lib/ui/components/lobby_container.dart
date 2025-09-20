import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:multipacman/clients/lobby_api.dart';
import 'package:multipacman/ui/components/lobby_list.dart';

class LobbyGridView extends ConsumerWidget {
  const LobbyGridView({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final lobbyList = ref.watch(lobbyListProvider);
    final reconnectDelay = ref.watch(reconnectProvider);

    return reconnectDelay != Duration.zero
        ? Center(
            child: Column(
              children: [
                Text(
                  'Connection lost to server, attempting reconnection in ${reconnectDelay.inSeconds} seconds',
                  style: TextStyle(fontSize: 25),
                ),
                SizedBox(height: 50),
                ElevatedButton(
                  onPressed: () => ref.invalidate(lobbyListProvider),
                  child: Text('Reconnect now'),
                ),
              ],
            ),
          )
        : lobbyList.when(
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
