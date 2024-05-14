import React from 'react';
import { StyledInput } from './Input.style';

export type InputProps = {
  value: string;
  onChange: (value: string) => void;
  placeholder?: string;
  style?: React.CSSProperties;
};

function Input({ value, onChange, placeholder, style }: InputProps) {
  return (
    <StyledInput
      style={{
        background: 'none',
        fontSize: 14,
        border: '1px solid #e1e1e1',
        borderRadius: 7,
        padding: '10px 15px',
        ...style,
      }}
      value={value}
      placeholder={placeholder}
      onChange={(e) => onChange(e.target.value)}
    />
  );
}

Input.defaultProps = {
  placeholder: '',
  style: {},
};

export default Input;
