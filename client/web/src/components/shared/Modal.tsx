import React, { useState } from 'react'
import { Modal as MantineModal, Button, Group } from '@mantine/core'

interface Props {
  title: string
  opened: boolean
  setOpened: React.Dispatch<React.SetStateAction<boolean>>
  children: React.ReactNode
}

const Modal = ({ title, opened, setOpened, children }: Props) => {
  return (
    <>
      <MantineModal opened={opened} onClose={() => setOpened(false)} title={title}>
        {children}
      </MantineModal>
    </>
  )
}

export default Modal
