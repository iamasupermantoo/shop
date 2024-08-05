<template>
  <div class="menu">
    <q-list v-for="(item, index) in data" :key="index" :class="listClass">
      <q-expansion-item
        :icon="item.data.icon"
        :label="item.name"
        :default-opened="expansionOpenedFunc(item.children)"
        header-class="text-body2 menu-expansion-icon"
        expand-icon-class="text-grey"
        expand-icon="arrow_drop_down"
        onselectstart="return false"
        v-if="
          item.hasOwnProperty('children') &&
          item.children !== null &&
          item.children.length > 0
        "
      >
        <template v-slot:header>
          <q-item-section avatar style="min-width: 0">
            <div class="text-light-blue">
              <q-icon :name="item.data.icon" size="1.5rem" />
            </div>
          </q-item-section>
          <q-item-section>
            <div class="text-black text-body2">
              {{ item.name }}
            </div>
          </q-item-section>
        </template>

        <ComponentsMenu :data="item.children" :inset-level="insetLevel + 0.5" :list-class="listClass" :active-class="activeClass"></ComponentsMenu>
      </q-expansion-item>
      <q-item
        clickable
        v-ripple
        v-else
        :inset-level="insetLevel"
        @click="menuLinkRoute(item)"
        :active="item.route === $route.fullPath"
        :active-class="activeClass"
        onselectstart="return false"
      >
        <q-item-section
          avatar
          v-if="item.data.icon !== ''"
          style="min-width: 0"
        >
          <div class="text-light-blue">
            <q-icon :name="item.data.icon" size="1.5rem"></q-icon>
          </div>
        </q-item-section>
        <q-item-section>
          <div class="text-body2 text-black">
            {{ item.name }}
          </div>
        </q-item-section>
      </q-item>
    </q-list>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import ComponentsMenu from 'src/components/menu.vue'

defineOptions({
  name: 'ComponentsMenu'
})

defineProps({
  data: { type: Array as any, default: Map },
  insetLevel: { type: Number, default: 0 },
  listClass: { type: String, default: '' },
  activeClass: { type: String, default: 'bg-light-blue-1' },
})

const $router = useRouter()

//  点击跳转路由
const menuLinkRoute = (item: any) => {
  $router.push(item.route);
};

//  判断折叠是否打开
const expansionOpenedFunc = (children: any) => {
  if (children === null) {
    return false;
  }
  let findMenu = children.find((item: any) => {
    return item.route === $router.currentRoute.value.path;
  });
  return findMenu !== undefined;
};
</script>

<style scoped>
.menu .q-item__section--avatar {
  min-width: 0 !important;
}
</style>
