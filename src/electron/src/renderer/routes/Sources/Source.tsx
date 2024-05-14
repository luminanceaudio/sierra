import path from 'path';
import { FolderIcon, TrashIcon } from '@heroicons/react/24/outline';
import { FolderPlusIcon } from '@heroicons/react/24/solid';
import React, { useState } from 'react';
import NewSourceModal from './NewSourceModal';
import Button from '../../components/Button/Button';
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
          <Button>
            <FolderIcon width={20} />
          </Button>
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
          <Button
            type="secondary"
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
          </Button>
        </div>
      }
    />
  );
}

export function SourceNew() {
  const [isOpen, setOpen] = useState(false);
  return (
    <div>
      <Button onClick={() => setOpen(true)}>
        <FolderPlusIcon width="20" />
        <div style={{ fontSize: 15 }}>Add source</div>
      </Button>
      {isOpen && (
        <NewSourceModal isOpen={isOpen} onClose={() => setOpen(false)} />
      )}
    </div>
  );
}
