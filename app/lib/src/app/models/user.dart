import 'package:flutter/cupertino.dart';

class User {
  final String id; //Ok
  final String name; //Ok
  final String email; //Ok
  final String data; //
  final String sex; //Ok
  final String phone; //Ok
  final String born; //Ok
  final String state; //Ok

  const User({
    this.id,
    @required this.sex,
    @required this.phone,
    @required this.born,
    @required this.state,
    @required this.name,
    @required this.email,
    @required this.data,
  });
}
