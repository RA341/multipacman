import 'package:connectrpc/connect.dart';
import 'package:connectrpc/http2.dart';
import 'package:connectrpc/protobuf.dart';
import 'package:connectrpc/protocol/connect.dart' as protocol;

Transport setupClientChannel(String basePath, Interceptor auth) {
  return protocol.Transport(
    baseUrl: basePath,
    codec: const ProtoCodec(),
    httpClient: createHttpClient(),
    // statusParser: StatusParser(), // This is required for gRPC and gRPC-Web
    interceptors: [auth],
  );
}
