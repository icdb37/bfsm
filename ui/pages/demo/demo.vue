<template>
  <view>
    <uni-data-select v-model="value" :localdata="range" @change="change"></uni-data-select>
    <uni-collapse ref="collapse" v-model="value">
      <uni-collapse-item title="更多">
        <button>详情</button>
        <button>修改</button>
        <button type="warn">删除</button>
      </uni-collapse-item>
    </uni-collapse>
  </view>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  
  const change = (e) => {
    console.log(e);
  }
  
  const value = ref(0);
  const range = [{"value": 0,"text": "篮球"	},{"value": 1,"text": "足球"},{"value": 2,"text": "游泳"}];
</script>

<style>
         
</style>
