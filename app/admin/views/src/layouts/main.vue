<template>
  <q-layout view="hHh LpR lFr">
    <q-header bordered class="bg-primary text-white">
      <LayoutsHeader @drawer-left="drawerLeftFunc" @drawer-right="drawerRightFunc" ref="drawerRightRef"></LayoutsHeader>
    </q-header>

    <q-drawer
      show-if-above
      v-model="drawerLeft"
      side="left"
      bordered
      :behavior="$q.platform.is.mobile ? 'mobile' : 'desktop'"
      :width="220"
    >
      <LayoutsLeft></LayoutsLeft>
    </q-drawer>

    <q-drawer v-model="drawerRight" side="right" @update:model-value="drawerRightChangeFunc" bordered :width="380">
      <LayoutsRight></LayoutsRight>
    </q-drawer>

    <q-page-container>
      <router-view :key="$router.currentRoute.value.fullPath" />
    </q-page-container>
  </q-layout>
  <audio ref="audioRef" @play="audioPaying = true" @pause="audioPaying = false"></audio>
</template>

<script setup lang="ts">
defineOptions({
  name: 'LayoutsMain'
})

const $userStore = useInitStore()
const drawerLeft = ref(false);
const drawerRight = ref(false);
const drawerRightRef = ref(null) as any
const audioRef = ref(null) as any
const audioPaying = ref(false)

const drawerLeftFunc = () => {
  drawerLeft.value = !drawerLeft.value;
}
const drawerRightFunc = () => {
  drawerRight.value = !drawerRight.value;
  drawerRightChangeFunc()
}

const drawerRightChangeFunc = () => {
  drawerRightRef.value.drawerRightEventFunc(drawerRight.value)
}

onMounted(() => {
  initWebsocket().setOnMessageOpenFunc(() => {
    appWebsocket.sendMessageJsonFunc({ op: MessageOperateBindUser, data: $userStore.userToken })
  }).setOnMessageFunc((msg: any) => {
    switch (msg.op) {
      case MessageOperateAudio:
        // 提示窗口
        Notify.create({type: msg.data.type, position: msg.data.position, timeout: msg.data.timeout, message: msg.data.label});

        // 添加提示音
        setTimeout(() => {
          if (!audioPaying.value) {
            audioRef.value.src = imageSrc(msg.data.source)
            audioRef.value.play()
          }
        }, 200)
        break
    }
  }).connect()
});

import {Notify} from 'quasar';
import {imageSrc} from 'src/utils';
import LayoutsHeader from 'src/layouts/header.vue'
import LayoutsLeft from 'src/layouts/left.vue'
import LayoutsRight from 'src/layouts/right.vue'
import {useInitStore} from 'stores/init';
import {onMounted, ref} from 'vue';
import {initWebsocket, appWebsocket, MessageOperateAudio, MessageOperateBindUser} from 'boot/websocket';
</script>
