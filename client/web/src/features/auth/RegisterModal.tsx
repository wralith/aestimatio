import { Button, Group, PasswordInput, TextInput } from '@mantine/core'
import { useForm, zodResolver } from '@mantine/form'
import { useMutation } from '@tanstack/react-query'
import { useState } from 'react'

import Modal from '../../components/shared/Modal'
import client from '../../api/client'
import { registerSchema } from './registerSchema'
import { showNotification } from '@mantine/notifications'
import { IconX } from '@tabler/icons'

const RegisterModal = () => {
  const [opened, setOpened] = useState(false)
  const form = useForm({
    validate: zodResolver(registerSchema),
    initialValues: {
      username: '',
      email: '',
      password: '',
    },
  })

  const onRegister = (v: any) => client.post('/auth/register', v)

  const mutation = useMutation({
    mutationFn: onRegister,
    onError: () =>
      showNotification({
        title: 'Register Error',
        message: 'Unexpected Error',
        icon: <IconX size={18} />,
        color: 'red',
      }),
    onSuccess: () => {
      showNotification({
        title: 'Registered Successfully',
        message: 'Welcome',
      })
    },
  })

  return (
    <>
      <Modal title="Register" opened={opened} setOpened={setOpened}>
        <form onSubmit={form.onSubmit(async v => await mutation.mutateAsync(v))}>
          <TextInput label="Username" placeholder="Your Desired Name" {...form.getInputProps('username')} />
          <TextInput label="Email" placeholder="something@mail.com" {...form.getInputProps('email')} />
          <PasswordInput label="Password" placeholder="password" {...form.getInputProps('password')} />
          <Button mt="md" type="submit">
            Register
          </Button>
        </form>
      </Modal>

      <Group position="center">
        <Button onClick={() => setOpened(true)}>Register</Button>
      </Group>
    </>
  )
}

export default RegisterModal
