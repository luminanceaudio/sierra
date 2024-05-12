import React, { useEffect, useState } from 'react';
import { PersistQueryClientProvider as RQPersistQueryClientProvider } from '@tanstack/react-query-persist-client';
import { createSyncStoragePersister } from '@tanstack/query-sync-storage-persister';
import {
  MutationFunction,
  QueryClient,
  QueryFunction,
  QueryKey,
  useMutation as rqUseMutation,
  useQuery as rqUseQuery,
  UseQueryOptions as RQUseQueryOptions,
  UseMutationOptions as RQUseMutationOptions,
  useQueryClient,
} from '@tanstack/react-query';
import { AxiosResponse } from 'axios';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      gcTime: Infinity,
    },
  },
});

type PersistQueryClientProviderProps = {
  children: React.ReactNode;
};

export function PersistQueryClientProvider({
  children,
}: PersistQueryClientProviderProps): React.ReactElement {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  if (isClient) {
    const persistOptions = {
      persister: createSyncStoragePersister({
        storage: localStorage,
      }),
    };

    return (
      <RQPersistQueryClientProvider
        client={queryClient}
        persistOptions={persistOptions}
      >
        {children}
      </RQPersistQueryClientProvider>
    );
  }

  return <div />;
}

// Use wrapper function

export type UseQueryOptions<
  TResponse,
  TQueryKey extends QueryKey = QueryKey,
> = Omit<
  RQUseQueryOptions<TResponse, unknown, TResponse, TQueryKey>,
  'initialData'
> & { initialData?: () => undefined };

// useQuery is a thin react-query wrapper function
export function useQuery<TResponse, TQueryKey extends QueryKey = QueryKey>(
  queryFn: QueryFunction<AxiosResponse<TResponse>, TQueryKey>,
  options: UseQueryOptions<AxiosResponse<TResponse>, TQueryKey>,
) {
  return rqUseQuery<
    AxiosResponse<TResponse>,
    unknown,
    AxiosResponse<TResponse>,
    TQueryKey
  >({
    queryFn: async (ctx) => {
      const data = await queryFn(ctx);
      if (data == null) {
        throw new Error('failed fetching data');
      }

      if (data.status !== 200) {
        throw new Error(`failed fetching data, got status ${data.status}`);
      }

      return data;
    },
    ...options,
  });
}

export type UseMutationOptions<TRequest, TResponse> = Omit<
  RQUseMutationOptions<TResponse, unknown, TRequest, unknown>,
  'mutationFn'
>;

// useMutation is a thin react-query wrapper function
export function useMutation<TRequest, TResponse>(
  mutationFn: MutationFunction<TResponse, TRequest>,
  options?: UseMutationOptions<TRequest, TResponse>,
) {
  const usedQueryClient = useQueryClient();

  return rqUseMutation({
    ...options,
    mutationFn,
    onSettled: async (data, error, variables, context) => {
      await usedQueryClient.invalidateQueries({
        queryKey: options?.mutationKey,
      });
      if (options?.onSettled)
        await options?.onSettled(data, error, variables, context);
    },
  });
}
