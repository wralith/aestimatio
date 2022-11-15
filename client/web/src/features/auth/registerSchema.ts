import { object, string } from 'zod'

export const registerSchema = object({
  username: string().min(4, { message: 'Username should have at least 4 characters' }),
  email: string().email({ message: 'Invalid email' }),
  password: string().min(6, { message: 'Password should have at least 6 characters' }),
})

export const loginSchema = object({
  email: string().email({ message: 'Invalid email' }),
  password: string().min(6, { message: 'Password should have at least 6 characters' }),
})
