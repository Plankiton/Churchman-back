import 'package:flutter/material.dart';

import 'components/church_body.dart';

class ChurchView extends StatelessWidget {
  static String routeName = '/church';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      extendBody: true,
      extendBodyBehindAppBar: true,
      appBar: AppBar(
        title: Text('Igreja'),
        centerTitle: true,
      ),
      body: ChurchBody(),
    );
  }
}
