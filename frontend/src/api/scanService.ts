import axios from 'axios'

export async function sendBarcode(barcode: string) {
  const response = await axios.post('/api/scan', { barcode })  // no userId
  return response.data
}