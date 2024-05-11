import React, { useRef, KeyboardEvent } from 'react';

export type InputProps = {
  // eslint-disable-next-line react/no-unused-prop-types
  ref: any;

  value: string;
  onChange: (value: string) => void;
  onEnter?: () => void;
  placeholder?: string;
  style?: React.CSSProperties;
  focusOnEnter?: boolean;
};

function Input({
  value,
  onChange,
  onEnter,
  placeholder,
  style,
  focusOnEnter,
}: InputProps) {
  const ref: any = useRef();
  function keydown(event: KeyboardEvent<HTMLInputElement>) {
    if (event.key !== 'Enter') {
      return;
    }

    const current = ref.current as any;
    current.blur();
  }

  function blur() {
    if (value !== '') {
      onEnter?.();
      if (focusOnEnter) {
        (ref.current as any).focus();
      }
    }
  }

  return (
    <input
      ref={ref}
      style={{
        background: 'none',
        color: '#111',
        fontSize: 16,
        fontWeight: 600,
        ...style,
      }}
      value={value}
      placeholder={placeholder}
      onChange={(e) => onChange(e.target.value)}
      onKeyDown={keydown}
      onBlur={blur}
    />
  );
}

Input.defaultProps = {
  placeholder: '',
  style: {},
  onEnter: () => {},
  focusOnEnter: false,
  onLoad: () => {},
};

export default Input;
