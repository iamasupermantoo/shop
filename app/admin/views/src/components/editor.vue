<template>
  <div>
    <q-editor
      min-height="5rem"
      :readonly="readonly"
      ref="editorRef"
      :placeholder="label"
      :toolbar="[
        ['token'],
        [
          {
            label: $q.lang.editor.align,
            icon: $q.iconSet.editor.align,
            fixedLabel: true,
            options: ['left', 'center', 'right', 'justify'],
          },
        ],
        ['bold', 'italic', 'strike', 'underline', 'subscript', 'superscript'],
        ['hr', 'link', 'custom_btn'],
        [
          {
            label: $q.lang.editor.fontSize,
            icon: $q.iconSet.editor.fontSize,
            fixedLabel: true,
            fixedIcon: true,
            list: 'no-icons',
            options: [
              'size-1',
              'size-2',
              'size-3',
              'size-4',
              'size-5',
              'size-6',
              'size-7'
            ]
          },
          {
            label: $q.lang.editor.defaultFont,
            icon: $q.iconSet.editor.font,
            fixedIcon: true,
            list: 'no-icons',
            options: [
              'default_font',
              'arial',
              'arial_black',
              'comic_sans',
              'courier_new',
              'impact',
              'times_new_roman',
              'verdana'
            ]
          },
          'removeFormat',
        ],
        ['viewsource'],
      ]"
      :fonts="{
        arial: 'Arial',
        arial_black: 'Arial Black',
        comic_sans: 'Comic Sans MS',
        courier_new: 'Courier New',
        impact: 'Impact',
        times_new_roman: 'Times New Roman',
        verdana: 'Verdana',
      }"
      v-model="currentValue"
      @update:model-value="$emit('update:modelValue', currentValue)"
      :dense="$q.screen.lt.md"
    >
      <template v-slot:token>
        <q-btn-dropdown
          dense
          no-caps
          ref="editorTokenRef"
          no-wrap
          unelevated
          color="white"
          text-color="primary"
          label="图片｜颜色"
          size="sm"
        >
          <q-list dense>
            <q-item class="q-mb-md">
              <q-uploader
                flat
                auto-upload
                style="height: 36px"
                @uploaded="editorUploadedEventFunc"
                :headers="[
                  { name: 'Authorization', value: 'Bearer ' + $initStore.userToken },
                ]"
                :url="baseURL + '/upload'"
                field-name="file"
              >
                <template v-slot:header></template>
                <template v-slot:list="scope">
                  <div @click="scope.pickFiles">
                    <q-uploader-add-trigger />
                    <div
                      class="text-body2 text-grey"
                      style="height: 36px; line-height: 36px"
                    >
                      选择需要上传图片...
                    </div>
                  </div>
                </template>
              </q-uploader>
            </q-item>
            <q-item
              tag="label"
              clickable
              @click="
                editorEditTextColorFunc('backColor', editorBackgroundColor)
              "
            >
              <q-item-section side>
                <q-icon name="highlight" />
              </q-item-section>
              <q-item-section>
                <q-color
                  v-model="editorBackgroundColor"
                  default-view="palette"
                  no-header
                  no-footer
                  :palette="[
                    '#ffccccaa',
                    '#ffe6ccaa',
                    '#ffffccaa',
                    '#ccffccaa',
                    '#ccffe6aa',
                    '#ccffffaa',
                    '#cce6ffaa',
                    '#ccccffaa',
                    '#e6ccffaa',
                    '#ffccffaa',
                    '#ff0000aa',
                    '#ff8000aa',
                    '#ffff00aa',
                    '#00ff00aa',
                    '#00ff80aa',
                    '#00ffffaa',
                    '#0080ffaa',
                    '#0000ffaa',
                    '#8000ffaa',
                    '#ff00ffaa',
                    '#ff00ffaa',
                  ]"
                />
              </q-item-section>
            </q-item>

            <q-item
              tag="label"
              clickable
              @click="editorEditTextColorFunc('foreColor', editorTextColor)"
            >
              <q-item-section side>
                <q-icon name="format_paint" />
              </q-item-section>
              <q-item-section>
                <q-color
                  v-model="editorTextColor"
                  default-view="palette"
                  no-header
                  no-footer
                  :palette="[
                    '#000000',
                    '#ff0000',
                    '#ff8000',
                    '#ffff00',
                    '#00ff00',
                    '#00ffff',
                    '#0080ff',
                    '#0000ff',
                    '#8000ff',
                    '#ff00ff',
                  ]"
                />
              </q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>
      </template>
    </q-editor>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue';
import { imageSrc } from 'src/utils';
import { Loading } from 'quasar';
import { useInitStore } from 'src/stores/init';

defineOptions({
  name: 'ComponentsEditor'
})

const props = defineProps({
  modelValue: { type: undefined, default: '' },
  label: { type: String, default: '' },
  readonly: { type: Boolean, default: false },
})

const emits = defineEmits(['update:modelValue'])

const $initStore = useInitStore()
const baseURL = process.env.baseURL
const currentValue = ref('') as any
const editorTextColor = ref('') as any
const editorBackgroundColor = ref('') as any
const editorTokenRef = ref(null) as any
const editorRef = ref(null) as any

onMounted(() => {
  currentValue.value = props.modelValue
})


// 编辑文本颜色
const editorEditTextColorFunc = (cmd: string, color: string) => {
  editorTokenRef.value.hide();
  editorRef.value.runCmd(cmd, color);
  editorRef.value.focus();
};

// 富文本编辑器上传图片
const editorUploadedEventFunc = (info: any) => {
  const fileURL = JSON.parse(info.xhr.response).data[0];
  currentValue.value += '<img src="' + imageSrc(fileURL) + '" alt="" />';
  editorTokenRef.value.hide();
  Loading.hide();
  emits('update:modelValue', currentValue.value);
};


watch(() => props.modelValue, (val: any) => {
  currentValue.value = val
})
</script>
