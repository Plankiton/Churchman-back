import 'package:flutter/material.dart';
import '../../../constants.dart';
import '../../../size_config.dart';
import 'package:projeto_igreja/src/app/models/event.dart';
import 'package:projeto_igreja/src/app/views/home/home_view.dart';

import 'icon_card.dart';

class ImageAndIcons extends StatelessWidget {
  const ImageAndIcons({
    Key key,
    @required this.size,
    @required String dateBegin,
    @required String timeBegin,
    @required String dateEnd,
    @required String timeEnd,
    @required this.event,
  })  : _dateBegin = dateBegin,
        _timeBegin = timeBegin,
        _dateEnd = dateEnd,
        _timeEnd = timeEnd,
        super(key: key);

  final Size size;
  final String _dateBegin;
  final String _timeBegin;
  final String _dateEnd;
  final String _timeEnd;
  final Event event;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.only(bottom: 5),
      child: SizedBox(
        height: size.height * 0.8,
        child: Row(
          children: [
            Expanded(
              child: Padding(
                padding: EdgeInsets.symmetric(
                    vertical: getProportionateScreenWidth(50)),
                child: Column(
                  children: [
                    Align(
                      alignment: Alignment.topLeft,
                      child: IconButton(
                        icon: Icon(Icons.arrow_back_ios_outlined),
                        onPressed: () {
                          Navigator.pushReplacementNamed(
                              context, HomeView.routeName);
                        },
                      ),
                    ),
                    Spacer(),
                    Text(
                      'Início',
                      style: TextStyle(
                        color: Colors.black,
                        fontSize: 20,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    IconCard(
                      icon: 'assets/icons/calendar.svg',
                    ),
                    Text(
                      _dateBegin,
                      style: TextStyle(
                        color: Colors.black,
                        fontSize: 10,
                      ),
                    ),
                    IconCard(
                      icon: 'assets/icons/clock.svg',
                    ),
                    Text(
                      _timeBegin,
                      style: TextStyle(
                        color: Colors.black,
                        fontSize: 10,
                      ),
                    ),
                    Container(
                      margin: EdgeInsets.only(top: 20),
                      width: 50,
                      height: 2,
                      color: kPrimaryColor.withOpacity(0.8),
                    ),
                    Text(
                      'Fim',
                      style: TextStyle(
                        color: Colors.black,
                        fontSize: 20,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    IconCard(
                      icon: 'assets/icons/calendar.svg',
                    ),
                    Text(
                      _dateEnd,
                      style: TextStyle(
                        color: Colors.black,
                        fontSize: 10,
                      ),
                    ),
                    IconCard(
                      icon: 'assets/icons/clock.svg',
                    ),
                    Text(
                      _timeEnd,
                      style: TextStyle(
                        color: Colors.black,
                        fontSize: 10,
                      ),
                    ),
                  ],
                ),
              ),
            ),
            Container(
              margin: EdgeInsets.only(top: 70),
              width: size.width * 0.8,
              height: size.height * 0.60,
              decoration: BoxDecoration(
                  borderRadius: BorderRadius.only(
                      bottomLeft: Radius.circular(63),
                      topLeft: Radius.circular(63)),
                  boxShadow: [
                    BoxShadow(
                        offset: Offset(0, 10),
                        blurRadius: 60,
                        color: kPrimaryColor.withOpacity(0.29))
                  ],
                  image: DecorationImage(
                    fit: BoxFit.cover,
                    image: AssetImage(event.cover),
                  )),
            ),
          ],
        ),
      ),
    );
  }
}
