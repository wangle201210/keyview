<template>
  <div class="keyboard-container">
    <h3 class="keyboard-title">按键使用统计</h3>

    <div class="keyboard-wrapper">
      <div class="main-keyboard">
        <!-- 数字键行 -->
        <div class="row">
          <Key keyName="`" :count="getKeyCount('`')" :maxCount="maxCount" />
          <Key v-for="key in numberKeys" :key="key" :keyName="key" :count="getKeyCount(key)" :maxCount="maxCount" />
          <Key keyName="-" :count="getKeyCount('-')" :maxCount="maxCount" />
          <Key keyName="=" :count="getKeyCount('=')" :maxCount="maxCount" />
          <Key keyName="Backspace" displayName="⌫" :count="getKeyCount('Backspace')" :maxCount="maxCount" class="backspace" />
        </div>

        <!-- QWERTY 行 -->
        <div class="row">
          <Key keyName="Tab" displayName="⇥" :count="getKeyCount('Tab')" :maxCount="maxCount" class="tab" />
          <Key v-for="key in row2" :key="key" :keyName="key" :count="getKeyCount(key)" :maxCount="maxCount" />
          <Key keyName="[" :count="getKeyCount('[')" :maxCount="maxCount" />
          <Key keyName="]" :count="getKeyCount(']')" :maxCount="maxCount" />
          <Key keyName="\\" displayName="⌋" :count="getKeyCount('\\')" :maxCount="maxCount" />
        </div>

        <!-- ASDF 行 -->
        <div class="row">
          <Key keyName="Caps Lock" displayName="⇪" :count="getKeyCount('Caps Lock')" :maxCount="maxCount" class="caps-lock" />
          <Key v-for="key in row3" :key="key" :keyName="key" :count="getKeyCount(key)" :maxCount="maxCount" />
          <Key keyName=";" :count="getKeyCount(';')" :maxCount="maxCount" />
          <Key keyName="'" :count="getKeyCount(&quot;'&quot;)" :maxCount="maxCount" />
          <Key keyName="Return" displayName="↩" :count="getKeyCount('Return')" :maxCount="maxCount" class="return" />
        </div>

        <!-- ZXCV 行 -->
        <div class="row">
          <Key keyName="Shift" displayName="⇧" :count="getKeyCount('Shift')" :maxCount="maxCount" class="shift-left" />
          <Key v-for="key in row4" :key="key" :keyName="key" :count="getKeyCount(key)" :maxCount="maxCount" />
          <Key keyName="/" :count="getKeyCount('/')" :maxCount="maxCount" />
          <Key keyName="Shift" displayName="⇧" :count="getKeyCount('Shift')" :maxCount="maxCount" class="shift-right" />
        </div>

        <!-- 底部功能键行 -->
        <div class="row">
          <Key keyName="Fn" :count="getKeyCount('Fn')" :maxCount="maxCount" class="fn" />
          <Key keyName="Control" displayName="⌃" :count="getKeyCount('Control')" :maxCount="maxCount" />
          <Key keyName="Option" displayName="⌥" :count="getKeyCount('Option')" :maxCount="maxCount" />
          <Key keyName="Command" displayName="⌘" :count="getKeyCount('Command')" :maxCount="maxCount" class="command" />
          <Key keyName="Space" :count="getKeyCount('Space')" :maxCount="maxCount" class="space" />
          <Key keyName="Command" displayName="⌘" :count="getKeyCount('Command')" :maxCount="maxCount" class="command" />
          <Key keyName="Option" displayName="⌥" :count="getKeyCount('Option')" :maxCount="maxCount" />
        </div>
      </div>

      <!-- 功能键和方向键区域 -->
      <div class="side-keys">
        <!-- 功能键 -->
        <div class="function-keys">
          <Key keyName="Escape" displayName="ESC" :count="getKeyCount('Escape')" :maxCount="maxCount" class="esc" />
          <div class="f-keys">
            <div class="f-key-row">
              <Key keyName="F1" :count="getKeyCount('F1')" :maxCount="maxCount" />
              <Key keyName="F2" :count="getKeyCount('F2')" :maxCount="maxCount" />
              <Key keyName="F3" :count="getKeyCount('F3')" :maxCount="maxCount" />
              <Key keyName="F4" :count="getKeyCount('F4')" :maxCount="maxCount" />
            </div>
            <div class="f-key-row">
              <Key keyName="F5" :count="getKeyCount('F5')" :maxCount="maxCount" />
              <Key keyName="F6" :count="getKeyCount('F6')" :maxCount="maxCount" />
              <Key keyName="F7" :count="getKeyCount('F7')" :maxCount="maxCount" />
              <Key keyName="F8" :count="getKeyCount('F8')" :maxCount="maxCount" />
            </div>
            <div class="f-key-row">
              <Key keyName="F9" :count="getKeyCount('F9')" :maxCount="maxCount" />
              <Key keyName="F10" :count="getKeyCount('F10')" :maxCount="maxCount" />
              <Key keyName="F11" :count="getKeyCount('F11')" :maxCount="maxCount" />
              <Key keyName="F12" :count="getKeyCount('F12')" :maxCount="maxCount" />
            </div>
          </div>
        </div>

        <!-- 方向键 -->
        <div class="arrow-keys">
          <Key keyName="↑" displayName="↑" :count="getKeyCount('↑')" :maxCount="maxCount" />
          <div class="arrow-row">
            <Key keyName="←" displayName="←" :count="getKeyCount('←')" :maxCount="maxCount" />
            <Key keyName="↓" displayName="↓" :count="getKeyCount('↓')" :maxCount="maxCount" />
            <Key keyName="→" displayName="→" :count="getKeyCount('→')" :maxCount="maxCount" />
          </div>
        </div>
      </div>
    </div>

    <!-- 图例 -->
    <div class="legend">
      <div class="legend-item">
        <div class="legend-color" style="background: hsl(210, 80%, 95%);"></div>
        <span>少</span>
      </div>
      <div class="legend-item">
        <div class="legend-color" style="background: hsl(180, 80%, 30%);"></div>
        <span>多</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  keyStats: {
    type: Array,
    default: () => []
  }
})

