import 'dart:io';

import 'package:dio/dio.dart';

class DioHttp{
  // 单例模式
  static final DioHttp _instance = DioHttp._internal();
  factory DioHttp() => _instance;
  DioHttp._internal(){
    init();
  }

  Dio _dio = Dio();

  // 初始化请求配置
  init(){
    _dio.options.baseUrl = 'https://www.xxxx/api';
    _dio.options.connectTimeout = 5000;
    _dio.options.receiveTimeout = 5000;
    _dio.options.contentType = ContentType.parse("application/json;charset=UTF-8").toString();
  }

}