import React, { CSSProperties } from 'react';

export type TitleProps = {
  children: React.ReactNode;
  style?: CSSProperties;
};

function Subtitle({ children, style }: TitleProps) {
  return (
    <h4
      style={{
        margin: 0,
        marginBottom: 17,
        padding: 0,
        fontWeight: '400',
        color: '#787878',
        ...style,
      }}
    >
      {children}
    </h4>
  );
}

Subtitle.defaultProps = {
  style: {},
};

export default Subtitle;
