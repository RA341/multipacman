import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/ui/components/ghost_stack.dart';
import 'package:multipacman/ui/components/register_button.dart';
import 'package:multipacman/ui/register.dart';

import 'login.dart';

class AuthContainerPage extends ConsumerWidget {
  const AuthContainerPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final isRegister = ref.watch(goToRegisterProvider);
    return GhostStack(
      child: SingleChildScrollView(
        child: Center(
          child: SizedBox(
            width: 400,
            height: 700,
            child: isRegister ? RegisterPage() : LoginPage(),
          ),
        ),
      ),
    );
  }
}
