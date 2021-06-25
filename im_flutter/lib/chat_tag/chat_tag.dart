import 'package:flutter/material.dart';

import 'package:im_flutter/net/ws/websocket_manager.dart';

import 'package:im_flutter/protoim/app_os.pb.dart';
import 'package:im_flutter/protoim/cmd.pb.dart';
import 'package:im_flutter/protoim/message_base.pb.dart';
import 'package:im_flutter/protoim/user_login_request.pb.dart';

// 聊天
class ChatTag extends StatelessWidget {
  const ChatTag({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {

    var userLoginRequest = UserLoginRequest.create();
    userLoginRequest.appOsId = AppOs.ANDROID;
    userLoginRequest.account = "laixhe-flutter";
    userLoginRequest.password = "123456";

    var messageBase = MessageBase.create();
    messageBase.cmd = CMD.C_USER_LOGIN_REQUEST;
    messageBase.data = userLoginRequest.writeToBuffer();

    WebSocketManager().send(messageBase.writeToBuffer());

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

