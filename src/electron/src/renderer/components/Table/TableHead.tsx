import { ChevronUpIcon, ChevronDownIcon } from '@heroicons/react/24/solid';
import React from 'react';
import Clickable from '../Clickable/Clickable';
import { StyledTableHead } from './TableHead.style';
import {
  SortColumn_Enum,
  SortDirection_Enum,
} from '../../../proto/app/appbase';

export type TableHeadProps = {
  name: string;
  column: SortColumn_Enum;

  sortColumn: SortColumn_Enum;
  setSortColumn: (column: SortColumn_Enum) => void;

  sortDirection: SortDirection_Enum;
  setSortDirection: (direction: SortDirection_Enum) => void;

  setPage: (page: number) => void;
};

function TableHead({
  name,
  column,
  sortColumn,
  setSortColumn,
  sortDirection,
  setSortDirection,
  setPage,
}: TableHeadProps) {
  function changeSort() {
    if (column !== sortColumn) {
      setSortColumn(column);
      setSortDirection(SortDirection_Enum.Asc);
    } else {
      setSortDirection(
        sortDirection === SortDirection_Enum.Asc
          ? SortDirection_Enum.Desc
          : SortDirection_Enum.Asc,
      );
    }

    setPage(1);
  }

  let icon: React.ReactNode | null = null;

  if (sortColumn === column) {
    if (sortDirection === SortDirection_Enum.Desc) {
      icon = <ChevronDownIcon width={10} style={{ color: '#787878' }} />;
    } else {
      icon = <ChevronUpIcon width={10} style={{ color: '#787878' }} />;
    }
  }

  if (sortColumn !== column) {
    icon = null;
  }

  return (
    <StyledTableHead>
      <Clickable onClick={() => changeSort()}>
        {name}
        {icon}
      </Clickable>
    </StyledTableHead>
  );
}

export default TableHead;
