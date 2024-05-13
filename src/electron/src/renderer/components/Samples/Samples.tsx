import React from 'react';
import { useSamples } from '../../../api/samples';
import Sample from './Sample';

function Samples(): React.ReactElement {
  const { data: samples, isLoading } = useSamples();

  if (isLoading) {
    return <div />;
  }

  return (
    <div style={{ flex: 1, display: 'flex', flexDirection: 'column', gap: 15 }}>
      {samples?.data?.samples?.map((sample) => {
        return <Sample key={sample.uri} sample={sample} />;
      })}
    </div>
  );
}

export default Samples;
