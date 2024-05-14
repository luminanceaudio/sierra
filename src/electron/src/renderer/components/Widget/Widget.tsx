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
        flex: 1,
        maxWidth: 600,
        margin: '30px 30px',
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
