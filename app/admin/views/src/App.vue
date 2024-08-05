<template>
  <router-view />
</template>

<script setup lang="ts">
import {useInitStore} from 'stores/init';
import {dynamicRouter} from 'src/router';
import {useRouter} from 'vue-router';

defineOptions({
  name: 'App'
});

const $router = useRouter()
const $initStore = useInitStore()
dynamicRouter($router, $initStore.menuList)
</script>
