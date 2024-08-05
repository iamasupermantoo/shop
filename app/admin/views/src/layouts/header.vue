<template>
  <div>
    <q-toolbar>
      <q-btn dense flat round icon="menu" @click="drawerLeftFunc" />
      <div
        @click="$router.push('/')"
        class="row justify-start items-center cursor-pointer"
      >
        <div class="text-h6 q-ml-md">管理系统</div>
      </div>
      <q-toolbar-title></q-toolbar-title>
      <q-btn
        dense
        flat
        :label="'过期:' + $initStore.userInfo.expiredAt.substring(0, 10)"
        v-if="$q.screen.width > 428"
        color="deep-orange"
      ></q-btn>
      <q-btn-dropdown flat>
        <template v-slot:label>
          <div class="q-ml-xs">{{ $initStore.userInfo.userName }}</div>
        </template>
        <q-list class="bg-primary text-white text-body2">
          <q-item
            v-close-popup
            clickable
            @click="openDialogFunc(dialogUserInfoKey)"
          >
            <q-item-section>
              <div class="row q-gutter-sm items-center">
                <q-icon name="sym_o_contact_mail" size="sm"></q-icon>
                <div>用户信息</div>
              </div>
            </q-item-section>
          </q-item>
          <q-item
            v-close-popup
            clickable
            @click="openDialogFunc(dialogUpdatePassword)"
          >
            <q-item-section>
              <div class="row q-gutter-sm items-center">
                <q-icon name="sym_o_lock" size="sm"></q-icon>
                <div>登陆密码</div>
              </div>
            </q-item-section>
          </q-item>
          <q-item
            v-close-popup
            clickable
            @click="openDialogFunc(dialogUpdateSecurityPassword)"
          >
            <q-item-section>
              <div class="row q-gutter-sm items-center">
                <q-icon name="sym_o_security" size="sm"></q-icon>
                <div>安全密码</div>
              </div>
            </q-item-section>
          </q-item>
          <q-separator></q-separator>
          <q-item v-close-popup clickable @click="logoutFunc">
            <q-item-section>
              <div class="row q-gutter-sm items-center">
                <q-icon name="sym_o_logout" size="sm"></q-icon>
                <div>退出登陆</div>
              </div>
            </q-item-section>
          </q-item>
        </q-list>
      </q-btn-dropdown>
      <q-btn flat @click="drawerRightFunc" size="md" v-if="$initStore.userInfo.seatLink != ''">
        <q-avatar size="30px">
          <img :src="imageSrc('/assets/icon/online.png')" alt="">
          <q-badge floating color="red" v-if="drawerRightNums > 0">{{drawerRightNums}}</q-badge>
        </q-avatar>
      </q-btn>
      <q-btn
        dense
        flat
        round
        icon="sym_o_fullscreen"
        v-if="$q.screen.width > 428"
        @click="$q.fullscreen.toggle()"
      ></q-btn>
    </q-toolbar>

    <DialogComponent ref="dialogRef"></DialogComponent>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import {imageSrc} from 'src/utils';
import {useInitStore} from 'stores/init';
import {useRouter} from 'vue-router';
import {InputTypeList} from 'src/utils/define'
import DialogComponent from 'src/components/dialog.vue'
import {ShowNotify} from 'src/utils/notify'

defineOptions({
  name: 'LayoutsHeader'
})

