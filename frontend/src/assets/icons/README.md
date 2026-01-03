# 图标资源管理

## 目录结构

```
assets/icons/
├── github.svg       # GitHub 图标
└── README.md        # 本文档
```

## 使用方法

### 推荐：使用 `?raw` 后缀导入（内联 SVG）

```vue
<template>
  <div v-html="githubIcon" class="icon-container"></div>
</template>

<script setup>
import githubIcon from '@/assets/icons/github.svg?raw'
</script>

<style scoped>
.icon-container :deep(svg) {
  width: 24px;
  height: 24px;
}
</style>
```

**优点**：
- SVG 直接内联到 HTML 中
- 可以通过 CSS 的 `color` 属性控制颜色
- 不需要额外的 HTTP 请求

### 作为图片导入

```vue
<template>
  <img :src="githubIcon" alt="GitHub" class="icon" />
</template>

<script setup>
import githubIcon from '@/assets/icons/github.svg'
</script>

<style scoped>
.icon {
  width: 24px;
  height: 24px;
}
</style>
```

**注意**：这种方式在 Vite 中可能需要额外配置

### 作为背景图片使用

```vue
<style scoped>
.icon-bg {
  width: 24px;
  height: 24px;
  background-image: url('@/assets/icons/github.svg');
  background-size: contain;
  background-repeat: no-repeat;
}
</style>
```

## SVG 图标最佳实践

1. **使用 `currentColor`**
   - 在 SVG 中使用 `fill="currentColor"` 可以让图标继承父元素的 `color` 属性
   - 这样可以通过 CSS 轻松改变图标颜色

2. **设置合适的尺寸**
   - SVG 的 `width` 和 `height` 属性定义图标原始尺寸
   - 实际显示尺寸通过 CSS 控制

3. **添加必要的属性**
   - 装饰性图标添加 `aria-hidden="true"`
   - 有意义的图标添加 `alt` 或 `aria-label`

4. **优化 SVG**
   - 移除不必要的元数据
   - 简化路径
   - 压缩文件大小

## Vite 后缀说明

- `?raw` - 将文件作为字符串导入
- `?url` - 将文件作为 URL 导入（返回资源路径）
- 无后缀 - Vite 会根据文件类型自动处理

## 添加新图标

1. 将 SVG 文件放到 `assets/icons/` 目录
2. 使用描述性的文件名（如：`github.svg`, `settings.svg`）
3. 确保图标使用 `fill="currentColor"` 以支持颜色自定义
4. 在组件中通过 `@/assets/icons/xxx.svg?raw` 引用

## 现有图标

| 文件名 | 用途 | 尺寸 |
|--------|------|------|
| github.svg | GitHub 链接 | 24x24 |
