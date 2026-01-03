<template>
  <div class="app">
    <header class="header">
      <h1>KeyView - 键盘使用历史记录</h1>
      <div class="github-link" @click="openGitHub">
        <svg height="24" viewBox="0 0 16 16" version="1.1" width="24" aria-hidden="true">
          <path fill="white" d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"></path>
        </svg>
      </div>
    </header>

    <main class="main">
      <div class="controls">
        <div class="status-indicator">
          <span class="status-dot recording"></span>
          <span class="status-text">正在监听键盘</span>
        </div>

        <button @click="refreshData" class="refresh-btn">
          刷新数据
        </button>

        <div class="stats">
          <span>总记录数: {{ totalRecords }}</span>
          <span>今日按键: {{ todayKeystrokes }}</span>
        </div>
      </div>

      <div class="filters">
        <el-select v-model="filterKeyName" placeholder="所有按键" @change="applyFilters" style="width: 200px">
          <el-option label="所有按键" value="" />
          <el-option
            v-for="key in uniqueKeys"
            :key="key"
            :label="key"
            :value="key"
          />
        </el-select>

        <el-date-picker
          v-model="dateRange"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          @change="applyFilters"
          style="width: 240px"
        />
      </div>

      <!-- 虚拟键盘 -->
      <VirtualKeyboard :keyStats="keyStats" />

      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>时间</th>
              <th>按键</th>
              <th>修饰键</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="record in paginatedRecords" :key="record.id">
              <td>{{ formatTime(record.created_at) }}</td>
              <td>{{ record.key_name }}</td>
              <td>{{ formatModifiers(record.modifier_flags) }}</td>
            </tr>
            <tr v-if="paginatedRecords.length === 0">
              <td colspan="3" class="no-data">暂无数据</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="pagination">
        <button
          @click="prevPage"
          :disabled="currentPage === 1"
          class="page-btn"
        >
          上一页
        </button>

        <span class="page-info">
          第 {{ currentPage }} / {{ totalPages }} 页
        </span>

        <button
          @click="nextPage"
          :disabled="currentPage === totalPages"
          class="page-btn"
        >
          下一页
        </button>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { Browser } from '@wailsio/runtime'
import { api } from './services/wails.js'
import VirtualKeyboard from './components/VirtualKeyboard.vue'

const isRecording = ref(true) // 默认为监听状态
const records = ref([])
const currentPage = ref(1)
const pageSize = 50
const filterKeyName = ref('')
const dateRange = ref([]) // 日期范围 [开始日期, 结束日期]
const totalRecordCount = ref(0)
const todayKeyCount = ref(0)
const keyStats = ref([])

// 打开 GitHub 链接
async function openGitHub() {
  await Browser.OpenURL('https://github.com/wangle201210/keyview')
}

// 自动刷新定时器
let refreshTimer = null

// 初始化日期范围为当前月份
function initDateRange() {
  const now = new Date()
  const firstDay = new Date(now.getFullYear(), now.getMonth(), 1)
  const lastDay = new Date(now.getFullYear(), now.getMonth() + 1, 0)
  dateRange.value = [firstDay, lastDay]
}

const filteredRecords = computed(() => {
  let result = [...records.value]

  if (filterKeyName.value) {
    result = result.filter(r => r.key_name === filterKeyName.value)
  }

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

const totalPages = computed(() => {
  return Math.ceil(filteredRecords.value.length / pageSize) || 1
})

const paginatedRecords = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  const end = start + pageSize
  return filteredRecords.value.slice(start, end)
})

const totalRecords = computed(() => totalRecordCount.value)

const todayKeystrokes = computed(() => todayKeyCount.value)

const uniqueKeys = computed(() => {
  const keys = new Set(records.value.map(r => r.key_name))
  return Array.from(keys).sort()
})

function formatTime(timestamp) {
  const date = new Date(timestamp)
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: false
  })
}

function formatModifiers(flags) {
  if (!flags || flags === 0) return '-'

  const modifiers = []
  if (flags & 0x20000) modifiers.push('⇧')
  if (flags & 0x40000) modifiers.push('⌃')
  if (flags & 0x80000) modifiers.push('⌥')
  if (flags & 0x100000) modifiers.push('⌘')
  if (flags & 0x10000) modifiers.push('Caps')

  return modifiers.length > 0 ? modifiers.join(' + ') : '-'
}

