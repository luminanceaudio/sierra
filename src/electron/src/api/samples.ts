import { AxiosResponse } from 'axios';
import { UseQueryOptions, useQuery } from './reactquery';
import { GetSamplesResponse } from '../proto/app/apprequests';
import axios from './axios';

export function useSamples(
  search: string,
  options?: UseQueryOptions<AxiosResponse<GetSamplesResponse>>,
) {
  return useQuery<GetSamplesResponse>(
    async () =>
      axios.get(
        `/api/v1/app/sample/search${
          search ? `/${encodeURIComponent(search)}` : ''
        }`,
      ),
    {
      queryKey: ['samples'],
      ...options,
    },
  );
}
