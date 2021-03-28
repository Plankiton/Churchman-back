import 'package:dio/dio.dart';

class CustomInterceptors extends InterceptorsWrapper{
  @override
  Future onRequest(RequestOptions options) {
      // TODO: implement onRequest
      return super.onRequest(options);
    }
  
  @override
    Future onResponse(Response response) {
      // TODO: implement onResponse
      //200
      //201
      return super.onResponse(response);
    }
  
  @override
    Future onError(DioError err) {
      // TODO: implement onError
      //Exception
      return super.onError(err);
    }
}