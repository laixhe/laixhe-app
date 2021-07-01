import 'package:flutter/material.dart';

// 聊天
class ChatTag extends StatelessWidget {
  const ChatTag({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {

    return Scaffold(
      appBar: AppBar(
        elevation: 0,      // 去掉下边框的阴影
        centerTitle: true, // 标题居中
        title: Text('聊天'),
        actions: <Widget>[

          IconButton(
              padding: EdgeInsets.only(left: 20),
              iconSize: 28,
              color: Colors.grey,
              icon: Icon(Icons.search),
              onPressed: (){
              }
          ),
          IconButton(
              iconSize: 24,
              color: Colors.grey,
              icon: Icon(Icons.control_point),
              onPressed: (){
              }
          )
        ],
      ),
      body: Center(
        child: Text(
          '聊天',
          style: TextStyle(
              fontSize: 30
          ),
        ),
      ),
    );
  }

}

