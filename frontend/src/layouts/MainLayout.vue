<template>
  <div class="main-layout">
    <header class="header">
      <h1>KeyView - 键盘使用历史记录</h1>
      <div class="github-link" @click="openGitHub" v-html="githubIcon"></div>
    </header>

    <main class="main-content">
      <slot></slot>
    </main>
  </div>
</template>

<script setup>
import { Browser } from '@wailsio/runtime'
import githubIcon from '@/assets/icons/github.svg?raw'

function openGitHub() {
  Browser.OpenURL('https://github.com/wangle201210/keyview')
}
</script>

<style scoped>
.main-layout {
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

.github-link :deep(svg) {
  width: 24px;
  height: 24px;
}

.main-content {
  padding: 2rem;
  max-width: 1400px;
  margin: 0 auto;
}
</style>
