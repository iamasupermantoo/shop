<template>
  <div style="max-width: 280px" class="ellipsis">
    <div
      v-if="
        currentScope.col.type == ContentTypeList.InputEditText ||
        currentScope.col.type == ContentTypeList.InputEditNumber ||
        currentScope.col.type == ContentTypeList.InputEditTextArea
      "
    >
      <div class="text-primary">{{ currentValue === null || currentValue === '' ? '- -' : currentValue }}</div>
      <q-popup-edit v-model="currentValue">
        <template v-slot:default>
          <q-input
            v-model="currentValue" color="primary"
            dense
            autofocus
            counter
            v-if="currentScope.col.type === ContentTypeList.InputEditTextArea"
            type="textarea"
            @keyup.enter.stop
          />
          <q-input
            v-model="currentValue"
            dense
            autofocus
            counter
            v-else-if="currentScope.col.type === ContentTypeList.InputEditText"
            @keyup.enter.stop
          />
          <q-input
            v-model.number="currentValue" color="primary"
            dense
            autofocus
            counter
            v-else-if="
              currentScope.col.type === ContentTypeList.InputEditNumber
            "
            type="number"
            @keyup.enter.stop
          />
          <div class="row justify-end">
            <div>
              <q-btn label="取消" color="primary" v-close-popup flat></q-btn>
            </div>
            <div>
              <q-btn
                label="确定"
                color="primary"
                v-close-popup
                flat
                @click="updatePopupEditFunc()"
              ></q-btn>
            </div>
          </div>
        </template>
      </q-popup-edit>
    </div>

    <!-- 开关按钮 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.InputEditToggle">
      <q-toggle
        v-model="currentValue"
        :true-value="currentScope.col.data[0].value"
        :false-value="currentScope.col.data[1].value"
        @update:model-value="updatePopupEditFunc()"
      />
    </div>

    <!-- 显示图片 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.Image">
      <q-img
        :src="imageSrc(currentValue)"
        loading="lazy"
        spinner-color="white"
        @click="
          imageList = [currentValue];
          imageDialog = true;
        "
        style="max-height: 50px; max-width: 50px"
      ></q-img>
    </div>

    <!-- 显示图片组 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.Images">
      <div style="width: 100px;" v-if="currentValue .length > 0"  @click="imageDialog = true;imageList = currentValue">
        <div class="row">
          <template v-for="(image, imageIndex) in currentValue" :key="imageIndex">
            <q-avatar size="40px" v-if="imageIndex < 4">
              <img :src="imageSrc(image)" alt="">
            </q-avatar>
          </template>
        </div>
      </div>
    </div>

    <!-- 显示时间 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.DatePicker">
      {{ date.formatDate(currentValue.toString(), 'YYYY/MM/DD HH:mm:ss') }}
    </div>

    <!-- 显示选择框 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.Select">
      <div
        v-for="(selected, selectedIndex) in currentScope.col.data"
        :key="selectedIndex"
      >
        <q-badge outline :color="quasarColorsList[selectedIndex % 25]"
          v-if="selected.value === currentValue"
          :label="selected.label"
        />
      </div>
    </div>

    <!-- 多语言类型 -->
    <div v-else-if="currentScope.col.type === ContentTypeList.Translate">
      <div v-if="translateList.length > 1">
        <div class="text-caption text-grey">{{translateList[1]}}</div>
        <div @click="$router.push('/systems/translate/index?field=' + translateList[0])" class="text-caption cursor-pointer text-primary">点击修改</div>
      </div>
      <div v-else>
        <div class="text-caption text-grey">{{translateList[0]}}</div>
        <div @click="$router.push('/systems/translate/index?field=' + translateList[0])" class="text-caption cursor-pointer text-primary">点击修改</div>
      </div>
    </div>

    <!-- 正常显示文本 -->
    <div v-else class="ellipsis">
      {{ typeof currentValue == 'number' ? currentValue : currentValue == '' ? '- -' : currentValue }}
      <q-popup-proxy v-if="currentValue">
        <q-card>
          <q-card-section>
            {{ currentValue }}
          </q-card-section>
        </q-card>
      </q-popup-proxy>
    </div>

    <!-- 显示图dialog地方 -->
    <div
      v-if="
        currentScope.col.type == ContentTypeList.Image ||
        currentScope.col.type == ContentTypeList.Images
      "
    >
      <q-dialog v-model="imageDialog" full-width>
        <q-card class="q-pa-sm">
          <q-carousel
            swipeable
            animated
            v-model="imageSlide"
            infinite
            navigation
            control-color="accent"
          >
            <q-carousel-slide
              :name="imageIndex"
              v-for="(image, imageIndex) in imageList"
              :key="imageIndex"
              class="column no-wrap"
            >
              <template v-slot:default>
                <q-img :src="imageSrc(image)" fit="contain"></q-img>
              </template>
            </q-carousel-slide>
          </q-carousel>
        </q-card>
      </q-dialog>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from 'vue';
import { imageSrc } from 'src/utils';
import { ContentTypeList, quasarColorsObject } from 'src/utils/define';
import { api } from 'src/boot/axios';
import {useRouter} from 'vue-router';
import {date} from 'quasar';

defineOptions({
  name: 'ComponentsColumnContent'
})

const $router = useRouter()
const emits = defineEmits(['done'])
const quasarColorsList = Object.values(quasarColorsObject)
const props = defineProps({
  url: { type: String, default: '' },
  Primary: { type: String, default: 'id' },
  modelValue: { type: undefined, default: '' },
  currentScope: {type: Object, default: () => {return {}}}
})

const imageDialog = ref(false)
const imageSlide = ref(0)
const imageList = ref([]) as any
const currentValue = ref('') as any
const translateList = ref([]) as any

onMounted(() => {
  currentValue.value = props.modelValue

  // 如果是翻译类型, 那么切割
  if (props.currentScope.col.type == ContentTypeList.Translate) {
    translateList.value = currentValue.value.toString().split('_')
  }
})

//  编辑方法提交
const updatePopupEditFunc = () => {
  const params = {} as any;
  params[props.Primary] = props.currentScope.row[props.Primary];
  params[props.currentScope.col.field] = currentValue.value;

  api.post(props.url, params).then(() => {
    emits('done', props.currentScope);
  });
};
</script>
