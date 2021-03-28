import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/provider/events_provider.dart';
import 'package:projeto_igreja/src/app/provider/users_provider.dart';
import 'package:projeto_igreja/src/app/routes/app_routes.dart';
import 'package:projeto_igreja/src/app/views/sign_In/sing_in_module.dart';
import 'package:provider/provider.dart';

import 'components/theme_data.dart';
import 'provider/cells_provider.dart';

class AppWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    //Aqui podem ser adicionados vários Providers;
    //têm-se o para controle de usuários e o para controle de eventos;
    return MultiProvider(
      providers: [
        //Provider de Usuários,
        ChangeNotifierProvider(
          create: (BuildContext context) => UsersProvider(),
        ),
        ChangeNotifierProvider(
          create: (BuildContext context) => EventsProvider(),
        ),
        ChangeNotifierProvider(
          create: (BuildContext context) => CellsProvider(),
        ),
      ],

      //MaterialApp
      child: MaterialApp(
        title: 'Projeto Igreja',
        theme: theme(),
        routes: routes,
        home: SignInModule(),
      ),
    );
  }
}
