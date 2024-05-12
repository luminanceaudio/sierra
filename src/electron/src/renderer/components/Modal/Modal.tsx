import { Dialog, DialogPanel } from '@headlessui/react';
import React, { ReactNode } from 'react';
import { StyledDialogTitle } from './Modal.style';

type ModalProps = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
  style?: React.CSSProperties;
  title: string;
};

function Modal({ isOpen, onClose, children, style, title }: ModalProps) {
  return (
    <Dialog
      open={isOpen}
      onClose={onClose}
      style={{
        zIndex: 50,
      }}
    >
      <div
        style={{
          position: 'fixed',
          inset: 0,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          width: '100vw',
          backgroundColor: 'rgba(0, 0, 0, 0.4)',
        }}
      >
        <DialogPanel
          style={{
            border: '1px solid #e5e7eb',
            backgroundColor: '#ffffff',
            padding: '1em',
            borderRadius: 7,
            display: 'flex',
            flexDirection: 'column',
            gap: 15,
            ...style,
          }}
        >
          <StyledDialogTitle>{title}</StyledDialogTitle>
          {children}
        </DialogPanel>
      </div>
    </Dialog>
  );
}

Modal.defaultProps = {
  style: {},
};

export default Modal;
