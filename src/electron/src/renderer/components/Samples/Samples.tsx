import React, { useEffect } from 'react';
import { useSamples } from '../../../api/samples';
import Search from '../Search/Search';
import Sample from './Sample';

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

  return (
    <div
      style={{
        display: 'flex',
        flex: 1,
        gap: 30,
        flexDirection: 'column',
      }}
    >
      <div style={{ display: 'flex', flexDirection: 'column' }}>
        <Search value={search} onChange={setSearch} />
      </div>
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            width: '100%',
            maxWidth: 300,
            gap: 15,
          }}
        >
          {samples?.data?.samples?.map((sample) => {
            return <Sample key={sample.uri} sample={sample} />;
          })}
        </div>
      </div>
    </div>
  );
}

export default Samples;
