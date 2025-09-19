import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:multipacman/clients/auth_api.dart';
import 'package:multipacman/ui/components/action_button.dart';
import 'package:multipacman/ui/components/register_button.dart';
import 'package:multipacman/ui/components/utils.dart';

class RegisterPage extends HookConsumerWidget {
  const RegisterPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authApi = ref.watch(authApiProvider);

    final username = useTextEditingController();
    final password = useTextEditingController();
    final verifyPassword = useTextEditingController();

    return Column(
      spacing: 50,
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text(
          'Register',
          style: TextStyle(fontSize: 40),
        ),
        createUpdateButtons2("Username", username),
        createUpdateButtons2("Password", password, secure: true),
        createUpdateButtons2("Verify Password", verifyPassword, secure: true),
        ActionButton(
          () async {
            if (password.text.isEmpty ||
                username.text.isEmpty ||
                verifyPassword.text.isEmpty) {
              showErrorDialog(context, "One or more fields are empty");
            }

            if (password.text != verifyPassword.text) {
              showErrorDialog(
                context,
                'Mismatched passwords',
                message: 'Check your password and try again',
              );
            }

            await runGrpcRequest(
              context,
              () async {
                await authApi.register(
                  user: username.text,
                  pass: password.text,
                  passVerify: verifyPassword.text,
                );

                ref.read(goToRegisterProvider.notifier).state = false;
              },
            );
          },
          'Register',
        ),
        RegisterButton(),
      ],
    );
  }
}
