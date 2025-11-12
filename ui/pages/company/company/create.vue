<template>
  <view class="form-container">
    <uni-forms :modelValue="form">
      <uni-forms-item label="名称" name="name">
        <uni-easyinput v-model="form.name" placeholder="请输入名称" />
      </uni-forms-item>
      <uni-forms-item label="描述" name="desc">
        <uni-easyinput v-model="form.desc" placeholder="请输入描述" />
      </uni-forms-item>
      <uni-forms-item label="地址" name="address">
        <uni-easyinput v-model="form.address" placeholder="请输入地址" />
      </uni-forms-item>
      <uni-forms-item label="联系方式" name="contacts">
        <view class="contact-item-container">
          <button size="mini" type="primary" :disabled="disabledAddContact" @click="onAddContact">+</button>
        </view>
        <uni-list>
          <view class="contact-item-container" v-for="(item,index) in form.contacts" >
            <uni-list-item :key="index" :title="item.name" :note="item.phone" />
            <button class="contact-item-chg" type="success" size="mini" @click="onChgContact(index)">编辑</button>
            <button class="contact-item-del" type="warn" size="mini" @click="onDelContact(index)">删除</button>
          </view>
        </uni-list>
      </uni-forms-item>
    </uni-forms>

    <view class="actions">
      <button type="primary" @click="submit">保存</button>
      <button type="default" @click="cancel">取消</button>
    </view>
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

  const popupTip = ref(null);
  const tip = ref({
    type: 'success',
    message: '成功'
  })
  const popupContact = ref(null);
  const newContact = ref({
    name: '',
    phone: '',
    desc: '',
  });
  const form = ref({
    name: '',
    desc: '',
    spec: '',
    size: '',
    price: 0,
    contacts: [
      {
        name: '测试1',
        phone: '18877779999',
        desc: '嗯嗯1',
      },
      {
        name: '测试2',
        phone: '18877779999',
        desc: '嗯嗯2',
      }
    ],
  });
  const disabledAddContact = ref(false);
  let posUpdateContact = -1;
  
  function onAddContact() {
    disabledAddContact.value = true;
    newContact.value.name = '';
    newContact.value.phone = '';
    newContact.value.desc = '';
    posUpdateContact = -1;
    popupContact.value.open();
  }
  function onDelContact(pos:any){
    form.value.contacts.splice(pos,1)
    disabledAddContact.value = false;
  }
  
  function onChgContact(pos:any){
    posUpdateContact = pos;
    newContact.value.name = form.value.contacts[posUpdateContact].name;
    newContact.value.phone = form.value.contacts[posUpdateContact].phone;
    newContact.value.desc = form.value.contacts[posUpdateContact].desc;
    popupContact.value.open();
  }
  
  function onAddContactConfirm() {
    popupContact.value.close();
    if (posUpdateContact < 0) {
      form.value.contacts.push({...newContact.value});
    } else {
       form.value.contacts[posUpdateContact].name = newContact.value.name;
       form.value.contacts[posUpdateContact].phone = newContact.value.phone;
       form.value.contacts[posUpdateContact].desc = newContact.value.desc;
    }
    if (form.value.contacts.length < 3) {
      disabledAddContact.value = false;
    }
  }
  
  function onAddContactCancel() {
    popupContact.value.close();
    disabledAddContact.value = false;
  }
  
  function submit() {
    uni.request({
      url: `${BaseURL}/api/v1/company/company`,
      method: 'POST',
      data: form.value,
      success: (res) => {
        console.log("success", res)
        if (res.statusCode == 200) {
          tip.value.type = 'success'
          tip.value.message = '成功'
          setTimeout(
            () => {
              uni.navigateBack();
              uni.$emit('refreshCompanyCompany')
            },
            500);
        } else {
          tip.value.type = 'error'
          tip.value.message = res.data.message
        }
        popupTip.value.open();
      },
      fail: (err) => {
        tip.value.type = 'error'
        tip.value.message = err.data
        popupTip.value.open();
        console.log("fail", err)
      }
    })
  }

  function cancel() {
    uni.navigateBack();
  }
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