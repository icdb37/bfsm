<template>
    <view class="page">
        <view>
            <view class="more-wrapper">
                <button type="default" class="dropdown-btn" :plain="true" @click="onClick">更多</button>
                <view class="dropdown" v-show="showMore">
                    <uni-list>
                        <uni-list-item class="dropdown-item" :title="item.text" clickable :key="index" @click="onClickSelect(item)" v-for="(item,index) in localdata"></uni-list-item>
                    </uni-list>
                </view>
            </view>
        </view>
        <label>====================</label>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    
    const showMore = ref(false);
    const value = ref('更多');
    
    function onClick() {
        console.log('onClick');
        showMore.value = !showMore.value;
    }
    function onClickSelect(item:selItem) {
        console.log("selected ", item)
        showMore.value = false;
    }
    interface selItem{
        text: string,
        value: string,
    }
    const localdata:selItem[] = [
        {
            text: '修改',
            value: '0'
        },
        {
            text: '删除',
            value: '1'
        }
    ]
</script>

<style>
    .more-wrapper {
        position: relative;
        display: inline-block;
    }
    .dropdown-btn{
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