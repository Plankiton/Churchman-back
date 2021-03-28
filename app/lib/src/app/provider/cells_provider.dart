import 'dart:math';
import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/data/dummy_cell.dart';
import 'package:projeto_igreja/src/app/models/cell.dart';

class CellsProvider with ChangeNotifier {
  final Map<String, Cell> _items = {...DUMMY_CELL};

  List<Cell> get all {
    return [..._items.values];
  }

  int get count {
    return _items.length;
  }

  Cell byIndex(int i) {
    return _items.values.elementAt(i);
  }

  //Adicionar ou Alterar Cell
  void put(Cell cell) {
    //Se nulo, retorna um parâmeto vasio
    if (cell == null) {
      return;
    }

    //Se o Id mandado já constar => Somente Update
    //Alterar
    if (cell.id != null &&
        cell.id.trim().isNotEmpty &&
        _items.containsKey(cell.id)) {
      _items.update(cell.id, (_) => cell);
    } else {
      //adicionar
      final id = Random().nextDouble().toString();
      _items.putIfAbsent(
          id,
          () => Cell(
                id: id,
                name: cell.name,
                address: cell.address,
                date: cell.date,
                number: cell.number,
                time: cell.time,
              ));
    }

    notifyListeners();
  }

  void remove(Cell cell) {
    if (cell != null && cell.id != null) {
      _items.remove(cell.id);
      notifyListeners();
    }
  }
}
