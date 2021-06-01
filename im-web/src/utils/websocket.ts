import * as MessageBase from '@/protoim/message_base_pb';
import * as GetUserInfoResponse from '@/protoim/get_user_info_response_pb';

// const baseURL = 'ws://192.168.10.240:5050';
const baseURL = 'ws://127.0.0.1:5050';
let ws: WebSocket;

// 路由回调
const routers: Map<number, any> = new Map();

// 初始化
const initWS = function(): void {
    ws = new WebSocket(baseURL+"/v1/ws");
    ws.binaryType = 'arraybuffer';

    //连接打开时触发
    ws.onopen = onOpen;
    //接收到消息时触发
    ws.onmessage = onMessage;
    //连接关闭时触发
    ws.onclose = onClose;
}

function onOpen(evt: Event){
    console.log("onopen..");
}

function onMessage(evt: MessageEvent) {

    //console.log("onmessage evt.data:", new Uint8Array(evt.data));
    const messageBase = MessageBase.MessageBase.deserializeBinary(evt.data);
    console.log("onmessage messageBase.cmd:", messageBase.getCmd());

    const base = GetUserInfoResponse.GetUserInfoResponse.deserializeBinary(messageBase.getData());
    routers.get(messageBase.getCmd() as number)(base);
}

function onClose(evt: CloseEvent){
    console.log("onclose...");
}

const sendWS = function(msg: any){
    ws.send(msg);
}

const onRouter = function(router: number, func: any){
    routers.set(router, func);
}

export {initWS, sendWS, onRouter};