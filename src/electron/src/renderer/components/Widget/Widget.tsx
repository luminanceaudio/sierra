import React, { CSSProperties } from 'react';
import Title from '../Title/Title';

export type WidgetProps = {
  children: React.ReactNode;
  title: string;
  style?: CSSProperties;
};

function Widget({ children, title, style }: WidgetProps) {
  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: 'white',
        borderRadius: 15,
        ...style,
      }}
    >
      <Title>{title}</Title>
      {children}
    </div>
  );
}

Widget.defaultProps = {
  style: {},
};

export default Widget;
