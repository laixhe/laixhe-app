///
//  Generated code. Do not modify.
//  source: message_request.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use messageRequestDescriptor instead')
const MessageRequest$json = const {
  '1': 'MessageRequest',
  '2': const [
    const {'1': 'conversation_id', '3': 1, '4': 1, '5': 9, '10': 'conversationId'},
    const {'1': 'local_msg_id', '3': 2, '4': 1, '5': 9, '10': 'localMsgId'},
    const {'1': 'from_id', '3': 3, '4': 1, '5': 9, '10': 'fromId'},
    const {'1': 'to_id', '3': 4, '4': 1, '5': 9, '10': 'toId'},
    const {'1': 'chat_type_id', '3': 5, '4': 1, '5': 14, '6': '.protoim.ChatType', '10': 'chatTypeId'},
    const {'1': 'message_type_id', '3': 6, '4': 1, '5': 14, '6': '.protoim.MessageType', '10': 'messageTypeId'},
    const {'1': 'content', '3': 7, '4': 1, '5': 9, '10': 'content'},
    const {'1': 'entity', '3': 8, '4': 1, '5': 11, '6': '.protoim.MessageEntityUri', '10': 'entity'},
    const {'1': 'media', '3': 9, '4': 1, '5': 11, '6': '.protoim.MessageMedia', '10': 'media'},
    const {'1': 'thumbnail', '3': 10, '4': 1, '5': 11, '6': '.protoim.MessageMedia', '10': 'thumbnail'},
  ],
};

/// Descriptor for `MessageRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageRequestDescriptor = $convert.base64Decode('Cg5NZXNzYWdlUmVxdWVzdBInCg9jb252ZXJzYXRpb25faWQYASABKAlSDmNvbnZlcnNhdGlvbklkEiAKDGxvY2FsX21zZ19pZBgCIAEoCVIKbG9jYWxNc2dJZBIXCgdmcm9tX2lkGAMgASgJUgZmcm9tSWQSEwoFdG9faWQYBCABKAlSBHRvSWQSMwoMY2hhdF90eXBlX2lkGAUgASgOMhEucHJvdG9pbS5DaGF0VHlwZVIKY2hhdFR5cGVJZBI8Cg9tZXNzYWdlX3R5cGVfaWQYBiABKA4yFC5wcm90b2ltLk1lc3NhZ2VUeXBlUg1tZXNzYWdlVHlwZUlkEhgKB2NvbnRlbnQYByABKAlSB2NvbnRlbnQSMQoGZW50aXR5GAggASgLMhkucHJvdG9pbS5NZXNzYWdlRW50aXR5VXJpUgZlbnRpdHkSKwoFbWVkaWEYCSABKAsyFS5wcm90b2ltLk1lc3NhZ2VNZWRpYVIFbWVkaWESMwoJdGh1bWJuYWlsGAogASgLMhUucHJvdG9pbS5NZXNzYWdlTWVkaWFSCXRodW1ibmFpbA==');