async function refreshData() {
  await loadRecords()
  await loadStats()
}

function applyFilters() {
  currentPage.value = 1
  loadKeyStats() // 重新加载键盘统计
}

function prevPage() {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

function nextPage() {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

async function loadRecords() {
  console.log('loadRecords')
  try {
    const offset = 0
    const limit = 1000 // 获取更多记录用于前端筛选
    const data = await api.getRecords(offset, limit)
    console.log('loadRecords, data',data)

    records.value = data
  } catch (error) {
    console.error('Failed to load records:', error)
    records.value = []
  }
}

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

async function loadKeyStats() {
  try {
    let startDate = ''
    let endDate = ''

    if (dateRange.value && dateRange.value.length === 2) {
      const [start, end] = dateRange.value
      startDate = new Date(start).toISOString()
      endDate = new Date(end).toISOString()
    }

    const stats = await api.getKeyStats(startDate, endDate)
    keyStats.value = stats
  } catch (error) {
    console.error('Failed to load key stats:', error)
    keyStats.value = []
  }
}

onMounted(async () => {
  // 初始化日期范围为当前月份
  initDateRange()

  await loadRecords()
  await loadStats()
  await loadKeyStats()

  // 每5秒自动刷新数据
  refreshTimer = setInterval(async () => {
    await loadRecords()
    await loadStats()
    await loadKeyStats()
  }, 5000)
})

onUnmounted(() => {
  if (refreshTimer) {
    clearInterval(refreshTimer)
  }
})
</script>

<style scoped>
.app {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.header {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  padding: 1.5rem 2rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header h1 {
  color: white;
  font-size: 1.8rem;
  font-weight: 600;
  margin: 0;
}

.github-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  transition: all 0.3s;
  cursor: pointer;
}

.github-link:hover {
  background: rgba(255, 255, 255, 0.2);
  transform: scale(1.1);
}

.github-link svg {
  width: 24px;
  height: 24px;
}

.main {
  padding: 2rem;
  max-width: 1400px;
  margin: 0 auto;
}

.controls {
  display: flex;
  gap: 1rem;
  align-items: center;
  margin-bottom: 1.5rem;
  flex-wrap: wrap;
}

.status-indicator {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: rgba(76, 175, 80, 0.2);
  padding: 0.6rem 1.2rem;
  border-radius: 20px;
  backdrop-filter: blur(10px);
}

.status-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #ccc;
}

.status-dot.recording {
  background: #4caf50;
  animation: pulse 2s infinite;
}

.status-text {
  color: white;
  font-weight: 600;
  font-size: 0.95rem;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}

.refresh-btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  background: #2196f3;
  color: white;
  transition: all 0.3s;
}

.refresh-btn:hover {
  background: #0b7dda;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.stats {
  display: flex;
  gap: 2rem;
  margin-left: auto;
  color: white;
  font-size: 0.95rem;
}

.stats span {
  background: rgba(255, 255, 255, 0.2);
  padding: 0.5rem 1rem;
  border-radius: 20px;
  backdrop-filter: blur(10px);
}

.filters {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
  flex-wrap: wrap;
  align-items: center;
}

.table-container {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  margin-bottom: 1.5rem;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
}

.data-table thead {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.data-table th {
  padding: 1rem;
  text-align: left;
  font-weight: 600;
  font-size: 0.95rem;
}

.data-table td {
  padding: 0.85rem 1rem;
  border-bottom: 1px solid #e0e0e0;
}

.data-table tbody tr:hover {
  background: #f5f5f5;
}

.data-table tbody tr:last-child td {
  border-bottom: none;
}

.no-data {
  text-align: center;
  color: #999;
  padding: 2rem !important;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
}

.page-btn {
  padding: 0.6rem 1.2rem;
  border: none;
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.9);
  font-size: 0.95rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
}

.page-btn:hover:not(:disabled) {
  background: white;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.page-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.page-info {
  color: white;
  font-weight: 600;
}
</style>
