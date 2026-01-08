/**
 * 格式化时间戳为本地化字符串
 * @param {string|Date} timestamp - 时间戳
 * @returns {string} 格式化后的时间字符串
 */
export function formatTime(timestamp) {
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

/**
 * 格式化修饰键标志为可读字符串
 * @param {number} flags - 修饰键标志
 * @returns {string} 格式化后的修饰键字符串
 */
export function formatModifiers(flags) {
  if (!flags || flags === 0) return '-'

  const modifiers = []
  if (flags & 0x20000) modifiers.push('⇧')
  if (flags & 0x40000) modifiers.push('⌃')
  if (flags & 0x80000) modifiers.push('⌥')
  if (flags & 0x100000) modifiers.push('⌘')
  if (flags & 0x10000) modifiers.push('Caps')

  return modifiers.length > 0 ? modifiers.join(' + ') : '-'
}

/**
 * 格式化数字为带单位的字符串
 * @param {number} num - 要格式化的数字
 * @returns {string} 格式化后的字符串
 */
export function formatCount(num) {
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  } else if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'K'
  }
  return num.toString()
}

/**
 * 获取最近一个月的日期范围（从30天前到今天）
 * @returns {Date[]} [开始日期, 结束日期]
 */
export function getLastMonthRange() {
  const now = new Date()
  const endDate = new Date(now)
  const startDate = new Date(now)
  startDate.setDate(now.getDate() - 30)
  return [startDate, endDate]
}
