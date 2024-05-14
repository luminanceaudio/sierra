import React from 'react';
import { PlayIcon, PauseIcon } from '@heroicons/react/24/solid';
import path from 'path';
import { StyledAudioPlayerBar } from './AudioPlayerBar.style';
import Button from '../Button/Button';
import { useAudioPlayer } from '../AudioPlayerProvider/AudioPlayerProvider';

function AudioPlayerBar() {
  const { audioUri, isPlaying, setIsPlaying } = useAudioPlayer();

  // useEffect(() => {
  //   console.log('PLAY');
  //   ref.current?.playAudioPromise();
  // }, [ref, audioUri]);

  return (
    <StyledAudioPlayerBar>
      <Button type="secondary" onClick={() => setIsPlaying(!isPlaying)}>
        {isPlaying ? <PauseIcon width={20} /> : <PlayIcon width={20} />}
      </Button>
      {path.basename(audioUri)}
    </StyledAudioPlayerBar>
  );
}

export default AudioPlayerBar;
