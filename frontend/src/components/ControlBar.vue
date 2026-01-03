<template>
  <div class="controls">
    <div class="status-indicator">
      <span class="status-dot recording"></span>
      <span class="status-text">正在监听键盘</span>
    </div>

    <button @click="$emit('refresh')" class="refresh-btn">
      刷新数据
    </button>

    <div class="stats">
      <span>总记录数: {{ totalRecords }}</span>
      <span>今日按键: {{ todayKeystrokes }}</span>
    </div>
  </div>
</template>

<script setup>
defineProps({
  totalRecords: {
    type: Number,
    default: 0
  },
  todayKeystrokes: {
    type: Number,
    default: 0
  }
})

defineEmits(['refresh'])
</script>

<style scoped>
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
</style>
