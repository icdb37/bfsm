<template>
  <view class="form-container">
    <uni-forms :modelValue="form">
      <uni-forms-item label="采购名称" name="name">
        <uni-easyinput v-model="form.name" placeholder="请输入名称" />
      </uni-forms-item>
      <uni-forms-item label="采购描述" name="desc">
        <uni-easyinput v-model="form.desc" placeholder="请输入描述" />
      </uni-forms-item>
      <uni-forms-item label="企业商品" name="companies">
        <view class="append-item-container">
          <button size="mini" type="primary" :disabled="disabledAddCompany" @click="onAddCompany">+</button>
        </view>
        <uni-list>
          <view class="append-item-container" v-for="(item,index) in form.companies" >
            <uni-list-item :key="index" :title="item.company.company_name" :note="item.amount" />
            <button class="append-item-chg" type="success" size="mini" @click="onChgCompany(index)">编辑</button>
            <button class="append-item-del" type="warn" size="mini" @click="onDelCompany(index)">删除</button>
          </view>
        </uni-list>
      </uni-forms-item>
      <uni-forms-item label="额外费用" name="extras">
        <view class="append-item-container">
          <button size="mini" type="primary" :disabled="disabledAddExpense" @click="onAddExpense">+</button>
        </view>
        <uni-list>
          <view class="append-item-container" v-for="(item,index) in form.extras" >
            <uni-list-item :key="index" :title="item.name" :note="item.amount" />
            <button class="append-item-chg" type="success" size="mini" @click="onChgExpense(index)">编辑</button>
            <button class="append-item-del" type="warn" size="mini" @click="onDelExpense(index)">删除</button>
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
    
    <!-- 商品列表 -->
    <view>
      <uni-popup ref="popupCompnay" type="dialog">
    		<uni-popup-dialog model="input" type="info" cancelText="取消" confirmText="确认" title="新增企业商品"  @confirm="onAddCompnayConfirm"@close="onAddCompnayCancel">
            <uni-forms :modelValue="newCompany">
              <uni-forms-item label="企业" name="company">
                <uni-data-select placeholder="选择企业" v-model="selectCompanyID" :localdata="selectCompanies" @close="onCloseSelectCompany"></uni-data-select>
              </uni-forms-item>
              <uni-forms-item label="描述" name="desc">
                <uni-easyinput v-model="newCompany.desc" placeholder="请输入描述" />
              </uni-forms-item>
              <uni-list class="company-commodity-container">
                <view class="company-commodity-container-item" v-for="(item,index) in selectCommodities">
                  <uni-list-item :key="index" :showSwitch="true" :switchChecked="item.count>0" @switchChange="onChooseCommodity($event,item)" :title="item.name" :note="item.spec+'/' + item.size + ' ' + item.desc"/>
                  <uni-easyinput class="company-commodity-container-count" maxlength="5" type="number" v-model.number="item.price" placeholder="价格" />
                  <uni-easyinput class="company-commodity-container-count" maxlength="5" type="number" v-model.number="item.count" placeholder="数量" />
                </view>
              </uni-list>
            </uni-forms>
        </uni-popup-dialog>
    	</uni-popup>
    </view>
    <!-- 费用列表 -->
    <view>
      <uni-popup ref="popupExpense" type="dialog">
				<uni-popup-dialog model="input" type="info" cancelText="取消" confirmText="确认" title="新增联系方式"  @confirm="onAddExpenseConfirm"@close="onAddExpenseCancel">
            <uni-forms :modelValue="newExpense">
              <uni-forms-item label="费用名称" name="name">
                <uni-easyinput v-model="newExpense.name" placeholder="请输入名称" />
              </uni-forms-item>
              <uni-forms-item label="费用描述" name="desc">
                <uni-easyinput v-model="newExpense.desc" placeholder="请输入描述" />
              </uni-forms-item>
              <uni-forms-item label="金额(分)" name="amount">
                <uni-easyinput maxlength="11" type="number" v-model.number="newExpense.amount" placeholder="请输入描述" />
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
  

  let purchaseURL = `${BaseURL}/api/v1/purchase`; // 页面加装时初始化
  let companyURL = `${BaseURL}/api/v1/company`; // 页面加装时初始化

  const popupTip = ref(null);
  const tip = ref({
    type: 'success',
    message: '成功'
  })
  const popupCompnay = ref(null);
  const newCompany = ref({
    company: '',
    desc: '',
    goods: [],
  });
  const disabledAddCompany = ref(false);
  let posUpdateCompany = -1;
  
  function onAddCompany() {
    disabledAddCompany.value = true;
    newCompany.value.goods = [];
    newCompany.value.company = {};
    newCompany.value.desc = '';
    posUpdateCompany = -1;
    selectCompanyID.value = '';
    popupCompnay.value.open();
  }
  function onChgCompany(pos:any) {
    posUpdateCompany = pos;
    newCompany.value = form.value.companies[posUpdateCompany];
    console.log(JSON.stringify(newCompany.value))
    console.log(JSON.stringify(form.value))
    selectCompanyID.value = newCompany.value.company.company_id;
    for (let i = 0; i < selectCompanies.value.length; i++){
      if (selectCompanies.value[i].value == selectCompanyID.value) {
        selectCompanies.value[i].disable = false;
        break;
      }
    }
    getCommoditySelect();
    popupCompnay.value.open();
  }
  function onDelCompany(pos:any){
    let dels = form.value.companies.splice(pos,1);
    for (let i = 0; i < dels.length; i++) {
      for (let j = 0; j < selectCompanies.value.length; j++) {
        if (selectCompanies.value[j].value == dels[i].company.company_id) {
          selectCompanies.value[j].disable = false;
          break;
        }
      }
    }
    disabledAddCompany.value = false;
  }
  
  function onAddCompnayConfirm() {
    disabledAddCompany.value = false;
    popupCompnay.value.close();
    selectCommodities.value = [];
    if (posUpdateCompany >= 0) {  
      return;
    }
    for (let i = 0; i < newCompany.value.goods.length; i++) {
      if (newCompany.value.goods[i].count <= 0) {
        //TODO 优化
        return;
      }
    }
    for (let i = 0; i < selectCompanies.value.length; i++) {
      if (selectCompanies.value[i].value == selectCompanyID.value) {
        newCompany.value.company = {
          company_id: selectCompanies.value[i].value,
          company_name: selectCompanies.value[i].text,
        }
        selectCompanies.value[i].disable = true;
        break;
      }
    }
    form.value.companies.push({
      desc: newCompany.value.desc,
      company: newCompany.value.company,
      goods: newCompany.value.goods,
    });
    console.log(JSON.stringify(newCompany.value))
    console.log(JSON.stringify(form.value))
  }
  function onAddCompnayCancel() {
    console.log("selectCompanyID", selectCompanyID.value)
    selectCommodities.value = [];
    disabledAddCompany.value = false;
    selectCompanyID.value = '';
    popupCompnay.value.close();
  }
  function onCloseSelectCompany(){
    getCommoditySelect();
  }
  
  
  const popupExpense = ref(null);
  const newExpense = ref({
    name: '',
    desc: '',
    amount: 0,
  });
  const form = ref({
    name: '',
    desc: '',
    companies: [],
    extras: [],
  });
  const disabledAddExpense = ref(false);
  let posUpdateExpense = -1;
  
  function onAddExpense() {
    disabledAddExpense.value = true;
    newExpense.value.name = '';
    newExpense.value.amount = 0;
    newExpense.value.desc = '';
    posUpdateExpense = -1;
    popupExpense.value.open();
  }
  function onDelExpense(pos:any){
    form.value.extras.splice(pos,1)
    disabledAddExpense.value = false;
  }
  
  function onChgExpense(pos:any){
    posUpdateExpense = pos;
    newExpense.value.name = form.value.extras[posUpdateExpense].name;
    newExpense.value.desc = form.value.extras[posUpdateExpense].desc;
    newExpense.value.amount = form.value.extras[posUpdateExpense].amount;
    popupExpense.value.open();
  }
  
  function onAddExpenseConfirm() {
    popupExpense.value.close();
    if (posUpdateExpense < 0) {
      newExpense.value.amount = newExpense.value.amount;
      form.value.extras.push({...newExpense.value});
    } else {
       form.value.extras[posUpdateExpense].name = newExpense.value.name;
       form.value.extras[posUpdateExpense].amount = newExpense.value.amount;
       form.value.extras[posUpdateExpense].desc = newExpense.value.desc;
    }
    if (form.value.extras.length < 5) {
      disabledAddExpense.value = false;
    }
  }
  
  function onAddExpenseCancel() {
    popupExpense.value.close();
    disabledAddExpense.value = false;
  }
  
  function submit() {
    uni.request({
      url: purchaseURL + "/batch",
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
              uni.$emit('purchaseBatchRefresh')
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
  onLoad((options) => {
    console.log(options);
    purchaseURL = `${BaseURL}/api/v1/purchase`;
    companyURL = `${BaseURL}/api/v1/company`;
    getCompanySelect();
  })
  const selectCompanies = ref([]);
  const selectCompanyID = ref("");
  
  function getCompanySelect() {
    uni.request({
      url: companyURL+"/select-all",
      method: 'GET',
      success: (res) => {
        console.log("success", res)
        if (res.statusCode == 200) {
          for (let i = 0; i < res.data.length; i++) {
            selectCompanies.value.push({"value": res.data[i].id, "text": res.data[i].name, disable: false});
          }
        } else {
          tip.value.type = 'error'
          tip.value.message = res.data.message
          popupTip.value.open();
        }
      },
      fail: (err) => {
        tip.value.type = 'error'
        tip.value.message = err.data
        popupTip.value.open();
        console.log("fail", err)
      }
    })
  }

  const selectCommodities = ref([]);
  function getCommoditySelect() {
    let requestURL = companyURL+"/" + selectCompanyID.value + "/commodity/search";
    console.log("requestURL", requestURL);
    uni.request({
      url: requestURL,
      method: 'POST',
      data: {size: 9999},
      success: (res) => {
        console.log("success", res)
        if (res.statusCode == 200) {
          selectCommodities.value = res.data.data;
          for (let i = 0; i < selectCommodities.value.length; i++) {
            for (let j = 0; j < newCompany.value.goods.length; j++) {
              if (newCompany.value.goods[j].id == selectCommodities.value[i].id) {
                selectCommodities.value[i] = newCompany.value.goods[j];
                break;
              }
            }
          }
        } else {
          tip.value.type = 'error'
          tip.value.message = res.data.message
          popupTip.value.open();
        }
      },
      fail: (err) => {
        tip.value.type = 'error'
        tip.value.message = err.data
        popupTip.value.open();
        console.log("fail", err)
      }
    })
  }
  
  function onChooseCommodity(e:{value: boolean},item:any) {
    console.log(e,item)
    for (let i = 0; i < newCompany.value.goods.length; i++) {
      if (newCompany.value.goods[i].id == item.id) {
        if (e.value) {
          return;
        }
        newCompany.value.goods.splice(i,1)
        return;
      }
    }
    if (e.value) {
      newCompany.value.goods.push(item);
    }
  }
</script>

<style>
  .form-container {
    padding: 16rpx;
    width: 600rpx;
  }

  .append-item-container {
    display: flex;
    align-items: center;
  }
   .append-item-chg {
      margin-right: 0px;
    }
  .append-item-del {
    margin-right: 0px;
    margin-left: 0px;
  }
  .actions {
    display: flex;
    justify-content: flex-end;
    gap: 12rpx;
    margin-top: 16rpx;
  }
  .company-commodity-container {
    /* 限制高度 */
  }
  .company-commodity-container-item {
    display: flex;
    align-items: center;
  }
  .company-commodity-container-count{
    width: 10%;
  }
</style>