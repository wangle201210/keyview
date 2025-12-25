// Wails runtime type definitions

export interface KeyRecord {
  id: number
  created_at: string
  key_code: number
  key_name: string
  is_down: boolean
  modifier_flags: number
}

export interface WailsAPI {
  StartRecording(): Promise<void>
  StopRecording(): Promise<void>
  IsRecording(): Promise<boolean>
  GetRecords(offset: number, limit: number): Promise<KeyRecord[]>
  GetRecordsByFilter(
    keyName: string,
    date: string,
    isDown: boolean | null,
    offset: number,
    limit: number
  ): Promise<KeyRecord[]>
  GetTotalCount(): Promise<number>
  GetTodayKeystrokes(): Promise<number>
  GetUniqueKeyNames(): Promise<string[]>
  DeleteRecordsBefore(date: string): Promise<number>
}

declare global {
  interface Window {
    wails?: WailsAPI
  }
}

export {}
