// This is a generated file - do not edit.
//
// Generated from lobby/v1/lobby.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class ListLobbiesRequest extends $pb.GeneratedMessage {
  factory ListLobbiesRequest() => create();

  ListLobbiesRequest._();

  factory ListLobbiesRequest.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory ListLobbiesRequest.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'ListLobbiesRequest',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListLobbiesRequest clone() => ListLobbiesRequest()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListLobbiesRequest copyWith(void Function(ListLobbiesRequest) updates) =>
      super.copyWith((message) => updates(message as ListLobbiesRequest))
          as ListLobbiesRequest;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListLobbiesRequest create() => ListLobbiesRequest._();
  @$core.override
  ListLobbiesRequest createEmptyInstance() => create();
  static $pb.PbList<ListLobbiesRequest> createRepeated() =>
      $pb.PbList<ListLobbiesRequest>();
  @$core.pragma('dart2js:noInline')
  static ListLobbiesRequest getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<ListLobbiesRequest>(create);
  static ListLobbiesRequest? _defaultInstance;
}

class ListLobbiesResponse extends $pb.GeneratedMessage {
  factory ListLobbiesResponse({
    $core.Iterable<Lobby>? lobbies,
  }) {
    final result = create();
    if (lobbies != null) result.lobbies.addAll(lobbies);
    return result;
  }

  ListLobbiesResponse._();

  factory ListLobbiesResponse.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory ListLobbiesResponse.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'ListLobbiesResponse',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'),
      createEmptyInstance: create)
    ..pc<Lobby>(1, _omitFieldNames ? '' : 'lobbies', $pb.PbFieldType.PM,
        subBuilder: Lobby.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListLobbiesResponse clone() => ListLobbiesResponse()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListLobbiesResponse copyWith(void Function(ListLobbiesResponse) updates) =>
      super.copyWith((message) => updates(message as ListLobbiesResponse))
          as ListLobbiesResponse;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListLobbiesResponse create() => ListLobbiesResponse._();
  @$core.override
  ListLobbiesResponse createEmptyInstance() => create();
  static $pb.PbList<ListLobbiesResponse> createRepeated() =>
      $pb.PbList<ListLobbiesResponse>();
  @$core.pragma('dart2js:noInline')
  static ListLobbiesResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<ListLobbiesResponse>(create);
  static ListLobbiesResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<Lobby> get lobbies => $_getList(0);
}

class AddLobbiesRequest extends $pb.GeneratedMessage {
  factory AddLobbiesRequest({
    $core.String? lobbyName,
  }) {
    final result = create();
    if (lobbyName != null) result.lobbyName = lobbyName;
    return result;
  }

  AddLobbiesRequest._();

  factory AddLobbiesRequest.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory AddLobbiesRequest.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'AddLobbiesRequest',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'lobbyName')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AddLobbiesRequest clone() => AddLobbiesRequest()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AddLobbiesRequest copyWith(void Function(AddLobbiesRequest) updates) =>
      super.copyWith((message) => updates(message as AddLobbiesRequest))
          as AddLobbiesRequest;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AddLobbiesRequest create() => AddLobbiesRequest._();
  @$core.override
  AddLobbiesRequest createEmptyInstance() => create();
  static $pb.PbList<AddLobbiesRequest> createRepeated() =>
      $pb.PbList<AddLobbiesRequest>();
  @$core.pragma('dart2js:noInline')
  static AddLobbiesRequest getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<AddLobbiesRequest>(create);
  static AddLobbiesRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get lobbyName => $_getSZ(0);
  @$pb.TagNumber(1)
  set lobbyName($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasLobbyName() => $_has(0);
  @$pb.TagNumber(1)
  void clearLobbyName() => $_clearField(1);
}

class AddLobbiesResponse extends $pb.GeneratedMessage {
  factory AddLobbiesResponse() => create();

  AddLobbiesResponse._();

