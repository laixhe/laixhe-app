///
//  Generated code. Do not modify.
//  source: error.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Error extends $pb.ProtobufEnum {
  static const Error E_UNKNOWN = Error._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_UNKNOWN');
  static const Error E_SERVER = Error._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_SERVER');
  static const Error E_WRITE_TIMEOUT = Error._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_WRITE_TIMEOUT');
  static const Error E_READ_TIMEOUT = Error._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_READ_TIMEOUT');
  static const Error E_ROUTE_NOT_EXIST = Error._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_ROUTE_NOT_EXIST');
  static const Error E_DECODE = Error._(5, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_DECODE');
  static const Error E_ENCODE = Error._(6, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_ENCODE');
  static const Error E_NOT_LOGIN = Error._(7, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_NOT_LOGIN');
  static const Error E_PARAMETER = Error._(8, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_PARAMETER');
  static const Error E_DB_NOT_DATA = Error._(9, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_DB_NOT_DATA');
  static const Error E_DB_SELECT = Error._(10, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_DB_SELECT');
  static const Error E_DB_INSERT = Error._(11, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_DB_INSERT');
  static const Error E_DB_DELETE = Error._(12, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_DB_DELETE');
  static const Error E_DB_UPATE = Error._(13, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_DB_UPATE');
  static const Error E_DB_OPERATION = Error._(14, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'E_DB_OPERATION');

  static const $core.List<Error> values = <Error> [
    E_UNKNOWN,
    E_SERVER,
    E_WRITE_TIMEOUT,
    E_READ_TIMEOUT,
    E_ROUTE_NOT_EXIST,
    E_DECODE,
    E_ENCODE,
    E_NOT_LOGIN,
    E_PARAMETER,
    E_DB_NOT_DATA,
    E_DB_SELECT,
    E_DB_INSERT,
    E_DB_DELETE,
    E_DB_UPATE,
    E_DB_OPERATION,
  ];

  static final $core.Map<$core.int, Error> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Error? valueOf($core.int value) => _byValue[value];

  const Error._($core.int v, $core.String n) : super(v, n);
}

