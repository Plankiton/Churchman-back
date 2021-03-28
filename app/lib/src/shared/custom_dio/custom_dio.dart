import 'package:dio/dio.dart';
import 'package:projeto_igreja/src/shared/custom_dio/interceptors.dart';
import '../constants.dart';

class CustomDio{
  final Dio client;

  CustomDio(this.client) {
    client.interceptors.add(CustomInterceptors());
    client.options.baseUrl = BASE_URL;
    client.options.connectTimeout = 5000;
  }

}
