import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/clients/lobby_api.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pb.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/ui/components/utils.dart';
import 'package:multipacman/utils.dart';
import 'package:timeago/timeago.dart' as timeago;
import 'package:universal_html/html.dart';

class LobbyList extends StatelessWidget {
  const LobbyList({
    super.key,
    required this.data,
  });

  final List<Lobby> data;

  @override
  Widget build(BuildContext context) {
    if (data.isEmpty) {
      return Center(
        child: Text(
          'No lobbies found, create one',
          style: TextStyle(fontSize: 35),
        ),
      );
    }

    return Expanded(
      child: SizedBox(
        width: 1000,
        child: GridView.builder(
          gridDelegate: SliverGridDelegateWithMaxCrossAxisExtent(
            maxCrossAxisExtent: 200, // Max width of each item
            mainAxisSpacing: 8, // Spacing between rows
            crossAxisSpacing: 8, // Spacing between columns
          ),
          itemCount: data.length, // Number of items
          itemBuilder: (context, index) => LobbyTile(
            item: data[index],
          ),
        ),
      ),
    );
  }
}

class LobbyTile extends StatelessWidget {
  const LobbyTile({super.key, required this.item});

  final Lobby item;

  @override
  Widget build(BuildContext context) {
    final createdAt = DateTime.parse(item.createdAt);
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(8),
        color: Colors.black54,
        shape: BoxShape.rectangle,
        border: Border.all(
          color: Colors.grey,
          width: 2,
        ),
      ),
      child: Padding(
        padding: const EdgeInsets.all(10),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          children: [
            Row(
              children: [
                Text(
                  'Name: ${item.lobbyName}',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 20
                  ),
                ),
              ],
            ),
            Row(
              children: [
                Text(
                  'Players joined: ${item.playerCount}/4',
                  style: TextStyle(color: Colors.white),
                ),
              ],
            ),
            smallDiv,
            Row(
              children: [
                Text(
                  'Added: ${timeago.format(createdAt)}',
                  style: TextStyle(
                      color: Colors.white, overflow: TextOverflow.ellipsis),
                ),
              ],
            ),
            smallDiv,
            Row(
              children: [
                Text(
                  'Created by ${item.ownerName}',
                  style: TextStyle(color: Colors.white),
                ),
              ],
            ),
            actionDivider,
            Row(
              children: [
                LobbyActionBar(item: item),
              ],
            ),
          ],
        ),
      ),
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
    final user = ref.watch(userDataProvider).value;

    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: [
        if (user?.username == item.ownerName)
          IconButton(
            onPressed: () async {
              await runGrpcRequest(
                context,
                () async {
                  await ref.read(lobbyApiProvider).deleteLobby(
                        DelLobbiesRequest(lobby: item),
                      );
                },
              );
            },
            icon: Icon(Icons.delete),
          ),
        item.playerCount >= 4
            ? Text('Lobby is full')
            : ElevatedButton(
                onPressed: () {
                  logger
                      .i("Going to lobby ${item.iD.toInt()} ${item.lobbyName}");
                  // ref.read(lobbyIDProvider.notifier).state = item.iD.toInt();
                  window.location.assign("/game/?lobby=${item.iD}");
                },
                style: globalButtonStyle,
                child: Text('Join'),
              )
      ],
    );
  }
}
