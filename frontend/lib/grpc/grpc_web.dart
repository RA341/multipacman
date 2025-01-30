import 'package:connectrpc/connect.dart';
import 'package:connectrpc/web.dart';
import 'package:connectrpc/protobuf.dart';
import 'package:connectrpc/protocol/grpc_web.dart' as protocol;

Transport setupClientChannel(String basePath, Interceptor auth) {
  return protocol.Transport(
    baseUrl: basePath,
    codec: const ProtoCodec(),
    statusParser: StatusParser(),
    httpClient: createHttpClient(),
    interceptors: [auth],
  );
}
