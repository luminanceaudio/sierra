import { AxiosResponse } from 'axios';
import { UseQueryOptions, useQuery } from './reactquery';
import {
  QuerySamplesCountRequest,
  QuerySamplesCountResponse,
  QuerySamplesRequest,
  QuerySamplesResponse,
} from '../proto/app/apprequests';
import axios from './axios';
import { SortColumn_Enum, SortDirection_Enum } from '../proto/app/appbase';

export function useSamples(
  query: string,
  page: number,
  size: number,
  sortColumn: SortColumn_Enum,
  sortDirection: SortDirection_Enum,
  options?: UseQueryOptions<AxiosResponse<QuerySamplesResponse>>,
) {
  return useQuery<QuerySamplesResponse>(
    async () =>
      axios.post(
        `/api/v1/app/sample/query`,
        QuerySamplesRequest.create({
          query,
          page,
          size,
          sortColumn,
          sortDirection,
        }),
      ),
    {
      queryKey: ['samples', query, page, size, sortColumn, sortDirection],
      ...options,
    },
  );
}

export function useSamplesCount(
  query: string,
  options?: UseQueryOptions<AxiosResponse<QuerySamplesCountResponse>>,
) {
  return useQuery<QuerySamplesCountResponse>(
    async () =>
      axios.post(
        `/api/v1/app/sample/query/count`,
        QuerySamplesCountRequest.create({
          query,
        }),
      ),
    {
      queryKey: ['samples', 'count', query],
      ...options,
    },
  );
}
