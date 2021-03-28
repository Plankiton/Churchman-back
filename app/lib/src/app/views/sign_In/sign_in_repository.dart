import 'package:dio/dio.dart';
import 'package:projeto_igreja/src/app/models/post_model.dart';
import 'package:projeto_igreja/src/shared/custom_dio/custom_dio.dart';

class SignInRepository {
  final CustomDio dio;

  SignInRepository(this.dio);

  Future<List<PostModel>> getUser() async {
      final prefs = await SharedPreferences.getInstance();

      final counter = prefs.getString('user_token') ?? '';
      try{
          var response = await dio.client.post('/verify');
          return (response.data as List).map((item) => PostModel.fromJson(item)).toList();
      } on DioError catch (e) {
          throw(e.message);
      }
  }
}
