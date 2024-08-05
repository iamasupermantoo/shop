<template>
  <div class="q-mb-sm">
    <!-- 文本域 -->
    <div v-if="type == InputTypeList.TextArea">
      <q-input
        dense
        outlined
        rows="5"
        :readonly="readonly"
        v-model="currentValue"
        :label="label"
        @update:model-value="$emit('refresh', currentValue, field)"
        type="textarea"
      ></q-input>
    </div>

    <!-- 富文本编辑器 -->
    <div v-else-if="type == InputTypeList.Editor">
      <input-editor-component
        v-model="currentValue"
        :label="label"
        @update:model-value="$emit('refresh', currentValue, field)"
        :readonly="readonly"
      ></input-editor-component>
    </div>

    <!-- 选择框 -->
    <div v-else-if="type == InputTypeList.Select">
      <q-select
        v-if="options && options.length >= 1"
        dense
        outlined
        v-model="currentValue"
        :label="label"
        :readonly="readonly"
        @update:model-value="$emit('refresh', currentValue, field)"
        :options="options"
        emit-value
        map-options
      >
      </q-select>
    </div>

    <!--    下级选项-->
    <div v-else-if="type == InputTypeList.InputChildren">
      <div class="q-mx-md">
        <ChildrenInputComponent
          :type="type"
          v-model="currentValue"
          :label="label"
          :options="options"
          :readonly="readonly"
          @refresh="$emit('refresh', currentValue, field)"
        ></ChildrenInputComponent>
      </div>
    </div>

    <!-- 单选框 -->
    <div v-else-if="type == InputTypeList.Radio">
      <div class="text-body2 text-bold text-grey">{{ label }}</div>
      <div class="q-gutter-sm">
        <q-radio
          v-model="currentValue"
          :val="radio.value"
          :label="radio.label"
          v-for="(radio, radioIndex) in options"
          :key="radioIndex"
        />
      </div>
    </div>

    <!-- 多选框 -->
    <div v-else-if="type == InputTypeList.Checkbox">
      <div class="text-body2 text-bold text-grey">{{ label }}</div>
      <div v-if="currentValue">
        <q-checkbox
          v-for="(checkbox, checkboxIndex) in options"
          :readonly="readonly"
          v-model="currentValue[checkboxIndex].value"
          :label="checkbox.label"
          class="no-margin"
          @update:model-value="$emit('refresh', currentValue, field)"
          :key="checkboxIndex"
        >
        </q-checkbox>
      </div>
    </div>

    <!-- 开关类型 -->
    <div v-else-if="type == InputTypeList.Toggle">
      <div>
        <div>
          <q-toggle
            v-model="currentValue"
            :label="label"
            :readonly="readonly"
            @update:model-value="$emit('refresh', currentValue, field)"
            :true-value="options[0].value"
            :false-value="options[1].value"
          ></q-toggle>
        </div>
      </div>
    </div>

    <!-- 时间类型 -->
    <div v-else-if="type == InputTypeList.DatePicker">
      <q-input
        dense
        outlined
        :label="label"
        :readonly="readonly"
        :model-value="currentValue"
      >
        <template v-slot:append>
          <q-icon name="event" class="cursor-pointer">
            <q-popup-proxy
              ref="qDateProxy"
              cover
              transition-show="scale"
              transition-hide="scale"
            >
              <q-date
                v-model="currentValue"
                mask="YYYY/MM/DD"
                @update:model-value="$emit('refresh', currentValue, field)"
              >
                <div class="row items-center justify-end">
                  <q-btn v-close-popup label="关闭" color="primary" flat />
                </div>
              </q-date>
            </q-popup-proxy>
          </q-icon>
        </template>
      </q-input>
    </div>

    <!-- 时间范围类型 -->
    <div v-else-if="type == InputTypeList.RangeDatePicker">
      <q-input
        dense
        outlined
        :label="label"
        :readonly="readonly"
        :model-value="currentValue ? currentValue.from + ' - ' + currentValue.to : ''"
      >
        <template v-slot:append>
          <q-icon name="event" class="cursor-pointer">
            <q-popup-proxy
              ref="qDateProxy"
              cover
              transition-show="scale"
              transition-hide="scale"
            >
              <q-date
                range
                v-model="currentValue"
                @update:model-value="$emit('refresh', typeof currentValue == 'string' ? {from: '', to: ''} : currentValue, field)"
              >
                <div class="row items-center justify-end">
                  <q-btn v-close-popup label="关闭" color="primary" flat />
                </div>
              </q-date>
            </q-popup-proxy>
          </q-icon>
        </template>
      </q-input>
    </div>

    <!-- 文件上传 -->
    <div v-else-if="type == InputTypeList.File || type == InputTypeList.Image">
      <uploader-component
        :label="label"
        v-model="currentValue"
        @uploaded="updateValueFunc"
      ></uploader-component>
    </div>
    <div v-else-if="type == InputTypeList.Icon">
      <div v-if="currentValue">
        <q-img :src="imageSrc(currentValue)" width="60px" height="60px" v-if="currentValue != ''"></q-img>
      </div>
      <uploader-component
        :label="label"
        v-model="currentValue"
        @uploaded="updateValueFunc"
        style="display: none"
      ></uploader-component>
    </div>
    <!-- 多文件, 多图上传 -->
    <div v-else-if="type == InputTypeList.Images">
      <uploader-component
        :label="label"
        v-model="currentValue"
        :multiple="true"
        @uploaded="updateValueFunc"
      ></uploader-component>
    </div>

    <!-- 数字input -->
    <div v-else-if="type == InputTypeList.Number">
      <q-input
        dense
        outlined
        v-model.number="currentValue"
        :label="label"
        type="number"
        :readonly="readonly"
        @change="$emit('refresh', currentValue, field)"
      ></q-input>
    </div>

    <!-- 密码格式 -->
    <div v-else-if="type == InputTypeList.Password">
      <q-input
        dense
        outlined
        v-model="currentValue"
        :label="label"
        type="password"
        :readonly="readonly"
        @change="$emit('refresh', currentValue, field)"
      ></q-input>
    </div>

    <div v-else-if="type == InputTypeList.InputJson">
      <JsonInputComponent
        v-if="type == InputTypeList.InputJson"
        :type="type"
        v-model="currentValue"
        :label="label"
        :options="options"
        :readonly="readonly"
        @refresh="$emit('refresh', currentValue, field)"
      ></JsonInputComponent>
    </div>

    <!-- 文本 -->
    <div v-else-if="type == InputTypeList.InputTranslate">
      <div class="text-caption text-primary cursor-pointer" @click="$router.push('/systems/translate/index?field=' + currentValue)">多语言翻译内容 >>>>> 点击跳转</div>
      <q-input
        dense
        outlined
        v-model="currentValue"
        :label="label"
        type="text"
        readonly
        @change="$emit('refresh', currentValue, field)"
      ></q-input>
    </div>

    <div v-else>
      <q-input
        dense
        outlined
        v-model="currentValue"
        :label="label"
        type="text"
        :readonly="readonly"
        @change="$emit('refresh', currentValue, field)"
      ></q-input>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue'
