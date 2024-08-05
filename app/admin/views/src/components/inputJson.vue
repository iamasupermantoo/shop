<template>
  <div>
    <div class="text-body2 text-grey q-mb-xs" v-if="!readonly">{{ label }}</div>
    <div v-for="(inputRow, inputRowIndex) in options" :key="inputRowIndex">
      <div class="row items-center q-gutter-xs">
        <div
          v-for="(inputCol, inputColIndex) in inputRow"
          :key="inputColIndex"
          class="col"
        >
          <InputComponent
            v-if="modelValue"
            :type="inputCol.type"
            :field="inputCol.field"
            v-model="currentValue[inputCol.field]"
            :label="inputCol.label"
            :options="inputCol.data"
            :readonly="inputCol.readonly"
            @refresh="componentsRefreshValue"
          ></InputComponent>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue';
import InputComponent from 'src/components/input.vue';
import { InputTypeList } from 'src/utils/define';

defineOptions({
  name: 'ComponentsJson'
})

const emits = defineEmits(['refresh'])
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

watch(() => props.modelValue, (val: any) => {
  currentValue.value = val
})

//  组件刷新值
const componentsRefreshValue = (value: any, field: string) => {
  currentValue.value[field] = value;
  emits('refresh', currentValue.value);
};
</script>
