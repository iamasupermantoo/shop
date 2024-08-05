<template>
  <div>
    <q-uploader
      :url="baseURL + '/auth/upload'"
      field-name="file"
      ref="uploaderRef"
      auto-upload
      style="box-shadow: none; width: auto"
      :multiple="multiple"
      :headers="[{ name: 'Authorization', value: 'Bearer ' + $initStore.userToken }]"
      @uploaded="uploadedFunc"
      @start="uploaderStartFunc"
      @finish="uploaderFinishFunc"
      @failed="uploadFailedFunc"
    >
      <template v-slot:header="scope">
        <div class="q-my-xs q-mx-sm">
          <div class="row justify-between items-center">
            <div>{{ label }}</div>
            <div>
              <q-btn
                v-if="scope.canAddFiles"
                type="a"
                icon="add_box"
                @click="scope.pickFiles"
                round
                dense
                flat
              >
                <q-uploader-add-trigger />
                <q-tooltip>Pick Files</q-tooltip>
              </q-btn>
            </div>
          </div>
        </div>
      </template>

      <template v-slot:list="scope">
        <q-list separator v-if="multiple">
          <q-item v-for="(image, imageIndex) in currentValue" :key="imageIndex">
            <q-item-section thumbnail>
              <img :src="imageSrc(image)" alt="" class="q-ml-sm" />
            </q-item-section>
            <q-item-section></q-item-section>
            <q-item-section side>
              <q-btn
                class="gt-xs"
                size="12px"
                flat
                dense
                round
                icon="delete"
                @click="deleteUploadedEventFunc(imageIndex)"
              />
            </q-item-section>
          </q-item>
        </q-list>

        <div
          @click="scope.pickFiles"
          class="row justify-center no-padding"
          v-else
        >
          <q-uploader-add-trigger />
          <img
            :src="imageSrc(currentValue)"
            alt=""
            v-if="currentValue && currentValue != ''"
          />
        </div>
      </template>
    </q-uploader>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue';
import { Loading, QSpinnerBars } from 'quasar';
import { useInitStore } from 'src/stores/init';
import { imageSrc } from 'src/utils';
import { NotifyNegative } from 'src/utils/notify';

defineOptions({
  name: 'ComponentsUploader'
})

const emits = defineEmits(['uploaded'])
const props = defineProps({
  label: { type: String, default: '' },
  modelValue: { type: undefined, default: '' },
  multiple: { type: Boolean, default: false },
})

const $initStore = useInitStore();
const uploaderRef = ref(null) as any
const baseURL = process.env.baseURL
const currentValue = ref('') as any

onMounted(() => {
  currentValue.value = props.modelValue
})

watch(() => props.modelValue, (val: any) => {
  currentValue.value = val
})

// 激活选择文件
const pickFilesFunc = () => {
  uploaderRef.value.pickFiles();
};

// 开始上传
const uploaderStartFunc = () => {
  Loading.show({
    spinner: QSpinnerBars,
    spinnerColor: 'secondary',
    spinnerSize: 50,
    message: 'Some important process is in progress. Hang on...',
  });
};

// 上传完成回调方法
const uploadedFunc = (info: any) => {
  const imagePath = JSON.parse(info.xhr.response).data[0];
  if (props.multiple) {
    currentValue.value.push(imagePath)
    emits('uploaded', currentValue.value);
  } else {
    (currentValue.value = imagePath)
    emits('uploaded', currentValue.value);
  }
};

// 上传完成
const uploaderFinishFunc = () => {
  Loading.hide();
};

// 删除上传的图片
const deleteUploadedEventFunc = (index: any) => {
  currentValue.value.splice(index, 1);
  emits('uploaded', currentValue.value);
};

// 上传失败方法
const uploadFailedFunc = () => {
  NotifyNegative('文件上传失败, 请检查文件是否符合上传...');
  Loading.hide();
};

defineExpose({
  pickFilesFunc
})
</script>
<style scoped></style>
