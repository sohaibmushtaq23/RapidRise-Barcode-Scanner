import { defineStore } from 'pinia'
import { ref } from 'vue'
import { sendBarcode, fetchScans } from '@/api/scanService'

export interface ScanLog {
  process: string
  component: string
  scannedBy: string
  scannedAt: string
}

export const useScanStore = defineStore('scan', () => {
  // State
  const scans = ref<ScanLog[]>([])

  // Actions
  async function processScan(barcode: string) {  
    try {
      const result = await sendBarcode(barcode)
      return result
    } catch (err: any) {
      return {
        success: false,
        message: err.message || 'Scan failed'
      }
    }
  }

  async function fetchScanningHistory() {
    try {
      scans.value = await fetchScans()
    } catch (err) {
      console.error('Fetch error:', err)
      scans.value = []
    }
  }

  return { 
    scans, 
    processScan, 
    fetchScanningHistory 
  }
})
