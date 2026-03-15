import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/authStore'

import LoginView from '../views/LoginView.vue'
import WelcomeView from '@/views/WelcomeView.vue'
import ScannerView from '../views/ScannerView.vue'
import HistoryView from '@/views/HistoryView.vue'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },

  {
    path: '/login',
    name: 'login',
    component: LoginView,
  },

  {
    path: '/welcome',
    name: 'welcome',
    component: WelcomeView,
    meta: { requiresAuth: true }
  },
  {
    path: '/scanner',
    name: 'scanner',
    component: ScannerView,
    meta: { requiresAuth: true }
  },
  {
    path: '/history',
    name: 'history',
    component: HistoryView,
    meta: { requiresAuth: true }
  }

]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to) => {

  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return '/login'
  }

  if (to.path === '/login' && authStore.isAuthenticated) {
    return '/welcome'
  }

  return true

})

export default router

