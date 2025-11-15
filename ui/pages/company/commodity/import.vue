<template>
  <view>
    <view>商品导入</view>
  </view>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { onLoad } from '@dcloudio/uni-app';
  import { BaseURL } from '@/xapi/xapi';
  
  let companyID = '';
  let baseURL = '';
  
  onLoad((option) => {
    companyID = option.company_id;
    if (companyID == undefined) {
      companyID = "1cfd04fe-d23c-46ab-9b4f-76050813ced5"
    }
    baseURL = `${BaseURL}/api/v1/company/${companyID}/commodity`
  })
</script>

<style>
</style>