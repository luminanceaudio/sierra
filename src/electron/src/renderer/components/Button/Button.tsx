import React from 'react';
import Clickable from '../Clickable/Clickable';

type IconButtonProps = {
  onClick?: (e: React.MouseEvent<HTMLButtonElement>) => void;
  style?: React.CSSProperties;
  children: React.ReactNode;
  type?: 'primary' | 'secondary';
  disabled?: boolean;
};

function Button({ onClick, style, children, type, disabled }: IconButtonProps) {
  return (
    <Clickable
      onClick={onClick}
      style={{
        padding: '10px 10px',
        borderRadius: 100,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        gap: 7,
        backgroundColor: type === 'primary' ? '#f1f1f1' : 'transparent',
        ...style,
      }}
      disabled={disabled}
    >
      {children}
    </Clickable>
  );
}

Button.defaultProps = {
  onClick: null,
  style: {},
  type: 'primary',
  disabled: false,
};

export default Button;
