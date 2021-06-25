import 'package:dio/dio.dart';

class AuthInterceptor extends Interceptor {
  @override
  Future onRequest(RequestOptions options, RequestInterceptorHandler handler) async {

    options.headers[Headers.acceptHeader] = Headers.jsonContentType;
    options.headers[Headers.contentTypeHeader] = Headers.jsonContentType;
    options.headers['Authorization'] = '-Authorization-';

    super.onRequest(options, handler);
  }

  @override
  Future onResponse(Response response, ResponseInterceptorHandler handler) async {
    super.onResponse(response, handler);
  }
}