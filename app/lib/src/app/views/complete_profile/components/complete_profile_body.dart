import 'package:flutter/material.dart';
import '../../../size_config.dart';

import 'package:projeto_igreja/src/app/components/default_open_text.dart';


import 'complete_profile_form.dart';

class CompleteProfileBody extends StatelessWidget {
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
                  title: 'Complete seu Perfil',
                  subtitle:
                      'Complete seus dados ou continue \n com as redes socias'),
              SizedBox(height: SizeConfig.screenHeight * 0.05),
              CompleteProfileForm(),
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

