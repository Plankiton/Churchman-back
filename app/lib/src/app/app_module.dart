import 'package:bloc_pattern/bloc_pattern.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/app_bloc.dart';
import 'package:projeto_igreja/src/app/app_widget.dart';
import 'package:projeto_igreja/src/shared/custom_dio/custom_dio.dart';

class AppModule extends ModuleWidget {
  @override
  // TODO: implement blocs
  List<Bloc> get blocs => [
    Bloc((i) => AppBloc()),
  ];

  @override
  // TODO: implement dependencies
  List<Dependency> get dependencies => [
    Dependency((i) => Dio()),
    Dependency((i) => CustomDio(i.getDependency<Dio>()))
  ];

  @override
  // TODO: implement view
  Widget get view => AppWidget();

  // TODO: implement inject
  static Inject get to => Inject<AppModule>.of();
  
}

  