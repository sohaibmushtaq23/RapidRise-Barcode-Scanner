import type { ScanLog } from '@/stores/scanStore'
import axios from 'axios'

export async function sendBarcode(barcode: string) {
  const response = await axios.post('/api/scan', { barcode })  // no userId
  return response.data
}

export async function fetchScans(): Promise<ScanLog[]> {
  const response = await fetch(`/api/history`)
  if (!response.ok) throw new Error('Failed to fetch scans log')
  return response.json()  // expects: [{ id: number, name: string }]
}