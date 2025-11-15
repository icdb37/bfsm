<template>
	<view>
		<view class="search-create">
			<uni-search-bar class="search-create-search" @confirm="onSearch" cancelButton=false :focus="true" v-model="searchValue"/>
			<button class="search-create-create" type="primary" size="mini" @click="onCraete">新建</button>
		</view>
		<!-- 企业列表 -->
    <view>
      <scroll-view class="scroll-view_H" scroll-x="true" scroll-left="100">
        	<uni-table border stripe emptyText="暂无更多数据">
        		<uni-tr border stripe >
        			<uni-th class="table-head" v-for="(item, index) in fields" :align="item.align" :key="index">{{item.name}}</uni-th>
        			<uni-th class="table-head" width="210" align="center">操作</uni-th>
        		</uni-tr>
        		<!-- 表格数据行 -->
        		<uni-tr v-for="(item, index) in searchResponse.data" :key="index">
        			<uni-td align="center">{{item.name}}</uni-td>
        			<uni-td align="center">{{item.desc}}</uni-td>
        			<uni-td align="center">
                <uni-tooltip v-for="(item,index) in item.simple_goods" :key="index" :content="item.count" placement="top">
                  <uni-tag :text="item.commodity_title" type="primary"/>
                </uni-tooltip>
              </uni-td>
              <uni-td align="center">{{status[item.status]}}</uni-td>
              <uni-td align="center">{{item.amount_total}}</uni-td>
        			<uni-td>
        				<view class="uni-group">
                  <button class="uni-button" size="mini" type="default" @click="onDetail(item)">详情</button>
        					<button class="uni-button" size="mini" type="primary" @click="onUpdate(item)">修改</button>
        					<button class="uni-button" size="mini" type="warn" @click="onTipDelete(item)">删除</button>
        				</view>
        			</uni-td>
        		</uni-tr>
        	</uni-table>
    
          <uni-popup ref="delpopup" type="dialog">
          	<uni-popup-dialog title="删除企业" :content="delItem.name" :duration="2000" :before-close="true" @close="onDeleteCancel" @confirm="onDeleteConfirm"></uni-popup-dialog>
          </uni-popup>
      </scroll-view>
      <view class="page-container">
        <uni-pagination class="page-container-total" show-icon :page-size="searchRequest.size" :current="searchRequest.index+1" :total="searchResponse.total" @change="onChangePage" />
        <uni-data-select class="page-container-page" :clear="false" v-model="searchRequest.size" :localdata="pageList" @change="onChangePage({'current': 1})" />
      </view>
    </view>
	</view>
</template>

<script lang="ts" setup>
import {ref} from 'vue';
import {onLoad,onUnload} from '@dcloudio/uni-app';
import {BaseURL} from '@/xapi/xapi';

let baseURL = `${BaseURL}/api/v1/purchase/batch`; // 页面加装时初始化
  
const searchValue = ref('');
const delpopup = ref(null);
const delItem = ref({
  id: '',
  name: '',
});

function onCraete(){
  uni.navigateTo({
    url: '/pages/purchase/batch/create',
    success: (res) => {
      console.log("success", res)
    }
  })
}

function onUpdate(item:any){
	console.log("onUpdate", item)
  uni.navigateTo({
    url: `/pages/purchase/batch/update?id=${item.id}`,
    success: (res) => {
      console.log("success", res)
    }
  })
}
function onTipDelete(e:any) {
  console.log("onTipDelete", e)
  delItem.value.id = e.id
  delItem.value.name = e.name
  delpopup.value.open()
}
function onDeleteConfirm(){
	console.log("onDeleteConfirm", delItem.value.id,delItem.value.name)
  uni.request({
    url: baseURL +"/" + delItem.value.id,
    method: 'DELETE',
    success: (res) => {
      console.log("success", res)
      onSearch()
    },
  })
  delpopup.value.close()
}
function onDeleteCancel(){
  console.log("onDeleteCancel", delItem)
  delpopup.value.close()
}
function onDetail(item:any){
	console.log("onDetail", item)
  uni.navigateTo({
    url: `/pages/purchase/batch/detail?id=${item.id}`,
    success: (res) => {
      console.log("success", res)
    }
  })
}

function onChangePage(e:any) {
  console.log("onChangePage", e)
  searchRequest.value.index=e.current-1
	onSearch({"value": searchValue.value})
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
      searchResponse.value.total = res.data.total
      searchResponse.value.data = res.data.data
    },
  })
}
function onClickContact(e:any){
  console.log("onClickContact", e.name,e.phone)
}
onLoad(() => {
  baseURL = `${BaseURL}/api/v1/purchase/batch`;
  searchRequest.value.index = 0
  searchRequest.value.size=pageList[0].value
  uni.$on('purchaseBatchRefresh', onSearch); //注册全局事件（创建/修改/删除）之后刷新列表
	onSearch()
})
onUnload(() => {
  uni.$off('purchaseBatchRefresh', onSearch) //注销全局事件
})
const fields = [
	{
		name: "名称",
		align: "center"
	},
	{
		name: "描述",
		align: "center"
	},
	{
		name: "商品",
		align: "center"
	},
  {
  	name: "状态",
  	align: "center"
  },
	{
		name: "金额(分)",
		align: "center"
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

const searchResponse = ref({
  total: 0,
  page: 0,
  size: 10,
  data: [],
});

const pageList=[
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

<style>
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


.page-container-total {
	width: 80%;
}

.table-head{
  background-color: lightblue;
}
</style>
