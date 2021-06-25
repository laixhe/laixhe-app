///
//  Generated code. Do not modify.
//  source: cmd.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use cMDDescriptor instead')
const CMD$json = const {
  '1': 'CMD',
  '2': const [
    const {'1': 'C_UNKNOWN', '2': 0},
    const {'1': 'C_ERROR', '2': 1},
    const {'1': 'C_PING', '2': 2},
    const {'1': 'C_PONG', '2': 3},
    const {'1': 'C_USER_LOGIN_REQUEST', '2': 1000},
    const {'1': 'C_USER_LOGIN_RESPONSE', '2': 1001},
    const {'1': 'C_GET_USER_REQUEST', '2': 1002},
    const {'1': 'C_GET_USER_RESPONSE', '2': 1003},
    const {'1': 'C_GET_FRIENDS_REQUEST', '2': 1004},
    const {'1': 'C_GET_FRIENDS_RESPONSE', '2': 1005},
    const {'1': 'C_UPDATE_FRIENDS', '2': 1006},
    const {'1': 'C_GET_CONVERSATIONS_REQUEST', '2': 1007},
    const {'1': 'C_GET_CONVERSATIONS_RESPONSE', '2': 1008},
    const {'1': 'C_MESSAGE_REQUEST', '2': 1009},
    const {'1': 'C_MESSAGE_RESPONSE', '2': 1010},
    const {'1': 'C_MESSAGE_FEEDBACK', '2': 1011},
  ],
};

/// Descriptor for `CMD`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List cMDDescriptor = $convert.base64Decode('CgNDTUQSDQoJQ19VTktOT1dOEAASCwoHQ19FUlJPUhABEgoKBkNfUElORxACEgoKBkNfUE9ORxADEhkKFENfVVNFUl9MT0dJTl9SRVFVRVNUEOgHEhoKFUNfVVNFUl9MT0dJTl9SRVNQT05TRRDpBxIXChJDX0dFVF9VU0VSX1JFUVVFU1QQ6gcSGAoTQ19HRVRfVVNFUl9SRVNQT05TRRDrBxIaChVDX0dFVF9GUklFTkRTX1JFUVVFU1QQ7AcSGwoWQ19HRVRfRlJJRU5EU19SRVNQT05TRRDtBxIVChBDX1VQREFURV9GUklFTkRTEO4HEiAKG0NfR0VUX0NPTlZFUlNBVElPTlNfUkVRVUVTVBDvBxIhChxDX0dFVF9DT05WRVJTQVRJT05TX1JFU1BPTlNFEPAHEhYKEUNfTUVTU0FHRV9SRVFVRVNUEPEHEhcKEkNfTUVTU0FHRV9SRVNQT05TRRDyBxIXChJDX01FU1NBR0VfRkVFREJBQ0sQ8wc=');
