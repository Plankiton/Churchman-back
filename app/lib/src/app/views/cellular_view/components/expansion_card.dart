import 'package:expansion_card/expansion_card.dart';
import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/constants.dart';

// ignore: non_constant_identifier_names
Container ExapansionCardModule(
    {String title, String subTitle, List<Widget> expansion}) {
  return Container(
      margin: const EdgeInsets.only(bottom: 32),
      padding: EdgeInsets.symmetric(horizontal: 16, vertical: 8),
      decoration: BoxDecoration(
          gradient: kPrimaryGradientColor,
          boxShadow: [
            BoxShadow(
                color: Color(0xFFFF7643).withOpacity(0.4),
                blurRadius: 8,
                spreadRadius: 2,
                offset: Offset(4, 4)),
          ],
          borderRadius: BorderRadius.all(Radius.circular(24))),
      child: ExpansionCard(
        margin: EdgeInsets.only(top: 5),
        title: Container(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Row(
                    children: [
                      Icon(
                        Icons.label,
                        color: Colors.white,
                        size: 24,
                      ),
                      SizedBox(
                        width: 8,
                      ),
                      Text('Visão Celular',
                          style: TextStyle(
                              color: Colors.white, fontFamily: 'avenir'))
                    ],
                  ),
                ],
              ),
              SizedBox(
                height: 5,
              ),
              Text(subTitle,
                  style: TextStyle(color: Colors.white, fontFamily: 'avenir')),
              Text(
                title,
                style: TextStyle(
                  color: Colors.white,
                  fontFamily: 'avenir',
                  fontWeight: FontWeight.w700,
                  fontSize: 24,
                ),
              )
            ],
          ),
        ),
        children: expansion,
      ));
}
