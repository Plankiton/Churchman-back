import 'package:bloc_pattern/bloc_pattern.dart';
import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/app_module.dart';
import 'package:projeto_igreja/src/app/views/sign_In/sign_in_repository.dart';
import 'package:projeto_igreja/src/app/views/sign_In/sing_in_bloc.dart';
import 'package:projeto_igreja/src/app/views/sign_In/sing_in_view.dart';
import 'package:projeto_igreja/src/shared/custom_dio/custom_dio.dart';

class SignInModule extends ModuleWidget {
  @override
  // TODO: implement blocs
  List<Bloc> get blocs => [
        Bloc((i) =>
            SignInBloc(SignInModule.to.getDependency<SignInRepository>())),
      ];

  @override
  // TODO: implement dependencies
  List<Dependency> get dependencies => [
        Dependency(
            (i) => SignInRepository(AppModule.to.getDependency<CustomDio>()))
      ];

  @override
  // TODO: implement view
  Widget get view => SignInView();

  static Inject get to => Inject<SignInModule>.of();
}
