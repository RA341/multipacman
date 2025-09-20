// This is a generated file - do not edit.
//
// Generated from lobby/v1/lobby.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use listLobbiesRequestDescriptor instead')
const ListLobbiesRequest$json = {
  '1': 'ListLobbiesRequest',
};

/// Descriptor for `ListLobbiesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listLobbiesRequestDescriptor =
    $convert.base64Decode('ChJMaXN0TG9iYmllc1JlcXVlc3Q=');

@$core.Deprecated('Use listLobbiesResponseDescriptor instead')
const ListLobbiesResponse$json = {
  '1': 'ListLobbiesResponse',
  '2': [
    {
      '1': 'lobbies',
      '3': 1,
      '4': 3,
      '5': 11,
      '6': '.lobby.v1.Lobby',
      '10': 'lobbies'
    },
  ],
};

/// Descriptor for `ListLobbiesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listLobbiesResponseDescriptor = $convert.base64Decode(
    'ChNMaXN0TG9iYmllc1Jlc3BvbnNlEikKB2xvYmJpZXMYASADKAsyDy5sb2JieS52MS5Mb2JieV'
    'IHbG9iYmllcw==');

@$core.Deprecated('Use addLobbiesRequestDescriptor instead')
const AddLobbiesRequest$json = {
  '1': 'AddLobbiesRequest',
  '2': [
    {'1': 'lobby_name', '3': 1, '4': 1, '5': 9, '10': 'lobbyName'},
  ],
};

/// Descriptor for `AddLobbiesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addLobbiesRequestDescriptor = $convert.base64Decode(
    'ChFBZGRMb2JiaWVzUmVxdWVzdBIdCgpsb2JieV9uYW1lGAEgASgJUglsb2JieU5hbWU=');

@$core.Deprecated('Use addLobbiesResponseDescriptor instead')
const AddLobbiesResponse$json = {
  '1': 'AddLobbiesResponse',
};

/// Descriptor for `AddLobbiesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addLobbiesResponseDescriptor =
    $convert.base64Decode('ChJBZGRMb2JiaWVzUmVzcG9uc2U=');

@$core.Deprecated('Use delLobbiesRequestDescriptor instead')
const DelLobbiesRequest$json = {
  '1': 'DelLobbiesRequest',
  '2': [
    {
      '1': 'lobby',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.lobby.v1.Lobby',
      '10': 'lobby'
    },
  ],
};

/// Descriptor for `DelLobbiesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List delLobbiesRequestDescriptor = $convert.base64Decode(
    'ChFEZWxMb2JiaWVzUmVxdWVzdBIlCgVsb2JieRgBIAEoCzIPLmxvYmJ5LnYxLkxvYmJ5UgVsb2'
    'JieQ==');

@$core.Deprecated('Use delLobbiesResponseDescriptor instead')
const DelLobbiesResponse$json = {
  '1': 'DelLobbiesResponse',
};

/// Descriptor for `DelLobbiesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List delLobbiesResponseDescriptor =
    $convert.base64Decode('ChJEZWxMb2JiaWVzUmVzcG9uc2U=');

@$core.Deprecated('Use lobbyDescriptor instead')
const Lobby$json = {
  '1': 'Lobby',
  '2': [
    {'1': 'ID', '3': 1, '4': 1, '5': 4, '10': 'ID'},
    {'1': 'lobby_name', '3': 2, '4': 1, '5': 9, '10': 'lobbyName'},
    {'1': 'ownerName', '3': 4, '4': 1, '5': 9, '10': 'ownerName'},
    {'1': 'ownerId', '3': 5, '4': 1, '5': 4, '10': 'ownerId'},
    {'1': 'created_at', '3': 3, '4': 1, '5': 9, '10': 'createdAt'},
    {'1': 'playerCount', '3': 6, '4': 1, '5': 4, '10': 'playerCount'},
  ],
};

/// Descriptor for `Lobby`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List lobbyDescriptor = $convert.base64Decode(
    'CgVMb2JieRIOCgJJRBgBIAEoBFICSUQSHQoKbG9iYnlfbmFtZRgCIAEoCVIJbG9iYnlOYW1lEh'
    'wKCW93bmVyTmFtZRgEIAEoCVIJb3duZXJOYW1lEhgKB293bmVySWQYBSABKARSB293bmVySWQS'
    'HQoKY3JlYXRlZF9hdBgDIAEoCVIJY3JlYXRlZEF0EiAKC3BsYXllckNvdW50GAYgASgEUgtwbG'
    'F5ZXJDb3VudA==');

const $core.Map<$core.String, $core.dynamic> LobbyServiceBase$json = {
  '1': 'LobbyService',
  '2': [
    {
      '1': 'ListLobbies',
      '2': '.lobby.v1.ListLobbiesRequest',
      '3': '.lobby.v1.ListLobbiesResponse',
      '4': {}
    },
    {
      '1': 'AddLobby',
      '2': '.lobby.v1.AddLobbiesRequest',
      '3': '.lobby.v1.AddLobbiesResponse',
      '4': {}
    },
    {
      '1': 'DeleteLobby',
      '2': '.lobby.v1.DelLobbiesRequest',
      '3': '.lobby.v1.DelLobbiesResponse',
      '4': {}
    },
  ],
};

@$core.Deprecated('Use lobbyServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>>
    LobbyServiceBase$messageJson = {
  '.lobby.v1.ListLobbiesRequest': ListLobbiesRequest$json,
  '.lobby.v1.ListLobbiesResponse': ListLobbiesResponse$json,
  '.lobby.v1.Lobby': Lobby$json,
  '.lobby.v1.AddLobbiesRequest': AddLobbiesRequest$json,
  '.lobby.v1.AddLobbiesResponse': AddLobbiesResponse$json,
  '.lobby.v1.DelLobbiesRequest': DelLobbiesRequest$json,
  '.lobby.v1.DelLobbiesResponse': DelLobbiesResponse$json,
};

/// Descriptor for `LobbyService`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List lobbyServiceDescriptor = $convert.base64Decode(
    'CgxMb2JieVNlcnZpY2USTAoLTGlzdExvYmJpZXMSHC5sb2JieS52MS5MaXN0TG9iYmllc1JlcX'
    'Vlc3QaHS5sb2JieS52MS5MaXN0TG9iYmllc1Jlc3BvbnNlIgASRwoIQWRkTG9iYnkSGy5sb2Ji'
    'eS52MS5BZGRMb2JiaWVzUmVxdWVzdBocLmxvYmJ5LnYxLkFkZExvYmJpZXNSZXNwb25zZSIAEk'
    'oKC0RlbGV0ZUxvYmJ5EhsubG9iYnkudjEuRGVsTG9iYmllc1JlcXVlc3QaHC5sb2JieS52MS5E'
    'ZWxMb2JiaWVzUmVzcG9uc2UiAA==');
