<template>
    <view>
        <infra-select :localdata="localdata" @choose="onChoose"></infra-select>
        <label>====================</label>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';

    function onChoose(item : {value:string, text:string}) {
        console.log(item.value);
    }
    const localdata = ref([
        {
            text: "修改",
            value: "aaa",
            disabled: true,
        }, {
            text: "创建",
            value: "bbb"
        }
    ])
</script>

<style>
    .more-wrapper {
        position: relative;
        display: inline-block;
    }

    .dropdown-btn {
        height: 30px;
        background-color: azure;
        display: flex;
        align-items: center;
    }

    .dropdown {
        position: absolute;
        top: 100%;
        left: 0;
        margin-top: 8px;
        z-index: 1000;
        background-color: #fff;
        box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
        border-radius: 6px;
        max-width: 160px;
        overflow: hidden;
    }

    .page {
        padding: 12px;
    }
</style>