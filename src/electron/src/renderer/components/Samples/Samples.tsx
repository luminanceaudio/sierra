import React, { useEffect } from 'react';
import { useSamples, useSamplesCount } from '../../../api/samples';
import Search from '../Search/Search';
import Sample from './Sample';
import Paginate from '../Paginate/Paginate';
/* eslint-disable camelcase */
import {
  SortColumn_Enum,
  SortDirection_Enum,
} from '../../../proto/app/appbase';
import TableHead from '../Table/TableHead';

const pageSize = 15;

function Samples(): React.ReactElement {
  const [search, setSearch] = React.useState('');
  const [page, setPage] = React.useState(1);
  const [sortColumn, setSortColumn] = React.useState(SortColumn_Enum.Name);
  const [sortDirection, setSortDirection] = React.useState<SortDirection_Enum>(
    SortDirection_Enum.Asc,
  );
  const {
    data: samples,
    isLoading,
    refetch,
  } = useSamples(search, page, pageSize, sortColumn, sortDirection);
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
        }}
      >
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            gap: 15,
          }}
        >
          <table style={{ flex: 1, textAlign: 'left', fontSize: 14 }}>
            <thead>
              <tr>
                <th />
                <TableHead
                  name="Name"
                  column={SortColumn_Enum.Name}
                  setPage={setPage}
                  sortColumn={sortColumn}
                  setSortColumn={setSortColumn}
                  sortDirection={sortDirection}
                  setSortDirection={setSortDirection}
                />
                <TableHead
                  name="Duration"
                  column={SortColumn_Enum.Duration}
                  setPage={setPage}
                  sortColumn={sortColumn}
                  setSortColumn={setSortColumn}
                  sortDirection={sortDirection}
                  setSortDirection={setSortDirection}
                />
              </tr>
            </thead>
            <tbody>
              {isLoading
                ? null
                : samples?.data?.samples?.map((sample) => {
                    return <Sample key={sample.uri} sample={sample} />;
                  })}
            </tbody>
          </table>
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
