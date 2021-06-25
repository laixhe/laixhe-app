///
//  Generated code. Do not modify.
//  source: message_type.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class MessageType extends $pb.ProtobufEnum {
  static const MessageType TEXT_TYPE = MessageType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TEXT_TYPE');
  static const MessageType IMAGE_TYPE = MessageType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'IMAGE_TYPE');
  static const MessageType FILE_TYPE = MessageType._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'FILE_TYPE');
  static const MessageType VOICE_TYPE = MessageType._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'VOICE_TYPE');
  static const MessageType VIDEO_TYPE = MessageType._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'VIDEO_TYPE');

  static const $core.List<MessageType> values = <MessageType> [
    TEXT_TYPE,
    IMAGE_TYPE,
    FILE_TYPE,
    VOICE_TYPE,
    VIDEO_TYPE,
  ];

  static final $core.Map<$core.int, MessageType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static MessageType? valueOf($core.int value) => _byValue[value];

  const MessageType._($core.int v, $core.String n) : super(v, n);
}

