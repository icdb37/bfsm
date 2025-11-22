<template>
    <view>
        <scroll-view scroll-y="true" class="commodity-scroll">
            <uni-list>
                <view class="commodity-container" v-for="(item,index) in list.items" :key="index">
                    <uni-list-item :title="item.getTitle()" :note="item.getNote()" :showSwitch="true"
                        :switchChecked="item.checked" @switchChange="onSwitchChange($event,index)"></uni-list-item>
                    <uni-easyinput class="commodity-input" maxlength="5" type="number" v-model.number="item.data.price"
                        placeholder="价格" />
                    <uni-easyinput class="commodity-input" maxlength="5" type="number" v-model.number="item.data.count"
                        placeholder="数量" />
                </view>
            </uni-list>
        </scroll-view>
        <view class="bottom-list-button ">
            <button type="default" @click="onCancel">取消</button>
            <button type="primary" @click="onSubmit">提交</button>
        </view>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import { onLoad } from '@dcloudio/uni-app';
    import { Commodity } from '@/xapi/xapi';

    const $emit = defineEmits(['select-submit', 'select-cancel']);
    const props = defineProps({
        listItems: {
            type: Array as () => Commodity[],
            default: []
        },
        checkedIndexs: {
            type: Array as () => number[],
            default: []
        }
    });

    // 先定义使用到的类，避免“Cannot access 'SelectList' before initialization”错误
    class SelectItem {
        data: Commodity;
        checked: boolean;

        constructor(data: Commodity) {
            this.data = data;
            this.checked = false;
        }

        getTitle(): string {
            return this.data.name;
        }
        getNote(): string {
            return this.data.sepc + "-" + this.data.size;
        }
    }

    class SelectList {
        items: SelectItem[];

        constructor() {
            this.items = [];
        }

        reset(datas: Commodity[], choose: number[]) {
            this.items = [];
            for (let i = 0; i < datas.length; i++) {
                this.items.push(new SelectItem(datas[i]));
            }
            for (let i = 0; i < choose.length; i++) {
                const idx = choose[i];
                if (idx >= 0 && idx < this.items.length) {
                    this.items[idx].checked = true;
                }
            }
        }
    }

    const checkedItems = ref<number[]>([]);
    const list = ref<SelectList>(new SelectList());

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
    function onSubmit() {
        console.log("checkedItems", checkedItems.value.join(","));
        $emit('select-submit', [...checkedItems.value]);
    }
    function onCancel() {
        $emit('select-cancel');
    }
    onLoad(() => {
        checkedItems.value.push(...props.checkedIndexs)
        list.value.reset(props.listItems, props.checkedIndexs);
        console.log("listItems", props.listItems);
    })


    // 上面已定义 SelectItem 与 SelectList
</script>

<style scoped>
    .commodity-scroll {
        max-height: 300px;
    }

    .bottom-list-button {
        display: flex;
    }

    .commodity-container {
        display: flex;
        align-items: center;
    }

    .commodity-input {
        width: 10%;
    }
</style>