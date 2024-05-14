import React, { useEffect } from 'react';
import { useSamples, useSamplesCount } from '../../../api/samples';
import Search from '../Search/Search';
import Sample from './Sample';
import Paginate from '../Paginate/Paginate';

const pageSize = 8;

function Samples(): React.ReactElement {
  const [search, setSearch] = React.useState('');
  const [page, setPage] = React.useState(1);
  const {
    data: samples,
    isLoading,
    refetch,
  } = useSamples(search, page, pageSize);
  const { data: totalCount } = useSamplesCount(search);

  useEffect(() => {
    (async () => {
      await refetch();
    })();
  }, [refetch, search, page]);

  useEffect(() => {
    setPage(1);
  }, [search]);

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
          {isLoading
            ? ''
            : samples?.data?.samples?.map((sample) => {
                return <Sample key={sample.uri} sample={sample} />;
              })}

          <Paginate
            page={page}
            onChangePage={setPage}
            totalItems={totalCount?.data?.count}
            pageSize={pageSize}
          />
        </div>
      </div>
    </div>
  );
}

export default Samples;
