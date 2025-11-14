<template>
  <view class="form-container">
    <uni-forms :modelValue="form">
      <uni-forms-item label="标识" name="id">
        <uni-easyinput v-model="form.id" disabled />
      </uni-forms-item>
      <uni-forms-item label="名称" name="name">
        <uni-easyinput v-model="form.name" disabled />
      </uni-forms-item>
      <uni-forms-item label="描述" name="desc">
        <uni-easyinput v-model="form.desc" disabled />
      </uni-forms-item>
      <uni-forms-item label="规格" name="spec">
        <uni-easyinput v-model="form.spec" disabled />
      </uni-forms-item>
      <uni-forms-item label="尺寸" name="size">
        <uni-easyinput v-model="form.size" disabled />
      </uni-forms-item>
      <uni-forms-item label="价格" name="price">
        <uni-easyinput type="number" v-model.number="form.price" disabled />
      </uni-forms-item>
      <uni-forms-item label="创建时间" name="created_at">
        <uni-easyinput v-model="form.created_at" disabled />
      </uni-forms-item>
      <uni-forms-item label="修改时间" name="updated_at">
        <uni-easyinput v-model="form.updated_at" disabled />
      </uni-forms-item>
    </uni-forms>
    
    <view>
      <uni-popup ref="popup" type="message">
        <uni-popup-message :type="tip.type" :message="tip.message" :duration="2000"></uni-popup-message>
      </uni-popup>
    </view>
  
  </view>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { onLoad } from '@dcloudio/uni-app';
  import { BaseURL } from '@/xapi/xapi';

  
  let baseURL = ""; // 页面加装时初始化
  let companyID = ""; // 页面加装时初始化
  
  const popup = ref(null);
  const tip = ref({
    type: 'success',
    message: '成功'
  })
  const form = ref({
    id: '',
    name: '',
    desc: '',
    spec: '',
    size: '',
    price: 0,
    created_at: '',
    updated_at: ''
  })
  
  function getDetail(id : string) {
    uni.request({
      url: `${baseURL}/${id}`,
      method: 'GET',
      success: (res) => {
        console.log("success", res)
        if (res.statusCode != 200) {
          tip.value.type = 'error'
          tip.value.message = res.data.message
          popup.value.open();
        }
        form.value = res.data
      },
      fail: (err) => {
        tip.value.type = 'error'
        tip.value.message = err.data
        popup.value.open();
      }
    })
  }
  
  onLoad((option) => {
    companyID = option.company_id;
    if (companyID == undefined) {
      companyID = "1cfd04fe-d23c-46ab-9b4f-76050813ced5"
    }
    baseURL = `${BaseURL}/api/v1/company/${companyID}/commodity`
    if (option.id) {
      getDetail(option.id)
      return
    }
    tip.value.type = 'error'
    tip.value.message = "商品标识无效"
    popup.value.open();
    setTimeout(
      () => {
        uni.navigateBack();
      },
      500);
  })
</script>

<style>
  .form-container {
    padding: 16rpx;
    width: 600rpx;
  }

  .actions {
    display: flex;
    justify-content: flex-end;
    gap: 12rpx;
    margin-top: 16rpx;
  }
</style>