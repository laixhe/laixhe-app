///
//  Generated code. Do not modify.
//  source: message_base.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'cmd.pbenum.dart' as $2;

class MessageBase extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MessageBase', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..e<$2.CMD>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'cmd', $pb.PbFieldType.OE, defaultOrMaker: $2.CMD.C_UNKNOWN, valueOf: $2.CMD.valueOf, enumValues: $2.CMD.values)
    ..a<$core.List<$core.int>>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'data', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  MessageBase._() : super();
  factory MessageBase({
    $2.CMD? cmd,
    $core.List<$core.int>? data,
  }) {
    final _result = create();
    if (cmd != null) {
      _result.cmd = cmd;
    }
    if (data != null) {
      _result.data = data;
    }
    return _result;
  }
  factory MessageBase.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MessageBase.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MessageBase clone() => MessageBase()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MessageBase copyWith(void Function(MessageBase) updates) => super.copyWith((message) => updates(message as MessageBase)) as MessageBase; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MessageBase create() => MessageBase._();
  MessageBase createEmptyInstance() => create();
  static $pb.PbList<MessageBase> createRepeated() => $pb.PbList<MessageBase>();
  @$core.pragma('dart2js:noInline')
  static MessageBase getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MessageBase>(create);
  static MessageBase? _defaultInstance;

  @$pb.TagNumber(1)
  $2.CMD get cmd => $_getN(0);
  @$pb.TagNumber(1)
  set cmd($2.CMD v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasCmd() => $_has(0);
  @$pb.TagNumber(1)
  void clearCmd() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get data => $_getN(1);
  @$pb.TagNumber(2)
  set data($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasData() => $_has(1);
  @$pb.TagNumber(2)
  void clearData() => clearField(2);
}