// Mac 键盘布局
const numberKeys = ['1', '2', '3', '4', '5', '6', '7', '8', '9', '0']
const row2 = ['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P']
const row3 = ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L']
const row4 = ['Z', 'X', 'C', 'V', 'B', 'N', 'M', ',', '.', '/']

// 计算最大次数用于颜色缩放
const maxCount = computed(() => {
  if (props.keyStats.length === 0) return 1
  return Math.max(...props.keyStats.map(s => s.count))
})

// 创建按键名到次数的映射
const statsMap = computed(() => {
  const map = {}
  props.keyStats.forEach(stat => {
    map[stat.key_name] = stat.count
  })
  return map
})

// 获取按键次数
function getKeyCount(keyName) {
  return statsMap.value[keyName] || 0
}
</script>

<script>
import Key from './Key.vue'

export default {
  components: {
    Key
  }
}
</script>

<style scoped>
.keyboard-container {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 12px;
  padding: 2rem;
  margin-bottom: 1.5rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.keyboard-title {
  color: #333;
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
  text-align: center;
}

.keyboard-wrapper {
  display: flex;
  gap: 1rem;
  background: #f0f0f0;
  padding: 1rem;
  border-radius: 8px;
}

.main-keyboard {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.side-keys {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.row {
  display: flex;
  gap: 0.35rem;
  justify-content: flex-start;
}

.function-keys {
  display: flex;
  flex-direction: column;
  gap: 0.35rem;
}

.f-keys {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.f-key-row {
  display: flex;
  gap: 0.35rem;
}

.arrow-keys {
  display: flex;
  flex-direction: column;
  gap: 0.35rem;
}

.arrow-row {
  display: flex;
  gap: 0.35rem;
}

/* 右侧边键区域统一设置为较小高度 */
:deep(.side-keys .key) {
  height: 40px;
  min-height: 40px;
  max-height: 40px;
}

/* 特殊按键宽度 */
:deep(.backspace) {
  flex: 1.5;
}

:deep(.tab) {
  flex: 1.3;
}

:deep(.caps-lock) {
  flex: 1.5;
}

:deep(.return) {
  flex: 1.8;
}

:deep(.shift-left) {
  flex: 1.8;
}

:deep(.shift-right) {
  flex: 2.2;
}

:deep(.space) {
  flex: 4;
}

:deep(.command) {
  flex: 1.3;
}

:deep(.fn) {
  flex: 1.1;
}

:deep(.esc) {
  flex: 1.2;
}

.legend {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 2rem;
  margin-top: 1.5rem;
  padding-top: 1rem;
  border-top: 1px solid #e0e0e0;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #666;
  font-size: 0.9rem;
}

.legend-color {
  width: 40px;
  height: 25px;
  border-radius: 4px;
  border: 1px solid #ddd;
}
</style>
