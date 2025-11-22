<template>
    <view class="form-container">
        <uni-forms :model="batch">
            <uni-forms-item label="采购名称" name="name">
                <uni-easyinput v-model="batch.name" :clearable="false" :disabled="true" />
            </uni-forms-item>
            <uni-forms-item label="采购描述" name="desc">
                <uni-easyinput v-model="batch.desc" :clearable="false" :disabled="true" />
            </uni-forms-item>
            <uni-forms-item label="企业商品" name="companies">
                <uni-list>
                    <view class="append-item-container" v-for="(item,index) in batch.companies">
                        <uni-list-item :key="index" :title="item.company.company_name"
                            :note="'总金额：'+item.amount_total + '  总类型：'+item.goods.length"/>
                    </view>
                </uni-list>
            </uni-forms-item>
            <uni-forms-item label="额外费用" name="extras">
                <uni-list>
                    <view class="append-item-container" v-for="(item,index) in batch.extras">
                        <uni-list-item :key="index" :title="item.name" :note="'金额：'+item.amount" />
                    </view>
                </uni-list>
            </uni-forms-item>
        </uni-forms>
        <view>
            <uni-popup ref="popupTip" type="message">
                <uni-popup-message :type="tip.type" :message="tip.message" :duration="2000"></uni-popup-message>
            </uni-popup>
        </view>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import { BaseURL } from '@/xapi/xapi';
    import { onLoad } from '@dcloudio/uni-app';
    import { PurchaseBatch } from '@/xapi/model/purchase';

    const batch = ref<PurchaseBatch>(new PurchaseBatch());
    let purchaseURL = `${BaseURL}/api/v1/purchase`; // 页面加装时初始化


    function getDetail(id : string) {
        uni.request({
            url: purchaseURL + "/batch/" + id,
            method: 'GET',
            success: (res) => {
                batch.value = { ...res.data };
            }
        })
    }

    function cancel() {
        uni.navigateBack();
    }
    onLoad((options) => {
        let id = "4525b2ac-5b04-4b1c-9ab9-a9a091d7f1b6";//TODO 提示报错
        if (options.id) {
            id = options.id;
        }
        getDetail(id);
    })
</script>

<style scoped>
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

    .company-commodity-container-count {
        width: 10%;
    }
</style>