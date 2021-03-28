import 'package:flutter/cupertino.dart';

class Report {
  final String id;
  final DateTime date;
  final String place;
  final String theme;
  final String food;
  final List<String> participants;

  const Report(
      {this.id,
      @required this.date,
      @required this.place,
      @required this.theme,
      @required this.food,
      @required this.participants});
}
