import 'package:flutter/material.dart';
import 'package:grpc/grpc_or_grpcweb.dart';
import 'package:multipacman/utils.dart';

void showErrorDialog(
  BuildContext context,
  String title, {
  String message = '',
  String errorMessage = '',
}) {
  showDialog(
    context: context,
    barrierDismissible: false,
    builder: (BuildContext context) {
      return AlertDialog(
        title: Text(
          title,
          style: const TextStyle(
            color: Colors.green,
            fontWeight: FontWeight.bold,
          ),
        ),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            ...message.isEmpty
                ? [
                    Text(
                      message,
                      style: const TextStyle(
                        fontSize: 16,
                        fontWeight: FontWeight.w500,
                      ),
                    ),
                    const SizedBox(height: 16)
                  ]
                : [SizedBox()],
            ...errorMessage.isNotEmpty
                ? [
                    Text(
                      'The following was the error message',
                      style: const TextStyle(
                        fontSize: 16,
                        fontWeight: FontWeight.w300,
                      ),
                    ),
                    const SizedBox(height: 16),
                    Container(
                      width: double.infinity,
                      padding: const EdgeInsets.all(12),
                      decoration: BoxDecoration(
                        color: Colors.grey[900],
                        borderRadius: BorderRadius.circular(8),
                        border: Border.all(
                          color: Colors.black,
                          width: 1,
                        ),
                      ),
                      child: SelectableText(
                        errorMessage,
                        style: TextStyle(
                          fontFamily: 'monospace',
                          fontSize: 14,
                          color: Colors.white,
                        ),
                      ),
                    ),
                  ]
                : [SizedBox()],
          ],
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.of(context).pop(),
            child: const Text('OK'),
          ),
        ],
      );
    },
  );
}

Widget createUpdateButtons2(
  String input,
  TextEditingController controller, {
  void Function()? editingController,
  String hintText = '',
  bool secure = false,
  List<String> autofillHints = const [],
}) {
  return TextField(
    autofillHints: autofillHints,
    controller: controller,
    onEditingComplete: editingController,
    obscureText: secure,
    decoration: InputDecoration(
      border: OutlineInputBorder(),
      labelText: input,
      hintText: hintText,
    ),
  );
}

Widget createDropDown2(
  List<String> options,
  TextEditingController controller,
  String hintText,
) {
  final selectedIndex = options.indexOf(controller.text);
  if (controller.text.isEmpty) {
    controller.text = options.first;
  }

  return DropdownMenu<String>(
    label: Text(hintText),
    requestFocusOnTap: false,
    // disable text editing
    initialSelection: options[selectedIndex == -1 ? 0 : selectedIndex],
    onSelected: (String? value) => controller.text = value ?? options.first,
    dropdownMenuEntries: options.map<DropdownMenuEntry<String>>((String value) {
      return DropdownMenuEntry<String>(
        value: value,
        label: value,
      );
    }).toList(),
  );
}

Future<void> runGrpcRequest(
  BuildContext context,
  Future<void> Function() onPress,
) async {
  try {
    await onPress();
  } on GrpcError catch (e) {
    if (!context.mounted) return;
    logger.e(
      "Error while making a grpc request",
      error: e,
    );

    showErrorDialog(
      context,
      e.message ?? "Unknown error",
      message: e.details
              ?.fold<String>(
                '',
                (value, element) => value = '$value\n${element.toString()}',
              )
              .toString() ??
          "",
      errorMessage: e.toString(),
    );
  } catch (e) {
    if (!context.mounted) return;
    logger.e(
      "Unknown error",
      error: e,
    );
    showErrorDialog(
      context,
      "An unexpected error occurred",
      errorMessage: e.toString(),
    );
  }
}

final globalButtonStyle = ElevatedButton.styleFrom(
  shape: RoundedRectangleBorder(
    borderRadius: BorderRadius.circular(5.0), // Rounded corners
    side: const BorderSide(
      color: Colors.grey,
      width: 2.0,
    ),
  ),
  backgroundColor: Colors.black,
  padding: EdgeInsets.symmetric(
    vertical: 15,
    horizontal: 15,
  ),
);

const actionDivider = Divider(
  height: 15,
  color: Colors.grey,
);

const smallDiv = Divider(
  height: 15,
  color: Colors.grey,
  endIndent: 30.0,
);
