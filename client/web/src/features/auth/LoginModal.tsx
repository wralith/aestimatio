import { Button, Group, PasswordInput, TextInput } from '@mantine/core'
import { useForm, zodResolver } from '@mantine/form'
import { showNotification } from '@mantine/notifications'
import { IconX } from '@tabler/icons'
import { useMutation } from '@tanstack/react-query'
import { useState } from 'react'
import client from '../../api/client'
import Modal from '../../components/shared/Modal'
import { useAuthStore } from '../../stores/authStore'
import { loginSchema, registerSchema } from './registerSchema'

const LoginModal = () => {
  const [opened, setOpened] = useState(false)
  const auth = useAuthStore()

  const form = useForm({
    validate: zodResolver(loginSchema),
    initialValues: {
      email: '',
      password: '',
    },
  })

  const onLogin = (v: any) => client.post('/auth/login', v)

  const mutation = useMutation({
    mutationFn: onLogin,
    onError: () =>
      showNotification({
        title: 'Login Error',
        message: 'Unexpected Error',
        icon: <IconX size={18} />,
        color: 'red',
      }),
    onSuccess: data => {
      showNotification({
        title: 'Login Successful',
        message: 'Welcome',
      })

      auth.login(data.data.jwt)
      auth.checkExp()
    },
  })

  return (
    <>
      <Modal title="Login" opened={opened} setOpened={setOpened}>
        <form onSubmit={form.onSubmit(async v => await mutation.mutateAsync(v))}>
          <TextInput label="Email" placeholder="something@mail.com" {...form.getInputProps('email')} />
          <PasswordInput label="Password" placeholder="password" {...form.getInputProps('password')} />
          <Button mt="md" type="submit">
            Login
          </Button>
        </form>
      </Modal>

      <Group position="center">
        <Button onClick={() => setOpened(true)}>Login</Button>
      </Group>
    </>
  )
}

export default LoginModal
