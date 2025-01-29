import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/clients/lobby_api.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pb.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/ui/components/utils.dart';
import 'package:multipacman/utils.dart';
import 'package:timeago/timeago.dart' as timeago;

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
    return GridTile(
      footer: LobbyActionBar(item: item),
      child: Container(
        decoration: BoxDecoration(
          color: Colors.blue,
          borderRadius: BorderRadius.circular(8),
        ),
        child: Padding(
          padding: const EdgeInsets.all(10),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              Text(
                'Name: ${item.lobbyName}',
                style: TextStyle(color: Colors.white),
              ),
              Text(
                'Created: ${timeago.format(createdAt)}',
                style: TextStyle(color: Colors.white),
              ),
              Text(
                'Players joined: ${item.playerCount}',
                style: TextStyle(color: Colors.white),
              ),
              Text(
                'Created by ${item.ownerName}',
                style: TextStyle(color: Colors.white),
              ),
              SizedBox(height: 15),
              Row(
                children: [
                  Text(
                    'Created by ${item.ownerName}',
                    style: TextStyle(color: Colors.white),
                  ),
                ],
              ),
            ],
          ),
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
        ElevatedButton(
          onPressed: () {
            logger.i("Going to lobby ${item.iD.toInt()} ${item.lobbyName}");
            ref.read(lobbyIDProvider.notifier).state = item.iD.toInt();
          },
          child: Text('Join'),
        )
      ],
    );
  }
}
