import { ref, computed } from 'vue'
import { getCurrentMonthRange } from '@/utils/formatters'

/**
 * 筛选管理 Composable
 */
export function useFilters(records) {
  const filterKeyName = ref('')
  const dateRange = ref([])

  // 初始化日期范围为当前月份
  function initDateRange() {
    dateRange.value = getCurrentMonthRange()
  }

  // 筛选后的记录
  const filteredRecords = computed(() => {
    let result = [...records.value]

    // 按按键名筛选
    if (filterKeyName.value) {
      result = result.filter(r => r.key_name === filterKeyName.value)
    }

    // 按日期范围筛选
    if (dateRange.value && dateRange.value.length === 2) {
      const [start, end] = dateRange.value
      const startDate = new Date(start)
      startDate.setHours(0, 0, 0, 0)
      const endDate = new Date(end)
      endDate.setHours(23, 59, 59, 999)

      result = result.filter(r => {
        const recordDate = new Date(r.created_at)
        return recordDate >= startDate && recordDate <= endDate
      })
    }

    return result
  })

  // 获取唯一的按键列表
  const uniqueKeys = computed(() => {
    const keys = new Set(records.value.map(r => r.key_name))
    return Array.from(keys).sort()
  })

  // 重置筛选
  function resetFilters() {
    filterKeyName.value = ''
    initDateRange()
  }

  return {
    filterKeyName,
    dateRange,
    filteredRecords,
    uniqueKeys,
    initDateRange,
    resetFilters
  }
}
