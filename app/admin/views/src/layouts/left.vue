<template>
  <div>
    <q-item
      clickable
      v-ripple
      @click="$router.push('/')"
      :active="$route.fullPath == '/'"
      class="q-mt-xs"
      :active-class="'bg-light-blue-1'"
      onselectstart="return false"
    >
      <q-item-section avatar style="min-width: 0">
        <div class="text-light-blue">
          <q-icon name="dashboard" size="1.5rem"></q-icon>
        </div>
      </q-item-section>
      <q-item-section>
        <div class="text-body2 text-black">控制台</div>
      </q-item-section>
    </q-item>
    <ComponentsMenu :data="$initStore.menuList"></ComponentsMenu>
    <div style="height: 100px"></div>
  </div>
</template>

<script setup lang="ts">
import ComponentsMenu from 'src/components/menu.vue'
import {useInitStore} from 'stores/init';

defineOptions({
  name: 'LayoutsLeft'
})

const $initStore = useInitStore()

</script>

<style scoped>

</style>
