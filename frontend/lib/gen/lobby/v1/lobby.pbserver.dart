//
//  Generated code. Do not modify.
//  source: lobby/v1/lobby.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'lobby.pb.dart' as $1;
import 'lobby.pbjson.dart';

export 'lobby.pb.dart';

abstract class LobbyServiceBase extends $pb.GeneratedService {
  $async.Future<$1.ListLobbiesResponse> listLobbies($pb.ServerContext ctx, $1.ListLobbiesRequest request);
  $async.Future<$1.AddLobbiesResponse> addLobby($pb.ServerContext ctx, $1.AddLobbiesRequest request);
  $async.Future<$1.DelLobbiesResponse> deleteLobby($pb.ServerContext ctx, $1.DelLobbiesRequest request);

  $pb.GeneratedMessage createRequest($core.String methodName) {
    switch (methodName) {
      case 'ListLobbies': return $1.ListLobbiesRequest();
      case 'AddLobby': return $1.AddLobbiesRequest();
      case 'DeleteLobby': return $1.DelLobbiesRequest();
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String methodName, $pb.GeneratedMessage request) {
    switch (methodName) {
      case 'ListLobbies': return this.listLobbies(ctx, request as $1.ListLobbiesRequest);
      case 'AddLobby': return this.addLobby(ctx, request as $1.AddLobbiesRequest);
      case 'DeleteLobby': return this.deleteLobby(ctx, request as $1.DelLobbiesRequest);
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => LobbyServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => LobbyServiceBase$messageJson;
}

