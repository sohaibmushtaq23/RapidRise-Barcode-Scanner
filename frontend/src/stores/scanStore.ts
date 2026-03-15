import { defineStore } from 'pinia'
import { sendBarcode } from '@/api/scanService'

export const useScanStore = defineStore('scan', {
  actions: {
    async processScan(barcode: string) {  
      try {
        const result = await sendBarcode(barcode)
        return result
      } catch (err: any) {
        return {
          success: false,
          message: err.message
        }
      }
    }
  }
})