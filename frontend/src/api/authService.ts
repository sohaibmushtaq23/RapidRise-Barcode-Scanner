import type { Employee } from '@/stores/employeeStore'

const API_BASE = 'https://192.168.100.13:8080' 

export async function fetchEmployees(): Promise<Employee[]> {
  const response = await fetch(`/api/employees`)
  if (!response.ok) throw new Error('Failed to fetch employees')
  return response.json()  // expects: [{ id: number, name: string }]
}

export interface LoginPayload {
  employeeId: number
  password: string
}

export interface LoginResponse {
  success: boolean
  message: string
  token?: string   // token from backend
}

export async function login(payload: LoginPayload): Promise<LoginResponse> {
  const response = await fetch(`/api/login`, {  // ensure leading slash
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  })
  if (!response.ok) {
    const err = await response.json()
    throw new Error(err.message || 'Login failed')
  }
  return response.json()
}