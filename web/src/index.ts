import template from 'lodash/template';
import $ from "jquery";

let CMD = require("./protoim/cmd_pb")
let MessageBase = require("./protoim/message_base_pb")
let GetUserInfoRequest = require("./protoim/get_user_info_pb")
let User = require("./protoim/user_pb")

const outputElement = document.getElementById('app');
if (outputElement) {
  var compiled = template(`
    <h2><%- wsHeading %></h2>
    <div>
      <button type="button" id="ws_connection"><%- wsConnection %></button>
    </div>
    <div>
      <button type="button" id="ws_close"><%- wsClose %></button>
    </div>
    <div>
    <button type="button" id="ws_send">发送</button>
    </div>
  `.trim());

  outputElement.innerHTML = compiled({
    wsHeading: 'WebSocket',
    wsConnection: '连接',
    wsClose: '关闭'
  });

  let i: number = 100;
  let ws: WebSocket;
  // 连接事件
  $('#ws_connection').on('click', function(){
    console.log('连接...');
    ws = new WebSocket("ws://192.168.10.240:5050/v1/ws");
    // 可以传递 String、ArrayBuffer 和 Blob 三种数据类型，默认为 string
    ws.binaryType = 'arraybuffer';
    //连接打开时触发
    ws.onopen = function(evt: Event) {
      console.log("onopen..");
    }
    //接收到消息时触发
    ws.onmessage = function(evt: MessageEvent) {
      console.log("onmessage evt:", evt);
      console.log("onmessage evt.data:", new Uint8Array(evt.data));
      let messageBase = MessageBase.MessageBase.deserializeBinary(evt.data);
      console.log("onmessage messageBase:", messageBase);
      console.log("onmessage messageBase.cmd:", messageBase.getCmd());
      console.log("onmessage messageBase.data:", messageBase.getData());
      let user = User.User.deserializeBinary(new Uint8Array(evt.data));
      console.log("onmessage user:", user);
      console.log("onmessage user.uid:", user.getUid());
      console.log("onmessage user.name:", user.getName());
      console.log("onmessage user.age:", user.getAge());
      console.log("onmessage user.isok:", user.getIsok());
    }
    //连接关闭时触发
    ws.onclose = function(evt) {
      console.log("onclose...");
    }
  });
  // 关闭事件
  $('#ws_close').on('click', ()=>{
    console.log("click close...");
    ws?.close();
  });

  // 发送事件
  $('#ws_send').on('click', ()=>{
    console.log("click send...");

    let getUserInfoRequest = new GetUserInfoRequest.GetUserInfoRequest();
    getUserInfoRequest.setUid(i);
    // 将 protobuf 序列化为字节数组 Uint8Array
    let getUserInfoRequestBuffer = getUserInfoRequest.serializeBinary();

    let messageBase = new MessageBase.MessageBase();
    messageBase.setCmd(CMD.CMD.GET_USER_INFO);
    messageBase.setData(getUserInfoRequestBuffer);
    let messageBaseBuffer = messageBase.serializeBinary();

    ws.send(messageBaseBuffer);

    i++;
    console.log(messageBase);
  });

}
