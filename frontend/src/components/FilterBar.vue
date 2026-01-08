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
      v-model="localDateRange"
      type="daterange"
      range-separator="至"
      start-placeholder="开始日期"
      end-placeholder="结束日期"
      style="width: 240px"
      clearable
      editable
      @change="handleDateChange"
      @focus="$emit('focus')"
      @blur="$emit('blur')"
    />
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
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

const emit = defineEmits(['update:keyName', 'update:dateRange', 'focus', 'blur'])

// 使用本地 ref 来管理日期选择器的值
const localDateRange = ref([])

// 监听外部 dateRange 变化，同步到本地
watch(() => props.dateRange, (newVal) => {
  if (newVal && newVal.length === 2) {
    localDateRange.value = [...newVal]
  } else {
    localDateRange.value = []
  }
}, { immediate: true })

// 处理日期变化
function handleDateChange(value) {
  emit('update:dateRange', value)
}
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
