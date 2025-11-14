<template>
  <view class="form-container">
    <uni-forms :modelValue="form">
      <uni-forms-item label="名称" name="name">
        <uni-easyinput v-model="form.name" disabled placeholder="请输入名称" />
      </uni-forms-item>
      <uni-forms-item label="描述" name="desc">
        <uni-easyinput v-model="form.desc" disabled placeholder="请输入描述" />
      </uni-forms-item>
      <uni-forms-item label="地址" name="address">
        <uni-easyinput v-model="form.address" disabled placeholder="请输入地址" />
      </uni-forms-item>
      <uni-forms-item label="联系方式" name="contacts">
        <uni-list>
          <view class="contact-item-container" v-for="(item,index) in form.contacts" >
            <uni-list-item :key="index" :title="item.name" :note="item.phone" />
          </view>
        </uni-list>
      </uni-forms-item>
    </uni-forms>
    <view>
      <uni-popup ref="popupTip" type="message">
        <uni-popup-message :type="tip.type" :message="tip.message" :duration="2000"></uni-popup-message>
      </uni-popup>
    </view>
    
    <view>
      <uni-popup ref="popupContact" type="dialog">
				<uni-popup-dialog model="input" type="info" cancelText="取消" confirmText="确认" title="新增联系方式"  @confirm="onAddContactConfirm"@close="onAddContactCancel">
            <uni-forms :modelValue="newContact">
              <uni-forms-item label="姓名" name="name">
                <uni-easyinput v-model="newContact.name" placeholder="请输入姓名" />
              </uni-forms-item>
              <uni-forms-item label="电话" name="phone">
                <uni-easyinput v-model="newContact.phone" placeholder="请输入电话" />
              </uni-forms-item>
              <uni-forms-item label="描述" name="desc">
                <uni-easyinput v-model="newContact.desc" placeholder="请输入描述" />
              </uni-forms-item>
            </uni-forms>
        </uni-popup-dialog>
			</uni-popup>
    </view>
  </view>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { BaseURL } from '@/xapi/xapi';
  import { onLoad } from '@dcloudio/uni-app';
  
  let baseURL = ""; // 页面加装时初始化
  let companyID = ""; // 页面加装时初始化
  
  const popupTip = ref(null);
  const tip = ref({
    type: 'success',
    message: '成功'
  })
  const popupContact = ref(null);
  const form = ref({
    name: '',
    desc: '',
    spec: '',
    size: '',
    price: 0,
    contacts: [],
  });
  
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
    companyID = option.id;
    if (companyID == undefined) {
      companyID = "1cfd04fe-d23c-46ab-9b4f-76050813ced5"
    }
    baseURL = `${BaseURL}/api/v1/company`
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

  .contact-item-container {
    display: flex;
    align-items: center;
  }
   .contact-item-chg {
      margin-right: 0px;
    }
  .contact-item-del {
    margin-right: 0px;
    margin-left: 0px;
  }
  .actions {
    display: flex;
    justify-content: flex-end;
    gap: 12rpx;
    margin-top: 16rpx;
  }
</style>