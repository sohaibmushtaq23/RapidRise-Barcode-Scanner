<script setup lang="ts">

  import { ref, onMounted } from 'vue'
  import { useRouter } from 'vue-router'
  import { useScanStore } from '@/stores/scanStore'

  const router = useRouter()
  const scanStore = useScanStore()

  interface SqlNullString {
    String: string;
    Valid: boolean;
  }

  // Helper to format: DD/MM/YY HH:mm
  const formatScanDate = (dateStr: string | null | undefined): string => {
    if (!dateStr) return '---'
    const d = new Date(dateStr)
    
    // Pad with leading zeros
    const pad = (n: number) => n.toString().padStart(2, '0')
    
    const day = pad(d.getDate())
    const month = pad(d.getMonth() + 1)
    const year = d.getFullYear().toString().slice(-2)
    const hours = pad(d.getHours())
    const mins = pad(d.getMinutes())

    return `${day}/${month}/${year} ${hours}:${mins}`
  }

  const getEmployeeName = (scannedBy: string | SqlNullString | null): string => {
    if (scannedBy && typeof scannedBy === 'object') {
      return (scannedBy as SqlNullString).Valid ? (scannedBy as SqlNullString).String : 'Unknown'
    }
    return (scannedBy as string) || 'Unknown'
  }

  onMounted(async () => {
    await scanStore.fetchScanningHistory()
  })

</script>

<template>
    <v-container>
  
      <v-row>
        <v-col>
  
          <v-btn
            class="mb-4"
            @click="$router.push('/welcome')"
            prepend-icon="mdi-arrow-left"
            color="error"
          >
            Back
          </v-btn>
  
          <h2>Scan History</h2>
  
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-list density="compact">
            <v-list-item
              v-for="(scan, index) in scanStore.scans"
              :key="index"
              :value="index"
              active-color="primary"
            >
              <!-- Title: Component Name (Bold Black) -->
              <v-list-item-title class="font-weight-bold opacity-100">
                {{ scan.component }}
              </v-list-item-title>

              <!-- Subtitle: Process, Name, and Date (Solid Black) -->
              <v-list-item-subtitle class="opacity-100 mt-1">
                {{ scan.process }} • 
                {{ getEmployeeName(scan.scannedBy) }} • 
                {{ formatScanDate(scan.scannedAt) }}
              </v-list-item-subtitle>
            </v-list-item>
          </v-list>
        </v-col>
      </v-row>
      
  
    </v-container>
  </template>
