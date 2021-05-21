<template>
<button v-on:click="onSend()">发送</button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

import {sendWS} from '@/utils/websocket';

import * as CMD from '@/protoim/cmd_pb';
import * as MessageBase from '@/protoim/message_base_pb';
import * as GetUserInfo from '@/protoim/get_user_info_pb';


export default defineComponent({
    // 当前组件的名字
  name: 'Home',
  
  setup() {

    let i = 100;
    // 点击发送按钮
    const onSend = () => {

      let getUserInfoRequest = new GetUserInfo.GetUserInfoRequest();
      getUserInfoRequest.setUid(i);
      // 将 protobuf 序列化为字节数组 Uint8Array
      let getUserInfoRequestBuffer = getUserInfoRequest.serializeBinary();

      let messageBase = new MessageBase.MessageBase();
      messageBase.setCmd(CMD.CMD.GETUSERINFOREQ);
      messageBase.setData(getUserInfoRequestBuffer);
      let messageBaseBuffer = messageBase.serializeBinary();

      sendWS(messageBaseBuffer);
      i++;
    }

    return {onSend};
  }

});
</script>

<style scoped>
</style>