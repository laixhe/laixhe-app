import 'package:flutter/material.dart';

// 我的
class MyTag extends StatelessWidget {
  const MyTag({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        elevation: 0,      // 去掉下边框的阴影
        centerTitle: true, // 标题居中
        title: Text('我的'),
      ),
      body: Center(
        child: Text(
          '我的',
          style: TextStyle(
              fontSize: 30
          ),
        ),
      ),
    );
  }

}
