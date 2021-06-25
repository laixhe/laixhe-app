///
//  Generated code. Do not modify.
//  source: user_login_request.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'app_os.pbenum.dart' as $9;

class UserLoginRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserLoginRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..e<$9.AppOs>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'appOsId', $pb.PbFieldType.OE, defaultOrMaker: $9.AppOs.OS_UNKNOWN, valueOf: $9.AppOs.valueOf, enumValues: $9.AppOs.values)
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'account')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'password')
    ..hasRequiredFields = false
  ;

  UserLoginRequest._() : super();
  factory UserLoginRequest({
    $9.AppOs? appOsId,
    $core.String? account,
    $core.String? password,
  }) {
    final _result = create();
    if (appOsId != null) {
      _result.appOsId = appOsId;
    }
    if (account != null) {
      _result.account = account;
    }
    if (password != null) {
      _result.password = password;
    }
    return _result;
  }
  factory UserLoginRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserLoginRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserLoginRequest clone() => UserLoginRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserLoginRequest copyWith(void Function(UserLoginRequest) updates) => super.copyWith((message) => updates(message as UserLoginRequest)) as UserLoginRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserLoginRequest create() => UserLoginRequest._();
  UserLoginRequest createEmptyInstance() => create();
  static $pb.PbList<UserLoginRequest> createRepeated() => $pb.PbList<UserLoginRequest>();
  @$core.pragma('dart2js:noInline')
  static UserLoginRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserLoginRequest>(create);
  static UserLoginRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $9.AppOs get appOsId => $_getN(0);
  @$pb.TagNumber(1)
  set appOsId($9.AppOs v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasAppOsId() => $_has(0);
  @$pb.TagNumber(1)
  void clearAppOsId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get account => $_getSZ(1);
  @$pb.TagNumber(2)
  set account($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasAccount() => $_has(1);
  @$pb.TagNumber(2)
  void clearAccount() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get password => $_getSZ(2);
  @$pb.TagNumber(3)
  set password($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasPassword() => $_has(2);
  @$pb.TagNumber(3)
  void clearPassword() => clearField(3);
}

