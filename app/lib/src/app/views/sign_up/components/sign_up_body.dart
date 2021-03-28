import 'package:flutter/material.dart';
import '../../../size_config.dart';
import 'package:projeto_igreja/src/app/components/default_open_text.dart';
import 'package:projeto_igreja/src/app/components/social_card_component.dart';

import 'sign_up_form.dart';

class SignUpBody extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: SingleChildScrollView(
        child: Column(
          children: [
            SizedBox(height: SizeConfig.screenHeight * 0.02),
            DefaultOpenText(
              title: 'Registrar Conta',
              subtitle:
                  'Complete suas informações ou continue \ncom suas redes sociais',
            ),
            SizedBox(height: SizeConfig.screenHeight * 0.06), //8% da altura total,
            SignUpForm(),
            SizedBox(height: SizeConfig.screenHeight * 0.06),
            SizedBox(height: getProportionateScreenHeight(20)),
            Text(
              'Continuando com o cadastro você concorda \n com nossos Termos e Condições',
              textAlign: TextAlign.center,
            ),
            SizedBox(height: SizeConfig.screenHeight * 0.02),
          ],
        ),
      ),
    );
  }
}

