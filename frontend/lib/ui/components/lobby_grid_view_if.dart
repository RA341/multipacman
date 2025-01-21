import 'package:flutter/material.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pb.dart';
import 'package:multipacman/model/lobby.dart';
import 'package:multipacman/ui/components/lobby_view.dart';
import 'package:timeago/timeago.dart' as timeago;

class LobbyGridViewIf extends StatelessWidget {
  const LobbyGridViewIf({
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
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 30),
        child: GridView.builder(
          gridDelegate: SliverGridDelegateWithMaxCrossAxisExtent(
            maxCrossAxisExtent: 150, // Max width of each item
            mainAxisSpacing: 8, // Spacing between rows
            crossAxisSpacing: 8, // Spacing between columns
          ),
          itemCount: data.length, // Number of items
          itemBuilder: (context, index) {
            final item = data[index];
            final createdAt = DateTime.parse(item.createdAt);

            return GridTile(
              footer: LobbyActionBar(item: item),
              child: Container(
                alignment: Alignment.center,
                decoration: BoxDecoration(
                  color: Colors.blueAccent,
                  borderRadius: BorderRadius.circular(8),
                ),
                child: Column(
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
                      'Created by ${item.ownerName}',
                      style: TextStyle(color: Colors.white),
                    ),
                  ],
                ),
              ),
            );
          },
        ),
      ),
    );
  }
}
