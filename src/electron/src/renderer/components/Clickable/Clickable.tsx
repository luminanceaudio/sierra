import React, { CSSProperties } from 'react';
import { StyledClickable } from './Clickable.style';

type ClickableProps = {
  children: React.ReactNode;
  style?: CSSProperties;
  onClick: () => void;
  disabled?: boolean;
};

function Clickable({ children, style, onClick, disabled }: ClickableProps) {
  return (
    <StyledClickable onClick={onClick} style={{ ...style }} disabled={disabled}>
      {children}
    </StyledClickable>
  );
}

Clickable.defaultProps = {
  style: {},
  disabled: false,
};

export default Clickable;
