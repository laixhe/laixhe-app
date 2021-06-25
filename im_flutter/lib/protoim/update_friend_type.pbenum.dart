///
//  Generated code. Do not modify.
//  source: update_friend_type.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class UpdateFriendType extends $pb.ProtobufEnum {
  static const UpdateFriendType UF_ONLINE_FALSE = UpdateFriendType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'UF_ONLINE_FALSE');
  static const UpdateFriendType UF_ONLINE_TRUE = UpdateFriendType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'UF_ONLINE_TRUE');
  static const UpdateFriendType UF_ADD = UpdateFriendType._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'UF_ADD');
  static const UpdateFriendType UF_UPDATE = UpdateFriendType._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'UF_UPDATE');
  static const UpdateFriendType UF_DELETE = UpdateFriendType._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'UF_DELETE');

  static const $core.List<UpdateFriendType> values = <UpdateFriendType> [
    UF_ONLINE_FALSE,
    UF_ONLINE_TRUE,
    UF_ADD,
    UF_UPDATE,
    UF_DELETE,
  ];

  static final $core.Map<$core.int, UpdateFriendType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static UpdateFriendType? valueOf($core.int value) => _byValue[value];

  const UpdateFriendType._($core.int v, $core.String n) : super(v, n);
}

