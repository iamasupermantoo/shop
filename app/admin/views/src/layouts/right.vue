<script setup lang="ts">
defineOptions({
  name: 'LayoutsRight'
})

const $initStore = useInitStore()

import {useInitStore} from 'stores/init';
</script>

<template>
  <div :style="{height: 'calc(100vh - 60px)'}">
    <iframe :src="$initStore.userInfo.SeatLink" width="100%" height="100%" frameborder="0"></iframe>
  </div>
</template>

<style scoped>

</style>
