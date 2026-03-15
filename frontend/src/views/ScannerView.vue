<script setup lang="ts">
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'
  import { useScanStore } from '@/stores/scanStore'

  import CameraFeed from '@/components/scanner/CameraFeed.vue'
  import ScanOverlay from '@/components/scanner/ScanOverlay.vue'

  const router = useRouter()
  const scanStore = useScanStore()

  const cameraRef = ref<any>(null)

  const snackbar = ref(false)
  const snackbarText = ref('')

  function showMessage(msg: string) {
    snackbarText.value = msg
    snackbar.value = true
  }

  async function handleScan(text: string) {
    showMessage(`Scanned: ${text}`)
    const result = await scanStore.processScan(text)  // no userId
    if (result.success) {
      showMessage(result.message)
    } else {
      showMessage(`Error: ${result.message}`)
    }
  }

</script>

<template>

  <v-container fluid class="pa-0 fill-height">

    <v-sheet height="100vh" width="100vw">

      <CameraFeed
        ref="cameraRef"
        @scanned="handleScan"
      />

      <ScanOverlay />

      <v-snackbar
        v-model="snackbar"
        location="bottom"
      >
      {{ snackbarText }}
      </v-snackbar>

    </v-sheet>

  </v-container>

</template>