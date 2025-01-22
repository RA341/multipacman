import 'package:web_socket_channel/web_socket_channel.dart';

void createConnection(Uri wsUrl) {
  final channel = WebSocketChannel.connect(
    wsUrl,
    // https://github.com/orgs/deepgram/discussions/175
    // sends token via the Sec-WebSocket-Protocol header
    // protocols: ["token", token],
  );

}