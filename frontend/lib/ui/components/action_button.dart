import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:grpc/grpc_web.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:multipacman/ui/components/utils.dart';

class ActionButton extends HookConsumerWidget {
  const ActionButton(this.onPress, this.buttonText, {super.key});

  final Future<void> Function() onPress;
  final String buttonText;

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final isLoading = useState(false);

    return isLoading.value
        ? CircularProgressIndicator()
        : ElevatedButton(
            onPressed: () async {
              try {
                await onPress();
              } on GrpcError catch (e) {
                if (!context.mounted) return;
                showErrorDialog(context, "An error occurred: ${e.message}");
                showErrorDialog(
                  context,
                  e.message ?? "Unknown error",
                  message: e.details
                          ?.fold<String>(
                            '',
                            (value, element) =>
                                value = '$value\n${element.toString()}',
                          )
                          .toString() ??
                      "",
                  errorMessage: e.toString(),
                );
              } catch (e) {
                if (!context.mounted) return;
                showErrorDialog(
                  context,
                  "An unexpected error occurred",
                  errorMessage: e.toString(),
                );
              }
            },
            child: Text(
              buttonText,
              style: TextStyle(fontSize: 15),
            ),
          );
  }
}
