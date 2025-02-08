import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:multipacman/ui/components/utils.dart';

class ActionButton extends HookConsumerWidget {
  const ActionButton(this.onPress, this.buttonText, {super.key});

  final Future<void> Function() onPress;
  final String buttonText;

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final isLoading = useState(false);

    return ElevatedButton(
      style: globalButtonStyle,
      onPressed: isLoading.value
          ? null
          : () async {
              isLoading.value = true;
              await onPress();
              isLoading.value = false;
            },
      child: isLoading.value
          ? CircularProgressIndicator()
          : Text(
              buttonText,
              style: TextStyle(fontSize: 20),
            ),
    );
  }
}
