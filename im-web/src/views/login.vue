<template>
<div>

  <el-form ref="loginForm" :model="userForm" :rules="userRules" label-width="80px" class="login-box">
    <h3 class="login-title">欢迎登录</h3>
    <el-form-item label="账号" prop="username">
      <el-input type="text" placeholder="请输入账号" v-model="userForm.username"/>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input type="password" placeholder="请输入密码" v-model="userForm.password"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" v-on:click="onSubmit()">登录</el-button>
    </el-form-item>
  </el-form>

</div>
</template>

<script lang="ts">
import { defineComponent, ref, reactive, toRefs } from 'vue'

// 自定义组件并导出
export default defineComponent({
  // 当前组件的名字
  name: 'Login',

  // 在创建组件之前执行，是组合API的入口函数
  // 返回对象中的属性或方法，在模板中可以直接使用
  setup(){

    const state = reactive({
      userForm: {
        username: '',
        password: ''
      },
      loginForm: ref<any>(null)
    });

    const userRules = {
      username: [
        {required: true, message: '账号不可为空', trigger: 'blur'}
      ],
      password: [
        {required: true, message: '密码不可为空', trigger: 'blur'}
      ]
    }

    const onSubmit = () => {
      state.loginForm.validate().then((valid: boolean) => {
        if (valid) {
            console.log('userForm...')
            console.log(state.userForm)
        }
      });
    }

    return {...toRefs(state), userRules, onSubmit}
  }
})
</script>

<style scoped>
.login-box {
  border: 1px solid #DCDFE6;
  width: 350px;
  margin: 180px auto;
  padding: 35px 35px 15px 35px;
  border-radius: 5px;
  -webkit-border-radius: 5px;
  -moz-border-radius: 5px;
  box-shadow: 0 0 25px #909399;
}
.login-title {
  text-align: center;
  margin: 0 auto 40px auto;
  color: #303133;
}
</style>
