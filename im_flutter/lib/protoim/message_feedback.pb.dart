///
//  Generated code. Do not modify.
//  source: message_feedback.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class MessageFeedback extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MessageFeedback', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'conversationId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'msgId')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'localMsgId')
    ..hasRequiredFields = false
  ;

  MessageFeedback._() : super();
  factory MessageFeedback({
    $core.String? conversationId,
    $core.String? msgId,
    $core.String? localMsgId,
  }) {
    final _result = create();
    if (conversationId != null) {
      _result.conversationId = conversationId;
    }
    if (msgId != null) {
      _result.msgId = msgId;
    }
    if (localMsgId != null) {
      _result.localMsgId = localMsgId;
    }
    return _result;
  }
  factory MessageFeedback.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MessageFeedback.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MessageFeedback clone() => MessageFeedback()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MessageFeedback copyWith(void Function(MessageFeedback) updates) => super.copyWith((message) => updates(message as MessageFeedback)) as MessageFeedback; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MessageFeedback create() => MessageFeedback._();
  MessageFeedback createEmptyInstance() => create();
  static $pb.PbList<MessageFeedback> createRepeated() => $pb.PbList<MessageFeedback>();
  @$core.pragma('dart2js:noInline')
  static MessageFeedback getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MessageFeedback>(create);
  static MessageFeedback? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get conversationId => $_getSZ(0);
  @$pb.TagNumber(1)
  set conversationId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasConversationId() => $_has(0);
  @$pb.TagNumber(1)
  void clearConversationId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get msgId => $_getSZ(1);
  @$pb.TagNumber(2)
  set msgId($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsgId() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsgId() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get localMsgId => $_getSZ(2);
  @$pb.TagNumber(3)
  set localMsgId($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLocalMsgId() => $_has(2);
  @$pb.TagNumber(3)
  void clearLocalMsgId() => clearField(3);
}

