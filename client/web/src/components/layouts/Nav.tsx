// import { Navbar, Text } from '@mantine/core'

interface Props {
  opened: boolean
}

import React, { useState } from 'react'
import { createStyles, Navbar } from '@mantine/core'
import { IconBellRinging, IconSettings, IconBrandGithub, IconHome, IconFilePlus, IconListDetails, TablerIcon } from '@tabler/icons'
import { useAuthStore } from '../../stores/authStore'
import { string } from 'zod'

const useStyles = createStyles((theme, _params, getRef) => {
  const icon = getRef('icon')
  return {
    header: {
      paddingBottom: theme.spacing.md,
      marginBottom: theme.spacing.md * 1.5,
      borderBottom: `1px solid ${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[2]}`,
    },

    footer: {
      paddingTop: theme.spacing.md,
      marginTop: theme.spacing.md,
      borderTop: `1px solid ${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[2]}`,
    },

    link: {
      ...theme.fn.focusStyles(),
      display: 'flex',
      alignItems: 'center',
      textDecoration: 'none',
      fontSize: theme.fontSizes.sm,
      color: theme.colorScheme === 'dark' ? theme.colors.dark[1] : theme.colors.gray[7],
      padding: `${theme.spacing.xs}px ${theme.spacing.sm}px`,
      borderRadius: theme.radius.sm,
      fontWeight: 500,

      '&:hover': {
        backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
        color: theme.colorScheme === 'dark' ? theme.white : theme.black,

        [`& .${icon}`]: {
          color: theme.colorScheme === 'dark' ? theme.white : theme.black,
        },
      },
    },

    linkIcon: {
      ref: icon,
      color: theme.colorScheme === 'dark' ? theme.colors.dark[2] : theme.colors.gray[6],
      marginRight: theme.spacing.sm,
    },

    linkActive: {
      '&, &:hover': {
        backgroundColor: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).background,
        color: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).color,
        [`& .${icon}`]: {
          color: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).color,
        },
      },
    },
  }
})

interface LinkItem {
  link: string
  label: string
  icon: TablerIcon
}

const shared: LinkItem[] = [{ link: '', label: 'Home', icon: IconHome }]

const restricted: LinkItem[] = [
  { link: '', label: 'My Tasks', icon: IconListDetails },
  { link: '', label: 'Create Task', icon: IconFilePlus },
  { link: '', label: 'Incoming Deadlines', icon: IconBellRinging },
  { link: '', label: 'Account Settings', icon: IconSettings },
]

export default function Nav({ opened }: Props) {
  const { classes, cx } = useStyles()
  const [active, setActive] = useState('Home')
  const isLoggedIn = useAuthStore(s => s.isLoggedIn)

  const linkArray: LinkItem[] = isLoggedIn ? [...shared, ...restricted] : shared

  const links = linkArray.map(item => (
    <a
      className={cx(classes.link, { [classes.linkActive]: item.label === active })}
      href={item.link}
      key={item.label}
      onClick={event => {
        event.preventDefault()
        setActive(item.label)
      }}>
      <item.icon className={classes.linkIcon} stroke={1.5} />
      <span>{item.label}</span>
    </a>
  ))

  return (
    <Navbar width={{ sm: 300 }} p="md" hidden={!opened}>
      <Navbar.Section grow>{links}</Navbar.Section>

      <Navbar.Section className={classes.footer}>
        <a href="https://github.com/wralith/aestimatio" target="_blank" className={classes.link}>
          <IconBrandGithub className={classes.linkIcon} stroke={1.5} />
          <span>Source Code</span>
        </a>
      </Navbar.Section>
    </Navbar>
  )
}
