import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/views/add_cell/add_cell_view.dart';
import 'package:projeto_igreja/src/app/views/all_cells/all_cells_view.dart';
import 'package:projeto_igreja/src/app/views/ganhar/ganhar_view.dart';
import 'expansion_card.dart';

class CellularViewBody extends StatefulWidget {
  @override
  _CellularViewBodyState createState() => _CellularViewBodyState();
}

class _CellularViewBodyState extends State<CellularViewBody> {
  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(horizontal: 32, vertical: 30),
      child: Expanded(
        child: ListView(
          children: [
            //Ganhar
            ExapansionCardModule(
                title: 'Ganhar',
                subTitle: 'Adicionar Discípulos',
                expansion: <Widget>[
                  Material(
                    color: Colors.transparent,
                    child: InkWell(
                      child: Padding(
                        padding: const EdgeInsets.only(left: 16, top: 10),
                        child: Align(
                          alignment: Alignment.centerLeft,
                          child: Row(
                            children: [
                              Icon(
                                Icons.arrow_right_rounded,
                                color: Colors.white,
                              ),
                              Text(
                                'Cadastrar Novos Membros',
                                style: TextStyle(
                                  fontSize: 20,
                                  color: Colors.white,
                                  fontFamily: 'avenir',
                                ),
                              ),
                            ],
                          ),
                        ),
                      ),
                      onTap: () {
                        Navigator.pushNamed(context, GanharView.routeName);
                      },
                    ),
                  ),
                ]),
            //Enviar
            ExapansionCardModule(
                title: 'Enviar',
                subTitle: 'Gerenciar Células',
                expansion: <Widget>[
                  Material(
                    color: Colors.transparent,
                    child: InkWell(
                      child: Padding(
                        padding: const EdgeInsets.only(left: 16, top: 10),
                        child: Align(
                          alignment: Alignment.centerLeft,
                          child: Row(
                            children: [
                              Icon(
                                Icons.arrow_right_rounded,
                                color: Colors.white,
                              ),
                              Text(
                                'Cadastrar Células',
                                style: TextStyle(
                                  fontSize: 20,
                                  color: Colors.white,
                                  fontFamily: 'avenir',
                                ),
                              ),
                            ],
                          ),
                        ),
                      ),
                      onTap: () {
                        Navigator.pushNamed(context, AddCellView.routeName);
                      },
                    ),
                  ),
                  Material(
                    color: Colors.transparent,
                    child: InkWell(
                      child: Padding(
                        padding: const EdgeInsets.only(left: 16, top: 10),
                        child: Align(
                          alignment: Alignment.centerLeft,
                          child: Row(
                            children: [
                              Icon(
                                Icons.arrow_right_rounded,
                                color: Colors.white,
                              ),
                              Text(
                                'Enviar Relatório',
                                style: TextStyle(
                                  fontSize: 20,
                                  color: Colors.white,
                                  fontFamily: 'avenir',
                                ),
                              ),
                            ],
                          ),
                        ),
                      ),
                      onTap: () {
                        Navigator.pushNamed(context, AllCellsView.routeName);
                      },
                    ),
                  ),
                ])
          ],
        ),
      ),
    );
  }
}
