import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final goToRegisterProvider = StateProvider.autoDispose<bool>((ref) {
  return false;
});

class RegisterButton extends ConsumerWidget {
  const RegisterButton({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final isRegister = ref.watch(goToRegisterProvider);

    return ElevatedButton(
      onPressed: () {
        final prevState = ref.read(goToRegisterProvider);
        ref.read(goToRegisterProvider.notifier).state = !prevState;
      },
      child: Text(
        isRegister ? 'Back to Login' : 'Register new account',
        style: TextStyle(fontSize: 15),
      ),
    );
  }
}
