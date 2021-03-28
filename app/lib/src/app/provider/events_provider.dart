import 'dart:math';
import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/data/dummy_event.dart';
import 'package:projeto_igreja/src/app/models/event.dart';

class EventsProvider with ChangeNotifier {
  final Map<String, Event> _items = {...DUMMY_EVENT};

  List<Event> get all {
    return [..._items.values];
  }

  int get count {
    return _items.length;
  }

  Event byIndex(int i) {
    return _items.values.elementAt(i);
  }

  //Adicionar ou Alterar Eventos
  void put(Event event) {
    //Se nulo, retorna um parâmeto vasio
    if (event == null) {
      return;
    }

    //Se o Id mandado já constar => Somente Update
    //Alterar
    if (event.id != null &&
        event.id.trim().isNotEmpty &&
        _items.containsKey(event.id)) {
      _items.update(event.id, (_) => event);
    } else {
      //adicionar
      final id = Random().nextDouble().toString();
      _items.putIfAbsent(
          id,
          () => Event(
                id: id,
                name: event.name,
                description: event.description,
                cover: event.cover,
                begin: event.begin,
                end: event.end,
              ));
    }

    notifyListeners();
  }

  void remove(Event event) {
    if (event != null && event.id != null) {
      _items.remove(event.id);
      notifyListeners();
    }
  }
}
