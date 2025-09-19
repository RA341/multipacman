// This is a generated file - do not edit.
//
// Generated from auth/v1/auth.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use emptyDescriptor instead')
const Empty$json = {
  '1': 'Empty',
};

/// Descriptor for `Empty`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List emptyDescriptor =
    $convert.base64Decode('CgVFbXB0eQ==');

@$core.Deprecated('Use authRequestDescriptor instead')
const AuthRequest$json = {
  '1': 'AuthRequest',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `AuthRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List authRequestDescriptor = $convert.base64Decode(
    'CgtBdXRoUmVxdWVzdBIaCgh1c2VybmFtZRgBIAEoCVIIdXNlcm5hbWUSGgoIcGFzc3dvcmQYAi'
    'ABKAlSCHBhc3N3b3Jk');

@$core.Deprecated('Use userResponseDescriptor instead')
const UserResponse$json = {
  '1': 'UserResponse',
  '2': [
    {'1': 'ID', '3': 1, '4': 1, '5': 4, '10': 'ID'},
    {'1': 'username', '3': 2, '4': 1, '5': 9, '10': 'username'},
    {'1': 'authToken', '3': 3, '4': 1, '5': 9, '10': 'authToken'},
    {'1': 'isGuest', '3': 4, '4': 1, '5': 8, '10': 'isGuest'},
  ],
};

/// Descriptor for `UserResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userResponseDescriptor = $convert.base64Decode(
    'CgxVc2VyUmVzcG9uc2USDgoCSUQYASABKARSAklEEhoKCHVzZXJuYW1lGAIgASgJUgh1c2Vybm'
    'FtZRIcCglhdXRoVG9rZW4YAyABKAlSCWF1dGhUb2tlbhIYCgdpc0d1ZXN0GAQgASgIUgdpc0d1'
    'ZXN0');

@$core.Deprecated('Use authResponseDescriptor instead')
const AuthResponse$json = {
  '1': 'AuthResponse',
  '2': [
    {'1': 'authToken', '3': 1, '4': 1, '5': 9, '10': 'authToken'},
  ],
};

/// Descriptor for `AuthResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List authResponseDescriptor = $convert.base64Decode(
    'CgxBdXRoUmVzcG9uc2USHAoJYXV0aFRva2VuGAEgASgJUglhdXRoVG9rZW4=');

@$core.Deprecated('Use testResponseDescriptor instead')
const TestResponse$json = {
  '1': 'TestResponse',
};

/// Descriptor for `TestResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List testResponseDescriptor =
    $convert.base64Decode('CgxUZXN0UmVzcG9uc2U=');

@$core.Deprecated('Use registerUserRequestDescriptor instead')
const RegisterUserRequest$json = {
  '1': 'RegisterUserRequest',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
    {'1': 'passwordVerify', '3': 3, '4': 1, '5': 9, '10': 'passwordVerify'},
  ],
};

/// Descriptor for `RegisterUserRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List registerUserRequestDescriptor = $convert.base64Decode(
    'ChNSZWdpc3RlclVzZXJSZXF1ZXN0EhoKCHVzZXJuYW1lGAEgASgJUgh1c2VybmFtZRIaCghwYX'
    'Nzd29yZBgCIAEoCVIIcGFzc3dvcmQSJgoOcGFzc3dvcmRWZXJpZnkYAyABKAlSDnBhc3N3b3Jk'
    'VmVyaWZ5');

@$core.Deprecated('Use registerUserResponseDescriptor instead')
const RegisterUserResponse$json = {
  '1': 'RegisterUserResponse',
};

/// Descriptor for `RegisterUserResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List registerUserResponseDescriptor =
    $convert.base64Decode('ChRSZWdpc3RlclVzZXJSZXNwb25zZQ==');

const $core.Map<$core.String, $core.dynamic> AuthServiceBase$json = {
  '1': 'AuthService',
  '2': [
    {
      '1': 'Login',
      '2': '.auth.v1.AuthRequest',
      '3': '.auth.v1.UserResponse',
      '4': {}
    },
    {
      '1': 'Register',
      '2': '.auth.v1.RegisterUserRequest',
      '3': '.auth.v1.RegisterUserResponse',
      '4': {}
    },
    {
      '1': 'Test',
      '2': '.auth.v1.AuthResponse',
      '3': '.auth.v1.UserResponse',
      '4': {}
    },
    {
      '1': 'GuestLogin',
      '2': '.auth.v1.Empty',
      '3': '.auth.v1.UserResponse',
      '4': {}
    },
    {'1': 'Logout', '2': '.auth.v1.Empty', '3': '.auth.v1.Empty', '4': {}},
  ],
};

@$core.Deprecated('Use authServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>>
    AuthServiceBase$messageJson = {
  '.auth.v1.AuthRequest': AuthRequest$json,
  '.auth.v1.UserResponse': UserResponse$json,
  '.auth.v1.RegisterUserRequest': RegisterUserRequest$json,
  '.auth.v1.RegisterUserResponse': RegisterUserResponse$json,
  '.auth.v1.AuthResponse': AuthResponse$json,
  '.auth.v1.Empty': Empty$json,
};

/// Descriptor for `AuthService`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List authServiceDescriptor = $convert.base64Decode(
    'CgtBdXRoU2VydmljZRI2CgVMb2dpbhIULmF1dGgudjEuQXV0aFJlcXVlc3QaFS5hdXRoLnYxLl'
    'VzZXJSZXNwb25zZSIAEkkKCFJlZ2lzdGVyEhwuYXV0aC52MS5SZWdpc3RlclVzZXJSZXF1ZXN0'
    'Gh0uYXV0aC52MS5SZWdpc3RlclVzZXJSZXNwb25zZSIAEjYKBFRlc3QSFS5hdXRoLnYxLkF1dG'
    'hSZXNwb25zZRoVLmF1dGgudjEuVXNlclJlc3BvbnNlIgASNQoKR3Vlc3RMb2dpbhIOLmF1dGgu'
    'djEuRW1wdHkaFS5hdXRoLnYxLlVzZXJSZXNwb25zZSIAEioKBkxvZ291dBIOLmF1dGgudjEuRW'
    '1wdHkaDi5hdXRoLnYxLkVtcHR5IgA=');
