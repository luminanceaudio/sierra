import React from 'react';
import path from 'path';
import { PlayIcon, PauseIcon } from '@heroicons/react/24/solid';
import { PlusIcon, CheckIcon } from '@heroicons/react/24/outline';
import { Sample as SampleModel } from '../../../proto/app/appbase';
import Button from '../Button/Button';
import { useAudioPlayer } from '../AudioPlayerProvider/AudioPlayerProvider';
import { StyledProgress } from './Sample.style';

export type SampleProps = {
  sample: SampleModel;
};

function Sample({ sample }: SampleProps): React.ReactNode {
  const { audioUri, setAudioUri, isPlaying, setIsPlaying } = useAudioPlayer();

  const handleDragStart = (e: React.DragEvent<HTMLDivElement>, uri: string) => {
    e.preventDefault();
    window.electron.startDrag(uri);
  };

  const isCurrentAudio = audioUri === sample.uri;

  return (
    <tr
      key={sample.sha256}
      style={{
        position: 'relative',
      }}
      draggable
      onDragStart={(e) => handleDragStart(e, sample.uri)}
    >
      {/* eslint-disable-next-line jsx-a11y/media-has-caption */}
      <td style={{ width: '10%' }}>
        {isCurrentAudio && isPlaying ? (
          <StyledProgress
            style={{
              animationPlayState: 'running',
              animationDuration: `${sample.duration / 1000.0}s`,
            }}
          />
        ) : (
          ''
        )}
        <Button
          onClick={() => {
            if (!isCurrentAudio) {
              setAudioUri(sample.uri);
              return;
            }

            setIsPlaying(!isPlaying);
          }}
        >
          {isCurrentAudio && isPlaying ? (
            <PauseIcon width={15} />
          ) : (
            <PlayIcon width={15} />
          )}
        </Button>
      </td>

      <td
        style={{ color: isCurrentAudio ? '#d13482' : '#333333', width: '60%' }}
      >
        {path.basename(sample.uri)}
      </td>
      <td style={{ width: '15%' }}>{Math.ceil(sample.duration / 1000.0)}s</td>
      <td style={{ width: '15%' }}>
        <Button type="secondary" onClick={() => {}}>
          <PlusIcon width={15} />
        </Button>
      </td>
    </tr>
  );
}

export default Sample;
