import 'package:flutter/material.dart';
import 'package:flutter_animate/flutter_animate.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/config.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/ui/auth.dart';
import 'package:multipacman/ui/lobby.dart';
import 'package:multipacman/utils.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await PreferencesService.init();
  runApp(ProviderScope(child: const MyApp()));
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Multipacman',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        fontFamily: 'Play',
        colorScheme: ColorScheme.fromSeed(
          seedColor: Colors.deepPurple,
          brightness: Brightness.dark,
        ),
        useMaterial3: true,
      ),
      home: const Root(),
    ).animate().fadeIn(duration: 400.ms);
  }
}

class Root extends ConsumerWidget {
  const Root({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authStatus = ref.watch(userDataProvider);

    return Scaffold(
      backgroundColor: Colors.black,
      body: authStatus.when(
        data: (status) =>
            status != null ? HomeContainerPage() : AuthContainerPage(),
        error: (error, stackTrace) {
          logger.e('Error occurred while checking auth', error: error);

          return Center(
            child: Column(
              children: [
                Text('Error occurred while authenticating'),
                SizedBox(height: 50),
                Text(error.toString()),
              ],
            ),
          );
        },
        loading: () => Center(child: CircularProgressIndicator()),
      ),
    );
  }
}

class HomeContainerPage extends ConsumerWidget {
  const HomeContainerPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final lobbyId = ref.watch(lobbyIDProvider);
    Future(() {
      if (!context.mounted) return;
      ref.modState(deviceWidthProvider, MediaQuery.of(context).size.width);
    });

    return LobbyPage();
  }
}
