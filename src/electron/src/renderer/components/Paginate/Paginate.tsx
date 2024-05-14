import React from 'react';
import { ArrowLeftIcon, ArrowRightIcon } from '@heroicons/react/24/solid';
import Button from '../Button/Button';

type PaginateProps = {
  page: number;
  onChangePage: (page: number) => void;
  totalItems: number | undefined;
  pageSize: number;
};

function Paginate({
  page,
  onChangePage,
  totalItems,
  pageSize,
}: PaginateProps): React.ReactNode {
  let maxPage = 0;
  if (totalItems) {
    maxPage = Math.ceil(totalItems / pageSize);
  }

  return (
    <div
      style={{
        display: 'flex',
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
      }}
    >
      <div style={{ display: 'flex', flexDirection: 'column', width: 30 }}>
        {page === 1 ? (
          ''
        ) : (
          <Button type="secondary" onClick={() => onChangePage(page - 1)}>
            <ArrowLeftIcon width={13} />
          </Button>
        )}
      </div>

      <span>{totalItems ? `${page}/${maxPage}` : ''}</span>

      <div style={{ display: 'flex', flexDirection: 'column', width: 30 }}>
        {page === maxPage ? (
          ''
        ) : (
          <div style={{ display: 'flex', flexDirection: 'column' }}>
            <Button type="secondary" onClick={() => onChangePage(page + 1)}>
              <ArrowRightIcon width={13} />
            </Button>
          </div>
        )}
      </div>
    </div>
  );
}

export default Paginate;
