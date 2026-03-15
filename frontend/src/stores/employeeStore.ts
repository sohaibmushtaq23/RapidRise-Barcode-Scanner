import { defineStore } from 'pinia'
import { ref } from 'vue'
import { fetchEmployees as fetchEmployeesAPI } from '@/api/authService'

export interface Employee {
  id: number
  name: string
}

export const useEmployeeStore = defineStore('employee', () => {
  const employees = ref<Employee[]>([])

  // Fetch employees from API later
  async function fetchEmployees() {
    try {
      employees.value = await fetchEmployeesAPI()
    } catch (err) {
      console.error(err)
      employees.value = []
    }
  }

  return {
    employees,
    fetchEmployees
  }
})

