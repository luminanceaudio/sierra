import React, { CSSProperties } from 'react';
import Title from '../Title/Title';

export type WidgetProps = {
  children: React.ReactNode;
  title?: string;
  style?: CSSProperties;
};

function Widget({ children, title, style }: WidgetProps) {
  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        margin: '30px 30px',
        borderRadius: 7,
        ...style,
      }}
    >
      {title && <Title>{title}</Title>}
      {children}
    </div>
  );
}

Widget.defaultProps = {
  title: null,
  style: {},
};

export default Widget;