  factory AddLobbiesResponse.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory AddLobbiesResponse.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'AddLobbiesResponse',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AddLobbiesResponse clone() => AddLobbiesResponse()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AddLobbiesResponse copyWith(void Function(AddLobbiesResponse) updates) =>
      super.copyWith((message) => updates(message as AddLobbiesResponse))
          as AddLobbiesResponse;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AddLobbiesResponse create() => AddLobbiesResponse._();
  @$core.override
  AddLobbiesResponse createEmptyInstance() => create();
  static $pb.PbList<AddLobbiesResponse> createRepeated() =>
      $pb.PbList<AddLobbiesResponse>();
  @$core.pragma('dart2js:noInline')
  static AddLobbiesResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<AddLobbiesResponse>(create);
  static AddLobbiesResponse? _defaultInstance;
}

class DelLobbiesRequest extends $pb.GeneratedMessage {
  factory DelLobbiesRequest({
    Lobby? lobby,
  }) {
    final result = create();
    if (lobby != null) result.lobby = lobby;
    return result;
  }

  DelLobbiesRequest._();

  factory DelLobbiesRequest.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory DelLobbiesRequest.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'DelLobbiesRequest',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'),
      createEmptyInstance: create)
    ..aOM<Lobby>(1, _omitFieldNames ? '' : 'lobby', subBuilder: Lobby.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DelLobbiesRequest clone() => DelLobbiesRequest()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DelLobbiesRequest copyWith(void Function(DelLobbiesRequest) updates) =>
      super.copyWith((message) => updates(message as DelLobbiesRequest))
          as DelLobbiesRequest;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DelLobbiesRequest create() => DelLobbiesRequest._();
  @$core.override
  DelLobbiesRequest createEmptyInstance() => create();
  static $pb.PbList<DelLobbiesRequest> createRepeated() =>
      $pb.PbList<DelLobbiesRequest>();
  @$core.pragma('dart2js:noInline')
  static DelLobbiesRequest getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<DelLobbiesRequest>(create);
  static DelLobbiesRequest? _defaultInstance;

  @$pb.TagNumber(1)
  Lobby get lobby => $_getN(0);
  @$pb.TagNumber(1)
  set lobby(Lobby value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasLobby() => $_has(0);
  @$pb.TagNumber(1)
  void clearLobby() => $_clearField(1);
  @$pb.TagNumber(1)
  Lobby ensureLobby() => $_ensure(0);
}

class DelLobbiesResponse extends $pb.GeneratedMessage {
  factory DelLobbiesResponse() => create();

  DelLobbiesResponse._();

  factory DelLobbiesResponse.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory DelLobbiesResponse.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'DelLobbiesResponse',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DelLobbiesResponse clone() => DelLobbiesResponse()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DelLobbiesResponse copyWith(void Function(DelLobbiesResponse) updates) =>
      super.copyWith((message) => updates(message as DelLobbiesResponse))
          as DelLobbiesResponse;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DelLobbiesResponse create() => DelLobbiesResponse._();
  @$core.override
  DelLobbiesResponse createEmptyInstance() => create();
  static $pb.PbList<DelLobbiesResponse> createRepeated() =>
      $pb.PbList<DelLobbiesResponse>();
  @$core.pragma('dart2js:noInline')
  static DelLobbiesResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<DelLobbiesResponse>(create);
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
    final result = create();
    if (iD != null) result.iD = iD;
    if (lobbyName != null) result.lobbyName = lobbyName;
    if (createdAt != null) result.createdAt = createdAt;
    if (ownerName != null) result.ownerName = ownerName;
    if (ownerId != null) result.ownerId = ownerId;
    if (playerCount != null) result.playerCount = playerCount;
    return result;
  }

  Lobby._();

  factory Lobby.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory Lobby.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Lobby',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'lobby.v1'),
      createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, _omitFieldNames ? '' : 'ID', $pb.PbFieldType.OU6,
        protoName: 'ID', defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(2, _omitFieldNames ? '' : 'lobbyName')
    ..aOS(3, _omitFieldNames ? '' : 'createdAt')
    ..aOS(4, _omitFieldNames ? '' : 'ownerName', protoName: 'ownerName')
    ..a<$fixnum.Int64>(5, _omitFieldNames ? '' : 'ownerId', $pb.PbFieldType.OU6,
        protoName: 'ownerId', defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(
        6, _omitFieldNames ? '' : 'playerCount', $pb.PbFieldType.OU6,
        protoName: 'playerCount', defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Lobby clone() => Lobby()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Lobby copyWith(void Function(Lobby) updates) =>
      super.copyWith((message) => updates(message as Lobby)) as Lobby;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Lobby create() => Lobby._();
  @$core.override
  Lobby createEmptyInstance() => create();
  static $pb.PbList<Lobby> createRepeated() => $pb.PbList<Lobby>();
  @$core.pragma('dart2js:noInline')
  static Lobby getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Lobby>(create);
  static Lobby? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get iD => $_getI64(0);
  @$pb.TagNumber(1)
  set iD($fixnum.Int64 value) => $_setInt64(0, value);
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get lobbyName => $_getSZ(1);
  @$pb.TagNumber(2)
  set lobbyName($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasLobbyName() => $_has(1);
  @$pb.TagNumber(2)
  void clearLobbyName() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get createdAt => $_getSZ(2);
  @$pb.TagNumber(3)
  set createdAt($core.String value) => $_setString(2, value);
  @$pb.TagNumber(3)
  $core.bool hasCreatedAt() => $_has(2);
  @$pb.TagNumber(3)
  void clearCreatedAt() => $_clearField(3);

  @$pb.TagNumber(4)
  $core.String get ownerName => $_getSZ(3);
  @$pb.TagNumber(4)
  set ownerName($core.String value) => $_setString(3, value);
  @$pb.TagNumber(4)
  $core.bool hasOwnerName() => $_has(3);
  @$pb.TagNumber(4)
  void clearOwnerName() => $_clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get ownerId => $_getI64(4);
  @$pb.TagNumber(5)
  set ownerId($fixnum.Int64 value) => $_setInt64(4, value);
  @$pb.TagNumber(5)
  $core.bool hasOwnerId() => $_has(4);
  @$pb.TagNumber(5)
  void clearOwnerId() => $_clearField(5);

  @$pb.TagNumber(6)
  $fixnum.Int64 get playerCount => $_getI64(5);
  @$pb.TagNumber(6)
  set playerCount($fixnum.Int64 value) => $_setInt64(5, value);
  @$pb.TagNumber(6)
  $core.bool hasPlayerCount() => $_has(5);
  @$pb.TagNumber(6)
  void clearPlayerCount() => $_clearField(6);
}

class LobbyServiceApi {
  final $pb.RpcClient _client;

  LobbyServiceApi(this._client);

  /// todo figure out lobby streaming
  $async.Future<ListLobbiesResponse> listLobbies(
          $pb.ClientContext? ctx, ListLobbiesRequest request) =>
      _client.invoke<ListLobbiesResponse>(
          ctx, 'LobbyService', 'ListLobbies', request, ListLobbiesResponse());
  $async.Future<AddLobbiesResponse> addLobby(
          $pb.ClientContext? ctx, AddLobbiesRequest request) =>
      _client.invoke<AddLobbiesResponse>(
          ctx, 'LobbyService', 'AddLobby', request, AddLobbiesResponse());
  $async.Future<DelLobbiesResponse> deleteLobby(
          $pb.ClientContext? ctx, DelLobbiesRequest request) =>
      _client.invoke<DelLobbiesResponse>(
          ctx, 'LobbyService', 'DeleteLobby', request, DelLobbiesResponse());
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
