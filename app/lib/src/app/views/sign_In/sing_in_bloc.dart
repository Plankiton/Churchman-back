import 'package:bloc_pattern/bloc_pattern.dart';
import 'package:projeto_igreja/src/app/models/post_model.dart';
import 'package:projeto_igreja/src/app/views/sign_In/sign_in_repository.dart';
import 'package:rxdart/rxdart.dart';

class SignInBloc extends BlocBase {
  final SignInRepository repo;

  SignInBloc(this.repo);

  var listPost = BehaviorSubject<List<PostModel>>();
  Sink<List<PostModel>> get responseIn => listPost.sink;
  Stream<List<PostModel>> get responseOut => listPost.stream;

  void getPosts() async {
    try {
      var res = await repo.getPosts();
      responseIn.add(res);
    } catch (e) {
      listPost.addError(e);
    }
  }

  @override
  void dispose() {
    listPost.close();
    super.dispose();
  }
}
