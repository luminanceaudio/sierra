import React from 'react';
import DragHeader from '../DragHeader/DragHeader';
import SideMenu from '../SideMenu/SideMenu';
import Header from '../Header/Header';
import AudioPlayerBar from '../AudioPlayerBar/AudioPlayerBar';

type LayoutProps = {
  title: string;
  children?: React.ReactNode;
  style?: React.CSSProperties;
};

function Layout({ title, children, style }: LayoutProps) {
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
            maxHeight: '100vh',
            overflow: 'auto',
          }}
        >
          <Header title={title} />
          <div
            style={{
              display: 'flex',
              flex: 1,
              paddingTop: 95,
              paddingBottom: 95,
              justifyContent: 'center',
              ...style,
            }}
          >
            {children}
          </div>
          <AudioPlayerBar />
        </div>
      </div>
    </div>
  );
}

Layout.defaultProps = {
  children: null,
  style: {},
};

export default Layout;
