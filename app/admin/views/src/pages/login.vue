<template>
  <div class="row items-center q-pa-md login-bg">
    <div class="col login-img text-center">
      <q-img src="/admin/images/bg.svg"></q-img>
    </div>
    <div class="col text-center row justify-center">
      <div class="full-width" style="max-width: 320px">
        <q-img :src="imageSrc('')" height="60px" width="60px" no-spinner></q-img>
        <div class="text-h4 text-bold text-grey q-pt-md q-pb-lg">
          BaJie Admin
        </div>
        <div class="q-gutter-md q-mt-lg">
          <q-input outlined dense v-model="params.username" label="账号">
            <template v-slot:prepend>
              <q-icon name="sym_o_person" />
            </template>
          </q-input>
          <q-input
            outlined
            dense
            v-model="params.password"
            type="password"
            clearable
            label="密码"
          >
            <template v-slot:prepend>
              <q-icon name="sym_o_lock" />
            </template>
          </q-input>
          <q-input outlined dense v-model="params.captchaVal" label="验证码">
            <template v-slot:prepend>
              <q-icon name="sym_o_security" />
            </template>
            <template v-slot:append>
              <q-img
                no-spinner
                v-if="params.captchaId !== ''"
                :src="
                  imageSrc(baseURL + '/captcha/' + params.captchaId + '/120-48')
                "
                width="120px"
                height="32px"
                @click="refreshCaptchaFunc"
              ></q-img>
            </template>
          </q-input>
          <div class="row justify-between items-center">
            <q-checkbox
              v-model="params.remember"
              size="xs"
              class="text-grey-8"
              label="记住密码"
            />
            <div class="text-primary">忘记密码？</div>
          </div>

          <div>
            <q-btn
              class="full-width bg-primary text-white"
              label="登陆"
              @click="submitFunc"
            ></q-btn>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { captchaCreateAPI, adminLoginAPI } from 'src/apis';
import { imageSrc } from 'src/utils';
import { useInitStore } from 'src/stores/init';
import {onMounted, ref} from 'vue';
import {dynamicRouter} from 'src/router';
import {useRouter} from 'vue-router';

defineOptions({
  name: 'IndexLogin'
})

const $initStore = useInitStore()
const $router = useRouter()
const baseURL = process.env.baseURL
const params = ref({username: '', password: '', captchaId: '', captchaVal: '', remember: true})

onMounted(() => {
  refreshCaptchaFunc()
})

// 刷新验证码方法
const refreshCaptchaFunc = () => {
  captchaCreateAPI().then((res: any) => {
    params.value.captchaId = res
  })
}

// 登录操作
const submitFunc = () => {
  adminLoginAPI(params.value).then((res: any) => {
    $initStore.updateUserToken(res.token);
    $initStore.updateUserInfo(res.info)
    $initStore.updateRouterList(res.routerList)
    $initStore.updateMenuList(res.menuList)

    // 加载路由, 跳转到首页
    dynamicRouter($router, res.menuList).then(() => {
      $router.push('/')
    })
  }).catch(() => {
    refreshCaptchaFunc()
  })
}
</script>

<style scoped>
.login-bg {
  height: 100vh;
  padding-bottom: 30%;
}

.login-img {
  display: none;
}

@media screen and (min-width: 1000px) {
  .login-bg {
    background-size: 100% 100%;
    padding-bottom: 10%;
  }

  .login-img {
    display: block;
  }
}
</style>
