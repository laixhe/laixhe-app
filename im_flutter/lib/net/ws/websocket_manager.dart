import 'dart:async';

import 'package:im_flutter/protoim/user_login_response.pb.dart';
import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;

import 'package:im_flutter/protoim/cmd.pb.dart';
import 'package:im_flutter/protoim/message_base.pb.dart';

enum StatusEnum{
  connect,     // 已连接
  connecting,  // 连接中
  closing,     // 关闭中
  close        // 已关闭
}

class WebSocketManager {
  // 单例模式
  static final WebSocketManager _instance = WebSocketManager._internal();
  factory WebSocketManager() => _instance;
  WebSocketManager._internal();

  String _url="ws://192.168.10.240:5050/v1/ws";

  WebSocketChannel? channel;
  StatusEnum isConnect = StatusEnum.close ;  // 默认为未连接

  // 连接
  Future connect() async{

    if(isConnect == StatusEnum.close){
      isConnect = StatusEnum.connecting;
      printStatus();

      channel= IOWebSocketChannel.connect(Uri.parse(_url));
      channel?.stream.listen( (data) => _onMessage(data),
          onError: _onError,
          onDone: _onDone
      );

      isConnect = StatusEnum.connect;
      printStatus();
      return true;
    }

  }

  /// WebSocket连接错误回调
  _onError(e) {
    print('ws - error:{$e}');
  }

  _onDone() {
    print('ws - done');
  }

  _onMessage(data) {

    var messageBase = MessageBase.fromBuffer(data);
    print('ws - messageBase.cmd=${messageBase.cmd}');
    switch(messageBase.cmd) {
      case CMD.C_USER_LOGIN_RESPONSE:
        var userLoginResponse = UserLoginResponse.fromBuffer(messageBase.data);
        print('ws - message: $data');
        print('ws - userLoginResponse userid=${userLoginResponse.userId} nickName=${userLoginResponse.nickName}');
        break;
    }
  }

  // 断开
  Future disconnect() async{
    if(isConnect == StatusEnum.connect){
      isConnect = StatusEnum.closing;
      printStatus();

      await channel?.sink.close(status.goingAway);

      isConnect = StatusEnum.close;
      printStatus();
    }

  }

  // 发送
  bool send(text){
    if(isConnect == StatusEnum.connect) {
      channel?.sink.add(text);
      return true;
    }
    return false;
  }

  void printStatus(){
    if(isConnect == StatusEnum.connect){
      print("websocket 已连接");
    }else if(isConnect == StatusEnum.connecting){
      print("websocket 连接中");
    }else if(isConnect == StatusEnum.closing){
      print("websocket 关闭中");
    }else if(isConnect == StatusEnum.close){
      print("websocket 已关闭");
    }
  }

}
