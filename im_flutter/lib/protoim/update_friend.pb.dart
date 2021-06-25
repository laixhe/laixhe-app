///
//  Generated code. Do not modify.
//  source: update_friend.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'user_info.pb.dart' as $1;

import 'update_friend_type.pbenum.dart' as $7;

class UpdateFriend extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UpdateFriend', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..e<$7.UpdateFriendType>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'tag', $pb.PbFieldType.OE, defaultOrMaker: $7.UpdateFriendType.UF_ONLINE_FALSE, valueOf: $7.UpdateFriendType.valueOf, enumValues: $7.UpdateFriendType.values)
    ..aOM<$1.UserInfo>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'user', subBuilder: $1.UserInfo.create)
    ..hasRequiredFields = false
  ;

  UpdateFriend._() : super();
  factory UpdateFriend({
    $7.UpdateFriendType? tag,
    $1.UserInfo? user,
  }) {
    final _result = create();
    if (tag != null) {
      _result.tag = tag;
    }
    if (user != null) {
      _result.user = user;
    }
    return _result;
  }
  factory UpdateFriend.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateFriend.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateFriend clone() => UpdateFriend()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateFriend copyWith(void Function(UpdateFriend) updates) => super.copyWith((message) => updates(message as UpdateFriend)) as UpdateFriend; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UpdateFriend create() => UpdateFriend._();
  UpdateFriend createEmptyInstance() => create();
  static $pb.PbList<UpdateFriend> createRepeated() => $pb.PbList<UpdateFriend>();
  @$core.pragma('dart2js:noInline')
  static UpdateFriend getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateFriend>(create);
  static UpdateFriend? _defaultInstance;

  @$pb.TagNumber(1)
  $7.UpdateFriendType get tag => $_getN(0);
  @$pb.TagNumber(1)
  set tag($7.UpdateFriendType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTag() => $_has(0);
  @$pb.TagNumber(1)
  void clearTag() => clearField(1);

  @$pb.TagNumber(2)
  $1.UserInfo get user => $_getN(1);
  @$pb.TagNumber(2)
  set user($1.UserInfo v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasUser() => $_has(1);
  @$pb.TagNumber(2)
  void clearUser() => clearField(2);
  @$pb.TagNumber(2)
  $1.UserInfo ensureUser() => $_ensure(1);
}

