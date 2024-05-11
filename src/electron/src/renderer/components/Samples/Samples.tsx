import React from 'react';
import path from 'path';
import { useSamples } from '../../../api/samples';
import config from '../../../config/config';

function Samples(): React.ReactElement {
  const { data: samples, isLoading } = useSamples();

  if (isLoading) {
    return <div />;
  }

  return (
    <div>
      {samples?.data?.samples?.map((sample) => {
        const audioEndpoint = `
        ${config.API_URL}/api/v1/app/audio/load/${encodeURIComponent(
          sample.uri,
        )}`;

        return (
          <div key={sample.sha256}>
            {/* eslint-disable-next-line jsx-a11y/media-has-caption */}
            <audio controls>
              <source src={audioEndpoint} type={`audio/${sample.format}`} />
              Your browser does not support the audio element.
            </audio>
          </div>
        );
      })}
    </div>
  );
}

export default Samples;
