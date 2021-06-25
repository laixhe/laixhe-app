///
//  Generated code. Do not modify.
//  source: message_type.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use messageTypeDescriptor instead')
const MessageType$json = const {
  '1': 'MessageType',
  '2': const [
    const {'1': 'TEXT_TYPE', '2': 0},
    const {'1': 'IMAGE_TYPE', '2': 1},
    const {'1': 'FILE_TYPE', '2': 2},
    const {'1': 'VOICE_TYPE', '2': 3},
    const {'1': 'VIDEO_TYPE', '2': 4},
  ],
};

/// Descriptor for `MessageType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List messageTypeDescriptor = $convert.base64Decode('CgtNZXNzYWdlVHlwZRINCglURVhUX1RZUEUQABIOCgpJTUFHRV9UWVBFEAESDQoJRklMRV9UWVBFEAISDgoKVk9JQ0VfVFlQRRADEg4KClZJREVPX1RZUEUQBA==');
