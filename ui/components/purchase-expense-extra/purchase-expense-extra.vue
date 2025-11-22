<template>
    <view class="main-container">
        <view>
            <uni-forms :modelValue="expense">
                <uni-forms-item label="费用名称" name="name">
                    <uni-easyinput v-model="expense.name" placeholder="请输入名称" />
                </uni-forms-item>
                <uni-forms-item label="费用描述" name="desc">
                    <uni-easyinput v-model="expense.desc" placeholder="请输入描述" />
                </uni-forms-item>
                <uni-forms-item label="费用金额" name="expense.amount">
                    <uni-easyinput v-model.number="expense.amount" placeholder="请输入金额" />
                </uni-forms-item>
            </uni-forms>
        </view>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import { onLoad } from '@dcloudio/uni-app';
    import { PurchaseExpense, PurchaseBatch } from '@/xapi/model/purchase';

    const props = defineProps({
        batch: {
            type: Object as () => PurchaseBatch,
            default: new PurchaseBatch(),
        },
        expense: {
            type: Object as () => PurchaseExpense,
            default: new PurchaseExpense(),
        }
    });
</script>

<style scoped>
    .main-container {
        background-color: #f8f8f8;
    }
</style>