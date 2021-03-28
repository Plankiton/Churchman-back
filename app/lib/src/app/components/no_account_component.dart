import 'package:flutter/material.dart';
import '../constants.dart';
import '../size_config.dart';
import 'package:projeto_igreja/src/app/views/sign_up/sign_up_view.dart';


class NoAccountText extends StatelessWidget {
  const NoAccountText({
    Key key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text(
          'Não possui uma conta? ',
          style: TextStyle(fontSize: getProportionateScreenWidth(16)),

        ),
        GestureDetector(
          onTap: () => Navigator.pushNamed(context, SignUpView.routeName),
          child: Text(
            'Cadastrar',
            style: TextStyle(
              fontSize: getProportionateScreenWidth(16),
              color: kPrimaryColor,
            ),
            
          ),
        )
      ],
    );
  }
}