<template>
  <div class="filters">
    <el-select
      :model-value="keyName"
      @change="$emit('update:keyName', $event)"
      placeholder="所有按键"
      style="width: 200px"
    >
      <el-option label="所有按键" value="" />
      <el-option
        v-for="key in uniqueKeys"
        :key="key"
        :label="key"
        :value="key"
      />
    </el-select>

    <el-date-picker
      :model-value="dateRange"
      @change="$emit('update:dateRange', $event)"
      type="daterange"
      range-separator="至"
      start-placeholder="开始日期"
      end-placeholder="结束日期"
      style="width: 240px"
    />
  </div>
</template>

<script setup>
defineProps({
  keyName: {
    type: String,
    default: ''
  },
  dateRange: {
    type: Array,
    default: () => []
  },
  uniqueKeys: {
    type: Array,
    default: () => []
  }
})

defineEmits(['update:keyName', 'update:dateRange'])
</script>

<style scoped>
.filters {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
  flex-wrap: wrap;
  align-items: center;
}
</style>
