import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/components/default_open_text.dart';

import '../../../size_config.dart';
import 'ganhar_form.dart';

class GanharBody extends StatefulWidget {
  @override
  _GanharBodyState createState() => _GanharBodyState();
}

class _GanharBodyState extends State<GanharBody> {
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
                  title: 'Cadastrar Discípulo',
                  subtitle: 'Complete com os dados\n do Discípulo'),
              SizedBox(height: SizeConfig.screenHeight * 0.05),
              GanharForm(),
              SizedBox(height: getProportionateScreenHeight(10)),
              Text(
                'Continuando com o cadastro você concorda \n com nossos Termos e Condições',
                textAlign: TextAlign.center,
              ),
              SizedBox(height: SizeConfig.screenHeight * 0.02),
            ],
          ),
        ),
      ),
    );
  }
}
