///
//  Generated code. Do not modify.
//  source: get_friends_response.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'user_info.pb.dart' as $1;

class GetFriendsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetFriendsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..pc<$1.UserInfo>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'users', $pb.PbFieldType.PM, subBuilder: $1.UserInfo.create)
    ..hasRequiredFields = false
  ;

  GetFriendsResponse._() : super();
  factory GetFriendsResponse({
    $core.Iterable<$1.UserInfo>? users,
  }) {
    final _result = create();
    if (users != null) {
      _result.users.addAll(users);
    }
    return _result;
  }
  factory GetFriendsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetFriendsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetFriendsResponse clone() => GetFriendsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetFriendsResponse copyWith(void Function(GetFriendsResponse) updates) => super.copyWith((message) => updates(message as GetFriendsResponse)) as GetFriendsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetFriendsResponse create() => GetFriendsResponse._();
  GetFriendsResponse createEmptyInstance() => create();
  static $pb.PbList<GetFriendsResponse> createRepeated() => $pb.PbList<GetFriendsResponse>();
  @$core.pragma('dart2js:noInline')
  static GetFriendsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetFriendsResponse>(create);
  static GetFriendsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$1.UserInfo> get users => $_getList(0);
}

