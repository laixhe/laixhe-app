const baseURL = 'ws://192.168.10.240:5050';
let ws: WebSocket;

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
    console.log("onmessage evt:", evt);
    console.log("onmessage evt.data:", new Uint8Array(evt.data));
}

function onClose(evt: CloseEvent){
    console.log("onclose...");
}

const sendWS = function(msg: any){
    ws.send(msg);
}

export {initWS, sendWS};