///
//  Generated code. Do not modify.
//  source: get_friends_request.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class GetFriendsRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetFriendsRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  GetFriendsRequest._() : super();
  factory GetFriendsRequest() => create();
  factory GetFriendsRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetFriendsRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetFriendsRequest clone() => GetFriendsRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetFriendsRequest copyWith(void Function(GetFriendsRequest) updates) => super.copyWith((message) => updates(message as GetFriendsRequest)) as GetFriendsRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetFriendsRequest create() => GetFriendsRequest._();
  GetFriendsRequest createEmptyInstance() => create();
  static $pb.PbList<GetFriendsRequest> createRepeated() => $pb.PbList<GetFriendsRequest>();
  @$core.pragma('dart2js:noInline')
  static GetFriendsRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetFriendsRequest>(create);
  static GetFriendsRequest? _defaultInstance;
}

