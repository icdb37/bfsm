<template>
    <view>

        <Demo2 :list-items="listItems" :checked-indexs="[0,3]" @select-submit="onSelectSumibt"
            @select-cancel="onSelectCancel" />
    </view>
    <view>
        <select-list :list-items="listItems" :checked-indexs="[0,3]" @select-submit="onSelectSumibt"
            @select-cancel="onSelectCancel" />
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import Demo1 from './demo1';
    import Demo2 from './demo2';
    import { examples } from './model';
    import { Commodity } from '@/xapi/xapi';
    

    const listItems = ref<ListSelecter[]>(examples);
    const listCommodities = ref<Commodity[]>([
        {
            
        }
    ])
    function onSelectSumibt(indexs : number[]) {
        console.log('select-submit', "indexs", indexs);
    }

    function onSelectCancel() {
        console.log('select-cancel');
    }
</script>

<style>
    .scroll-Y {
        height: 300px;
    }

    .bottom-list-button {
        display: flex;
    }
</style>