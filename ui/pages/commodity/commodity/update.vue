<template>
  <view class="form-container">
    <uni-forms :modelValue="form">
      <uni-forms-item label="标识" name="id">
        <uni-easyinput v-model="form.id" disabled />
      </uni-forms-item>
      <uni-forms-item label="名称" name="name">
        <uni-easyinput v-model="form.name" placeholder="请输入名称" />
      </uni-forms-item>
      <uni-forms-item label="描述" name="desc">
        <uni-easyinput v-model="form.desc" placeholder="请输入描述" />
      </uni-forms-item>
      <uni-forms-item label="规格" name="spec">
        <uni-easyinput v-model="form.spec" placeholder="请输入规格" />
      </uni-forms-item>
      <uni-forms-item label="尺寸" name="size">
        <uni-easyinput v-model="form.size" placeholder="请输入尺寸" />
      </uni-forms-item>
      <uni-forms-item label="价格" name="price">
        <uni-easyinput type="number" v-model.number="form.price" placeholder="请输入价格" />
      </uni-forms-item>
    </uni-forms>

    <view class="actions">
      <button type="primary" @click="submit">保存</button>
      <button type="default" @click="cancel">取消</button>
    </view>
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
  })

  function submit() {
    uni.request({
      url: `http://localhost:8080/api/v1/commodity/commodity/${form.value.id}`,
      method: 'PUT',
      data: form.value,
      success: (res) => {
        console.log("success", res)
        if (res.statusCode == 200) {
          tip.value.type = 'success'
          tip.value.message = '成功'
          setTimeout(
            () => {
              uni.navigateBack();
              uni.$emit('refreshCommodityCommodity')
            },
            500);
        } else {
          tip.value.type = 'error'
          tip.value.message = res.data.message
        }
        popup.value.open();
      },
      fail: (err) => {
        tip.value.type = 'error'
        tip.value.message = err.data
        popup.value.open();
        console.log("fail", err)
      }
    })
  }

  function cancel() {
    uni.navigateBack();
  }
  function getDetail(id : string) {
    uni.request({
      url: `http://localhost:8080/api/v1/commodity/commodity/${id}`,
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
    if (option.id) {
      getDetail(option.id)
    }
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