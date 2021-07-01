import 'package:flutter/material.dart';

import 'package:im_flutter/net/ws/websocket_manager.dart';

import 'package:im_flutter/protoim/app_os.pb.dart';
import 'package:im_flutter/protoim/cmd.pb.dart';
import 'package:im_flutter/protoim/message_base.pb.dart';
import 'package:im_flutter/protoim/user_login_request.pb.dart';

// 好友
class FriendTag extends StatelessWidget {
  const FriendTag({Key? key}) : super(key: key);

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
