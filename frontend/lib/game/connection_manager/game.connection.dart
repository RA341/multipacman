import 'package:flutter/foundation.dart';
import 'package:multipacman/utils.dart';
import 'package:universal_html/html.dart' as html;
import 'package:web_socket_channel/web_socket_channel.dart';

Future<WebSocketChannel> createConnection(Uri wsUrl, String token) async {
  if (kIsWeb) {
    // set cookie on web platform
    setCookie(name: 'auth', value: token);
  }

  final channel = WebSocketChannel.connect(
    wsUrl,
    // https://github.com/orgs/deepgram/discussions/175
    // sends token via the Sec-WebSocket-Protocol header
    // do not add on web platform, since it breaks websockets for some reason
    protocols: kIsWeb ? null : [token],
  );

  await channel.ready;
  logger.i("Channel connected and ready");
  return channel;
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

// void setCookie({
//   required String name,
//   required String value,
// }) {
//   html.document.cookie = '$name=$value; httpOnly; secure';
// }
