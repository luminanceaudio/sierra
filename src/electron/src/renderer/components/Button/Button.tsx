import React from 'react';
import Clickable, { ClickableProps } from '../Clickable/Clickable';

export type ButtonProps = ClickableProps & {};

function Button({ children, style, onClick, disabled }: ButtonProps) {
  return (
    <Clickable
      style={{
        backgroundColor: '#2e2e2e',
        padding: '7px 18px',
        borderRadius: 30,
        cursor: 'pointer',
        color: 'white',
        fontWeight: 500,
        ...style,
      }}
      disabled={disabled}
      onClick={onClick}
    >
      {children}
    </Clickable>
  );
}

Button.defaultProps = {
  style: {},
  disabled: false,
  onClick: null,
};

export default Button;
