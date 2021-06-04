<template>
  
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue';

import {onRouter} from '@/utils/websocket';

import * as CMD from '@/protoim/cmd_pb';


export default defineComponent({
    // 当前组件的名字
  name: 'Home',
  
  setup() {

    const userList = reactive([])

    onRouter(CMD.CMD.UPDATE_USER_FRIEND_INFO as number, function(data: any){
      console.log(data.getTag());
      console.log(data.getUid());
      console.log(data.getName());
    });

    return {userList};
  }

});
</script>

<style scoped>
</style>