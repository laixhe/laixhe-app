///
//  Generated code. Do not modify.
//  source: update_friends.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'update_friend.pb.dart' as $8;

class UpdateFriends extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateFriends', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..pc<$8.UpdateFriend>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'users', $pb.PbFieldType.PM, subBuilder: $8.UpdateFriend.create)
    ..hasRequiredFields = false
  ;

  UpdateFriends._() : super();
  factory UpdateFriends({
    $core.Iterable<$8.UpdateFriend>? users,
  }) {
    final _result = create();
    if (users != null) {
      _result.users.addAll(users);
    }
    return _result;
  }
  factory UpdateFriends.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateFriends.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateFriends clone() => UpdateFriends()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateFriends copyWith(void Function(UpdateFriends) updates) => super.copyWith((message) => updates(message as UpdateFriends)) as UpdateFriends; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateFriends create() => UpdateFriends._();
  UpdateFriends createEmptyInstance() => create();
  static $pb.PbList<UpdateFriends> createRepeated() => $pb.PbList<UpdateFriends>();
  @$core.pragma('dart2js:noInline')
  static UpdateFriends getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateFriends>(create);
  static UpdateFriends? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$8.UpdateFriend> get users => $_getList(0);
}

