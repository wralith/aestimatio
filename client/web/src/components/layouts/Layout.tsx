import { useState } from 'react'
import { AppShell, Box, useMantineTheme } from '@mantine/core'
import Nav from './Nav'
import AppHeader from './AppHeader'

interface Props {
  children: React.ReactNode
}

const Layout = ({ children }: Props) => {
  const theme = useMantineTheme()
  const [opened, setOpened] = useState(false)
  return (
    <AppShell
      styles={{
        main: {
          background: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.colors.gray[0],
        },
      }}
      navbarOffsetBreakpoint="sm"
      asideOffsetBreakpoint="sm"
      navbar={<Nav opened={opened} />}
      header={<AppHeader color={theme.colors.gray[6]} opened={opened} setOpened={setOpened} />}>
      <Box>{children}</Box>
    </AppShell>
  )
}

export default Layout
