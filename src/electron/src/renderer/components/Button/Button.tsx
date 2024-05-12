import React, { CSSProperties } from 'react';
import config from '../../../config/config';

export type ButtonProps = {
  onClick: () => void;
  children: React.ReactNode;
  style?: CSSProperties;
};

function Button({ onClick, children, style }: ButtonProps) {
  return (
    <button
      type="button"
      onClick={onClick}
      style={{
        backgroundColor: '#2e2e2e',
        padding: '7px 18px',
        borderRadius: 30,
        cursor: 'pointer',
        color: 'white',
        fontWeight: 500,
        ...style,
      }}
    >
      {children}
    </button>
  );
}

Button.defaultProps = {
  style: {},
};

export default Button;
