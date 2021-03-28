import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import '../../../constants.dart';
import 'package:projeto_igreja/src/app/models/event.dart';

import 'image_icons.dart';
import 'name_description.dart';

class EventRegisterBody extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    Size size = MediaQuery.of(context).size;
    final Event event = ModalRoute.of(context).settings.arguments;

    String _dateBegin = event.begin != null
        ? '${event.begin.day} - ${event.begin.month} - ${event.begin.year}'
        : 'Não Definido';
    String _timeBegin = event.begin != null
        ? '${event.begin.hour} - ${event.begin.minute}'
        : 'Não Definido';
    String _dateEnd = event.end != null
        ? '${event.end.day} - ${event.end.month} - ${event.end.year}'
        : 'Não Definido';
    String _timeEnd = event.end != null
        ? '${event.end.hour} - ${event.end.minute}'
        : 'Não Definido';

    return SingleChildScrollView(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          ImageAndIcons(
              size: size,
              dateBegin: _dateBegin,
              timeBegin: _timeBegin,
              dateEnd: _dateEnd,
              timeEnd: _timeEnd,
              event: event),
          NameAndDescription(event: event),
          Row(
            children: [
              SizedBox(
                width: size.width,
                height: 84,
                child: FlatButton(
                  shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.only(
                          topRight: Radius.circular(20),
                          topLeft: Radius.circular(20))),
                  color: kPrimaryColor,
                  onPressed: () {},
                  child: Text(
                    'Inscrever',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 16,
                    ),
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
