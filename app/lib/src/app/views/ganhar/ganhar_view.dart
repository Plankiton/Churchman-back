import 'package:flutter/material.dart';
import 'components/ganhar_body.dart';

class GanharView extends StatelessWidget {
  static String routeName = '/ganhar_view';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Ganhar'),
        centerTitle: true,
      ),
      body: GanharBody(),
    );
  }
}
