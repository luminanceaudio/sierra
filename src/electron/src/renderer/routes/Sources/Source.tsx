import path from 'path';
import { FolderIcon } from '@heroicons/react/24/outline';
import { FolderPlusIcon } from '@heroicons/react/24/solid';
import React, { useState } from 'react';
import Clickable from '../../components/Clickable/Clickable';
import NewSourceModal from './NewSourceModal';

type SourceBaseProps = {
  render: React.ReactNode;
};

function SourceBase({ render }: SourceBaseProps) {
  return (
    <Clickable
      onClick={() => {}}
      style={{
        display: 'flex',
        gap: 10,
        alignItems: 'center',
        padding: 15,
        border: '1px solid #e6e6e6',
        borderRadius: 10,
        width: '100%',
      }}
    >
      {render}
    </Clickable>
  );
}

export type SourceProps = {
  uri: string;
};

export function Source({ uri }: SourceProps) {
  return (
    <SourceBase
      render={
        <>
          <div
            style={{
              width: 20,
              padding: '10px 10px',
              backgroundColor: '#f1f1f1',
              borderRadius: 7,
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
            }}
          >
            <FolderIcon width="100%" />
          </div>
          <div
            style={{
              display: 'flex',
              flexDirection: 'column',
              gap: 4,
            }}
          >
            <div style={{ fontSize: 17 }}>{path.basename(uri)}</div>
            <span style={{ color: '#a1a1a1', fontSize: 12 }}>{uri}</span>
          </div>
        </>
      }
    />
  );
}

export function SourceNew() {
  const [isOpen, setOpen] = useState(false);
  return (
    <div>
      <Clickable
        style={{
          height: 40,
          padding: '10px 10px',
          backgroundColor: '#f1f1f1',
          borderRadius: 7,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          gap: 10,
        }}
        onClick={() => setOpen(true)}
      >
        <FolderPlusIcon width="20" />
        <div style={{ fontSize: 16 }}>Add source</div>
      </Clickable>
      <NewSourceModal isOpen={isOpen} onClose={() => setOpen(false)} />
    </div>
  );
}
