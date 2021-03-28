import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/provider/events_provider.dart';
import 'package:projeto_igreja/src/app/views/event_register/event_register_view.dart';

class SlideTile extends StatelessWidget {
  final EventsProvider events;
  final int index;
  final bool activePage;

  const SlideTile({Key key, this.events, this.index, this.activePage})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    final double top = this.activePage ? 50 : 150;
    final double blur = this.activePage ? 30 : 0;
    final double offset = this.activePage ? 10 : 0;

    return InkWell(
      child: FractionallySizedBox(
        heightFactor: 1.03,
        child: AnimatedContainer(
          duration: Duration(microseconds: 500),
          margin: EdgeInsets.only(top: top, bottom: 50, right: 15),
          decoration: BoxDecoration(
              image: DecorationImage(
                image: AssetImage(events.all.elementAt(index).cover),
                fit: BoxFit.cover,
              ),
              borderRadius: BorderRadius.circular(20),
              boxShadow: [
                BoxShadow(
                  color: Colors.black87,
                  blurRadius: blur,
                  offset: Offset(offset, offset),
                )
              ]),
        ),
      ),
      onTap: () {
        Navigator.pushReplacementNamed(context, EventRegisterView.routeName,
            arguments: events.all.elementAt(index));
      },
    );
  }
}
