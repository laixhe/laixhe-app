import 'package:flutter/material.dart';

import 'main_navi.dart';

class MainApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        title: 'Laixhe IM',
        theme: ThemeData(
            primaryColor: Color(0xFFEBEBEB)
        ),
        home: MainNavi()
    );
  }
}