import React from 'react';
import Clickable from '../Clickable/Clickable';

type IconButtonProps = {
  onClick?: () => void;
  style?: React.CSSProperties;
  children: React.ReactNode;
  type?: 'primary' | 'secondary';
  disabled?: boolean;
};

function IconButton({
  onClick,
  style,
  children,
  type,
  disabled,
}: IconButtonProps) {
  return (
    <Clickable
      onClick={onClick}
      style={{
        padding: '10px 10px',
        backgroundColor: type === 'primary' ? '#f1f1f1' : '#fff',
        border: type === 'primary' ? 'inherit' : '1px solid #e1e1e1',
        borderRadius: 7,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        ...style,
      }}
      disabled={disabled}
    >
      {children}
    </Clickable>
  );
}

IconButton.defaultProps = {
  onClick: null,
  style: {},
  type: 'primary',
  disabled: false,
};

export default IconButton;
