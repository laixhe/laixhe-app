///
//  Generated code. Do not modify.
//  source: get_dialogs_response.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class GetDialogsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetDialogsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  GetDialogsResponse._() : super();
  factory GetDialogsResponse() => create();
  factory GetDialogsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetDialogsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetDialogsResponse clone() => GetDialogsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetDialogsResponse copyWith(void Function(GetDialogsResponse) updates) => super.copyWith((message) => updates(message as GetDialogsResponse)) as GetDialogsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetDialogsResponse create() => GetDialogsResponse._();
  GetDialogsResponse createEmptyInstance() => create();
  static $pb.PbList<GetDialogsResponse> createRepeated() => $pb.PbList<GetDialogsResponse>();
  @$core.pragma('dart2js:noInline')
  static GetDialogsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetDialogsResponse>(create);
  static GetDialogsResponse? _defaultInstance;
}

