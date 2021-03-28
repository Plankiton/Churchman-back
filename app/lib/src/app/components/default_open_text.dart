import 'package:flutter/material.dart';
import '../constants.dart';

class DefaultOpenText extends StatelessWidget {
  final String title;
  final String subtitle;

  const DefaultOpenText({
    Key key,
    @required this.title,
    @required this.subtitle
  }) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text(
          title,
          style: headingStyle,
        ),
        Text(
          subtitle,
          textAlign: TextAlign.center,
        ),
      ],
    );
  }
}

//Seja Bem Vindo
//'Faça login com seu email e senha \nou continue com as redes sociais'