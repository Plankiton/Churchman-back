import 'package:flutter/material.dart';

import 'components/add_cell_body.dart';

class AddCellView extends StatelessWidget {
  static String routeName = '/add_cellula';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Cadastrar Células'),
        centerTitle: true,
      ),
      body: AddCellBody(),
    );
  }
}
