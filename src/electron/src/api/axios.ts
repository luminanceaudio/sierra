import axios, { AxiosError, AxiosInstance } from 'axios';
import config from '../config/config';

function fetchClient() {
  return axios.create({
    baseURL: config.API_URL,
  });
}

const axiosInstance: AxiosInstance = fetchClient();

export function setAxiosToken(token: string, deleteTokenFn: () => void) {
  // Add an interceptor to inject the token in the Authorization header
  axiosInstance.interceptors.request.use((reqConfig) => {
    reqConfig.headers.Authorization = token ? `Bearer ${token}` : '';
    return reqConfig;
  });

  // Add an interceptor to handle 401 responses
  axiosInstance.interceptors.response.use(
    (response) => response,
    (error: AxiosError<{ code: string }>) => {
      if (error.response && error.response.status === 401) {
        deleteTokenFn();
      }
    },
  );
}

export default axiosInstance;
