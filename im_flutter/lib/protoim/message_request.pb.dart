///
//  Generated code. Do not modify.
//  source: message_request.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'message_entity_uri.pb.dart' as $3;
import 'message_media.pb.dart' as $4;

import 'chat_type.pbenum.dart' as $5;
import 'message_type.pbenum.dart' as $6;

class MessageRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MessageRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'protoim'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'conversationId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'localMsgId')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'fromId')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'toId')
    ..e<$5.ChatType>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chatTypeId', $pb.PbFieldType.OE, defaultOrMaker: $5.ChatType.USE_CHAT, valueOf: $5.ChatType.valueOf, enumValues: $5.ChatType.values)
    ..e<$6.MessageType>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'messageTypeId', $pb.PbFieldType.OE, defaultOrMaker: $6.MessageType.TEXT_TYPE, valueOf: $6.MessageType.valueOf, enumValues: $6.MessageType.values)
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'content')
    ..aOM<$3.MessageEntityUri>(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'entity', subBuilder: $3.MessageEntityUri.create)
    ..aOM<$4.MessageMedia>(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'media', subBuilder: $4.MessageMedia.create)
    ..aOM<$4.MessageMedia>(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'thumbnail', subBuilder: $4.MessageMedia.create)
    ..hasRequiredFields = false
  ;

  MessageRequest._() : super();
  factory MessageRequest({
    $core.String? conversationId,
    $core.String? localMsgId,
    $core.String? fromId,
    $core.String? toId,
    $5.ChatType? chatTypeId,
    $6.MessageType? messageTypeId,
    $core.String? content,
    $3.MessageEntityUri? entity,
    $4.MessageMedia? media,
    $4.MessageMedia? thumbnail,
  }) {
    final _result = create();
    if (conversationId != null) {
      _result.conversationId = conversationId;
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
    if (chatTypeId != null) {
      _result.chatTypeId = chatTypeId;
    }
    if (messageTypeId != null) {
      _result.messageTypeId = messageTypeId;
    }
    if (content != null) {
      _result.content = content;
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
  factory MessageRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MessageRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MessageRequest clone() => MessageRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MessageRequest copyWith(void Function(MessageRequest) updates) => super.copyWith((message) => updates(message as MessageRequest)) as MessageRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MessageRequest create() => MessageRequest._();
  MessageRequest createEmptyInstance() => create();
  static $pb.PbList<MessageRequest> createRepeated() => $pb.PbList<MessageRequest>();
  @$core.pragma('dart2js:noInline')
  static MessageRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MessageRequest>(create);
  static MessageRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get conversationId => $_getSZ(0);
  @$pb.TagNumber(1)
  set conversationId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasConversationId() => $_has(0);
  @$pb.TagNumber(1)
  void clearConversationId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get localMsgId => $_getSZ(1);
  @$pb.TagNumber(2)
  set localMsgId($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLocalMsgId() => $_has(1);
  @$pb.TagNumber(2)
  void clearLocalMsgId() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get fromId => $_getSZ(2);
  @$pb.TagNumber(3)
  set fromId($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasFromId() => $_has(2);
  @$pb.TagNumber(3)
  void clearFromId() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get toId => $_getSZ(3);
  @$pb.TagNumber(4)
  set toId($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasToId() => $_has(3);
  @$pb.TagNumber(4)
  void clearToId() => clearField(4);

  @$pb.TagNumber(5)
  $5.ChatType get chatTypeId => $_getN(4);
  @$pb.TagNumber(5)
  set chatTypeId($5.ChatType v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasChatTypeId() => $_has(4);
  @$pb.TagNumber(5)
  void clearChatTypeId() => clearField(5);

  @$pb.TagNumber(6)
  $6.MessageType get messageTypeId => $_getN(5);
  @$pb.TagNumber(6)
  set messageTypeId($6.MessageType v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasMessageTypeId() => $_has(5);
  @$pb.TagNumber(6)
  void clearMessageTypeId() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get content => $_getSZ(6);
  @$pb.TagNumber(7)
  set content($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasContent() => $_has(6);
  @$pb.TagNumber(7)
  void clearContent() => clearField(7);

  @$pb.TagNumber(8)
  $3.MessageEntityUri get entity => $_getN(7);
  @$pb.TagNumber(8)
  set entity($3.MessageEntityUri v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasEntity() => $_has(7);
  @$pb.TagNumber(8)
  void clearEntity() => clearField(8);
  @$pb.TagNumber(8)
  $3.MessageEntityUri ensureEntity() => $_ensure(7);

  @$pb.TagNumber(9)
  $4.MessageMedia get media => $_getN(8);
  @$pb.TagNumber(9)
  set media($4.MessageMedia v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasMedia() => $_has(8);
  @$pb.TagNumber(9)
  void clearMedia() => clearField(9);
  @$pb.TagNumber(9)
  $4.MessageMedia ensureMedia() => $_ensure(8);

  @$pb.TagNumber(10)
  $4.MessageMedia get thumbnail => $_getN(9);
  @$pb.TagNumber(10)
  set thumbnail($4.MessageMedia v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasThumbnail() => $_has(9);
  @$pb.TagNumber(10)
  void clearThumbnail() => clearField(10);
  @$pb.TagNumber(10)
  $4.MessageMedia ensureThumbnail() => $_ensure(9);
}

