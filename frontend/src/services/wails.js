// Wails runtime API wrapper
// 直接调用 Wails 生成的绑定

import * as AppService from '../../bindings/github.com/wangle201210/keyview/internal/app/appservice.js';

// API 服务
export const api = {
  // 开始记录
  async startRecording() {
    return await AppService.StartRecording()
  },

  // 停止记录
  async stopRecording() {
    return await AppService.StopRecording()
  },

  // 获取记录状态
  async isRecording() {
    return await AppService.IsRecording()
  },

  // 获取记录列表
  async getRecords(offset = 0, limit = 50) {
    return await AppService.GetRecords(offset, limit)
  },

  // 根据条件筛选记录
  async getRecordsByFilter(keyName, date, isDown, offset = 0, limit = 50) {
    return await AppService.GetRecordsByFilter(keyName, date, isDown, offset, limit)
  },

  // 获取总记录数
  async getTotalCount() {
    return await AppService.GetTotalCount()
  },

  // 获取今日按键次数
  async getTodayKeystrokes() {
    return await AppService.GetTodayKeystrokes()
  },

  // 获取所有唯一按键名称
  async getUniqueKeyNames() {
    return await AppService.GetUniqueKeyNames()
  },

  // 获取按键统计（支持日期范围筛选）
  async getKeyStats(startDate = '', endDate = '') {
    return await AppService.GetKeyStats(startDate, endDate)
  },

  // 删除指定日期之前的记录
  async deleteRecordsBefore(date) {
    return await AppService.DeleteRecordsBefore(date)
  }
}

export default api
