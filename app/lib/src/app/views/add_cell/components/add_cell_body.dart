import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/components/default_open_text.dart';
import 'package:projeto_igreja/src/app/views/add_cell/components/add_cell_form.dart';

import '../../../size_config.dart';

class AddCellBody extends StatefulWidget {
  @override
  _AddCellulaBodyState createState() => _AddCellulaBodyState();
}

class _AddCellulaBodyState extends State<AddCellBody> {
  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: Padding(
        padding:
            EdgeInsets.symmetric(horizontal: getProportionateScreenWidth(20)),
        child: SingleChildScrollView(
          child: Column(
            children: [
              SizedBox(height: SizeConfig.screenHeight * 0.02),
              DefaultOpenText(
                  title: 'Cadastrar Célula',
                  subtitle: 'Complete com as informações\nreferentes à célula'),
              SizedBox(height: SizeConfig.screenHeight * 0.05),
              AddCellForm(),
            ],
          ),
        ),
      ),
    );
  }
}
