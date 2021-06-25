import 'dart:convert';

import 'package:convert/convert.dart';
import 'package:crypto/crypto.dart';
import 'package:encrypt/encrypt.dart';

/// encrypt util
class EncryptUtil {

  /// md5 加密
  static String encodeMd5(String data) {
    var content = new Utf8Encoder().convert(data);
    var digest = md5.convert(content);
    return hex.encode(digest.bytes);
  }

  /// aes加密
  static String encodeAes(String? data, String keyStr, String vector) {

    final plainText = data ?? '';
    final key = Key.fromUtf8(keyStr);
    final iv = IV.fromUtf8(vector);
    final encrypter = Encrypter(AES(key, mode: AESMode.cbc));
    final encrypted = encrypter.encrypt(plainText, iv: iv);

    return encrypted.base16;
  }

  /// Base64编码
  static String encodeBase64(String data) {
    return base64Encode(utf8.encode(data));
  }

  /// Base64解码
  static String decodeBase64(String data) {
    return String.fromCharCodes(base64Decode(data));
  }

}
