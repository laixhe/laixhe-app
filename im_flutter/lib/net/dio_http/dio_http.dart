import 'package:dio/dio.dart';

import 'auth_interceptor.dart';

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
    _dio.options.baseUrl = 'http://192.168.10.240:5050';
    _dio.options.connectTimeout = 3000;
    _dio.options.sendTimeout = 3000;
    _dio.options.receiveTimeout = 5000;

    _dio.interceptors.add(AuthInterceptor());
    // 是否开启请求日志
    _dio.interceptors.add(LogInterceptor(request:true, requestHeader:true, responseHeader: true, responseBody: true));
  }

  // 请求
  Future request(String url,{String method = "GET",Map<String,dynamic>? params}) async {
    Options options = Options(method: method);
    try{
      final result = await _dio.request(url, queryParameters: params, options: options);
      return result;
    } on DioError catch(error){
      throw error;
    }
  }

}