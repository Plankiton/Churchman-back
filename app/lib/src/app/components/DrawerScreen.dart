import 'dart:ui';
import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/models/user.dart';
import 'package:projeto_igreja/src/app/provider/users_provider.dart';
import 'package:projeto_igreja/src/app/views/home/home_view.dart';
import 'package:provider/provider.dart';

import '../constants.dart';
import '../size_config.dart';
import 'theme_data.dart';

class DrawerScreen extends StatefulWidget {
  @override
  _DrawerScreenState createState() => _DrawerScreenState();
}

class _DrawerScreenState extends State<DrawerScreen> {
  @override
  Widget build(BuildContext context) {
    final UsersProvider usersProvider = Provider.of(context);
    User user = usersProvider.byIndex(0);

    return Container(
      decoration: BoxDecoration(gradient: kPrimaryGradientColor),
      padding: EdgeInsets.only(top: 65, bottom: 150, left: 15),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          
          //Usuário
          Container(
            child: Row(
              children: [
                imgProfileCircle(user: user),
                SizedBox(width: getProportionateScreenWidth(10)),
                Text(
                  user.name,
                  style: TextStyle(
                    color: Colors.white,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
          ),

          //DrawerMenu
          Column(
            children: drawerItems
                .map((e) => Material(
                      color: Colors.transparent,
                      child: InkWell(
                        child: Padding(
                          padding: const EdgeInsets.only(
                            top: 10.0,
                            bottom: 10.0,
                            left: 10.0,
                            right: 174.0
                          ),
                          child: e['id'] == 1 ||
                                  e['id'] == 2 ||
                                  e['id'] == 4 ||
                                  e['id'] == 6
                              ? Row(
                                  children: [
                                    Icon(
                                      e['icon'],
                                      color: Colors.white70,
                                      size: 28,
                                    ),
                                    SizedBox(
                                        width: getProportionateScreenWidth(15)),
                                    Text(
                                      e['title'],
                                      style: TextStyle(
                                        color: Colors.white70,
                                        fontSize: 20,
                                      ),
                                    ),
                                  ],
                                )
                              : ExpansionTile(
                                  expandedCrossAxisAlignment: CrossAxisAlignment.start,
                                  expandedAlignment: Alignment.bottomLeft,
                                  tilePadding: EdgeInsets.all(0.0),
                                  title: Text(
                                    e['title'],
                                    style: TextStyle(
                                      color: Colors.white70,
                                      fontSize: 20,
                                    ),
                                  ),
                                  leading: Icon(
                                    e['icon'],
                                    color: Colors.white70,
                                    size: 28,
                                  ),
                                  trailing: Icon(Icons.keyboard_arrow_down_rounded, color: Colors.white70),
                                  children: e['id'] == 3
                                      ? [
                                          Material(
                                            color: Colors.transparent,
                                            child: InkWell(
                                              child: Text(
                                                'Células',
                                                style: TextStyle(
                                                  fontSize: 18,
                                                  color: Colors.white70,
                                                ),
                                              ),
                                              onTap: () {
                                                print('Células');
                                              },
                                            )
                                          ),
                                          
                                          Material(
                                            color: Colors.transparent,
                                            child: InkWell(
                                              child: Text(
                                                'Eventos',
                                                style: TextStyle(
                                                  fontSize: 18,
                                                  color: Colors.white70,
                                                ),
                                              ),
                                              onTap: () {
                                                print('Eventos');
                                              },
                                            )
                                          ),
                                        ]
                                      : [
                                          Material(
                                            color: Colors.transparent,
                                            child: InkWell(
                                              child: Text(
                                                'Igreja',
                                                style: TextStyle(
                                                  fontSize: 18,
                                                  color: Colors.white70,
                                                ),
                                              ),
                                              onTap: () {
                                                print('Igreja');
                                              },
                                            )
                                          ),

                                          Material(
                                            color: Colors.transparent,
                                            child: InkWell(
                                              child: Text(
                                                'Eventos',
                                                style: TextStyle(
                                                  fontSize: 18,
                                                  color: Colors.white70,
                                                ),
                                              ),
                                              onTap: () {
                                                print('Eventos');
                                              },
                                            )
                                          ),
                                          
                                          Material(
                                            color: Colors.transparent,
                                            child: InkWell(
                                              child: Text(
                                                'Notícias',
                                                style: TextStyle(
                                                  fontSize: 18,
                                                  color: Colors.white70,
                                                ),
                                              ),
                                              onTap: () {
                                                print('Notícias');
                                              },
                                            )
                                          ),

                                          Material(
                                            color: Colors.transparent,
                                            child: InkWell(
                                              child: Text(
                                                'Títulos',
                                                style: TextStyle(
                                                  fontSize: 18,
                                                  color: Colors.white70,
                                                ),
                                              ),
                                              onTap: () {
                                                print('Títulos');
                                              },
                                            )
                                          ),

                                        ]),
                        ),

                        //Funções de Cada Item
                        onTap: () {
                          if (e['id'] == 1) {
                            Navigator.pushReplacementNamed(context, HomeView.routeName);
                          }
                          else if (e['id'] == 2) {

                          }
                          else if (e['id'] == 4) {

                          }
                          else if (e['id'] == 6) {

                          }
                        },
                      ),
                    ))
                .toList(),
          ),

          //Logout

        ],
      ),
    );
  }
}
