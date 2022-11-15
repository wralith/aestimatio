import { Burger, Button, Container, Group, Header, MediaQuery, Text, Title } from '@mantine/core'
import LoginModal from '../../features/auth/LoginModal'
import RegisterModal from '../../features/auth/RegisterModal'
import { useAuthStore } from '../../stores/authStore'

interface Props {
  color: string
  opened: boolean
  setOpened: React.Dispatch<React.SetStateAction<boolean>>
}

const AppHeader = ({ opened, setOpened, color }: Props) => {
  const auth = useAuthStore()

  return (
    <Header height={{ base: 50, md: 70 }} p="md">
      <div style={{ display: 'flex', alignItems: 'center', height: '100%' }}>
        <MediaQuery largerThan="sm" styles={{ display: 'none' }}>
          <Burger opened={opened} onClick={() => setOpened(o => !o)} size="sm" color={color} mr="xl" />
        </MediaQuery>
        <Group position="apart" w={'100%'}>
          <Title order={3}>Aestimatio</Title>
          {auth.isLoggedIn ? (
            <Button onClick={auth.logout}>Logout</Button>
          ) : (
            <Group>
              <RegisterModal />
              <LoginModal />
            </Group>
          )}
        </Group>
      </div>
    </Header>
  )
}

export default AppHeader
