import 'package:flutter/cupertino.dart';

class Cell {
  final String id; //Ok
  final String name; //Ok
  final String address;
  final String number; //Ok
  final String date; //Ok
  final String time;

  const Cell({
    this.id,
    @required this.name,
    @required this.address,
    @required this.number,
    @required this.date,
    @required this.time,
  });
}
