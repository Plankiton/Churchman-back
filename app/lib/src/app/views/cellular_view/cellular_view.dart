import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/constants.dart';
import 'package:projeto_igreja/src/app/views/cellular_view/components/cellular_view_body.dart';

class CellularView extends StatelessWidget {
  static String routeName = '/cellular_view';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      extendBody: true,
      extendBodyBehindAppBar: true,
      appBar: AppBar(
        title: Text(
          'Visão Celular',
          style: TextStyle(
              color: Colors.white,
              fontFamily: 'avenir',
              fontWeight: FontWeight.w700,
              fontSize: 24,
              fontStyle: FontStyle.italic),
        ),
        centerTitle: true,
        backgroundColor: kPrimaryColor,
      ),
      body: CellularViewBody(),
    );
  }
}
