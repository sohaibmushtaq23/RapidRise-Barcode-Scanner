<template>  
  <v-container class="fill-height">
    <v-row  justify="center">
      <v-col cols="12">
        <v-img
          width="100%"
          aspect-ratio="16/9"
          contain
          :src="logo"
          class="mx-auto"
        ></v-img>
      </v-col>
    </v-row>
    <v-row  justify="center">
      <v-card width="400">

        <v-card-title>
          Login
        </v-card-title>

        <v-card-text>

          <v-select
              v-model="selectedEmployee"
              :items="employeeStore.employees"
              item-title="name"
              item-text="name"
              item-value="id"
              label="User"
              return-object
              density="comfortable"
              class="mb-4"
              prepend-inner-icon="mdi-account"
            />

            <!-- Password -->
            <v-text-field
              v-model="password"
              label="Password"
              type="password"
              density="comfortable"
              class="mb-4"
              prepend-inner-icon="mdi-lock"
              append-inner-icon="mdi-eye-off"
            />

        </v-card-text>

        <v-card-actions>
          <v-btn
            variant="elevated"
            height="60"
            color="#25374b" 
            size="large"
            class="text-h5 text-white font-weight-bold"
            block
            @click="login"
          >
            Login
          </v-btn>

        </v-card-actions>

      </v-card>
    </v-row>
  </v-container>
</template>
  
<script setup lang="ts">

import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useEmployeeStore } from '@/stores/employeeStore'
import type { Employee } from '@/stores/employeeStore'
import { useAuthStore } from '@/stores/authStore'
import logo from '@/assets/RapidRise_Dark_Logo.jpg'

const router = useRouter()
const employeeStore = useEmployeeStore()
const authStore = useAuthStore()

const selectedEmployee = ref<Employee | null>(null)
const password = ref('')

onMounted(async () => {
  await employeeStore.fetchEmployees()
})

async function login() {
  if (!selectedEmployee.value) {
    alert('Please select a user')
    return
  }
  if (!password.value) {
    alert('Enter password')
    return
  }

  const success = await authStore.login(selectedEmployee.value, password.value)
  if (success) {
    router.push('/welcome')
  } else {
    alert('Invalid password')
  }
}
</script>