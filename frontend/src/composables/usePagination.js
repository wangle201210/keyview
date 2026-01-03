import { ref, computed } from 'vue'

/**
 * 分页管理 Composable
 */
export function usePagination(pageSize = 50) {
  const currentPage = ref(1)
  const size = ref(pageSize)

  const totalPages = computed(() => (totalCount) => {
    return Math.ceil(totalCount / size.value) || 1
  })

  function nextPage(totalPages) {
    if (currentPage.value < totalPages) {
      currentPage.value++
    }
  }

  function prevPage() {
    if (currentPage.value > 1) {
      currentPage.value--
    }
  }

  function resetPage() {
    currentPage.value = 1
  }

  function getPaginatedData(data) {
    const start = (currentPage.value - 1) * size.value
    const end = start + size.value
    return data.slice(start, end)
  }

  return {
    currentPage,
    pageSize: size,
    totalPages,
    nextPage,
    prevPage,
    resetPage,
    getPaginatedData
  }
}
