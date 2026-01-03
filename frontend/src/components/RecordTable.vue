<template>
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
        <tr v-for="record in records" :key="record.id">
          <td>{{ formatTime(record.created_at) }}</td>
          <td>{{ record.key_name }}</td>
          <td>{{ formatModifiers(record.modifier_flags) }}</td>
        </tr>
        <tr v-if="records.length === 0">
          <td colspan="3" class="no-data">暂无数据</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { formatTime, formatModifiers } from '@/utils/formatters'

defineProps({
  records: {
    type: Array,
    default: () => []
  }
})
</script>

<style scoped>
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
</style>
