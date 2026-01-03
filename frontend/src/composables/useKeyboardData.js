import { ref, computed } from 'vue'
import { api } from '@/api'

/**
 * 键盘数据管理 Composable
 */
export function useKeyboardData() {
  const records = ref([])
  const totalRecordCount = ref(0)
  const todayKeyCount = ref(0)
  const keyStats = ref([])

  // 加载记录
  async function loadRecords(offset = 0, limit = 1000) {
    try {
      const data = await api.getRecords(offset, limit)
      records.value = data
    } catch (error) {
      console.error('Failed to load records:', error)
      records.value = []
    }
  }

  // 加载统计数据
  async function loadStats() {
    try {
      const [totalCount, todayCount] = await Promise.all([
        api.getTotalCount(),
        api.getTodayKeystrokes()
      ])
      totalRecordCount.value = totalCount
      todayKeyCount.value = todayCount
    } catch (error) {
      console.error('Failed to load stats:', error)
    }
  }

  // 加载键盘统计（支持日期范围）
  async function loadKeyStats(startDate = '', endDate = '') {
    try {
      const stats = await api.getKeyStats(startDate, endDate)
      keyStats.value = stats
    } catch (error) {
      console.error('Failed to load key stats:', error)
      keyStats.value = []
    }
  }

  // 刷新所有数据
  async function refreshAll() {
    await Promise.all([
      loadRecords(),
      loadStats()
    ])
  }

  return {
    records,
    totalRecordCount,
    todayKeyCount,
    keyStats,
    loadRecords,
    loadStats,
    loadKeyStats,
    refreshAll
  }
}
