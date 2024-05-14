import React from 'react';
import path from 'path';
import { PlayIcon, PauseIcon } from '@heroicons/react/24/solid';
import { Sample as SampleModel } from '../../../proto/app/appbase';
import Button from '../Button/Button';
import { useAudioPlayer } from '../AudioPlayerProvider/AudioPlayerProvider';

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
    <div
      key={sample.sha256}
      style={{
        display: 'flex',
        flex: 1,
        gap: 10,
        alignItems: 'center',
        border: '3px solid #f1f1f1',
        borderRadius: 50,
      }}
      draggable
      onDragStart={(e) => handleDragStart(e, sample.uri)}
    >
      {/* eslint-disable-next-line jsx-a11y/media-has-caption */}
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
          <PauseIcon width={20} />
        ) : (
          <PlayIcon width={20} />
        )}
      </Button>
      <span style={{ color: '#333333', fontSize: 13 }}>
        {path.basename(sample.uri)}
      </span>
    </div>
  );
}

export default Sample;