const dialogRef = ref(null) as any
const $initStore = useInitStore()
const $router = useRouter()
const dialogUserInfoKey = 'userInfo';
const dialogUpdatePassword = 'updatePassword';
const dialogUpdateSecurityPassword = 'updateSecurity';
const dialogList = new Map([
  [
    dialogUserInfoKey,
    {
      id: dialogUserInfoKey,
      url: '/auth/update',
      title: '更新管理信息',
      sizing: 'small',
      params: {
        avatar: $initStore.userInfo.avatar,
        nickName: $initStore.userInfo.nickName,
        email: $initStore.userInfo.email,
        domains: $initStore.userInfo.domains,
        online: $initStore.userInfo.online,
        seatLink: $initStore.userInfo.seatLink,
      } as any,
      inputList: [
        {
          label: '管理头像',
          field: 'avatar',
          type: InputTypeList.Image,
          default: '',
          readonly: false,
          data: [],
        },
        {
          label: '管理昵称',
          field: 'nickName',
          type: InputTypeList.Text,
          default: '',
          readonly: false,
          data: [],
        },
        {
          label: '管理邮箱',
          field: 'email',
          type: InputTypeList.Text,
          default: '',
          readonly: false,
          data: [],
        },
        {
          label: '坐席链接',
          field: 'seatLink',
          type: InputTypeList.Text,
          default: '',
          readonly: false,
          data: []
        },
        {
          label: '客服链接',
          field: 'online',
          type: InputTypeList.Text,
          default: '',
          readonly: false,
          data: []
        },
        {
          label: '绑定前端域名',
          field: 'domains',
          type: InputTypeList.TextArea,
          default: '',
          readonly: false,
          data: []
        }
      ],
      buttons: {
        cancel: { label: '取消', color: 'grey', size: 'md' },
        confirm: { label: '提交', color: 'primary', size: 'md', done: (conf: any) => {
            //  更新当前用户信息
            $initStore.userInfo.domains = conf.params.domains
            $initStore.userInfo.avatar = conf.params.avatar
            $initStore.userInfo.nickName = conf.params.nickName
            $initStore.userInfo.email = conf.params.email
            $initStore.userInfo.online = conf.params.online
            $initStore.userInfo.seatLink = conf.params.seatLink
            $initStore.updateUserInfo($initStore.userInfo)
          } },
      },
    },
  ],
  [
    dialogUpdatePassword,
    {
      id: dialogUpdatePassword,
      url: '/auth/update/password',
      title: '修改登录密码',
      sizing: 'small',
      params: {
        type: 1,
        oldPassword: '',
        newPassword: '',
        cmfPassword: '',
      } as any,
      inputList: [
        {
          label: '旧密码',
          field: 'oldPassword',
          type: InputTypeList.Password,
          default: '',
          readonly: false,
          data: [],
        },
        {
          label: '新密码',
          field: 'newPassword',
          type: InputTypeList.Password,
          default: '',
          readonly: false,
          data: [],
        },
        {
          label: '确认密码',
          field: 'cmfPassword',
          type: InputTypeList.Password,
          default: '',
          readonly: false,
          data: [],
        },
      ],
      buttons: {
        cancel: { label: '取消', color: 'grey', size: 'md' },
        confirm: { label: '提交', color: 'primary', size: 'md', done: () => {} },
      },
    },
  ],
  [
    dialogUpdateSecurityPassword,
    {
      id: dialogUpdateSecurityPassword,
      url: '/auth/update/password',
      title: '修改安全密码',
      sizing: 'small',
      params: {
        type: 2,
        oldPassword: '',
        newPassword: '',
        cmfPassword: '',
      } as any,
      inputList: [
        {
          label: '旧密码',
          field: 'oldPassword',
          type: InputTypeList.Password,
          default: '',
          readonly: false,
          data: [],
        },
        {
          label: '新密码',
          field: 'newPassword',
          type: InputTypeList.Password,
          default: '',
          readonly: false,
          data: [],
        },
        {
          label: '确认密码',
          field: 'cmfPassword',
          type: InputTypeList.Password,
          default: '',
          readonly: false,
          data: [],
        },
      ],
      buttons: {
        cancel: { label: '取消', color: 'grey', size: 'md' },
        confirm: { label: '提交', color: 'primary', size: 'md', done: () => {} },
      },
    },
  ],
])

// 打开模态框方法
const openDialogFunc = (key: string) => {
  const configInfo = dialogList.get(key);
  dialogRef.value.setDialogConfig(configInfo);
  dialogRef.value.dialogOpenFunc();
}

const emits = defineEmits(['drawer-left', 'drawer-right'])
const drawerRight = ref(false) as any
const drawerRightNums = ref(0)

// 左侧开关控制
const drawerLeftFunc = () => {
  emits('drawer-left')
}

const drawerRightFunc = () => {
  emits('drawer-right')
}

// 左侧按钮打开事件
const drawerRightEventFunc = (value: boolean) => {
  drawerRight.value = value
  drawerRightNums.value = 0
}

// 监听客服提示
window.addEventListener('message', (msg: any) => {
  if (!drawerRight.value && msg.data.hasOwnProperty('Data')) {
    ShowNotify({message: '客服消息[用户ID:'+msg.data.SenderId+'] => ' + msg.data.Data, color: 'accent', icon: 'headset_mic', position: 'top-right', timeout: 500})
    drawerRightNums.value++
  }
})

const logoutFunc = () => {
  $initStore.updateUserToken('')
  window.location.reload()
}

defineExpose({drawerRightEventFunc})
</script>

<style scoped>

</style>
