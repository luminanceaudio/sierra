import React, { useEffect, useState } from 'react';
import path from 'path';
import Modal from '../../components/Modal/Modal';
import RadioGroup from '../../components/RadioGroup/RadioGroup';
import RadioItem from '../../components/RadioGroup/RadioItem';
import Button from '../../components/Button/Button';
import { useCreateSource } from '../../../api/sources';
import { CreateSourceRequest } from '../../../proto/app/apprequests';
import { getErrorMessage } from '../../../api/reactquery';

type NewSourceModalProps = {
  isOpen: boolean;
  onClose: () => void;
};

function NewSourceModal({ isOpen, onClose }: NewSourceModalProps) {
  const { mutate: createSource, data, error } = useCreateSource();
  const [selected, setSelected] = useState('local');
  const [errorMessage, setErrorMessage] = useState('');
  const inputRef = React.createRef<HTMLInputElement>();

  useEffect(() => {
    if (error) {
      setErrorMessage(getErrorMessage(error));
      return;
    }

    if (data?.status === 200) {
      onClose();
      setErrorMessage('');
      return;
    }

    setErrorMessage(data?.statusText || '');
  }, [data, onClose, setErrorMessage, error]);

  return (
    <Modal
      isOpen={isOpen}
      onClose={onClose}
      style={{ width: '60%', maxWidth: 400 }}
      title="Add new source"
    >
      <RadioGroup selected={selected} onChange={setSelected}>
        <RadioItem value="local" title="Local source" />
        <RadioItem value="remote" title="Remote source" disabled />
      </RadioGroup>
      <div
        style={{
          marginTop: 15,
          display: 'flex',
          flex: 1,
          flexDirection: 'column',
        }}
      >
        <div
          style={{
            display: 'flex',
            flex: 1,
            justifyContent: 'space-between',
            alignItems: 'center',
          }}
        >
          <span
            style={{
              color: '#cf1212',
              fontSize: 13,
              height: 15,
            }}
          >
            {errorMessage}
          </span>
          <input
            ref={inputRef}
            directory=""
            webkitdirectory=""
            type="file"
            style={{ display: 'none' }}
            onInput={(f) => {
              setErrorMessage('');
              const firstFilePath = f.currentTarget.files?.[0];
              if (!firstFilePath) {
                setErrorMessage('Not allowed to select empty directory.');
                return;
              }

              createSource(
                CreateSourceRequest.create({
                  uri: `file://${path.dirname(firstFilePath.path)}`,
                }),
              );
            }}
          />
          <Button
            onClick={() => {
              inputRef.current?.click();
            }}
          >
            Next
          </Button>
        </div>
      </div>
    </Modal>
  );
}

export default NewSourceModal;
