<template>
    <view class="form-container">
        <uni-forms :model="info">
            <uni-forms-item class="cell-field" label="采购批次" name="purchase_name">
                <view>{{info.purchase_name}}</view>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="企业名称" name="company_name">
                <view>{{info.company_name}}</view>
            </uni-forms-item>
            <uni-forms-item label="商品名称" name="name">
                <view>{{info.name}}</view>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="商品描述" name="desc">
                <view>{{info.desc}}</view>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="商品规格" name="spec">
                <view>{{info.spec}}</view>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="商品尺寸" name="size">
                <view>{{info.size}}</view>
            </uni-forms-item>
            <uni-forms-item class="cell-field"  label="商品价格" name="price">
                <view>{{info.price}}</view>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="商品数量" name="count">
                <view>{{info.count}}</view>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="商品金额" name="amount">
                <view>{{info.amount}}</view>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="存储位置" name="storeage">
                 <view>{{info.storeage}}</view>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="生产日期" name="produced_at">
                <uni-dateformat format="yyyy-MM-dd hh:mm:ss" :date="info.produced_at"></uni-dateformat>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="到期日期" name="expired_at">
                <uni-dateformat format="yyyy-MM-dd hh:mm:ss" :date="info.expired_at"></uni-dateformat>
            </uni-forms-item>
            <uni-forms-item class="cell-field" label="创建日期" name="created_at">
                <uni-dateformat format="yyyy-MM-dd hh:mm:ss" :date="info.created_at"></uni-dateformat>
            </uni-forms-item>
        </uni-forms>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import { BaseURL,FormatUTC } from '@/xapi/xapi';
    import { onLoad } from '@dcloudio/uni-app';
    import { PurchaseGoodsX } from '@/xapi/model/purchase';

    const info = ref<PurchaseGoodsX>(new PurchaseGoodsX());
    let purchaseURL = `${BaseURL}/api/v1/purchase`; // 页面加装时初始化

    function getDetail(id : string) {
        uni.request({
            url: purchaseURL + "/goods/" + id,
            method: 'GET',
            success: (res) => {
                console.log(res)
                info.value.set({ ...res.data });
                console.log(JSON.stringify(info.value))
            }
        })
    }

    function cancel() {
        uni.navigateBack();
    }

    onLoad((options) => {
        let id = "84b302d6-b141-4955-a0bf-a81ae6bccb3c";
        if (options && options.id) {
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
    .cell-field {
        display: flex;
        align-items: center;
    }
    .cell-field-value {
    }
.cell-field {
        display: flex;
        align-items: center;
    }    
</style>