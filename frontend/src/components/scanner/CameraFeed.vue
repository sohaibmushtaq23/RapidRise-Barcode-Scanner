<script setup lang="ts">
import { onMounted, onBeforeUnmount, ref } from 'vue'
import { BrowserMultiFormatReader } from '@zxing/browser'
import { useRouter } from 'vue-router'

const emit = defineEmits(['scanned'])

const videoRef = ref<HTMLVideoElement | null>(null)

let codeReader: BrowserMultiFormatReader
let currentDeviceId: string | undefined
let controls: any

const router=useRouter()
const devices = ref<any[]>([])
const currentDeviceIndex = ref(0)
const scanLock = ref(false)  // Prevent multiple rapid scans

const snackbar = ref({
  show: false,
  message: ''
})

async function startCamera(deviceId?: string) {
  codeReader = new BrowserMultiFormatReader()
  
  devices.value = await BrowserMultiFormatReader.listVideoInputDevices()
  if (!devices.value.length) {
    alert('No camera found!')
    return
  }

  const selectedDevice = deviceId ?? devices.value[currentDeviceIndex.value].deviceId

  controls = await codeReader.decodeFromVideoDevice(
    selectedDevice,
    videoRef.value!,
    (result, err) => {
      if (result && !scanLock.value) {
        scanLock.value = true
        emit('scanned', result.getText())
        // Unlock after 1 second
        setTimeout(() => scanLock.value = false, 2000)
      }

      if (err && err.name !== 'NotFoundException' && err && err.name !== 'NotFoundException2') {
        console.error(err)
      }
    }
  )
}

function cancelScan(){
  stopCamera()
  router.push("/welcome")

}

function stopCamera() {
  
  if (controls) {
    controls.stop()
    controls = null
  }
}

function flipCamera() {
  if (!devices.value.length) return

  currentDeviceIndex.value = (currentDeviceIndex.value + 1) % devices.value.length

  startCamera(devices.value[currentDeviceIndex.value].deviceId)
}

onMounted(() => startCamera())

onBeforeUnmount(() => stopCamera())
</script>

<template>
    <v-card class="pa-2">
    <!-- Camera video feed -->
    <video
          ref="videoRef"
          width="100%"
          height="570"
          autoplay
          muted
          playsinline
        />

    <!-- Controls -->
    <v-row class="mt-1" justify="space-between">
      <v-btn color="primary" width="150" @click="flipCamera" prepend-icon="mdi-camera-flip">Flip</v-btn>
      <v-btn color="error" width="150" @click="cancelScan" prepend-icon="mdi-close-circle">Cancel</v-btn>
    </v-row>
  </v-card>
</template>