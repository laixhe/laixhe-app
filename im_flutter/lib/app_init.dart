import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter/foundation.dart';

import 'package:dokit/dokit.dart';
import 'package:get/get.dart';

import 'package:im_flutter/utils/date_util.dart';
import 'package:im_flutter/utils/path_util.dart';
import 'package:im_flutter/utils/device_util.dart';

import 'main_app.dart';
import 'net/ws/websocket_manager.dart';

// 应用初始化
class AppInit {

  // 获取 app
  static Future<Widget> getApp() async {
    // 程序初始化操作
    await initApp();

    return GetMaterialApp(
        home:MainApp()
    );

  }

  // 启动 app
  static Future<void> startApp() async {
    runApp(await getApp());
  }

  // 程序初始化操作
  static Future<void> initApp() async {
    WebSocketManager().connect();
  }

  // 运行 app
  static Future<void> run() async {

    if (!DeviceUtil.isMobile) {
      DoKit.runApp(
        appCreator: () async => DoKitApp(await getApp()),
        useInRelease: false,
        exceptionCallback: (obj, trace) {
          reportErrorAndLog(makeDetails(obj, trace));
        });

    } else {
      // 捕获异常
      catchException(() => startApp());
    }
  }

  // 构建错误信息
  static FlutterErrorDetails makeDetails(Object obj, StackTrace stack) {
    return FlutterErrorDetails(exception: obj, stack: stack);
  }

  /// 异常捕获处理
  static void catchException<T>(T callback()) {
    // 捕获异常的回调
    FlutterError.onError = (FlutterErrorDetails details) {
      reportErrorAndLog(details);
    };
    runZoned<Future<Null>>(
          () async {
        callback();
      },
      zoneSpecification: ZoneSpecification(
        print: (Zone self, ZoneDelegate parent, Zone zone, String line) {
          collectLog(parent, zone, line); // 收集日志
        },
      ),
      //未捕获的异常的回调
      onError: (Object obj, StackTrace stack) {
        var details = makeDetails(obj, stack);
        reportErrorAndLog(details);
      },
    );
  }

  // 日志拦截, 收集日志
  static void collectLog(ZoneDelegate parent, Zone zone, String line) {
    parent.print(zone, "$line");
  }

  // 上报错误和日志逻辑
  static void reportErrorAndLog(FlutterErrorDetails details) {
    print(details);
    saveErrorToFile(details.toString());

    // kReleaseMode => isReleaseMode
    if (DeviceUtil.isMobile && kReleaseMode) {
      // 可使用第三方如：腾讯Bugly应用更新统计及异常上报插件
    }

  }

  static Future<void> saveErrorToFile(String error) async {

    String? dirPath = '';
    if (DeviceUtil.isIOS) {
      dirPath = await PathUtil.getCacheDirPath();
      if (dirPath.isEmpty) {
        return null;
      }
    } else {
      dirPath = await PathUtil.getExternalCacheDirPath();
      if (dirPath == null || dirPath.isEmpty) {
        return null;
      }
    }

    Directory? crashDir = PathUtil.createDirSync('$dirPath/crash');
    String? crashDirPath = crashDir?.path;
    if (dirPath.isEmpty) {
      return null;
    }

    String fileName = 'crash_' + DateUtil.getNowDateStr()!.replaceAll(' ', '_') + '.txt';
    File file = new File('$crashDirPath/$fileName');
    if (!file.existsSync()) {
      file.createSync();
    }

    file.writeAsString(error);
  }


}