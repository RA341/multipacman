//
//  Generated code. Do not modify.
//  source: lobby/v1/lobby.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class ListLobbiesRequest extends $pb.GeneratedMessage {
  factory ListLobbiesRequest() => create();
  ListLobbiesRequest._() : super();
  factory ListLobbiesRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListLobbiesRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ListLobbiesRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListLobbiesRequest clone() => ListLobbiesRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListLobbiesRequest copyWith(void Function(ListLobbiesRequest) updates) => super.copyWith((message) => updates(message as ListLobbiesRequest)) as ListLobbiesRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListLobbiesRequest create() => ListLobbiesRequest._();
  ListLobbiesRequest createEmptyInstance() => create();
  static $pb.PbList<ListLobbiesRequest> createRepeated() => $pb.PbList<ListLobbiesRequest>();
  @$core.pragma('dart2js:noInline')
  static ListLobbiesRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListLobbiesRequest>(create);
  static ListLobbiesRequest? _defaultInstance;
}

class ListLobbiesResponse extends $pb.GeneratedMessage {
  factory ListLobbiesResponse({
    $core.Iterable<Lobby>? lobbies,
  }) {
    final $result = create();
    if (lobbies != null) {
      $result.lobbies.addAll(lobbies);
    }
    return $result;
  }
  ListLobbiesResponse._() : super();
  factory ListLobbiesResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListLobbiesResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ListLobbiesResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'), createEmptyInstance: create)
    ..pc<Lobby>(1, _omitFieldNames ? '' : 'lobbies', $pb.PbFieldType.PM, subBuilder: Lobby.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListLobbiesResponse clone() => ListLobbiesResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListLobbiesResponse copyWith(void Function(ListLobbiesResponse) updates) => super.copyWith((message) => updates(message as ListLobbiesResponse)) as ListLobbiesResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListLobbiesResponse create() => ListLobbiesResponse._();
  ListLobbiesResponse createEmptyInstance() => create();
  static $pb.PbList<ListLobbiesResponse> createRepeated() => $pb.PbList<ListLobbiesResponse>();
  @$core.pragma('dart2js:noInline')
  static ListLobbiesResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListLobbiesResponse>(create);
  static ListLobbiesResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<Lobby> get lobbies => $_getList(0);
}

class AddLobbiesRequest extends $pb.GeneratedMessage {
  factory AddLobbiesRequest({
    $core.String? lobbyName,
  }) {
    final $result = create();
    if (lobbyName != null) {
      $result.lobbyName = lobbyName;
    }
    return $result;
  }
  AddLobbiesRequest._() : super();
  factory AddLobbiesRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddLobbiesRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AddLobbiesRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'lobbyName')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddLobbiesRequest clone() => AddLobbiesRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddLobbiesRequest copyWith(void Function(AddLobbiesRequest) updates) => super.copyWith((message) => updates(message as AddLobbiesRequest)) as AddLobbiesRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AddLobbiesRequest create() => AddLobbiesRequest._();
  AddLobbiesRequest createEmptyInstance() => create();
  static $pb.PbList<AddLobbiesRequest> createRepeated() => $pb.PbList<AddLobbiesRequest>();
  @$core.pragma('dart2js:noInline')
  static AddLobbiesRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddLobbiesRequest>(create);
  static AddLobbiesRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get lobbyName => $_getSZ(0);
  @$pb.TagNumber(1)
  set lobbyName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLobbyName() => $_has(0);
  @$pb.TagNumber(1)
  void clearLobbyName() => $_clearField(1);
}

class AddLobbiesResponse extends $pb.GeneratedMessage {
  factory AddLobbiesResponse() => create();
  AddLobbiesResponse._() : super();
  factory AddLobbiesResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddLobbiesResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AddLobbiesResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddLobbiesResponse clone() => AddLobbiesResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddLobbiesResponse copyWith(void Function(AddLobbiesResponse) updates) => super.copyWith((message) => updates(message as AddLobbiesResponse)) as AddLobbiesResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AddLobbiesResponse create() => AddLobbiesResponse._();
  AddLobbiesResponse createEmptyInstance() => create();
  static $pb.PbList<AddLobbiesResponse> createRepeated() => $pb.PbList<AddLobbiesResponse>();
  @$core.pragma('dart2js:noInline')
  static AddLobbiesResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddLobbiesResponse>(create);
  static AddLobbiesResponse? _defaultInstance;
}

