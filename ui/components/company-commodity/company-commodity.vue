<template>
    <view class="main-container">
        <view>
            <uni-popup ref="popupTip" type="message">
                <uni-popup-message :type="tip.type" :message="tip.message" :duration="2000" />
            </uni-popup>
        </view>
        <view>
            <uni-forms :modelValue="company">
                <uni-forms-item label="采购企业" name="company">
                    <uni-data-select placeholder="选择企业" v-model="checkedCompanyID" :localdata="selectCompanyList"
                        @close="onSelectCommodity"></uni-data-select>
                </uni-forms-item>
                <uni-forms-item label="采购描述" name="desc">
                    <uni-easyinput v-model="company.desc" placeholder="请输入描述" />
                </uni-forms-item>
                <uni-forms-item label="额外费用" name="amount_extra">
                    <uni-easyinput v-model.number="company.amount_extra" placeholder="请输入额外费用" />
                </uni-forms-item>
                <uni-forms-item label="计算费用" name="amount_total">
                    <button type="primary" size="mini" @click="onClickSum">{{company.amount_total}}</button>
                </uni-forms-item>
            </uni-forms>
        </view>
        <view>
            <scroll-view scroll-y="true" class="commodity-scroll">
                <uni-list>
                    <view class="commodity-container" v-for="(item,index) in selectCommodityList" :key="index">
                        <uni-list-item :title="item.name" :note="item.spec + '-' + item.size" :showSwitch="true"
                            :switchChecked="checkedCommodities[index]" @switchChange="onSwitchChange($event,index)"></uni-list-item>
                        <uni-easyinput class="commodity-input" :clearable="false" maxlength="5" type="number" v-model.number="item.price"
                            placeholder="价格" />
                        <uni-easyinput class="commodity-input" :clearable="false" maxlength="5" type="number" v-model.number="item.count"
                            placeholder="数量" />
                    </view>
                </uni-list>
            </scroll-view>
        </view>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import { onLoad } from '@dcloudio/uni-app';
    import { Commodity, PurchaseCompany, PurchaseGoods,PurchaseBatch } from '@/xapi/model/purchase';
    import { BaseURL } from '@/xapi/xapi';


    let purchaseURL = `${BaseURL}/api/v1/purchase`; // 页面加装时初始化
    let companyURL = `${BaseURL}/api/v1/company`; // 页面加装时初始化
    const modeCreate = "create";
    const modeUpdate = "update";
    const props = defineProps({
        batch: {
            type: Object as  () => PurchaseBatch,
            default: new PurchaseBatch(),
        },
        mode:{
            type: String,
            default: modeCreate
        },
        company: {
            type: Object as () => PurchaseCompany,
            default: new PurchaseCompany(),
        }
    });

    const selectCompanyList = ref<SelectItem[]>([]);
    const selectCommodityList = ref<PurchaseGoods[]>([]);
    const checkedCompanyID = ref('');
    const checkedCommodities = ref<boolean[]>([]);
    
    function onSwitchChange(e : { value : boolean }, index : number) {
        checkedCommodities.value[index] = e.value;
        props.company.goods = [];
        props.company.amount_total = props.company.amount_extra;
        for (let i = 0; i < selectCommodityList.value.length; i++) {
            if (checkedCommodities.value[i]) {
                props.company.goods.push(selectCommodityList.value[i]);
                props.company.amount_total += selectCommodityList.value[i].price * selectCommodityList.value[i].count;
            }
        }
        console.log(JSON.stringify(checkedCommodities.value));
        console.log(JSON.stringify(props.company));
    }
    // onSelectCommodity 查询企业商品
    function onSelectCommodity(){
        console.log("======", JSON.stringify(props.batch))
        props.company.company.company_id = checkedCompanyID.value;
        for (let i = 0; i < selectCompanyList.value.length; i++) {
            if (checkedCompanyID.value == selectCompanyList.value[i].value) {
                props.company.company.company_name = selectCompanyList.value[i].text;
                break;
            }
        }
        let requestURL = companyURL + "/" + checkedCompanyID.value + "/commodity/search";
        console.log("requestURL", requestURL);
        uni.request({
            url: requestURL,
            method: 'POST',
            data: { size: 99999 },
            success: (res) => {
                console.log("success", res)
                if (res.statusCode == 200) {
                    checkedCommodities.value = [];
                    selectCommodityList.value.goods = [];
                    if (res.data.total == 0) {
                        return;
                    }
                    let existsCompany:PurchaseCompany = null;
                    for (let i = 0;i < props.batch.companies.length; i++) {
                        if (checkedCompanyID.value == props.batch.companies[i].company.company_id) {
                            existsCompany = props.batch.companies[i];
                            break;
                        }
                    }
                    for (let i = 0; i < res.data.data.length; i++) {
                        selectCommodityList.value.push(new PurchaseGoods(res.data.data[i]));
                        let checkedGoods = false;
                        if (existsCompany != null){
                            for (let j = 0; j < existsCompany.goods.length; j++) {
                                if (existsCompany.goods[j].hash == res.data.data[i].hash) {
                                    checkedGoods = true;
                                    break;
                                }
                            }   
                        }
                        checkedCommodities.value.push(checkedGoods);
                    }
                    console.log("get select-commodity result", props.company.goods.length)
                } else {
                    tip.value.type = 'error'
                    tip.value.message = res.data.message
                    popupTip.value.open();
                }
            }
        })
    }
    // onSelectCommodity 查询下拉企业
    function onSelectCompany(){
        uni.request({
            url: companyURL + "/select-all",
            method: 'GET',
            success: (res) => {
                console.log("success", res)
                if (res.statusCode == 200) {
                    for (let i = 0; i < res.data.length; i++) {
                        let disableCompany = false;
                        for (let j = 0; j < props.batch.companies.length; j++) {
                            if (props.batch.companies[j].company.company_id == res.data[i].id) {
                                disableCompany = true;
                                break;
                            }
                        }
                        selectCompanyList.value.push({
                            value: res.data[i].id,
                            text: res.data[i].name,
                            disable: disableCompany && props.mode == modeCreate,
                        });
                    }
                    console.log("get select-company result", selectCompanyList.value.length);
                } else {
                    tip.value.type = 'error'
                    tip.value.message = res.data.message
                    popupTip.value.open();
                }
            }
        })
    }
    function onClickSum(){
        props.company.amount_total = props.company.amount_extra;
        for (let i = 0; i < props.company.goods.length; i++) {
            props.company.amount_total += props.company.goods[i].price * props.company.goods[i].count;
        }
    }
    onLoad(() => {
        onSelectCompany();
    })
    interface SelectItem{
        value: string,
        text: string,
        disable: boolean,
    }
    
    const popupTip = ref(null);
    const tip = ref({
        type: 'success',
        message: '成功'
    });
</script>

<style scoped>
    .main-container{
        background-color: #f8f8f8;
    }
    .commodity-scroll {
        max-height: 300px;
    }

    .bottom-list-button {
        display: flex;
    }

    .commodity-container {
        display: flex;
        align-items: center;
    }

    .commodity-input {
        width: 20%;
    }
    
    .company-actions {
        display: flex;
        justify-content: flex-end;
        gap: 12rpx;
        margin-top: 16rpx;
    }
    
</style>