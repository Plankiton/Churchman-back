import 'dart:math';

import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/data/dummy_user.dart';
import 'package:projeto_igreja/src/app/models/user.dart';

class UsersProvider with ChangeNotifier {
  final Map<String, User> _items = {...DUMMY_USERS};

  List<User> get all {
    return [..._items.values];
  }

  int get count {
    return _items.length;
  }

  User byIndex(int i) {
    return _items.values.elementAt(i);
  }

  //Adicionar ou Alterar Usuários
  void put(User user) {
    //Se nulo, retorna um parâmeto vasio
    if (user == null) {
      return;
    }

    //Se o Id mandado já constar => Somente Update
    //Alterar
    if (user.id != null &&
        user.id.trim().isNotEmpty &&
        _items.containsKey(user.id)) {
      _items.update(user.id, (_) => user);
    } else {
      //adicionar
      final id = Random().nextDouble().toString();
      _items.putIfAbsent(
          id,
          () => User(
                id: id,
                name: user.name,
                email: user.email,
                data: user.data,
                born: user.born,
                state: user.state,
                phone: user.phone,
                sex: user.sex,
              ));
    }

    notifyListeners();
  }

  void remove(User user) {
    if (user != null && user.id != null) {
      _items.remove(user.id);
      notifyListeners();
    }
  }
}