import InputEditorComponent from 'src/components/editor.vue';
import JsonInputComponent from 'src/components/inputJson.vue';
import ChildrenInputComponent from 'src/components/inputChildren.vue'
import UploaderComponent from 'src/components/uploader.vue';
import { InputTypeList } from 'src/utils/define';
import {imageSrc} from 'src/utils';
import {date} from 'quasar';

defineOptions({
  name: 'ComponentsInput'
})

const emits = defineEmits(['refresh'])
const props = defineProps({
  type: { type: Number, default: InputTypeList.Text },
  field: { type: String, default: ''},
  label: { type: String, default: '' },
  modelValue: { type: undefined, default: '' },
  options: { type: null, default: [] },
  readonly: { type: Boolean, default: false },
})

const currentValue = ref(null) as any

// 上传完成刷新
const updateValueFunc = (val: string) => {
  emits('refresh', val, props.field);
}

onMounted(() => {
  currentValue.value = filterModelValueFunc(props.modelValue)
})

watch(() => props.modelValue, (val: any) => {
  currentValue.value = filterModelValueFunc(val)
})

// 过滤值
const filterModelValueFunc = (value: any) => {
  if (value == null) {
    return value
  }

  switch (props.type) {
    case InputTypeList.DatePicker:
      value = date.formatDate(value, 'YYYY/MM/DD')
      break
  }

  emits('refresh', value, props.field);
  return value
}

</script>
