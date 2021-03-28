import 'package:flutter/material.dart';
import 'components/sign_in_body.dart';

class SignInView extends StatelessWidget {
  static String routeName = '/sign_in';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      extendBody: true,
      extendBodyBehindAppBar: true,
      appBar: AppBar(
        title: Text('Login'),
        centerTitle: true,
      ),
      body: SignInBody(),
    );
  }
}
