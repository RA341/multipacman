import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:multipacman/clients/lobby_api.dart';
import 'package:multipacman/config.dart';
import 'package:multipacman/gen/lobby/v1/lobby.pb.dart';
import 'package:multipacman/grpc/api.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/ui/components/action_button.dart';
import 'package:multipacman/ui/components/ghost_stack.dart';
import 'package:multipacman/ui/components/lobby_container.dart';
import 'package:multipacman/ui/components/utils.dart';
import 'package:multipacman/utils.dart';

class LobbyPage extends StatelessWidget {
  const LobbyPage({super.key});

  @override
  Widget build(BuildContext context) {
    return GhostStack(
      child: Center(
        child: Column(
          children: [
            LobbyBar(),
            SizedBox(height: 50),
            AddLobbyBar(),
            SizedBox(height: 50),
            LobbyGridView(),
          ],
        ),
      ),
    );
  }
}

class AddLobbyBar extends HookConsumerWidget {
  const AddLobbyBar({super.key});

  @override
  Widget build(BuildContext context, ref) {
    final lobbyName = useTextEditingController();

    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        IconButton(
          icon: Icon(Icons.refresh),
          onPressed: () => ref.invalidate(lobbyListProvider),
        ),
        SizedBox(width: 20),
        SizedBox(
          width: 300,
          child: createUpdateButtons2('New lobby name', lobbyName),
        ),
        SizedBox(width: 20),
        ActionButton(
          () async {
            if (lobbyName.text.isEmpty) {
              showErrorDialog(context, 'Please enter a lobby name');
              return;
            }

            await runGrpcRequest(
              context,
              () async {
                await ref
                    .read(lobbyApiProvider)
                    .addLobby(AddLobbiesRequest(lobbyName: lobbyName.text));
              },
            );

            lobbyName.clear();
            // ref.invalidate(lobbyListProvider);
          },
          'Add lobby',
        )
      ],
    );
  }
}

class LobbyBar extends ConsumerWidget {
  const LobbyBar({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.watch(userDataProvider).value!;

    return Padding(
      padding: const EdgeInsets.only(top: 20),
      child: Stack(
        children: [
          Center(
            child: Text(
              'Lobby',
              style: TextStyle(fontSize: 40, fontWeight: FontWeight.bold),
            ),
          ),
          Positioned(
            right: 20,
            child: Column(
              spacing: 10,
              children: [
                Text('Hello ${user.username}'),
                ActionButton(
                  () async {
                    logger.i('Logging out');

                    try {
                      await prefs.setString("apikey", "");
                    } catch (e) {
                      if (!context.mounted) return;
                      logger.e(e.toString());
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          content: Text('An error occurred while logging out'),
                          duration: Duration(seconds: 3),
                        ),
                      );
                    }

                    ref.invalidate(apiTokenProvider);
                  },
                  'Logout',
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
