<template>
    <view>
        <company-commodity v-model:info="info" v-model:mode="mode" ></company-commodity>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import { examples } from './model';
    import { PurchaseBatch } from '@/xapi/model/purchase';
    
    const info = ref<PurchaseBatch>(new PurchaseBatch());
    const mode = ref('insert');
    
</script>

<style>
    .scroll-Y {
        height: 300px;
    }

    .bottom-list-button {
        display: flex;
    }
</style>