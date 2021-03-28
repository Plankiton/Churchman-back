import 'package:flutter/material.dart';
import 'components/event_register_body.dart';

class EventRegisterView extends StatelessWidget {
  static String routeName = '/event_register';

  //final List<String> arguments = ModalRoute.of(context).settings.arguments;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[200],
      body: EventRegisterBody(),
    );
  }
}
