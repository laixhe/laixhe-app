///
//  Generated code. Do not modify.
//  source: app_os.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use appOsDescriptor instead')
const AppOs$json = const {
  '1': 'AppOs',
  '2': const [
    const {'1': 'OS_UNKNOWN', '2': 0},
    const {'1': 'WEB', '2': 1},
    const {'1': 'IOS', '2': 2},
    const {'1': 'ANDROID', '2': 3},
    const {'1': 'PC', '2': 4},
  ],
};

/// Descriptor for `AppOs`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List appOsDescriptor = $convert.base64Decode('CgVBcHBPcxIOCgpPU19VTktOT1dOEAASBwoDV0VCEAESBwoDSU9TEAISCwoHQU5EUk9JRBADEgYKAlBDEAQ=');
