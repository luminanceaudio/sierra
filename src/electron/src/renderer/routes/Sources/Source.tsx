import path from 'path';
import { FolderIcon, TrashIcon } from '@heroicons/react/24/outline';
import { FolderPlusIcon } from '@heroicons/react/24/solid';
import React, { useState } from 'react';
import Clickable from '../../components/Clickable/Clickable';
import NewSourceModal from './NewSourceModal';
import IconButton from '../../components/IconButton/IconButton';
import { useDeleteSource } from '../../../api/sources';
import { DeleteSourceRequest } from '../../../proto/app/apprequests';
import Toast from '../../components/Toast/Toast';

type SourceBaseProps = {
  render: React.ReactNode;
};

function SourceBase({ render }: SourceBaseProps) {
  return (
    <div
      style={{
        display: 'flex',
        gap: 10,
        alignItems: 'center',
        padding: 15,
        border: '1px solid #e6e6e6',
        borderRadius: 10,
        cursor: 'default',
      }}
    >
      {render}
    </div>
  );
}

export type SourceProps = {
  uri: string;
};

export function Source({ uri }: SourceProps) {
  const { mutate: deleteSource, error, isPending } = useDeleteSource();

  return (
    <SourceBase
      render={
        <div
          style={{ display: 'flex', flex: 1, gap: 10, alignItems: 'center' }}
        >
          <Toast error={error} />
          <IconButton>
            <FolderIcon width={20} />
          </IconButton>
          <div
            style={{
              display: 'flex',
              flex: 1,
              flexDirection: 'column',
              gap: 4,
            }}
          >
            <div style={{ fontSize: 17 }}>{path.basename(uri)}</div>
            <span
              style={{
                color: '#a1a1a1',
                fontSize: 12,
                width: '100%',
                textOverflow: 'clip',
              }}
            >
              {uri}
            </span>
          </div>
          <IconButton
            type="secondary"
            style={{ border: 'none' }}
            onClick={() =>
              deleteSource(
                DeleteSourceRequest.create({
                  uri,
                }),
              )
            }
            disabled={isPending}
          >
            <TrashIcon width={20} style={{ color: '#cf1212' }} />
          </IconButton>
        </div>
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
        <div style={{ fontSize: 15 }}>Add source</div>
      </Clickable>
      {isOpen && (
        <NewSourceModal isOpen={isOpen} onClose={() => setOpen(false)} />
      )}
    </div>
  );
}
