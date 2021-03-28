import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/constants.dart';
import 'package:projeto_igreja/src/app/models/cell.dart';
import 'components/send_report_body.dart';

class SendReportView extends StatelessWidget {
  static String routeName = '/send_report';

  @override
  Widget build(BuildContext context) {
    final Cell cell = ModalRoute.of(context).settings.arguments;
    return Scaffold(
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(90),
        child: AppBar(
          backgroundColor: kPrimaryColor,
          flexibleSpace: Container(
            padding: EdgeInsets.only(top: 65, left: 30, right: 30),
            height: MediaQuery.of(context).size.height,
            child: Center(
              child: Padding(
                padding:
                    const EdgeInsets.symmetric(vertical: 10, horizontal: 20),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  mainAxisAlignment: MainAxisAlignment.start,
                  children: [
                    Text(
                      cell.name,
                      style: TextStyle(
                          color: Colors.white,
                          fontFamily: 'Circular',
                          fontWeight: FontWeight.w700,
                          fontSize: 20),
                    ),
                    Text(
                      '${cell.address}, ${cell.number}',
                      style: TextStyle(
                          color: Colors.white,
                          fontFamily: 'Circular',
                          fontWeight: FontWeight.w500,
                          fontSize: 15),
                    ),
                    Text(
                      '${cell.date} às ${cell.time}h',
                      style: TextStyle(
                          color: Colors.white,
                          fontFamily: 'Circular',
                          fontWeight: FontWeight.w500,
                          fontSize: 15),
                    ),
                  ],
                ),
              ),
            ),
          ),
        ),
      ),
      body: SendReportBody(),
    );
  }
}
