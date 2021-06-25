///
//  Generated code. Do not modify.
//  source: app_os.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class AppOs extends $pb.ProtobufEnum {
  static const AppOs OS_UNKNOWN = AppOs._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'OS_UNKNOWN');
  static const AppOs WEB = AppOs._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'WEB');
  static const AppOs IOS = AppOs._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'IOS');
  static const AppOs ANDROID = AppOs._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ANDROID');
  static const AppOs PC = AppOs._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'PC');

  static const $core.List<AppOs> values = <AppOs> [
    OS_UNKNOWN,
    WEB,
    IOS,
    ANDROID,
    PC,
  ];

  static final $core.Map<$core.int, AppOs> _byValue = $pb.ProtobufEnum.initByValue(values);
  static AppOs? valueOf($core.int value) => _byValue[value];

  const AppOs._($core.int v, $core.String n) : super(v, n);
}

