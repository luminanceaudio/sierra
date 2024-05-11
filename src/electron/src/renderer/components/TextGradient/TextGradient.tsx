import React, { CSSProperties } from 'react';

export type TitleProps = {
  children: React.ReactNode;
  gradient?: string;
  style?: CSSProperties;
};

function TextGradient({ children, gradient, style }: TitleProps) {
  return (
    <span
      style={{
        background: gradient,
        WebkitBackgroundClip: 'text',
        WebkitTextFillColor: 'transparent',
        ...style,
      }}
    >
      {children}
    </span>
  );
}

TextGradient.defaultProps = {
  gradient:
    'linear-gradient(200.96deg, #ebfc4e -29.09%, #ed618f 51.77%, #00b7ff 129.35%)',
  style: {},
};

export default TextGradient;
