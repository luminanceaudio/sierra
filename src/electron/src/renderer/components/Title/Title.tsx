import React, { CSSProperties } from 'react';

export type TitleProps = {
  children: React.ReactNode;
  style?: CSSProperties;
};

function Title({ children, style }: TitleProps) {
  return (
    <h2
      style={{
        margin: 0,
        marginBottom: 17,
        padding: 0,
        fontSize: 23,
        ...style,
      }}
    >
      {children}
    </h2>
  );
}

Title.defaultProps = {
  style: {},
};

export default Title;
