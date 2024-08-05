<template>
  <div>
    <div class="row justify-between q-mb-xs" v-if="!readonly">
      <div class="text-body2 text-grey">{{ label }}</div>
      <div class="row q-gutter-xs">
        <div>
          <q-btn
            dense
            flat
            outline
            class="bg-secondary text-white q-pa-xs"
            size="xs"
            @click="addChildrenFunc()"
            label="新增"
          ></q-btn>
        </div>
        <div>
          <q-btn
            dense
            flat
            outline
            class="bg-red text-white q-pa-xs"
            size="xs"
            @click="delChildrenFunc()"
            label="删除"
          ></q-btn>
        </div>
      </div>
    </div>

    <div v-for="(item, itemIndex) in currentValue" :key="itemIndex">
      <div class="text-caption text-grey">
        #<span style="margin-left: 1px">{{ itemIndex }}</span>
      </div>
      <div v-for="(inputRow, inputRowIndex) in options" :key="inputRowIndex">
        <div class="row items-center q-gutter-xs">
          <div
            v-for="(inputCol, inputColIndex) in inputRow"
            :key="inputColIndex"
            class="col"
          >
            <InputComponent
              :type="inputCol.type"
              :field="itemIndex+'-'+inputCol.field"
              v-model="currentValue[itemIndex][inputCol.field]"
              :label="inputCol.label"
              :options="inputCol.data"
              :readonly="inputCol.readonly"
              @refresh="componentsRefreshValue"
            ></InputComponent>
          </div>
        </div>
      </div>
    </div>
    <div
      v-if="currentValue.length == 0"
      class="text-caption text-center text-grey q-my-md"
    >
      暂无数据
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue';
import InputComponent from 'src/components/input.vue';
import { InputTypeList } from 'src/utils/define';

defineOptions({
  name: 'ComponentsChildren'
})

const emits = defineEmits(['refresh'])
const props = defineProps({
  type: { type: Number, default: InputTypeList.Text },
  label: { type: String, default: '' },
  modelValue: { type: null, default: [] as any },
  options: { type: null, default: [] },
  readonly: { type: Boolean, default: false },
})

const currentValue = ref('') as any

onMounted(() => {
  currentValue.value = props.modelValue == null ? [] :props.modelValue
})

// 组件数据刷新
const componentsRefreshValue = (value: any, field: string) => {
  const fieldList = field.split('-')
  currentValue.value[fieldList[0]][fieldList[1]] = value
  emits('refresh', currentValue.value);
}

// 添加子级
const addChildrenFunc = () => {
  const tmpValue = {} as any;
  for (let i = 0; i < props.options.length; i++) {
    for (let j = 0; j < props.options[i].length; j++) {
      switch (props.options[i][j].type) {
        case InputTypeList.InputChildren:
          tmpValue[props.options[i][j].field] = []
          break
        case InputTypeList.Checkbox:
          tmpValue[props.options[i][j].field] = props.options[i][j].data
          break
        default:
          tmpValue[props.options[i][j].field] = '';
      }
    }
  }

  currentValue.value.push(tmpValue);
  emits('refresh', currentValue.value);
};

// 删除子级
const delChildrenFunc = () => {
  currentValue.value.pop();
  emits('refresh', currentValue.value);
};

watch(() => props.modelValue, (val: any) => {
  currentValue.value = val
})
</script>
