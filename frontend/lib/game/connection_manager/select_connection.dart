import 'package:flutter/foundation.dart';
import 'package:universal_html/html.dart' as html;
import 'package:web_socket_channel/web_socket_channel.dart';

void createConnection(Uri wsUrl, String token) {
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
}

void setCookie({
  required String name,
  required String value,
}) {
  html.document.cookie = '$name=$value; httpOnly; secure';
}
