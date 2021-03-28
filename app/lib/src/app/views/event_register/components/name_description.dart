import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/models/event.dart';

class NameAndDescription extends StatelessWidget {
  const NameAndDescription({
    Key key,
    @required this.event,
  }) : super(key: key);

  final Event event;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: 10),
      child: RichText(
        text: TextSpan(
          children: [
            TextSpan(
                text: '${event.name}\n',
                style: Theme.of(context).textTheme.headline5.copyWith(
                    color: Colors.black, fontWeight: FontWeight.bold)),
            TextSpan(
                text: '${event.description}\n',
                style: TextStyle(
                  color: Colors.grey[800],
                  fontWeight: FontWeight.w300,
                  fontSize: 18,
                )),
          ],
        ),
      ),
    );
  }
}
