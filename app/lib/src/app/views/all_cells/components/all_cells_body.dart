import 'dart:math';

import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/constants.dart';
import 'package:projeto_igreja/src/app/provider/cells_provider.dart';
import 'package:projeto_igreja/src/app/views/send_report/send_report_view.dart';
import 'package:provider/provider.dart';

class AllCellsBody extends StatefulWidget {
  @override
  _AllCellsBodyState createState() => _AllCellsBodyState();
}

class _AllCellsBodyState extends State<AllCellsBody> {
  final _random = new Random();
  List<Color> colors = gradientColors[0];
  @override
  Widget build(BuildContext context) {
    final CellsProvider cells = Provider.of(context);
    return Column(
      children: [
        SizedBox(
          height: 10,
        ),
        Expanded(
          child: ListView.builder(
            itemCount: cells.count,
            itemBuilder: (BuildContext context, int index) {
              colors = gradientColors[_random.nextInt(gradientColors.length)];
              return Stack(
                children: [
                  Padding(
                    padding: const EdgeInsets.symmetric(
                      vertical: 10,
                      horizontal: 10,
                    ),
                    child: InkWell(
                      child: Container(
                        width: double.infinity,
                        decoration: BoxDecoration(
                            borderRadius: BorderRadius.circular(24),
                            gradient: LinearGradient(
                              colors: colors,
                              begin: Alignment.topLeft,
                              end: Alignment.topRight,
                            ),
                            boxShadow: [
                              BoxShadow(
                                  color: colors[1].withOpacity(0.4),
                                  blurRadius: 8,
                                  spreadRadius: 2,
                                  offset: Offset(4, 4))
                            ]),
                        child: Padding(
                          padding: const EdgeInsets.symmetric(
                              vertical: 10, horizontal: 20),
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            mainAxisAlignment: MainAxisAlignment.start,
                            children: [
                              Text(
                                cells.all.elementAt(index).name,
                                style: TextStyle(
                                    color: Colors.white,
                                    fontFamily: 'Circular',
                                    fontWeight: FontWeight.w700,
                                    fontSize: 25),
                              ),
                              Text(
                                '${cells.all.elementAt(index).address}, ${cells.all.elementAt(index).number}',
                                style: TextStyle(
                                    color: Colors.white,
                                    fontFamily: 'Circular',
                                    fontWeight: FontWeight.w500,
                                    fontSize: 20),
                              ),
                              Text(
                                '${cells.all.elementAt(index).date} às ${cells.all.elementAt(index).time}h',
                                style: TextStyle(
                                    color: Colors.white,
                                    fontFamily: 'Circular',
                                    fontWeight: FontWeight.w500,
                                    fontSize: 20),
                              ),
                            ],
                          ),
                        ),
                      ),
                      onTap: () {
                        Navigator.pushReplacementNamed(
                            context, SendReportView.routeName,
                            arguments: cells.all.elementAt(index));
                      },
                    ),
                  ),
                ],
              );
            },
          ),
        ),
      ],
    );
  }
}
