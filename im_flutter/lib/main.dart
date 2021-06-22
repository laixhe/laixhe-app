import 'package:flutter/material.dart';
import 'package:get/get.dart';

import 'main_navi.dart';

void main() {
  runApp(GetMaterialApp(
    home: MyApp(),
  ));
}

class MyApp extends StatelessWidget {
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
