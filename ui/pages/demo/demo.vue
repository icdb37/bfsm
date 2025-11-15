<template>
  <view>
    <!-- “更多”下拉列表，始终显示“更多”，不随选择变化 -->
    <uni-data-select
      class="more-select"
      placeholder="更多"
      :clear="false"
      v-model="moreValue"
      :localdata="moreRange"
      @change="onMoreChange"
    />
  </view>
</template>

<script setup>
  import { ref } from 'vue';

  // 操作函数
  const onDetail = () => {
    console.log('详情');
  };
  const onEdit = () => {
    console.log('修改');
  };
  const onDelete = () => {
    console.log('删除');
  };

  // “更多”下拉值与选项
  const moreValue = ref('');
  const moreRange = [
    { value: 'detail', text: '详情' },
    { value: 'edit', text: '修改' },
    { value: 'delete', text: '删除' },
  ];

  // 选择后执行对应操作，并重置为占位显示“更多”
  const onMoreChange = (e) => {
    const val = (e && typeof e === 'object') ? (e.detail?.value ?? e.value) : e;
    switch (val) {
      case 'detail':
        onDetail();
        break;
      case 'edit':
        onEdit();
        break;
      case 'delete':
        onDelete();
        break;
      default:
        console.log('未知操作', val);
    }
    // 立即清空以保持显示占位符“更多”
    moreValue.value = '';
  };
</script>

<style>
/* 下拉样式可按需调整 */
.more-select {
  margin: 12px;
}
</style>
