import 'package:flutter/material.dart';
import 'components/complete_profile_body.dart';

class CompleteProfileView extends StatelessWidget {
  static String routeName = '/complete_profile';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Perfil'),
        centerTitle: true,
      ),
      body: CompleteProfileBody(),
    );
  }
}
