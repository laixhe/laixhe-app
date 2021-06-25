///
//  Generated code. Do not modify.
//  source: message_response.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'message_entity_uri.pb.dart' as $3;
import 'message_media.pb.dart' as $4;

import 'chat_type.pbenum.dart' as $5;
import 'message_type.pbenum.dart' as $6;

class MessageResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MessageResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'conversationId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'msgId')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'localMsgId')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'fromId')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'toId')
    ..a<$core.int>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pts', $pb.PbFieldType.O3)
    ..e<$5.ChatType>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatTypeId', $pb.PbFieldType.OE, defaultOrMaker: $5.ChatType.USE_CHAT, valueOf: $5.ChatType.valueOf, enumValues: $5.ChatType.values)
    ..e<$6.MessageType>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'messageTypeId', $pb.PbFieldType.OE, defaultOrMaker: $6.MessageType.TEXT_TYPE, valueOf: $6.MessageType.valueOf, enumValues: $6.MessageType.values)
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'content')
    ..aInt64(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'latestTime')
    ..aOM<$3.MessageEntityUri>(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'entity', subBuilder: $3.MessageEntityUri.create)
    ..aOM<$4.MessageMedia>(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'media', subBuilder: $4.MessageMedia.create)
    ..aOM<$4.MessageMedia>(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'thumbnail', subBuilder: $4.MessageMedia.create)
    ..hasRequiredFields = false
  ;

  MessageResponse._() : super();
  factory MessageResponse({
    $core.String? conversationId,
    $core.String? msgId,
    $core.String? localMsgId,
    $core.String? fromId,
    $core.String? toId,
    $core.int? pts,
    $5.ChatType? chatTypeId,
    $6.MessageType? messageTypeId,
    $core.String? content,
    $fixnum.Int64? latestTime,
    $3.MessageEntityUri? entity,
    $4.MessageMedia? media,
    $4.MessageMedia? thumbnail,
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
    if (fromId != null) {
      _result.fromId = fromId;
    }
    if (toId != null) {
      _result.toId = toId;
    }
    if (pts != null) {
      _result.pts = pts;
    }
    if (chatTypeId != null) {
      _result.chatTypeId = chatTypeId;
    }
    if (messageTypeId != null) {
      _result.messageTypeId = messageTypeId;
    }
    if (content != null) {
      _result.content = content;
    }
    if (latestTime != null) {
      _result.latestTime = latestTime;
    }
    if (entity != null) {
      _result.entity = entity;
    }
    if (media != null) {
      _result.media = media;
    }
    if (thumbnail != null) {
      _result.thumbnail = thumbnail;
    }
    return _result;
  }
  factory MessageResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MessageResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MessageResponse clone() => MessageResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MessageResponse copyWith(void Function(MessageResponse) updates) => super.copyWith((message) => updates(message as MessageResponse)) as MessageResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MessageResponse create() => MessageResponse._();
  MessageResponse createEmptyInstance() => create();
  static $pb.PbList<MessageResponse> createRepeated() => $pb.PbList<MessageResponse>();
  @$core.pragma('dart2js:noInline')
  static MessageResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MessageResponse>(create);
  static MessageResponse? _defaultInstance;

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

  @$pb.TagNumber(4)
  $core.String get fromId => $_getSZ(3);
  @$pb.TagNumber(4)
  set fromId($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFromId() => $_has(3);
  @$pb.TagNumber(4)
  void clearFromId() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get toId => $_getSZ(4);
  @$pb.TagNumber(5)
  set toId($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasToId() => $_has(4);
  @$pb.TagNumber(5)
  void clearToId() => clearField(5);

  @$pb.TagNumber(6)
  $core.int get pts => $_getIZ(5);
  @$pb.TagNumber(6)
  set pts($core.int v) { $_setSignedInt32(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasPts() => $_has(5);
  @$pb.TagNumber(6)
  void clearPts() => clearField(6);

  @$pb.TagNumber(7)
  $5.ChatType get chatTypeId => $_getN(6);
  @$pb.TagNumber(7)
  set chatTypeId($5.ChatType v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasChatTypeId() => $_has(6);
  @$pb.TagNumber(7)
  void clearChatTypeId() => clearField(7);

  @$pb.TagNumber(8)
  $6.MessageType get messageTypeId => $_getN(7);
  @$pb.TagNumber(8)
  set messageTypeId($6.MessageType v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasMessageTypeId() => $_has(7);
  @$pb.TagNumber(8)
  void clearMessageTypeId() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get content => $_getSZ(8);
  @$pb.TagNumber(9)
  set content($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasContent() => $_has(8);
  @$pb.TagNumber(9)
  void clearContent() => clearField(9);

  @$pb.TagNumber(10)
  $fixnum.Int64 get latestTime => $_getI64(9);
  @$pb.TagNumber(10)
  set latestTime($fixnum.Int64 v) { $_setInt64(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasLatestTime() => $_has(9);
  @$pb.TagNumber(10)
  void clearLatestTime() => clearField(10);

  @$pb.TagNumber(11)
  $3.MessageEntityUri get entity => $_getN(10);
  @$pb.TagNumber(11)
  set entity($3.MessageEntityUri v) { setField(11, v); }
  @$pb.TagNumber(11)
  $core.bool hasEntity() => $_has(10);
  @$pb.TagNumber(11)
  void clearEntity() => clearField(11);
  @$pb.TagNumber(11)
  $3.MessageEntityUri ensureEntity() => $_ensure(10);

  @$pb.TagNumber(12)
  $4.MessageMedia get media => $_getN(11);
  @$pb.TagNumber(12)
  set media($4.MessageMedia v) { setField(12, v); }
  @$pb.TagNumber(12)
  $core.bool hasMedia() => $_has(11);
  @$pb.TagNumber(12)
  void clearMedia() => clearField(12);
  @$pb.TagNumber(12)
  $4.MessageMedia ensureMedia() => $_ensure(11);

  @$pb.TagNumber(13)
  $4.MessageMedia get thumbnail => $_getN(12);
  @$pb.TagNumber(13)
  set thumbnail($4.MessageMedia v) { setField(13, v); }
  @$pb.TagNumber(13)
  $core.bool hasThumbnail() => $_has(12);
  @$pb.TagNumber(13)
  void clearThumbnail() => clearField(13);
  @$pb.TagNumber(13)
  $4.MessageMedia ensureThumbnail() => $_ensure(12);
}

