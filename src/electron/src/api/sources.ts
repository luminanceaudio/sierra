import { AxiosResponse } from 'axios';
import {
  UseQueryOptions,
  useQuery,
  useMutation,
  UseMutationOptions,
} from './reactquery';
import {
  GetSourcesResponse,
  CreateSourceRequest,
  CreateSourceResponse,
} from '../proto/app/apprequests';
import axios from './axios';

export function useSources(
  options?: UseQueryOptions<AxiosResponse<GetSourcesResponse>>,
) {
  return useQuery<GetSourcesResponse>(
    async () => axios.get(`/api/v1/app/source`),
    {
      queryKey: ['sources'],
      ...options,
    },
  );
}

export function useCreateSource(
  options?: UseMutationOptions<
    CreateSourceRequest,
    AxiosResponse<CreateSourceResponse>
  >,
) {
  return useMutation<CreateSourceRequest, AxiosResponse<CreateSourceResponse>>(
    async (req: CreateSourceRequest) => axios.post('/api/v1/app/source', req),
    {
      mutationKey: ['sources'],
      ...options,
    },
  );
}
