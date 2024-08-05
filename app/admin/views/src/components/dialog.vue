<template>
  <div>
    <q-dialog
      v-model="dialog"
      :full-width="currentConfig.fullWidth"
      :full-height="currentConfig.fullHeight"
    >
      <q-card
        :style="sizingList.get(currentConfig.sizing)"
        :class="currentConfig.fullHeight ? 'column full-height' : ''"
      >
        <q-card-section v-if="currentConfig.title != ''">
          <div class="text-h6">{{ currentConfig.title }}</div>
          <div class="text-caption text-grey" style="white-space: pre-wrap">{{currentConfig.small}}</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <div class="text-body1" v-if="currentConfig.content != ''" style="word-break:break-all">
            {{ currentConfig.content }}
          </div>

          <div class="q-mt-md">
            <div
              v-for="(input, inputIndex) in currentConfig.inputList"
              :key="inputIndex"
            >
              <InputLayoutsComponent
                v-model="currentConfig.params[input.field]"
                :label="input.label"
                :type="input.type"
                :options="input.data"
                :readonly="input.readonly"
              ></InputLayoutsComponent>
            </div>
          </div>
        </q-card-section>

        <q-card-actions
          align="right"
          v-if="currentConfig.buttons.cancel || currentConfig.buttons.confirm"
        >
          <q-btn
            flat
            :color="
              currentConfig.buttons.cancel.color == ''
                ? 'grey'
                : currentConfig.buttons.cancel.color
            "
            :label="currentConfig.buttons.cancel.label"
            :size="currentConfig.buttons.cancel.size"
            v-close-popup
          />
          <q-btn
            flat
            :color="
              currentConfig.buttons.confirm.color == ''
                ? 'primary'
                : currentConfig.buttons.confirm.color
            "
            :label="currentConfig.buttons.confirm.label"
            :size="currentConfig.buttons.confirm.size"
            @click="submitFunc()"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import { api } from 'src/boot/axios';
import {NotifyPositive} from 'src/utils/notify';
import InputLayoutsComponent from 'src/components/inputLayouts.vue';

defineOptions({
  name: 'ComponentsDialog'
})

const emits = defineEmits(['done'])
const sizingList = new Map([
  ['small', 'width: 300px'],
  ['medium', 'width: 700px; max-width: 80vw'],
])

const dialog = ref(false)
let currentConfig = {id: 'none', url: '', title: '', content: '', sizing: 'medium', fullWidth: false, fullHeight: false, params: {}, inputList: [], buttons: {cancel: {}, confirm: {}}} as any

const setDialogConfig = (config: any) => {
  currentConfig = config
}

// 打开模态框
const dialogOpenFunc = () => {
  dialog.value = true
}

// 提交请求
const submitFunc = () => {
  api.post(currentConfig.url, currentConfig.params).then(() => {
    emits('done', currentConfig.id);
    dialog.value = false

    // 执行完成方法
    NotifyPositive('提交成功')
    if (currentConfig.buttons.confirm.hasOwnProperty('done')) {
      currentConfig.buttons.confirm.done(currentConfig)
    }
  })
}

defineExpose({
  setDialogConfig, dialogOpenFunc
})
</script>
