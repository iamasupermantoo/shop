<template>
  <div class="q-pa-md" v-if="!initPage">
    <!-- 头部数据查询 -->
    <div class="row q-gutter-sm items-center" @keyup.enter="enterSearchFunc">
      <div
        v-for="(input, inputIndex) in currentConfig.search.inputList"
        :key="inputIndex"
      >
        <InputLayoutsComponent
          v-model="currentConfig.search.params[input.field]"
          :label="input.label"
          :type="input.type"
          :options="input.data"
        ></InputLayoutsComponent>
      </div>
      <div v-if="currentConfig.search.inputList.length > 0" style="margin-top: 0">
        <q-btn
          icon="search"
          color="primary"
          @click="requestTableFunc({ pagination: currentConfig.pagination })"
        ></q-btn>
      </div>
    </div>

    <!-- 数据表格 -->
    <q-card flat bordered class="q-mt-md">
      <!-- 工具栏 -->
      <q-card-section>
        <div class="row">
          <div
            v-for="(tool, toolIndex) in currentConfig.table.tools"
            :key="toolIndex"
            class="q-mr-sm q-mb-sm"
          >
            <q-btn
              :label="tool.label"
              :color="tool.color"
              :size="tool.size"
              v-if="$initStore.hasRoute(tool.config.url)"
              @click="dialogOpenFunc(tool.config, null)"
            ></q-btn>
          </div>
        </div>
      </q-card-section>
      <q-card-section>
        <q-table
          flat
          :rows="currentRows"
          :columns="currentConfig.table.columns"
          :row-key="currentConfig.table.key"
          selection="multiple"
          @request="requestTableFunc"
          v-model:selected="checkboxList"
          v-model:pagination="currentConfig.pagination"
        >
          <template v-slot:top>
            <q-space />
            <q-btn
              color="secondary"
              icon-right="archive"
              label="下载CSV文件"
              no-caps
              @click="exportFileFunc"
            />
          </template>
          <template v-slot:body-cell="scope">
            <q-td>
              <ColumnsContent
                v-if="scope.col.type != 'options'"
                :url="currentConfig.table.updateUrl"
                :primary="currentConfig.table.key"
                :model-value="scopeFind(scope.row, scope.col.field)"
                @done="dialogDoneFunc"
                :current-scope="scope"
              ></ColumnsContent>

              <!-- 额外的操作 -->
              <div v-if="scope.col.type == 'options'">
                <div class="row">
                  <div
                    v-for="(option, optionIndex) in currentConfig.table.options"
                    :key="optionIndex"
                    class="q-mr-sm q-mb-sm"
                  >
                    <q-btn
                      :label="option.label"
                      :color="option.color"
                      :size="option.size"
                      v-if="
                        $initStore.hasRoute(option.config.url) &&
                        (option.eval == '' ||
                          executeEval(option.eval, {
                            scope: scope,
                          }))
                      "
                      @click="dialogOpenFunc(option.config, scope)"
                    ></q-btn>
                  </div>
                </div>
              </div>
            </q-td>
          </template>
        </q-table>
      </q-card-section>
    </q-card>
  </div>
  <DialogComponent ref="domDialogRef" @done="dialogDoneFunc"></DialogComponent>
</template>

<script setup lang="ts">
import {onMounted, ref} from 'vue';
import { api } from 'src/boot/axios';
import { exportCSVFile } from 'src/utils/export';
import InputLayoutsComponent from 'src/components/inputLayouts.vue';
import DialogComponent from 'src/components/dialog.vue';
import ColumnsContent from 'src/components/columnsContent.vue';
import { useRouter } from 'vue-router';
import { useInitStore } from 'src/stores/init';
import { inputOperateSettingFun, executeEval } from 'src/utils/input';
import {LocalStorage} from 'quasar';

defineOptions({
  name: 'IndexViews'
})

const $initStore = useInitStore()
const initPage = ref(true) as any
const domDialogRef = ref(null) as any
const $router = useRouter()

const currentRows = ref([]) as any
const checkboxList = ref([]) as any
const currentConfig = ref({
  url: '',
  pagination: {sortBy: 'id', descending: true, page: 0, rowsPerPage: 50, rowsNumber: 0} as any,
  search: {params: {} as any, inputList: [] as any},
  table: {key: 'id', updateUrl: '', tools: [] as any, columns: [] as any, options: [] as any}
})

onMounted(() => {
  const fullPathList = $router.currentRoute.value.fullPath.split('?')
  api.post(<string>$router.currentRoute.value.meta.views + (fullPathList.length > 1 ? '?' + fullPathList[1] : ''), {}, {showLoading: false} as any).then((conf: any) => {
    initPage.value = false
    currentConfig.value = conf

    if (currentConfig.value.table.options.length > 0) {
      currentConfig.value.table.columns.push({label: '操作栏', field: '_options', type: 'options', name: '_options', align: 'left'})
    }

    // 如果有缓存, 那么获取缓存中的请求参数
    const storageParams = LocalStorage.getItem(currentConfig.value.url)
    if (storageParams != null) {
      currentConfig.value.search.params = Object.assign(storageParams, currentConfig.value.search.params)
    }

    requestTableFunc({ pagination: currentConfig.value.pagination })
  })
})

