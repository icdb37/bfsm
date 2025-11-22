<!-- 采购商品列表 -->
<template>
    <view>
        <view class="search-create">
            <uni-search-bar class="search-create-search" @confirm="onSearch" cancelButton=false :focus="true"
                v-model="searchValue" />
        </view>
        <!-- 企业列表 -->
        <view>
            <scroll-view class="scroll-view_H" scroll-x="true" scroll-left="100">
                <uni-table border stripe emptyText="暂无更多数据">
                    <uni-tr border stripe>
                        <uni-th class="table-head" v-for="(item, index) in fields" :align="item.align"
                            :width="item.width" :key="index">{{item.name}}</uni-th>
                    </uni-tr>
                    <!-- 表格数据行 -->
                    <uni-tr v-for="(item, index) in searchResponse.data" :key="index">
                        <uni-td class="table-cell" align="center">{{item.name}}</uni-td>
                        <uni-td align="center">{{item.spec}}</uni-td>
                        <uni-td align="center">{{item.size}}</uni-td>
                        <uni-td align="center">{{item.price}}</uni-td>
                        <uni-td align="center">{{item.count}}</uni-td>
                        <uni-td align="center">{{item.amount}}</uni-td>
                        <uni-td class="table-cell" align="center">{{item.company_name}}</uni-td>
                        <uni-td class="table-cell" align="center">{{item.purchase_name}}</uni-td>
                        <uni-td align="center">
                            <uni-dateformat format="yyyy-MM-dd hh:mm:ss" :date="item.created_at"/>
                        </uni-td>
                        <uni-td>
                            <view class="uni-group">
                                <button class="uni-button" size="mini" type="default"
                                    @click="onDetail(item)">详情</button>
                            </view>
                        </uni-td>
                    </uni-tr>
                </uni-table>
            </scroll-view>
        </view>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import { onLoad, onUnload } from '@dcloudio/uni-app';
    import { BaseURL, FormatUTC } from '@/xapi/xapi';
    const searchValue = ref('');
    const searchRequest = ref({
        index: 0,
        size: 10,
        sorts: ['created_at'],
        query: {
            name: "",
            desc: "",
            address: "",
        },
    });
    const searchResponse = ref({
        total: 0,
        page: 0,
        size: 10,
        data: [],
    });
    function onSearch() {
        console.log("onSearch", searchRequest.value.index, searchRequest.value.size)
        searchRequest.value.query.name = searchValue.value
        searchRequest.value.query.desc = searchValue.value
        searchRequest.value.query.spec = searchValue.value
        searchRequest.value.query.size = searchValue.value
        uni.request({
            url: baseURL + "/search",
            method: 'POST',
            data: searchRequest.value,
            success: (res) => {
                console.log("success", res)
                searchResponse.value.total = res.data.total
                searchResponse.value.data = res.data.data
            },
        })
    }
    function onDetail(item) {
        uni.navigateTo({
            url: '/pages/purchase/goods/detail?id=' + item.id,
        })
    }
    onLoad(() => {
        baseURL = `${BaseURL}/api/v1/purchase/goods`;
        searchRequest.value.index = 0
        searchRequest.value.size = pageList[0].value
        uni.$on('purchaseGoodsRefresh', onSearch); //注册全局事件（创建/修改/删除）之后刷新列表
        onSearch()
    })
    onUnload(() => {
        uni.$off('purchaseGoodsRefresh', onSearch) //注销全局事件
    })

    let baseURL = `${BaseURL}/api/v1/purchase/goods`; // 页面加装时初始化
    const fields = [
        {
            name: "名称",
            align: "center",
            width: 100,
        },
        {
            name: "规格",
            align: "center",
            width: 100,
        },
        {
            name: "尺寸",
            align: "center",
            width: 80,
        },
        {
            name: "价格",
            align: "center",
            width: 60,
        },
        {
            name: "数量",
            align: "center",
            width: 60,
        },
        {
            name: "金额(分)",
            align: "center",
            width: 80,
        },
        {
            name: "采购企业",
            align: "center",
            width: 100,
        },
        {
            name: "采购批次",
            align: "center",
            width: 100,
        },
        {
            name: "创建时间",
            align: "center",
            width: 150,
        },
        {
            name: "操作",
            align: "center",
            width: 100,
        }
    ];
    const pageList = [
        {
            value: 10,
            text: '10条/页'
        },
        {
            value: 20,
            text: '20条/页'
        },
        {
            value: 50,
            text: '50条/页'
        },
        {
            value: 100,
            text: '100条/页'
        }
    ];
</script>

<style scoped>
    .table-head {
        background-color: lightblue;
    }

    .table-cell {
        overflow: hidden;
        /* 隐藏超出部分 */
        white-space: nowrap;
        /* 禁止文本换行 */
        text-overflow: ellipsis;
        /* 文本溢出时显示省略号 */
        /* 根据实际布局设置一个合适的宽度 */
        /* width: 200px; */
    }
</style>