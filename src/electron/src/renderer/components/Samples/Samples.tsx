import React, { useEffect } from 'react';
import path from 'path';
import { useSamples } from '../../../api/samples';
import config from '../../../config/config';
import Search from '../Search/Search';

function Samples(): React.ReactElement {
  const [search, setSearch] = React.useState('');
  const { data: samples, isLoading, refetch } = useSamples(search);

  useEffect(() => {
    (async () => {
      await refetch();
    })();
  }, [refetch, search]);

  if (isLoading) {
    return <div />;
  }

  const handleDragStart = (
    e: React.DragEvent<HTMLAudioElement>,
    uri: string,
  ) => {
    e.preventDefault();
    window.electron.startDrag(uri);
  };

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: 30 }}>
      <Search value={search} onChange={setSearch} />
      {samples?.data?.samples?.map((sample) => {
        const audioEndpoint = `
        ${config.API_URL}/api/v1/app/sample/load/${encodeURIComponent(
          sample.uri,
        )}`;

        return (
          <div
            key={sample.sha256}
            style={{
              display: 'flex',
              flex: 1,
              flexDirection: 'column',
              paddingTop: 10,
            }}
          >
            <span style={{ marginBottom: 10, color: '#393939' }}>
              {path.basename(sample.uri)}
            </span>
            {/* eslint-disable-next-line jsx-a11y/media-has-caption */}
            <div style={{ flex: 1, display: 'flex' }}>
              <audio
                controls
                preload="none"
                draggable
                onDragStart={(e) => handleDragStart(e, sample.uri)}
                style={{ flex: 1 }}
              >
                <source src={audioEndpoint} type={`audio/${sample.format}`} />
                Your browser does not support the audio element.
              </audio>
            </div>
          </div>
        );
      })}
    </div>
  );
}

export default Samples;
