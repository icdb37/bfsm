<template>
	<view>
		<Demo1 />
	</view>
	<view>
		<button type="primary" @click="onPrint">测试</button>
		<scroll-view scroll-y="true" class="scroll-Y">
			<uni-list class="scroll-Y">
				<uni-list-item v-for="(item,index) in listItems" :key="index" 
					:title="item.getTitle()"
					:note="item.getNote()"
					:showSwitch="true" :switchChecked="false"
					@switchChange="onSwitchChange($event,index)"></uni-list-item>
			</uni-list>
		</scroll-view>
		<view class="bottom-list-button ">
		<button type="default" @click="onCancel">取消</button>
		<button type="primary" @click="onSubmit">提交</button>
		</view>
	</view>
</template>

<script lang="ts" setup>
	import Demo1 from './demo1';
	import { ref } from 'vue';
	import {examples,ListSelecter} from './model';
	import {onLoad} from '@dcloudio/uni-app';

	const checkedItems = ref<number[]>([]);
	const listItems = ref<ListSelecter[]>([]);

	function onSwitchChange(e : { value : boolean }, checkedIndex : number) {
		console.log(e, checkedIndex);
		let posExists = -1;
		for (let i = 0; i < checkedItems.value.length; i++) {
			if (checkedIndex == checkedItems.value[i]) {
				posExists = i;
				break;
			}
		}
		if (e.value && posExists == -1) {
			checkedItems.value.push(checkedIndex)
		} else if (!e.value && posExists != -1) {
			checkedItems.value.splice(posExists, 1)
		}
	}
	function onPrint() {
		console.log("checkedItems", checkedItems.value.join(","));
	}
	function onSubmit() {
		console.log("checkedItems", checkedItems.value.join(","));
	}
	function onCancel(){
		
	}
onLoad(() => {
	listItems.value = examples;
})
</script>

<style>
	.scroll-Y {
		height: 300px;
	}

	.bottom-list-button {
		display: flex;
	}
</style>