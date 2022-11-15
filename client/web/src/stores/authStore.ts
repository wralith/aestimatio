import create from 'zustand'
import { persist } from 'zustand/middleware'
import jwt_decode from 'jwt-decode'
import { date } from 'zod'

interface AuthState {
  isLoggedIn: boolean
  jwt: string
  login: (payload: string) => void
  logout: () => void
  checkExp: () => void
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      isLoggedIn: false,
      jwt: '',
      exp: 0,
      login: (payload: string) => set({ jwt: payload, isLoggedIn: true }),
      logout: () => set({ jwt: '', isLoggedIn: false }),
      checkExp: () => {
        const token = get().jwt
        try {
          var decoded: any = jwt_decode(token)
          var exp: number = decoded.exp
          if (Date.now() > exp * 1000) {
            get().logout()
          }
        } catch (error) {
          console.log(error)
          get().logout
        }
      },
    }),
    {
      name: 'auth-store',
    }
  )
)
