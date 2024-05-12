import React, { CSSProperties } from 'react';
import { StyledRadioGroup } from './RadioGroup.style';

type RadioGroupProps = {
  selected: string;
  onChange: (value: string) => void;
  children: React.ReactNode;
  style?: CSSProperties;
};

function RadioGroup({ children, style, selected, onChange }: RadioGroupProps) {
  return (
    <StyledRadioGroup value={selected} onChange={onChange} style={style}>
      {children}
    </StyledRadioGroup>
  );
}

RadioGroup.defaultProps = {
  style: {},
};

export default RadioGroup;
