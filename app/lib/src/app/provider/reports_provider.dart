import 'dart:math';

import 'package:flutter/cupertino.dart';
import 'package:projeto_igreja/src/app/data/dummy_report.dart';
import 'package:projeto_igreja/src/app/models/report.dart';

class ReportsProvider with ChangeNotifier {
  final Map<String, Report> _items = {...DUMMY_REPORT};

  List<Report> get all {
    return [..._items.values];
  }

  int get count {
    return _items.length;
  }

  Report byIndex(int i) {
    return _items.values.elementAt(i);
  }

  //Adicionar ou Alterar Cell
  void put(Report report) {
    //Se nulo, retorna um parâmeto vasio
    if (report == null) {
      return;
    }

    //Se o Id mandado já constar => Somente Update
    //Alterar
    if (report.id != null &&
        report.id.trim().isNotEmpty &&
        _items.containsKey(report.id)) {
      _items.update(report.id, (_) => report);
    } else {
      //adicionar
      final id = Random().nextDouble().toString();
      _items.putIfAbsent(
          id,
          () => Report(
                id: id,
                date: report.date,
                food: report.food,
                participants: report.participants,
                place: report.place,
                theme: report.theme,
              ));
    }

    notifyListeners();
  }

  void remove(Report cell) {
    if (cell != null && cell.id != null) {
      _items.remove(cell.id);
      notifyListeners();
    }
  }
}
