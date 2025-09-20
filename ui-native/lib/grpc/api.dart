import 'dart:io';

import 'package:connectrpc/connect.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/clients/auth_api.dart';
import 'package:multipacman/config.dart';
import 'package:multipacman/utils.dart';
import 'package:universal_html/html.dart' as html;

import 'grpc_native.dart' if (dart.library.html) 'grpc_web.dart';

final apiTokenProvider = Provider<String>((ref) {
  return prefs.getString(authTokenKey) ?? '';
});

final basePathProvider = Provider<String>((ref) {
  // setup for future feature to modify base path from within the client
  final basePath = prefs.getString('basePath');

  final finalPath = kDebugMode
      ? isMobilePlatform()
          ? 'http://192.168.50.111:11200'
          : 'http://localhost:11200'
      : basePath ??
          (kIsWeb
              ? html.window.location.toString()
              : 'https://multipacman.dumbapps.org');

  logger.i('Base path is: $finalPath');

  return finalPath;
});

final grpcChannelProvider = Provider<Transport>((ref) {
  final apiBasePath = ref.watch(basePathProvider);
  final auth = ref.watch(apiTokenProvider);

  final Interceptor logger = <I extends Object, O extends Object>(next) {
    return (req) {
      req.headers.add('Authorization', auth);
      return next(req);
    };
  };

  return setupClientChannel(apiBasePath, logger);
});
