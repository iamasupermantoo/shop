<template>
  <div>
    <!-- Json对象 -->
    <JsonInputComponent
      v-if="type == InputTypeList.InputJson"
      :type="type"
      v-model="currentValue"
      :label="label"
      :options="options"
      :readonly="readonly"
      @refresh="componentsRefreshValue"
    ></JsonInputComponent>

    <!-- 子级数组对象 -->
    <ChildrenInputComponent
      v-else-if="type == InputTypeList.InputChildren"
      :type="type"
      v-model="currentValue"
      :label="label"
      :options="options"
      :readonly="readonly"
      @refresh="componentsRefreshValue"
    ></ChildrenInputComponent>

    <!-- 正常input对象 -->
    <InputComponent
      v-else
      :type="type"
      v-model="currentValue"
      :label="label"
      :options="options"
      :readonly="readonly"
      @refresh="componentsRefreshValue"
    ></InputComponent>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from 'vue';
import { InputTypeList } from 'src/utils/define';
import InputComponent from 'src/components/input.vue';
import JsonInputComponent from 'src/components/inputJson.vue';
import ChildrenInputComponent from 'src/components/inputChildren.vue';

const emits = defineEmits(['update:modelValue'])
const props = defineProps({
  type: { type: Number, default: InputTypeList.Text },
  label: { type: String, default: '' },
  modelValue: { type: undefined, default: '' },
  options: { type: null, default: [] },
  readonly: { type: Boolean, default: false },
})

const currentValue = ref(null) as any
onMounted(() => {
  currentValue.value = props.modelValue
})

//  组件刷新值
const componentsRefreshValue = (value: any) => {
  emits('update:modelValue', value);
};
</script>
