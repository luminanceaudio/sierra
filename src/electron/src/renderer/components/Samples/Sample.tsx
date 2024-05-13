import React from 'react';
import config from '../../../config/config';
import { Sample as SampleModel } from '../../../proto/app/appbase';

type SampleProps = {
  sample: SampleModel;
};

function Sample({ sample }: SampleProps): React.ReactNode {
  const audioEndpoint = `
        ${config.API_URL}/api/v1/app/sample/load/${encodeURIComponent(
          sample.uri,
        )}`;

  const handleDragStart = (
    e: React.DragEvent<HTMLAudioElement>,
    uri: string,
  ) => {
    e.preventDefault();
    window.electron.startDrag(uri);
  };

  return (
    <div key={sample.sha256}>
      {/* eslint-disable-next-line jsx-a11y/media-has-caption */}
      <audio
        controls
        preload="none"
        draggable
        onDragStart={(e) => handleDragStart(e, sample.uri)}
      >
        <source src={audioEndpoint} type={`audio/${sample.format}`} />
        Your browser does not support the audio element.
      </audio>
    </div>
  );
}

export default Sample;
