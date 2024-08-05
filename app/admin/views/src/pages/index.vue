<template>
  <div class="q-ma-md">
    <div v-for="(rows, rowsIndex) in items" :key="rowsIndex">
      <div class="row items-center">
        <div
          v-for="(row, rowIndex) in rows"
          :key="rowIndex"
          class="col-md col-sm col-xs-12 text-white"
        >
          <q-card class="q-ma-xs" :class="row.color">
            <q-card-actions>
              <div class="row full-width">
                <div class="text-bold col-4 ellipsis">{{ row.name }}</div>
                <div class="text-caption text-right col-8 ellipsis">
                  合计 {{ row.total.toLocaleString() }}
                </div>
              </div>
            </q-card-actions>
            <q-card-section>
              <div class="row">
                <div class="col-8">
                  <div class="text-h4 ellipsis">
                    {{ row.today.toLocaleString() }}
                  </div>
                  <div class="text-caption ellipsis">
                    昨日 {{ row.yesterday.toLocaleString() }}
                  </div>
                </div>
                <div class="col-4 text-right">
                  <q-icon :name="row.icon" size="60px"></q-icon>
                </div>
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </div>

    <div class="q-mt-xl">
      <div
        :id="echartsDomId"
        style="height: 400px; width: 100%"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import * as echarts from 'echarts';
import { onMounted } from 'vue';
import {homeIndexAPI} from 'src/apis';

defineOptions({
  name: 'IndexHome'
})

const echartsDomId = 'echarts'
const items = ref([]) as any
let echartsInfo = {} as any

onMounted(() => {
  homeIndexAPI().then((res: any) => {
    items.value = res.items
    echartsInfo = res.echarts

    chartSetOptions()
  })
})

// 设置数据统计图
const chartSetOptions = () => {
  const chartDom = document.getElementById(echartsDomId) as HTMLElement;
  const myChart = echarts.init(chartDom);
  let option: echarts.EChartsOption;
  const legendList = [];
  for (let i = 0; i < echartsInfo.series.length; i++) {
    legendList.push(echartsInfo.series[i].name);
  }
  option = {
    tooltip: { trigger: 'axis' },
    legend: { data: legendList },
    grid: { left: '0', right: '0', bottom: '0', containLabel: true },
    toolbox: { feature: { saveAsImage: {} } },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: echartsInfo.category,
    },
    yAxis: { type: 'value' },
    series: echartsInfo.series as any,
  };
  myChart.setOption(option);
  window.addEventListener('resize', () => {
    setTimeout(myChart.resize, 300);
  });
};
</script>
