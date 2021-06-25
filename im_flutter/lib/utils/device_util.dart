import 'package:flutter/foundation.dart';

/// 当前设备
class DeviceUtil {
  static bool get isAndroid => (defaultTargetPlatform == TargetPlatform.android);
  static bool get isIOS => (defaultTargetPlatform == TargetPlatform.iOS);
  static bool get isMobile => isAndroid || isIOS;
}