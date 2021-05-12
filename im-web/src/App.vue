<template>
  <router-link to="/login">登录</router-link>
  <router-view ></router-view>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  // 当前组件的名字
  name: 'App',
})
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
