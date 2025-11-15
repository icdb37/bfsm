<template>
	<view>
		<view class="search-create">
			<uni-search-bar class="search-create-search" @confirm="onSearch" cancelButton=false :focus="true" v-model="searchValue"/>
			<button class="search-create-create" type="primary" @click="onCraete">新建</button>
		</view>
		<!-- 商品列表 -->
    <view>
      <scroll-view class="scroll-view_H" scroll-x="true" scroll-left="100">
        	<uni-table border stripe emptyText="暂无更多数据">
        		<uni-tr border stripe >
        			<uni-th class="table-head" v-for="(item, index) in fields" :align="item.align" :key="index">{{item.name}}</uni-th>
        			<uni-th class="table-head" width="240" align="center">操作</uni-th>
        		</uni-tr>
        		<!-- 表格数据行 -->
        		<uni-tr v-for="(item, index) in searchResponse.data" :key="index">
        			<uni-td align="center">{{item.name}}</uni-td>
        			<uni-td align="center">{{item.desc}}</uni-td>
        			<uni-td align="center">{{item.spec}}</uni-td>
        			<uni-td align="center">{{item.size}}</uni-td>
        			<uni-td align="center">{{item.price}}</uni-td>
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
          	<uni-popup-dialog title="删除商品" :content="delItem.name" :duration="2000" :before-close="true" @close="onDeleteCancel" @confirm="onDeleteConfirm"></uni-popup-dialog>
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
import {BaseURL} from '../../xapi/xapi';


const searchValue = ref('');
const delpopup = ref(null);
const delItem = ref({
  id: '',
  name: '',
});

function onCraete(){
  uni.navigateTo({
    url: '/pages/commodity/commodity/create',
    success: (res) => {
      console.log("success", res)
    }
  })
}

function onUpdate(item:any){
	console.log("onUpdate", item)
  uni.navigateTo({
    url: `/pages/commodity/commodity/update?id=${item.id}`,
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
    url: BaseURL + "/api/v1/commodity/commodity/" + delItem.value.id,
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
    url: `/pages/commodity/commodity/detail?id=${item.id}`,
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
    url: BaseURL + "/api/v1/commodity/commodity/search",
    method: 'POST',
    data: searchRequest.value,
    success: (res) => {
      console.log("success", res)
      searchResponse.value.total = res.data.total
      searchResponse.value.data = res.data.data
    },
  })
}

onLoad(() => {
  searchRequest.value.index = 0
  searchRequest.value.size=pageList[0].value
  uni.$on('refreshCommodityCommodity', onSearch); //注册全局事件（创建/修改/删除）之后刷新列表
	onSearch()
})
onUnload(() => {
  uni.$off('refreshCommodityCommodity', onSearch) //注销全局事件
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
		name: "规格",
		align: "center"
	},
	{
		name: "尺寸",
		align: "center"
	},
	{
		name: "价格",
		align: "center"
	}
]

const datas = [
	{
		id: "a1",
		name: "a1",
		desc: "a1",
		spec: "a1",
		size: "a1",
		price: 11
	},
	{
		id: "a2",
		name: "a2",
		desc: "a2",
		spec: "a2",
		size: "a2",
		price: 12
	}
];

const searchRequest = ref({
  index: 0,
  size: 10,
  sorts: ['-updated_at'],
  query: {
    name: "",
    desc: "",
    spec: "",
    size: ""
  },
});

const searchResponse = ref({
  total: 0,
  page: 0,
  size: 10,
  data: datas,
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
}
.search-create-search {
	margin-left: 2px;
}
.search-create-create {
	margin-right: 2px;
  margin-top: 3px;
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
