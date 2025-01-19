//
//  Generated code. Do not modify.
//  source: lobby/v1/lobby.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'lobby.pb.dart' as $1;

export 'lobby.pb.dart';

@$pb.GrpcServiceName('lobby.v1.LobbyService')
class LobbyServiceClient extends $grpc.Client {
  static final _$listLobbies = $grpc.ClientMethod<$1.ListLobbiesRequest, $1.ListLobbiesResponse>(
      '/lobby.v1.LobbyService/ListLobbies',
      ($1.ListLobbiesRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.ListLobbiesResponse.fromBuffer(value));
  static final _$addLobby = $grpc.ClientMethod<$1.AddLobbiesRequest, $1.AddLobbiesResponse>(
      '/lobby.v1.LobbyService/AddLobby',
      ($1.AddLobbiesRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.AddLobbiesResponse.fromBuffer(value));
  static final _$deleteLobby = $grpc.ClientMethod<$1.DelLobbiesRequest, $1.DelLobbiesResponse>(
      '/lobby.v1.LobbyService/DeleteLobby',
      ($1.DelLobbiesRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.DelLobbiesResponse.fromBuffer(value));

  LobbyServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$1.ListLobbiesResponse> listLobbies($1.ListLobbiesRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$listLobbies, request, options: options);
  }

  $grpc.ResponseFuture<$1.AddLobbiesResponse> addLobby($1.AddLobbiesRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$addLobby, request, options: options);
  }

  $grpc.ResponseFuture<$1.DelLobbiesResponse> deleteLobby($1.DelLobbiesRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteLobby, request, options: options);
  }
}

@$pb.GrpcServiceName('lobby.v1.LobbyService')
abstract class LobbyServiceBase extends $grpc.Service {
  $core.String get $name => 'lobby.v1.LobbyService';

  LobbyServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.ListLobbiesRequest, $1.ListLobbiesResponse>(
        'ListLobbies',
        listLobbies_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.ListLobbiesRequest.fromBuffer(value),
        ($1.ListLobbiesResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.AddLobbiesRequest, $1.AddLobbiesResponse>(
        'AddLobby',
        addLobby_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.AddLobbiesRequest.fromBuffer(value),
        ($1.AddLobbiesResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.DelLobbiesRequest, $1.DelLobbiesResponse>(
        'DeleteLobby',
        deleteLobby_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.DelLobbiesRequest.fromBuffer(value),
        ($1.DelLobbiesResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.ListLobbiesResponse> listLobbies_Pre($grpc.ServiceCall call, $async.Future<$1.ListLobbiesRequest> request) async {
    return listLobbies(call, await request);
  }

  $async.Future<$1.AddLobbiesResponse> addLobby_Pre($grpc.ServiceCall call, $async.Future<$1.AddLobbiesRequest> request) async {
    return addLobby(call, await request);
  }

  $async.Future<$1.DelLobbiesResponse> deleteLobby_Pre($grpc.ServiceCall call, $async.Future<$1.DelLobbiesRequest> request) async {
    return deleteLobby(call, await request);
  }

  $async.Future<$1.ListLobbiesResponse> listLobbies($grpc.ServiceCall call, $1.ListLobbiesRequest request);
  $async.Future<$1.AddLobbiesResponse> addLobby($grpc.ServiceCall call, $1.AddLobbiesRequest request);
  $async.Future<$1.DelLobbiesResponse> deleteLobby($grpc.ServiceCall call, $1.DelLobbiesRequest request);
}
