import 'package:flutter/material.dart';

import 'components/sign_up_body.dart';

class SignUpView extends StatelessWidget {
  static String routeName = '/sign_up';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      extendBody: true,
      extendBodyBehindAppBar: true,
      appBar: AppBar(
        title: Text('Cadastrar'),
        centerTitle: true,
      ),
      body: SignUpBody(),
    );
  }
}
