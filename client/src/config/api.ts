export const API_CONFIG = {
  BASE_URL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080',
  ENDPOINTS: {
    AUTH: {
      LOGIN: '/api/login',
      REGISTER: '/api/register',
      LOGOUT: '/api/logout'
    },
    PRODUCTS: {
      BASE: '/api/v1/products',
      GET_ALL: '/api/v1/products',
      CREATE: '/api/v1/products',
      DELETE: (id: number) => `/api/v1/products/${id}`
    }
  }
} as const; 