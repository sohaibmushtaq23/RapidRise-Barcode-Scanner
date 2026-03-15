<template>
  <v-container
    class="fill-height"
  >
  <v-row  justify="center">
      <v-col cols="12">
        <v-img
          :width="400"
          aspect-ratio="16/9"
          cover
          :src="logo"
          class="mx-auto"
        ></v-img>
      </v-col>
    </v-row>

    <v-row
      align="center"
      justify="center"
    >
      <v-col
        cols="12"
        sm="8"
        md="5"
      >

        <v-card elevation="6" height="550" class="d-flex flex-column">

          <v-card-title class="text-h5 text-center">
            Welcome {{ userName }}!
          </v-card-title>

          <v-card-text class="flex-grow-1">

            <v-row class="mt-4">

              <!-- Start Scan Button -->

              <v-col cols="12">
                <v-btn
                  block
                  size="x-large"
                  color="primary"
                  prepend-icon="mdi-barcode-scan"
                  @click="startScan"
                >
                  Start Scan
                </v-btn>
              </v-col>

              <!-- History Button -->

              <v-col cols="12">
                <v-btn
                  block
                  size="x-large"
                  color="secondary"
                  variant="outlined"
                  prepend-icon="mdi-history"
                  @click="openHistory"
                >
                  Scan History
                </v-btn>
              </v-col>

            </v-row>
          </v-card-text>

          <v-divider></v-divider>
          <v-card-actions>
            <v-btn
              color="red"
              variant="elevated"
              block
              size="large"
              prepend-icon="mdi-logout"
              class="justify-start"
              spaced="both"
              @click="handleLogout">
              Logout
            </v-btn>
          </v-card-actions>

        </v-card>

      </v-col>
    </v-row>

  </v-container>
</template>

<script setup lang="ts">
  import { useRouter } from 'vue-router'
  import { useAuthStore } from '@/stores/authStore'
  import logo from '@/assets/RapidRise_Dark_Logo.jpg'
  
  const router = useRouter()
  const authStore = useAuthStore()

  const userName=authStore.user?.name

  function startScan() {
    router.push('/scanner')
  }

  function openHistory() {
    router.push('/history')
  }

  function handleLogout() {
    authStore.logout()
    router.push('/login')
  }
</script>