class DelLobbiesRequest extends $pb.GeneratedMessage {
  factory DelLobbiesRequest({
    Lobby? lobby,
  }) {
    final $result = create();
    if (lobby != null) {
      $result.lobby = lobby;
    }
    return $result;
  }
  DelLobbiesRequest._() : super();
  factory DelLobbiesRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DelLobbiesRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DelLobbiesRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'), createEmptyInstance: create)
    ..aOM<Lobby>(1, _omitFieldNames ? '' : 'lobby', subBuilder: Lobby.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DelLobbiesRequest clone() => DelLobbiesRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DelLobbiesRequest copyWith(void Function(DelLobbiesRequest) updates) => super.copyWith((message) => updates(message as DelLobbiesRequest)) as DelLobbiesRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DelLobbiesRequest create() => DelLobbiesRequest._();
  DelLobbiesRequest createEmptyInstance() => create();
  static $pb.PbList<DelLobbiesRequest> createRepeated() => $pb.PbList<DelLobbiesRequest>();
  @$core.pragma('dart2js:noInline')
  static DelLobbiesRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DelLobbiesRequest>(create);
  static DelLobbiesRequest? _defaultInstance;

  @$pb.TagNumber(1)
  Lobby get lobby => $_getN(0);
  @$pb.TagNumber(1)
  set lobby(Lobby v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasLobby() => $_has(0);
  @$pb.TagNumber(1)
  void clearLobby() => $_clearField(1);
  @$pb.TagNumber(1)
  Lobby ensureLobby() => $_ensure(0);
}

class DelLobbiesResponse extends $pb.GeneratedMessage {
  factory DelLobbiesResponse() => create();
  DelLobbiesResponse._() : super();
  factory DelLobbiesResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DelLobbiesResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DelLobbiesResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DelLobbiesResponse clone() => DelLobbiesResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DelLobbiesResponse copyWith(void Function(DelLobbiesResponse) updates) => super.copyWith((message) => updates(message as DelLobbiesResponse)) as DelLobbiesResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DelLobbiesResponse create() => DelLobbiesResponse._();
  DelLobbiesResponse createEmptyInstance() => create();
  static $pb.PbList<DelLobbiesResponse> createRepeated() => $pb.PbList<DelLobbiesResponse>();
  @$core.pragma('dart2js:noInline')
  static DelLobbiesResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DelLobbiesResponse>(create);
  static DelLobbiesResponse? _defaultInstance;
}

class Lobby extends $pb.GeneratedMessage {
  factory Lobby({
    $fixnum.Int64? iD,
    $core.String? lobbyName,
    $core.String? createdAt,
    $core.String? ownerName,
    $fixnum.Int64? ownerId,
    $fixnum.Int64? playerCount,
  }) {
    final $result = create();
    if (iD != null) {
      $result.iD = iD;
    }
    if (lobbyName != null) {
      $result.lobbyName = lobbyName;
    }
    if (createdAt != null) {
      $result.createdAt = createdAt;
    }
    if (ownerName != null) {
      $result.ownerName = ownerName;
    }
    if (ownerId != null) {
      $result.ownerId = ownerId;
    }
    if (playerCount != null) {
      $result.playerCount = playerCount;
    }
    return $result;
  }
  Lobby._() : super();
  factory Lobby.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Lobby.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Lobby', package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, _omitFieldNames ? '' : 'ID', $pb.PbFieldType.OU6, protoName: 'ID', defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(2, _omitFieldNames ? '' : 'lobbyName')
    ..aOS(3, _omitFieldNames ? '' : 'createdAt')
    ..aOS(4, _omitFieldNames ? '' : 'ownerName', protoName: 'ownerName')
    ..a<$fixnum.Int64>(5, _omitFieldNames ? '' : 'ownerId', $pb.PbFieldType.OU6, protoName: 'ownerId', defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(6, _omitFieldNames ? '' : 'playerCount', $pb.PbFieldType.OU6, protoName: 'playerCount', defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Lobby clone() => Lobby()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Lobby copyWith(void Function(Lobby) updates) => super.copyWith((message) => updates(message as Lobby)) as Lobby;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Lobby create() => Lobby._();
  Lobby createEmptyInstance() => create();
  static $pb.PbList<Lobby> createRepeated() => $pb.PbList<Lobby>();
  @$core.pragma('dart2js:noInline')
  static Lobby getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Lobby>(create);
  static Lobby? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get iD => $_getI64(0);
  @$pb.TagNumber(1)
  set iD($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get lobbyName => $_getSZ(1);
  @$pb.TagNumber(2)
  set lobbyName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLobbyName() => $_has(1);
  @$pb.TagNumber(2)
  void clearLobbyName() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get createdAt => $_getSZ(2);
  @$pb.TagNumber(3)
  set createdAt($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasCreatedAt() => $_has(2);
  @$pb.TagNumber(3)
  void clearCreatedAt() => $_clearField(3);

  @$pb.TagNumber(4)
  $core.String get ownerName => $_getSZ(3);
  @$pb.TagNumber(4)
  set ownerName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasOwnerName() => $_has(3);
  @$pb.TagNumber(4)
  void clearOwnerName() => $_clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get ownerId => $_getI64(4);
  @$pb.TagNumber(5)
  set ownerId($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasOwnerId() => $_has(4);
  @$pb.TagNumber(5)
  void clearOwnerId() => $_clearField(5);

  @$pb.TagNumber(6)
  $fixnum.Int64 get playerCount => $_getI64(5);
  @$pb.TagNumber(6)
  set playerCount($fixnum.Int64 v) { $_setInt64(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasPlayerCount() => $_has(5);
  @$pb.TagNumber(6)
  void clearPlayerCount() => $_clearField(6);
}

class LobbyServiceApi {
  $pb.RpcClient _client;
  LobbyServiceApi(this._client);

  /// todo figure out lobby streaming
  $async.Future<ListLobbiesResponse> listLobbies($pb.ClientContext? ctx, ListLobbiesRequest request) =>
    _client.invoke<ListLobbiesResponse>(ctx, 'LobbyService', 'ListLobbies', request, ListLobbiesResponse())
  ;
  $async.Future<AddLobbiesResponse> addLobby($pb.ClientContext? ctx, AddLobbiesRequest request) =>
    _client.invoke<AddLobbiesResponse>(ctx, 'LobbyService', 'AddLobby', request, AddLobbiesResponse())
  ;
  $async.Future<DelLobbiesResponse> deleteLobby($pb.ClientContext? ctx, DelLobbiesRequest request) =>
    _client.invoke<DelLobbiesResponse>(ctx, 'LobbyService', 'DeleteLobby', request, DelLobbiesResponse())
  ;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
