///
//  Generated code. Do not modify.
//  source: error.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use errorDescriptor instead')
const Error$json = const {
  '1': 'Error',
  '2': const [
    const {'1': 'E_UNKNOWN', '2': 0},
    const {'1': 'E_SERVER', '2': 1},
    const {'1': 'E_WRITE_TIMEOUT', '2': 2},
    const {'1': 'E_READ_TIMEOUT', '2': 3},
    const {'1': 'E_ROUTE_NOT_EXIST', '2': 4},
    const {'1': 'E_DECODE', '2': 5},
    const {'1': 'E_ENCODE', '2': 6},
    const {'1': 'E_NOT_LOGIN', '2': 7},
    const {'1': 'E_PARAMETER', '2': 8},
    const {'1': 'E_DB_NOT_DATA', '2': 9},
    const {'1': 'E_DB_SELECT', '2': 10},
    const {'1': 'E_DB_INSERT', '2': 11},
    const {'1': 'E_DB_DELETE', '2': 12},
    const {'1': 'E_DB_UPATE', '2': 13},
    const {'1': 'E_DB_OPERATION', '2': 14},
  ],
};

/// Descriptor for `Error`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List errorDescriptor = $convert.base64Decode('CgVFcnJvchINCglFX1VOS05PV04QABIMCghFX1NFUlZFUhABEhMKD0VfV1JJVEVfVElNRU9VVBACEhIKDkVfUkVBRF9USU1FT1VUEAMSFQoRRV9ST1VURV9OT1RfRVhJU1QQBBIMCghFX0RFQ09ERRAFEgwKCEVfRU5DT0RFEAYSDwoLRV9OT1RfTE9HSU4QBxIPCgtFX1BBUkFNRVRFUhAIEhEKDUVfREJfTk9UX0RBVEEQCRIPCgtFX0RCX1NFTEVDVBAKEg8KC0VfREJfSU5TRVJUEAsSDwoLRV9EQl9ERUxFVEUQDBIOCgpFX0RCX1VQQVRFEA0SEgoORV9EQl9PUEVSQVRJT04QDg==');
