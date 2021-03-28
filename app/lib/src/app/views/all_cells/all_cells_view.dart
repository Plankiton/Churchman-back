import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/components/default_open_text.dart';
import 'components/all_cells_body.dart';

class AllCellsView extends StatelessWidget {
  static String routeName = '/all_cells';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[200],
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(90),
        child: AppBar(
          backgroundColor: Colors.grey[200],
          flexibleSpace: Container(
            padding: EdgeInsets.only(top: 65, left: 30, right: 30),
            height: MediaQuery.of(context).size.height,
            child: DefaultOpenText(
                title: 'Células',
                subtitle: 'Escolha uma das células\npara envio de relatório'),
          ),
        ),
      ),
      body: AllCellsBody(),
    );
  }
}
