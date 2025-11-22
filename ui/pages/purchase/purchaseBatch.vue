<!-- 采购批次列表 -->
<template>
    <view>
        <view class="search-create">
            <uni-search-bar class="search-create-search" @confirm="onSearch" cancelButton=false :focus="true"
                v-model="searchValue" />
            <button class="search-create-create" type="primary" size="mini" @click="onCraete">新建</button>
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
                        <uni-td align="center">{{item.name}}</uni-td>
                        <uni-td align="center">{{item.desc}}</uni-td>
                        <uni-td align="center">
                            <uni-tooltip v-for="(item,index) in item.simple_goods" :key="index"
                                :content="'' + item.count" placement="top">
                                <uni-tag :text="item.commodity_title" type="primary" />
                            </uni-tooltip>
                        </uni-td>
                        <uni-td align="center">{{status[item.status]}}</uni-td>
                        <uni-td align="center">{{item.amount_total}}</uni-td>
                        <uni-td>
                            <view class="select-group">
                                <infra-select :actionname="action.curd.label" :localdata="action.curd.datas[index]"
                                    @choose="onChooseCurd"></infra-select>
                                <infra-select :actionname="action.status.label" :localdata="action.status.datas[index]"
                                    @choose="onChooseStatus"></infra-select>
                            </view>
                        </uni-td>
                    </uni-tr>
                </uni-table>

                <uni-popup ref="delpopup" type="dialog">
                    <uni-popup-dialog title="删除企业" :content="delItem.name" :duration="2000" :before-close="true"
                        @close="onDeleteCancel" @confirm="onDeleteConfirm"></uni-popup-dialog>
                </uni-popup>
            </scroll-view>
            <view class="page-container">
                <uni-pagination class="page-container-total" show-icon :page-size="searchRequest.size"
                    :current="searchRequest.index+1" :total="searchResponse.total" @change="onChangePage" />
                <uni-data-select class="page-container-page" :clear="false" v-model="searchRequest.size"
                    :localdata="pageList" @change="onChangePage({'current': 1})" />
            </view>
        </view>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import { onLoad, onUnload } from '@dcloudio/uni-app';
    import { BaseURL } from '@/xapi/xapi';

    let baseURL = `${BaseURL}/api/v1/purchase/batch`; // 页面加装时初始化

    const searchValue = ref('');
    const delpopup = ref(null);
    const delItem = ref({
        id: '',
        name: '',
    });

    function onCraete() {
        uni.navigateTo({
            url: '/pages/purchase/batch/create',
            success: (res) => {
                console.log("success", res)
            }
        })
    }
    function onDeleteConfirm() {
        console.log("onDeleteConfirm", delItem.value.id, delItem.value.name)
        uni.request({
            url: baseURL + "/" + delItem.value.id,
            method: 'DELETE',
            success: (res) => {
                console.log("success", res)
                onSearch()
            },
        })
        delpopup.value.close()
    }
    function onDeleteCancel() {
        console.log("onDeleteCancel", delItem)
        delpopup.value.close()
    }

    function onChangePage(e : any) {
        console.log("onChangePage", e)
        searchRequest.value.index = e.current - 1
        onSearch({ "value": searchValue.value })
    }

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
                searchResponse.value.total = res.data.total;
                searchResponse.value.data = res.data.data;
                action.value.status.datas = [];
                action.value.curd.datas = [];
                for (let i = 0; i < res.data.data.length; i++) {
                    let id = res.data.data[i].id;
                    let status = res.data.data[i].status;
                    action.value.curd.datas.push(initMoreCurd(id));
                    action.value.status.datas.push(initMoreStatus(id, status));
                }
            },
        })
    }
    function onChooseCurd(item : { text : string, value : string }) {
        console.log(item);
        switch (item.text) {
            case crudDetail:
                gotoDetail(item.value)
                break;
            case crudUpdate:
                gotoUpdate(item.value)
                break;
            case crudDelete:
                gotoDelete(item.value)
                break;
        }
    }
    function gotoDetail(id : string) {
        console.log("onDetail", id)
        uni.navigateTo({
            url: `/pages/purchase/batch/detail?id=${id}`,
            success: (res) => {
                console.log("success", res)
            }
        })
    }
    function gotoUpdate(id : string) {
        console.log("onUpdate", id)
        uni.navigateTo({
            url: `/pages/purchase/batch/update?id=${id}`,
            success: (res) => {
                console.log("success", res)
            }
        })
    }
    function gotoDelete(id : string) {
        console.log("onTipDelete", id)
        delItem.value.id = id
        for (let i = 0; i < searchResponse.value.data.length; i++) {
            if (searchResponse.value.data[i].id == id) {
                delItem.value.name = searchResponse.value.data[i].name;
                break;
            }
        }
        delpopup.value.open()
    }
    function onChooseStatus(item : { text : string, value : string }) {
        let req = { desc: "", status: 0 };
        console.log(item);
        switch (item.text) {
            case statusCheck:
                req.status = statusCodeCheck;
                break;
            case statusComplete:
                req.status = statusCodeComplete;
                break;
            case statusCancel:
                req.status = statusCodeCancel;
                break;
        }
        uni.request({
            url: baseURL + "/" + item.value + "/status",
            method: 'PATCH',
            data: req,
            success: (res) => {
                console.log("success", res)
                uni.showToast({
                    title: "操作成功",
                    icon: "success",
                    duration: 2000,
                })
                onSearch()
            },
            fail(res) {
                uni.showToast({
                    title: res.data,
                    icon: "success",
                    duration: 2000,
                })
            }
        })
    }
    function onClickContact(e : any) {
        console.log("onClickContact", e.name, e.phone)
    }
    onLoad(() => {
        baseURL = `${BaseURL}/api/v1/purchase/batch`;
        searchRequest.value.index = 0
        searchRequest.value.size = pageList[0].value
        uni.$on('purchaseBatchRefresh', onSearch); //注册全局事件（创建/修改/删除）之后刷新列表
        onSearch()
    })
    onUnload(() => {
        uni.$off('purchaseBatchRefresh', onSearch) //注销全局事件
    })
    const fields = [
        {
            name: "名称",
            align: "center",
            width: 100,
        },
        {
            name: "描述",
            align: "center",
            width: 100,
        },
        {
            name: "商品",
            align: "center",
            width: 200,
        },
        {
            name: "状态",
            align: "center",
            width: 100,
        },
        {
            name: "金额(分)",
            align: "center",
            width: 100,
        },
        {
            name: "操作",
            align: "center",
            width: 210,
        }
    ]
    const status = ["未定义", "已提交", "已审核", "已完成", "已取消", "已关闭"];
    const searchRequest = ref({
        index: 0,
        size: 10,
        sorts: ['-updated_at'],
        query: {
            name: "",
            desc: "",
            address: "",
        },
    });
    interface MoreItem {
        value : string,
        text : string,
        disabled : boolean,
    }
    const crudDetail = "详情";
    const crudUpdate = "修改";
    const crudDelete = "删除";
    const statusCheck = "审核";
    const statusComplete = "完成";
    const statusCancel = "取消";
    const statusCodeSubimt = 1;
    const statusCodeCheck = 2;
    const statusCodeComplete = 3;
    const statusCodeCancel = 4;
    function initMoreCurd(id : string) : MoreItem[] {
        return [
            {
                value: id,
                text: crudDetail
            }, {
                value: id,
                text: crudUpdate
            }, {
                value: id,
                text: crudDelete
            }]
    }
    function initMoreStatus(id : string, status : number) : MoreItem[] {
        return [
            {
                value: id,
                text: statusCheck,
                disabled: status > statusCodeSubimt,
            }, {
                value: id,
                text: statusComplete,
                disabled: status > statusCodeCheck,
            }, {
                value: id,
                text: statusCancel,
                disabled: (status != statusCodeSubimt && status != statusCodeCheck),
            }]
    }
    const searchResponse = ref({
        total: 0,
        page: 0,
        size: 10,
        data: [],
    });
    const action = ref({
        status: {
            label: "状态",
            datas: []
        },
        curd: {
            label: "操作",
            datas: []
        },
    })
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
    .search-create {
        display: flex;
        align-items: center;
    }

    .search-create-search {
        margin-left: 2px;
    }

    .search-create-create {
        margin-right: 3px;
        height: 80%;
    }

    .scroll-view-item_H {
        display: inline-block;
        width: 100%;
        height: 300rpx;
        line-height: 300rpx;
        text-align: center;
        font-size: 36rpx;
    }

    .page-container {
        display: flex;
    }

    .page-container-page {
        width: 20%;
    }

    .select-group {
        display: flex;
    }

    .page-container-total {
        width: 80%;
    }

    .table-head {
        background-color: lightblue;
    }
</style>