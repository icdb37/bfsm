<template>
    <view>
        <button type="primary" @click="onPrint">测试</button>
        <scroll-view scroll-y="true" class="scroll-Y">
            <uni-list>
                <uni-list-item v-for="(item,index) in list.items" :key="index" :title="item.getTitle()"
                    :note="item.getNote()" :showSwitch="true" :switchChecked="item.checked"
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
    import { ref } from 'vue';
    import { onLoad } from '@dcloudio/uni-app';

    const $emit = defineEmits(['select-submit', 'select-cancel']);
    const props = defineProps({
        listItems: {
            type: Array as () => SelectGetter[],
            default: []
        },
        checkedIndexs: {
            type: Array as () => number[],
            default: []
        }
    });

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
    function onPrint() {
        console.log("checkedItems", checkedItems.value.join(","));
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
        list.value.reset(props.listItems,props.checkedIndexs);
    })
    
    interface SelectGetter {
      getTitle : () => string,
      getNote : () => string
    };
    
    class SelectItem {
      data : SelectGetter;
      checked : boolean;
    
      constructor(data : SelectGetter) {
        this.data = data;
      }
    
      getTitle() : string {
        return this.data.getTitle();
      }
      getNote() : string {
        return this.data.getNote();
      }
    };
    
    class SelectList {
      items : SelectItem[];
    
      reset(datas : SelectGetter[], choose:number[]) {
        this.items = [];
        for (let i = 0; i < datas.length; i++) {
          this.items.push(new SelectItem(datas[i]));
        }
        for (let i = 0; i < choose.length; i++){
          this.items[choose[i]].checked = true;
        }
      }
    };
</script>

<style>
    .scroll-Y {
        max-height: 300px;
    }

    .bottom-list-button {
        display: flex;
    }
</style>