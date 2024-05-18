import React, { useEffect } from 'react';
import { ChevronUpDownIcon } from '@heroicons/react/24/solid';
import { useSamples, useSamplesCount } from '../../../api/samples';
import Search from '../Search/Search';
import Sample from './Sample';
import Paginate from '../Paginate/Paginate';
import { StyledTableHead } from './Samples.style';
import Clickable from '../Clickable/Clickable';
/* eslint-disable camelcase */
import {
  SortColumn_Enum,
  SortDirection_Enum,
} from '../../../proto/app/appbase';

const pageSize = 8;

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
                <StyledTableHead />
                <StyledTableHead>Name</StyledTableHead>
                <StyledTableHead>
                  <Clickable
                    style={{
                      display: 'flex',
                      alignItems: 'center',
                      gap: 2,
                      fontWeight: 'inherit',
                      fontSize: 'inherit',
                    }}
                    onClick={() => {
                      setSortColumn(SortColumn_Enum.Duration);
                      setSortDirection(
                        sortDirection === SortDirection_Enum.Asc
                          ? SortDirection_Enum.Desc
                          : SortDirection_Enum.Asc,
                      );
                    }}
                  >
                    Duration
                    <ChevronUpDownIcon
                      width={15}
                      style={{ color: '#787878' }}
                    />
                  </Clickable>
                </StyledTableHead>
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
