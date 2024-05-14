import React, { CSSProperties } from 'react';
import { StyledClickable } from './Clickable.style';

export type ClickableProps = {
  children: React.ReactNode;
  style?: CSSProperties;
  onClick?: (e: React.MouseEvent<HTMLButtonElement>) => void;
  disabled?: boolean;
};

function Clickable({ children, style, onClick, disabled }: ClickableProps) {
  return (
    <StyledClickable onClick={onClick} style={style} disabled={disabled}>
      {children}
    </StyledClickable>
  );
}

Clickable.defaultProps = {
  style: {},
  onClick: null,
  disabled: false,
};

export default Clickable;
