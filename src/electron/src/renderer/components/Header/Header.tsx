import React from 'react';
import Title from '../Title/Title';

type HeaderProps = {
  title: string;
};

function Header({ title }: HeaderProps) {
  return (
    <div
      style={{
        display: 'flex',
        padding: 20,
        boxSizing: 'border-box',
        height: 95,
        borderBottom: '1px solid #f1f1f1',
        alignItems: 'center',
      }}
    >
      <div>
        <div style={{ alignItems: 'center' }}>
          <Title
            style={{
              marginBottom: 0,
              alignItems: 'center',
              display: 'inline-block',
              fontSize: 19,
            }}
          >
            {title}
            {/* <TextGradient>{' // '}</TextGradient> */}
          </Title>
          {/* <Title */}
          {/*  style={{ */}
          {/*    paddingTop: 0, */}
          {/*    fontWeight: 100, */}
          {/*    marginLeft: 6, */}
          {/*  }} */}
          {/* > */}
          {/*  Library */}
          {/* </Title> */}
        </div>
        {/* <div style={{ fontSize: 13, color: '#8f8f8f' }}>All</div> */}
      </div>
    </div>
  );
}

export default Header;
