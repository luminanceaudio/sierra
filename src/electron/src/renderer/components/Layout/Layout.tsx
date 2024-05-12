import React from 'react';
import DragHeader from '../DragHeader/DragHeader';
import SideMenu from '../SideMenu/SideMenu';
import Header from '../Header/Header';

type LayoutProps = {
  title: string;
  children?: React.ReactNode;
};

function Layout({ title, children }: LayoutProps) {
  return (
    <div
      style={{
        display: 'flex',
        flex: 1,
        flexDirection: 'column',
      }}
    >
      <DragHeader />
      <div style={{ display: 'flex', flex: 1 }}>
        <SideMenu />
        <div
          style={{
            display: 'flex',
            flex: 1,
            flexDirection: 'column',
          }}
        >
          <Header title={title} />
          <div
            style={{
              display: 'flex',
              flex: 1,
              position: 'relative',
            }}
          >
            {children}
          </div>
          {/* <div style={{ flex: 1, backgroundColor: '#1f1f1f' }}>Right sidebar</div> */}
        </div>
      </div>
    </div>
  );
}

Layout.defaultProps = {
  children: null,
};

export default Layout;
