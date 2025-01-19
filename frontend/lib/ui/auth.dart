import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/ui/components/register_button.dart';
import 'package:multipacman/ui/login.dart';
import 'package:multipacman/ui/register.dart';

class AuthPage extends ConsumerWidget {
  const AuthPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final isRegister = ref.watch(goToRegisterProvider);
    return isRegister ? RegisterPage() : LoginPage();
  }
}
