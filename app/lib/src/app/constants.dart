import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

import 'size_config.dart';

const kPrimaryColor = Color(0xFFFF7643);
const kPrimaryLightColor = Color(0xFFFFECDF);
const kPrimaryGradientColor = LinearGradient(
  begin: Alignment.topRight,
  end: Alignment.topLeft,
  colors: [Color(0xFFFFA53E), Color(0xFFFF7643)],
);
const kSecondaryColor = Color(0xFF979797);
const kTextColor = Color(0xFF757575);

final headingStyle = TextStyle(
  color: Colors.black,
  fontSize: getProportionateScreenWidth(28),
  fontWeight: FontWeight.bold,
);

// Form Error
final RegExp emailValidatorRegExp =
    RegExp(r"^[a-zA-Z0-9.]+@[a-zA-Z0-9]+\.[a-zA-Z]+");
const String kEmailNullError = "Por favor, insira seu Email";
const String kInvalidEmailError = "Por favor, insira um Email válido";
const String kPassNullError = "Por favor, insira uma senha válida";
const String kShortPassError = "A senha deve conter mais de 8 caracteres";
const String kMatchPassError = "Senha inválida";
const String kNamelNullError = "Por favor, insira seu nome";
const String kPhoneNumberNullError = "Por favor, insira seu número de telefone";
const String kAddressNullError = "Por favor, insira seu endereço";
const String kCivilNullError = "Por favor, insira seu Estado Civil";
const String kSexNullError = "Por favor, insira seu Sexo";

//Map Drawer
List<Map> drawerItems = [
  {'id': 1, 'icon': Icons.home_rounded, 'title': 'Tela Inicial'},
  {'id': 2, 'icon': FontAwesomeIcons.church, 'title': 'Igreja'},
  {'id': 3, 'icon': FontAwesomeIcons.edit, 'title': 'Inscrições '},
  {'id': 4, 'icon': FontAwesomeIcons.user, 'title': 'Meus Dados'},
  {'id': 5, 'icon': FontAwesomeIcons.plus, 'title': 'Cadastrar'},
  {'id': 6, 'icon': FontAwesomeIcons.scroll, 'title': 'Relatórios'},
];

List<List<Color>> gradientColors = [
  [Colors.blue[800], Colors.blue],
  [Colors.pink, Colors.deepPurple[400]],
  [Color(0xFFFFA53E), Color(0xFFFF7643)],
  [Colors.deepPurpleAccent[400], Colors.blue[800]]
];
