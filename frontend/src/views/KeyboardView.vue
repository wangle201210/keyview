<template>
  <div class="keyboard-view">
    <!-- 控制栏 -->
    <ControlBar
      :total-records="totalRecordCount"
      :today-keystrokes="todayKeyCount"
      @refresh="handleRefresh"
    />

    <!-- 筛选器 -->
    <FilterBar
      v-model:key-name="filterKeyName"
      v-model:date-range="dateRange"
      :unique-keys="uniqueKeys"
      @focus="handleDateFocus"
      @blur="handleDateBlur"
    />

    <!-- 虚拟键盘 -->
    <VirtualKeyboard :key-stats="keyStats" />

    <!-- 数据表格 -->
    <RecordTable :records="paginatedRecords" />

    <!-- 分页 -->
    <Pagination
      :current-page="currentPage"
      :total-pages="totalPages"
      @prev="handlePrevPage"
      @next="handleNextPage"
    />
  </div>
</template>

<script setup>
import { onMounted, onUnmounted, watch } from 'vue'
import ControlBar from '@/components/ControlBar.vue'
import FilterBar from '@/components/FilterBar.vue'
import VirtualKeyboard from '@/components/VirtualKeyboard.vue'
import RecordTable from '@/components/RecordTable.vue'
import Pagination from '@/components/Pagination.vue'
import { useKeyboardData } from '@/composables/useKeyboardData'
import { usePagination } from '@/composables/usePagination'
import { useFilters } from '@/composables/useFilters'

// 使用 composables
const {
  records,
  totalRecordCount,
  todayKeyCount,
  keyStats,
  loadRecords,
  loadStats,
  loadKeyStats,
  refreshAll
} = useKeyboardData()

const { currentPage, nextPage, prevPage, resetPage, getPaginatedData } = usePagination(50)

const {
  filterKeyName,
  dateRange,
  filteredRecords,
  uniqueKeys,
  initDateRange
} = useFilters(records)

// 计算属性
const totalPages = computed(() => {
  return Math.ceil(filteredRecords.value.length / 50) || 1
})

const paginatedRecords = computed(() => {
  return getPaginatedData(filteredRecords.value)
})

// 事件处理
async function handleRefresh() {
  await refreshAll()
  await loadKeyStatsWithFilter()
}

function handlePrevPage() {
  prevPage()
}

function handleNextPage() {
  nextPage(totalPages.value)
}

// 根据日期范围加载键盘统计
async function loadKeyStatsWithFilter() {
  let startDate = ''
  let endDate = ''

  if (dateRange.value && dateRange.value.length === 2) {
    const [start, end] = dateRange.value
    startDate = new Date(start).toISOString()
    endDate = new Date(end).toISOString()
  }

  await loadKeyStats(startDate, endDate)
}

// 监听筛选条件变化
watch([filterKeyName, dateRange], () => {
  resetPage()
  loadKeyStatsWithFilter()
})

// 自动刷新定时器
let refreshTimer = null
let isDateFocused = false

onMounted(async () => {
  initDateRange()
  await refreshAll()
  await loadKeyStatsWithFilter()

  // 每5秒自动刷新数据
  refreshTimer = setInterval(async () => {
    // 当用户正在选择日期时，跳过自动刷新
    if (!isDateFocused) {
      await refreshAll()
      await loadKeyStatsWithFilter()
    }
  }, 5000)
})

// 处理日期选择器聚焦事件
function handleDateFocus() {
  isDateFocused = true
}

// 处理日期选择器失焦事件
function handleDateBlur() {
  isDateFocused = false
  // 失焦后立即刷新一次数据
  refreshAll()
  loadKeyStatsWithFilter()
}

onUnmounted(() => {
  if (refreshTimer) {
    clearInterval(refreshTimer)
  }
})
</script>

<script>
import { computed } from 'vue'
export default {
  name: 'KeyboardView'
}
</script>
