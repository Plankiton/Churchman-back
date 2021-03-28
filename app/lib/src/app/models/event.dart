import 'package:flutter/cupertino.dart';

class Event {
  final String id;
  final String name;
  final String description;
  final String cover;
  final DateTime begin;
  final DateTime end;

  const Event(
      {@required this.name,
      @required this.description,
      @required this.cover,
      @required this.begin,
      @required this.end,
      this.id});
}
