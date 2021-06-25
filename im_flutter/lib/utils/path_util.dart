import 'dart:io';

import 'package:path_provider/path_provider.dart';

/// 文件路径工具类
class PathUtil {

  /// 获取缓存目录路径
  static Future<String> getCacheDirPath() async {
    Directory directory = await getTemporaryDirectory();
    return directory.path;
  }

  /// 获取文件缓存目录路径
  static Future<String> getFilesDirPath() async {
    Directory directory = await getApplicationSupportDirectory();
    return directory.path;
  }

  /// 获取文档存储目录路径
  static Future<String> getDocumentsDirPath() async {
    Directory directory = await getApplicationDocumentsDirectory();
    return directory.path;
  }

  static Future<String?> getExternalFileDirPath() async {
    Directory? directory = await getExternalStorageDirectory();
    return directory?.path;
  }

  static Future<String?> getExternalCacheDirPath() async {
    List<Directory>? directories = await getExternalCacheDirectories();
    if (directories == null || directories.isEmpty) {
      return null;
    }
    return directories.first.path;
  }

  /// 同步创建文件夹
  static Directory? createDirSync(String path) {
    if (path.isEmpty) {
      return null;
    }

    Directory dir = new Directory(path);
    if (!dir.existsSync()) {
      dir.createSync(recursive: true);
    }
    return dir;
  }

  /// 异步创建文件夹
  static Future<Directory?> createDir(String path) async {
    if (path.isEmpty) {
      return null;
    }

    Directory dir = new Directory(path);
    bool exist = await dir.exists();
    if (!exist) {
      dir = await dir.create(recursive: true);
    }
    return dir;
  }

}