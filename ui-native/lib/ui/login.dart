import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:multipacman/clients/auth_api.dart';
import 'package:multipacman/grpc/api.dart';
import 'package:multipacman/ui/components/action_button.dart';
import 'package:multipacman/ui/components/register_button.dart';
import 'package:multipacman/ui/components/utils.dart';

class LoginPage extends HookConsumerWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final username = useTextEditingController();
    final password = useTextEditingController();
    final authApi = ref.watch(authApiProvider);

    return Column(
      spacing: 50,
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text(
          'Login',
          style: TextStyle(fontSize: 40),
        ),
        createUpdateButtons2("Username", username),
        createUpdateButtons2("Password", password, secure: true),
        ActionButton(
          () async {
            if (password.text.isEmpty || username.text.isEmpty) {
              showErrorDialog(context, "One or more fields are empty");
            }

            await runGrpcRequest(
              context,
              () async {
                await authApi.login(
                  user: username.text,
                  pass: password.text,
                );
              },
            );
            ref.invalidate(apiTokenProvider);
          },
          'Login',
        ),
        ActionButton(
          () async {
            await runGrpcRequest(
              context,
              authApi.guestLogin,
            );
            ref.invalidate(apiTokenProvider);
          },
          'Guest Login',
        ),
        RegisterButton(),
      ],
    );
  }
}
