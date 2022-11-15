import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { useEffect } from 'react'
import { GetTokenFromLocalStorage } from './api/client'
import Layout from './components/layouts/Layout'

function App() {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <Layout>Hello World</Layout>
    </QueryClientProvider>
  )
}

export default App
