///
//  Generated code. Do not modify.
//  source: message_response.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use messageResponseDescriptor instead')
const MessageResponse$json = const {
  '1': 'MessageResponse',
  '2': const [
    const {'1': 'conversation_id', '3': 1, '4': 1, '5': 9, '10': 'conversationId'},
    const {'1': 'msg_id', '3': 2, '4': 1, '5': 9, '10': 'msgId'},
    const {'1': 'local_msg_id', '3': 3, '4': 1, '5': 9, '10': 'localMsgId'},
    const {'1': 'from_id', '3': 4, '4': 1, '5': 9, '10': 'fromId'},
    const {'1': 'to_id', '3': 5, '4': 1, '5': 9, '10': 'toId'},
    const {'1': 'pts', '3': 6, '4': 1, '5': 5, '10': 'pts'},
    const {'1': 'chat_type_id', '3': 7, '4': 1, '5': 14, '6': '.protoim.ChatType', '10': 'chatTypeId'},
    const {'1': 'message_type_id', '3': 8, '4': 1, '5': 14, '6': '.protoim.MessageType', '10': 'messageTypeId'},
    const {'1': 'content', '3': 9, '4': 1, '5': 9, '10': 'content'},
    const {'1': 'latest_time', '3': 10, '4': 1, '5': 3, '10': 'latestTime'},
    const {'1': 'entity', '3': 11, '4': 1, '5': 11, '6': '.protoim.MessageEntityUri', '10': 'entity'},
    const {'1': 'media', '3': 12, '4': 1, '5': 11, '6': '.protoim.MessageMedia', '10': 'media'},
    const {'1': 'thumbnail', '3': 13, '4': 1, '5': 11, '6': '.protoim.MessageMedia', '10': 'thumbnail'},
  ],
};

/// Descriptor for `MessageResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageResponseDescriptor = $convert.base64Decode('Cg9NZXNzYWdlUmVzcG9uc2USJwoPY29udmVyc2F0aW9uX2lkGAEgASgJUg5jb252ZXJzYXRpb25JZBIVCgZtc2dfaWQYAiABKAlSBW1zZ0lkEiAKDGxvY2FsX21zZ19pZBgDIAEoCVIKbG9jYWxNc2dJZBIXCgdmcm9tX2lkGAQgASgJUgZmcm9tSWQSEwoFdG9faWQYBSABKAlSBHRvSWQSEAoDcHRzGAYgASgFUgNwdHMSMwoMY2hhdF90eXBlX2lkGAcgASgOMhEucHJvdG9pbS5DaGF0VHlwZVIKY2hhdFR5cGVJZBI8Cg9tZXNzYWdlX3R5cGVfaWQYCCABKA4yFC5wcm90b2ltLk1lc3NhZ2VUeXBlUg1tZXNzYWdlVHlwZUlkEhgKB2NvbnRlbnQYCSABKAlSB2NvbnRlbnQSHwoLbGF0ZXN0X3RpbWUYCiABKANSCmxhdGVzdFRpbWUSMQoGZW50aXR5GAsgASgLMhkucHJvdG9pbS5NZXNzYWdlRW50aXR5VXJpUgZlbnRpdHkSKwoFbWVkaWEYDCABKAsyFS5wcm90b2ltLk1lc3NhZ2VNZWRpYVIFbWVkaWESMwoJdGh1bWJuYWlsGA0gASgLMhUucHJvdG9pbS5NZXNzYWdlTWVkaWFSCXRodW1ibmFpbA==');
