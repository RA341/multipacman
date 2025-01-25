import 'package:flutter/foundation.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:multipacman/grpc/api.dart';
import 'package:multipacman/providers.dart';
import 'package:multipacman/utils.dart';
import 'package:universal_html/html.dart' as html;
import 'package:web_socket_channel/web_socket_channel.dart';

final gameStreamProvider =
    StreamProvider.autoDispose<List<Map<String, dynamic>>>((ref) async* {
  final lobbyId = ref.watch(lobbyIDProvider);
  if (lobbyId == 0) {
    throw Exception("Invalid lobby ID");
  }

  final token = ref.watch(apiTokenProvider);
  final baseUrl = Uri.parse(ref.watch(basePathProvider));

  final wsUrl = getWsUrl(baseUrl, lobbyId);

  final channel = createConnection(wsUrl, token);

  try {
    await channel.ready;
  } catch (e) {
    logger.e(e, error: e);
    rethrow;
  }

  logger.i("Channel connected and ready");

  channel.stream.listen(
    (message) {
      print('Received: $message');
    },
    onError: (error) {
      print('Error: $error');
      // Handle reconnection logic here if needed
    },
    onDone: () {
      print('Connection closed');
      // Handle reconnection logic here if needed
    },
  );

  yield List.empty();
});

WebSocketChannel createConnection(Uri wsUrl, String token) {
  if (kIsWeb) {
    // set cookie on web platform
    setCookie(name: 'auth', value: token);
  }

  return WebSocketChannel.connect(
    wsUrl,
    // https://github.com/orgs/deepgram/discussions/175
    // sends token via the Sec-WebSocket-Protocol header
    // do not add on web platform, since it breaks websockets for some reason
    protocols: kIsWeb ? null : [token],
  );
}

void setCookie({
  required String name,
  required String value,
}) {
  html.document.cookie = '$name=$value';
}

Uri getWsUrl(Uri baseUrl, int lobby) {
  var host = '';
  if (baseUrl.scheme == "https") {
    host = 'wss://${baseUrl.host}';
  } else {
    host = 'ws://${baseUrl.host}:${baseUrl.port}';
  }

  return Uri.parse('$host/api/game').replace(queryParameters: {
    'lobby': lobby.toString(),
  });
}
