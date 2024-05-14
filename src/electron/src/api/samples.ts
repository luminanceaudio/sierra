import { AxiosResponse } from 'axios';
import { UseQueryOptions, useQuery } from './reactquery';
import {
  GetSamplesCountResponse,
  GetSamplesResponse,
} from '../proto/app/apprequests';
import axios from './axios';

export function useSamples(
  query: string,
  page: number,
  size: number,
  options?: UseQueryOptions<AxiosResponse<GetSamplesResponse>>,
) {
  return useQuery<GetSamplesResponse>(
    async () =>
      axios.get(
        `/api/v1/app/sample/search/paginated/${page}${
          query ? `/${encodeURIComponent(query)}` : ''
        }?size=${size}`,
      ),
    {
      queryKey: ['samples', query],
      ...options,
    },
  );
}

export function useSamplesCount(
  query: string,
  options?: UseQueryOptions<AxiosResponse<GetSamplesCountResponse>>,
) {
  return useQuery<GetSamplesCountResponse>(
    async () =>
      axios.get(
        `/api/v1/app/sample/search/count${
          query ? `/${encodeURIComponent(query)}` : ''
        }`,
      ),
    {
      queryKey: ['samples', 'count', query],
      ...options,
    },
  );
}
