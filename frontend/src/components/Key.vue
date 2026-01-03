<template>
  <div
    v-if="!hidden"
    class="key"
    :style="{ backgroundColor: bgColor }"
  >
    <div class="key-name">{{ displayLabel }}</div>
    <div v-if="count > 0" class="key-count">{{ formatCount(count) }}</div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  keyName: {
    type: String,
    required: true
  },
  displayName: {
    type: String,
    default: null
  },
  count: {
    type: Number,
    default: 0
  },
  maxCount: {
    type: Number,
    default: 1
  },
  hidden: {
    type: Boolean,
    default: false
  }
})

// 显示名称转换
const displayLabel = computed(() => {
  if (props.displayName) {
    return props.displayName
  }

  const nameMap = {
    'Space': ' ',
    ',': ',',
    '.': '.',
    '/': '/',
    ';': ';',
    "'": "'",
    '`': '`',
    '-': '-',
    '=': '=',
    '[': '[',
    ']': ']'
  }
  return nameMap[props.keyName] || props.keyName
})

// 根据使用次数计算颜色
const bgColor = computed(() => {
  if (props.count === 0) {
    return '#ffffff'
  }

  // 从浅蓝到深蓝的热力图颜色
  const ratio = props.count / props.maxCount

  // 使用 HSL 颜色空间，从 210 (浅蓝) 到 240 (深蓝)
  // 亮度从 95% 到 30%
  const lightness = 95 - (ratio * 65)
  const hue = 210 - (ratio * 30)

  return `hsl(${hue}, 80%, ${lightness}%)`
})

// 格式化次数显示
function formatCount(num) {
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  } else if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'K'
  }
  return num.toString()
}
</script>

<style scoped>
.key {
  height: 50px;
  min-width: 40px;
  flex: 1;
  background: #ffffff;
  border: 1px solid #ccc;
  border-radius: 6px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  font-size: 0.95rem;
  font-weight: 600;
  color: #333;
  transition: all 0.2s;
  cursor: default;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: relative;
  padding: 0 4px;
}

.key:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
  border-color: #999;
}

.key-name {
  font-size: 0.85rem;
  font-weight: 600;
  line-height: 1.2;
}

.key-count {
  font-size: 0.65rem;
  color: #666;
  margin-top: 2px;
  font-weight: 500;
}
</style>
