///
//  Generated code. Do not modify.
//  source: cmd.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class CMD extends $pb.ProtobufEnum {
  static const CMD C_UNKNOWN = CMD._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_UNKNOWN');
  static const CMD C_ERROR = CMD._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_ERROR');
  static const CMD C_PING = CMD._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_PING');
  static const CMD C_PONG = CMD._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_PONG');
  static const CMD C_USER_LOGIN_REQUEST = CMD._(1000, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_USER_LOGIN_REQUEST');
  static const CMD C_USER_LOGIN_RESPONSE = CMD._(1001, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_USER_LOGIN_RESPONSE');
  static const CMD C_GET_USER_REQUEST = CMD._(1002, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_GET_USER_REQUEST');
  static const CMD C_GET_USER_RESPONSE = CMD._(1003, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_GET_USER_RESPONSE');
  static const CMD C_GET_FRIENDS_REQUEST = CMD._(1004, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_GET_FRIENDS_REQUEST');
  static const CMD C_GET_FRIENDS_RESPONSE = CMD._(1005, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_GET_FRIENDS_RESPONSE');
  static const CMD C_UPDATE_FRIENDS = CMD._(1006, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_UPDATE_FRIENDS');
  static const CMD C_GET_CONVERSATIONS_REQUEST = CMD._(1007, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_GET_CONVERSATIONS_REQUEST');
  static const CMD C_GET_CONVERSATIONS_RESPONSE = CMD._(1008, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_GET_CONVERSATIONS_RESPONSE');
  static const CMD C_MESSAGE_REQUEST = CMD._(1009, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_MESSAGE_REQUEST');
  static const CMD C_MESSAGE_RESPONSE = CMD._(1010, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_MESSAGE_RESPONSE');
  static const CMD C_MESSAGE_FEEDBACK = CMD._(1011, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'C_MESSAGE_FEEDBACK');

  static const $core.List<CMD> values = <CMD> [
    C_UNKNOWN,
    C_ERROR,
    C_PING,
    C_PONG,
    C_USER_LOGIN_REQUEST,
    C_USER_LOGIN_RESPONSE,
    C_GET_USER_REQUEST,
    C_GET_USER_RESPONSE,
    C_GET_FRIENDS_REQUEST,
    C_GET_FRIENDS_RESPONSE,
    C_UPDATE_FRIENDS,
    C_GET_CONVERSATIONS_REQUEST,
    C_GET_CONVERSATIONS_RESPONSE,
    C_MESSAGE_REQUEST,
    C_MESSAGE_RESPONSE,
    C_MESSAGE_FEEDBACK,
  ];

  static final $core.Map<$core.int, CMD> _byValue = $pb.ProtobufEnum.initByValue(values);
  static CMD? valueOf($core.int value) => _byValue[value];

  const CMD._($core.int v, $core.String n) : super(v, n);
}

