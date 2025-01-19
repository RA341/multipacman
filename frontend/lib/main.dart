import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/config.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/ui/auth.dart';
import 'package:multipacman/ui/lobby.dart';

Future<void> main() async {
  await PreferencesService.init();
  runApp(ProviderScope(child: const MyApp()));
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(
            seedColor: Colors.deepPurple, brightness: Brightness.dark),
        useMaterial3: true,
      ),
      home: const Root(),
    );
  }
}

class Root extends ConsumerWidget {
  const Root({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authStatus = ref.watch(userDataProvider);

    return Scaffold(
      body: authStatus.when(
        data: (status) => status != null ? LobbyPage() : AuthPage(),
        error: (error, stackTrace) => Center(
          child: Column(
            children: [
              Text('Unable to verify authentication status'),
              SizedBox(height: 50),
              Text(error.toString()),
            ],
          ),
        ),
        loading: () => Center(child: CircularProgressIndicator()),
      ),
    );
  }
}
