import React from 'react';
import Title from '../Title/Title';

function Header() {
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
      <span style={{ display: 'flex', flex: 1, alignItems: 'center' }}>
        <Title
          style={{
            marginBottom: 0,
            alignItems: 'center',
            display: 'inline-block',
          }}
        >
          Samples
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
      </span>
    </div>
  );
}

export default Header;
