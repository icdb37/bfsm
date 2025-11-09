<template>
	<view>
		<view class="search-create">
			<uni-search-bar @confirm="onSearch" cancelButton=false :focus="true" v-model="searchValue"/>
			<button type="primary">新建</button>
		</view>
		<!-- 商品列表 -->
		<view>
			<uni-table border stripe emptyText="暂无更多数据">
				<uni-tr border stripe >
					<uni-th  v-for="(item, index) in fields" :align="item.align">{{item.name}}</uni-th>
					<uni-th width="240" align="center">操作</uni-th>
				</uni-tr>
				<!-- 表格数据行 -->
				<uni-tr v-for="(item, index) in datas" :key="index">
					<uni-td align="center">{{item.name}}</uni-td>
					<uni-td align="center">{{item.desc}}</uni-td>
					<uni-td align="center">{{item.spec}}</uni-td>
					<uni-td align="center">{{item.size}}</uni-td>
					<uni-td align="center">{{item.price}}</uni-td>
					<uni-td>
						<view class="uni-group">
							<button class="uni-button" size="mini" type="primary" @click="onUpdate(item)">修改</button>
							<button class="uni-button" size="mini" type="warn" @click="onDelete(item)">删除</button>
							<button class="uni-button" size="mini" type="default" @click="onDetail(item)">x详情</button>
						</view>
					</uni-td><strong></strong>
				</uni-tr>
			</uni-table>
		</view>
	</view>
</template>

<script lang="ts" setup>
import {ref} from 'vue';

const searchValue = ref('');

function onUpdate(item:any){
	console.log("onUpdate", item)
}
function onDelete(item:any){
	console.log("onDelete", item)
}
function onDetail(item:any){
	console.log("onDetail", item)
}
function onSearch() {
	console.log("onSearch", searchValue)
}

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
]
</script>

<style>
.search-create {
	display: flex;
}
</style>
