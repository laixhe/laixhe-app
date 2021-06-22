import 'package:flutter/material.dart';

// 好友
class FriendTag extends StatelessWidget {
  const FriendTag({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        elevation: 0,      // 去掉下边框的阴影
        centerTitle: true, // 标题居中
        title: Text('好友'),
      ),
      body: Center(
        child: Text(
          '好友',
          style: TextStyle(
              fontSize: 30
          ),
        ),
      ),
    );
  }

}