//  请求表格数据
const requestTableFunc = (props: { pagination: any }) => {
  currentRows.value = [];
  currentConfig.value.pagination = props.pagination;
  currentConfig.value.search.params['pagination'] = currentConfig.value.pagination;

  // 设置请求参数缓存
  LocalStorage.set(currentConfig.value.url, currentConfig.value.search.params)
  api.post(currentConfig.value.url, currentConfig.value.search.params).then((res: any) => {
    currentRows.value = res.items;
    currentConfig.value.pagination.rowsNumber = res.count;
  });
};

//  弹窗操作
const dialogOpenFunc = (config: any, rawData: any) => {
  const dialogConfig = JSON.parse(JSON.stringify(config));

  switch (dialogConfig.params.operate) {
    case 'checkbox':
      //  批量操作｜操作方法
      if (checkboxList.value.length == 0) {
        return;
      }

      //  生成弹窗名称内容
      const content = [];
      const scanList = [];
      for (let i = 0; i < checkboxList.value.length; i++) {
        content.push(checkboxList.value[i][config.params['name']]);
        scanList.push(checkboxList.value[i][config.params['scan']]);
      }

      dialogConfig.content = content.toString();
      dialogConfig.params = {};
      dialogConfig.params[config.params['field']] = scanList;
      break
    case 'setting':
      //  设置操作
      const inputOperateInfo = inputOperateSettingFun(
        config.params['operate'],
        currentConfig.value,
        dialogConfig,
        rawData
      );

      dialogConfig.params = inputOperateInfo.params;
      dialogConfig.inputList = inputOperateInfo.inputList;
      break;
    default:
      // 如果有原数据, 那么使用原数据 - 先使用原数据, 后面使用默认数据
      if (rawData == null) {
        break
      }

      // 提交ID 用于判断是那条数据
      dialogConfig.params[currentConfig.value.table.key] = rawData.row[currentConfig.value.table.key]
      dialogConfig.params = mergeData(dialogConfig.params, rawData.row, dialogConfig.inputList, rawData.row)
  }

  domDialogRef.value.setDialogConfig(dialogConfig);
  domDialogRef.value.dialogOpenFunc();
};

//  完成操作
const dialogDoneFunc = () => {
  checkboxList.vlaue = [];
  requestTableFunc({ pagination: currentConfig.value.pagination });
};

// 下载csv文件
const exportFileFunc = () => {
  exportCSVFile(currentConfig.value.table.columns, currentRows);
};

// 合并对象
const mergeData = (leftData: any, rightData: any, inputList: any, sourceData: any): any => {
  if (Array.isArray(leftData) && Array.isArray(rightData)) {
    leftData = rightData
    return leftData
    // // 如果左边和右边都是数组，则遍历合并
    // return leftData.map((item, index) => {
    //   if (index < rightData.length) {
    //     return mergeData(item, rightData[index], inputList, sourceData);
    //   }
    //   return item;
    // });
  } else if (typeof leftData === 'object' && leftData !== null && typeof rightData === 'object' && rightData !== null) {
    // 如果左边和右边都是对象，则合并键值
    let result = { ...leftData };
    for (let key in leftData) {
      const valAlias = inputListFindAlias(inputList, key)
      if (valAlias != '') {
        result[key] = scopeFind(sourceData, valAlias)
      }else {
        if (rightData.hasOwnProperty(key)) {
          result[key] = mergeData(leftData[key], rightData[key], inputList, sourceData);
        }
      }
    }
    return result;
  } else {
    // 其他情况直接返回左边的数据
    leftData = rightData
    return leftData;
  }
}

const inputListFindAlias = (inputList: any, field: string): string => {
  let value = ''
  for (let i = 0; i < inputList.length; i++) {
    if (inputList[i].field == field) {
      return inputList[i].alias
    }

    if (inputList[i].data) {
      for (let j = 0; j < inputList[i].data.length; j++) {
        if (inputList[i].data[j]) {
          value = inputListFindAlias(inputList[i].data[j], field)
        }
      }
    }
  }
  return value
}

// 获取scope多层值
const scopeFind = (row: any, field: string): any => {
  const fieldList = field.split('.')
  let currentVal = row[fieldList[0]]
  for (let i = 1; i < fieldList.length; i++) {
    if (currentVal && currentVal.hasOwnProperty(fieldList[i])) {
      currentVal = currentVal[fieldList[i]]
    }
  }

  return currentVal
}

// enterSearchFunc 回车搜索功能
const enterSearchFunc = () => {
  requestTableFunc({ pagination: currentConfig.value.pagination })
}

</script>
