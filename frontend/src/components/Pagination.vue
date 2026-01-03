<template>
  <div class="pagination">
    <button
      @click="$emit('prev')"
      :disabled="currentPage === 1"
      class="page-btn"
    >
      上一页
    </button>

    <span class="page-info">
      第 {{ currentPage }} / {{ totalPages }} 页
    </span>

    <button
      @click="$emit('next')"
      :disabled="currentPage === totalPages"
      class="page-btn"
    >
      下一页
    </button>
  </div>
</template>

<script setup>
defineProps({
  currentPage: {
    type: Number,
    required: true
  },
  totalPages: {
    type: Number,
    required: true
  }
})

defineEmits(['prev', 'next'])
</script>

<style scoped>
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
