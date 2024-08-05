<script setup lang="ts">
import {onMounted, ref} from 'vue';
import {api} from 'src/boot/axios'
import {useRouter} from 'vue-router';
import InputLayoutsComponent from 'src/components/inputLayouts.vue';
import {LocalStorage} from 'quasar';
import {NotifyPositive} from 'src/utils/notify';

defineOptions({
  name: 'IndexSetting'
})

const storageGroupIdKey = '_storageGroupIdKey'
const currentRows = ref([]) as any
const currentConf = ref({}) as any
const groupTabs = ref([]) as any
const groupTab = ref('')
const $router = useRouter()

onMounted(() => {
  api.post(<string>$router.currentRoute.value.meta.views, {}, {showLoading: false} as any).then((conf: any) => {
    currentConf.value = conf
    groupTabs.value = conf.GroupOptions
    if (groupTabs.value.length > 0) {
      groupTab.value = LocalStorage.getItem(storageGroupIdKey) ? Number(LocalStorage.getItem(storageGroupIdKey)) : conf.GroupOptions[0].value

      // 请求配置列表
      switchGroupFunc(Number(groupTab.value))
    }
  })
})

const switchGroupFunc = (groupId: number) => {
  LocalStorage.set(storageGroupIdKey, groupId)
  currentRows.value = []
  api.post(currentConf.value.IndexURL, {groupId: groupId}).then((rows: any) => {
    currentRows.value = rows.items
  })
}

const submitFunc = () => {
  api.post(currentConf.value.UpdateURL, {items: currentRows.value}).then(() => {
    NotifyPositive('保存成功')
  })
}

</script>

<template>
  <div class="q-ma-md">
    <q-tabs
      v-model="groupTab"
      align="left" narrow-indicator
      class="bg-white text-primary"
    >
      <q-tab :name="group.value" :label="group.label" @click="switchGroupFunc(group.value)"
             v-for="(group, groupIndex) in groupTabs" :key="groupIndex" />
    </q-tabs>

    <q-card flat>
      <q-list>
        <q-item v-for="(row, rowIndex) in currentRows" :key="rowIndex">
          <q-item-section avatar top style="min-width: 140px">
            <div class="text-caption text-grey">管理ID:{{row.adminId}}</div>
            <div class="text-body1 text-weight-bold">{{row.name}}</div>
          </q-item-section>
          <q-item-section>
            <div>
              <InputLayoutsComponent
                v-model="row.valueJson"
                :label="row.name"
                :type="row.type"
                :options="row.dataJson"></InputLayoutsComponent>
            </div>
          </q-item-section>
        </q-item>
      </q-list>
    </q-card>
    <div style="height: 200px"></div>

    <div class="fixed-bottom-right q-mr-xl q-mb-xl">
      <div class="row justify-end">
        <q-btn label="保存" size="lg" rounded class="bg-red text-white" @click.stop="submitFunc"></q-btn>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
