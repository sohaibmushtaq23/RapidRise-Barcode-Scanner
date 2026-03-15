import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Employee } from '@/stores/employeeStore'
import { login as loginAPI } from '@/api/authService'
import type { LoginPayload } from '@/api/authService'
import axios from 'axios'

// Helper to parse user from localStorage
function parseUserStorage() {
  const item = localStorage.getItem('user')
  if (!item || item === 'undefined') return null
  try {
    return JSON.parse(item)
  } catch {
    return null
  }
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<Employee | null>(parseUserStorage())
  const token = ref<string | null>(localStorage.getItem('token'))

  // Set axios default Authorization header when token changes
  function setAuthHeader(t: string | null) {
    if (t) {
      axios.defaults.headers.common['Authorization'] = `Bearer ${t}`
    } else {
      delete axios.defaults.headers.common['Authorization']
    }
  }

  // Initialize header if token exists
  if (token.value) {
    setAuthHeader(token.value)
  }

  const isAuthenticated = computed(() => !!user.value && !!token.value)

  async function login(selectedEmployee: Employee, password: string) {
    try {
      const payload: LoginPayload = {
        employeeId: selectedEmployee.id,
        password
      }
      const response = await loginAPI(payload)

      if (response.success && response.token) {
        // Store user and token
        user.value = selectedEmployee
        token.value = response.token
        localStorage.setItem('user', JSON.stringify(selectedEmployee))
        localStorage.setItem('token', response.token)
        setAuthHeader(response.token)
        return true
      } else {
        console.warn(response.message)
        return false
      }
    } catch (err: any) {
      console.error(err)
      return false
    }
  }

  function logout() {
    user.value = null
    token.value = null
    localStorage.removeItem('user')
    localStorage.removeItem('token')
    setAuthHeader(null)
  }

  return {
    user,
    token,
    isAuthenticated,
    login,
    logout
  }
})