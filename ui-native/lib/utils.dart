import 'dart:io';

import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:logger/logger.dart';

extension StringExtension on String {
  String truncate([int maxLength = 10]) =>
      length > maxLength ? '${substring(0, maxLength)}...' : this;
}

var logger = Logger(
  printer: PrettyPrinter(
    // Number of method calls to be displayed
    // methodCount: 2,
    // Number of method calls if stacktrace is provided
    errorMethodCount: 8,
    // Width of the output
    lineLength: 120,
    colors: true,
    printEmojis: true,
    // Should each log print contain a timestamp
    dateTimeFormat: DateTimeFormat.onlyTimeAndSinceStart,
  ),
);

Future<void> landscapeModeOnly(double width, bool orient) async {
  if (isMobileWidth(width)) {
    logger.d("Mobile size detected setting screen orientation");
    if (orient) {
      await SystemChrome.setPreferredOrientations([
        DeviceOrientation.landscapeLeft,
        DeviceOrientation.landscapeRight,
      ]);
    } else {
      await SystemChrome.setPreferredOrientations([
        DeviceOrientation.portraitUp,
      ]);
    }
  }
}

extension ModStaeWidgetRef on WidgetRef {
  /// calls .state on provider to change its value
  void modState<T>(StateProvider<T> provider, T value) {
    read(provider.notifier).state = value;
  }
}

extension ModStateRef on Ref {
  /// calls .state on provider to change its value
  void modState<T>(StateProvider<T> provider, T value) {
    read(provider.notifier).state = value;
  }
}

// will fail on web since dart:io is not available on web,
// implying its on web, and returning false
bool isMobilePlatform() {
  try {
    return (Platform.isAndroid || Platform.isIOS);
  } catch (e) {
    return false;
  }
}

bool isMobileWidth(double width) => isMobilePlatform() && width < 450.0;
