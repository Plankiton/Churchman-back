import 'package:flutter/material.dart';
import '../constants.dart';
import 'package:projeto_igreja/src/app/models/user.dart';

ThemeData theme() {
  return ThemeData(
      scaffoldBackgroundColor: Colors.white,
      fontFamily: 'Muli',
      appBarTheme: appBarTheme(),
      textTheme: textTheme(),
      inputDecorationTheme: inputDecorationTheme(),
      primaryColor: Colors.blue,
      visualDensity: VisualDensity.adaptivePlatformDensity);
}

InputDecorationTheme inputDecorationTheme() {
  OutlineInputBorder outlineInputBorder = OutlineInputBorder(
    borderRadius: BorderRadius.circular(28),
    borderSide: BorderSide(color: kTextColor),
    gapPadding: 10,
  );
  return InputDecorationTheme(
    contentPadding: EdgeInsets.symmetric(horizontal: 42, vertical: 20),
    enabledBorder: outlineInputBorder,
    focusedBorder: outlineInputBorder,
    border: outlineInputBorder,
  );
}

InputDecoration inputDecorationRow({String labelText}) {
  OutlineInputBorder outlineInputBorder = OutlineInputBorder(
    borderRadius: BorderRadius.circular(28),
    borderSide: BorderSide(color: kTextColor),
    gapPadding: 10,
  );
  return InputDecoration(
    floatingLabelBehavior: FloatingLabelBehavior.always,
    labelText: labelText,
    contentPadding: EdgeInsets.symmetric(horizontal: 21, vertical: 10),
    enabledBorder: outlineInputBorder,
    focusedBorder: outlineInputBorder,
    border: outlineInputBorder,
  );
}

TextTheme textTheme() {
  return TextTheme(
    bodyText1: TextStyle(color: kTextColor),
    bodyText2: TextStyle(color: kTextColor),
  );
}

AppBarTheme appBarTheme() {
  return AppBarTheme(
      color: Colors.white,
      elevation: 0,
      brightness: Brightness.light,
      iconTheme: IconThemeData(color: Colors.black),
      textTheme: TextTheme(
          headline6: TextStyle(color: Color(0XFF8B8B8B), fontSize: 18)));
}

CircleAvatar imgProfileCircle({User user}) {
  final avatar = user.data == null || user.data.isEmpty
      ? CircleAvatar(
          radius: 25.0,
          backgroundColor: Colors.transparent,
          child: Icon(Icons.person),
        )
      : CircleAvatar(
          backgroundImage: NetworkImage(user.data),
          radius: 25.0,
          backgroundColor: Colors.transparent,
        );

  return avatar;
}
