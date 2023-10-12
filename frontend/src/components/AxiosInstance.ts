import axios, { AxiosError } from "axios";
// Initialize Axios Instance
axios.defaults.withCredentials = true;
const AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL as string,
  withCredentials: true,
});

type ApiRes<T> = {
  data: T;
  error: string;
};

export const AxiosGet = async <T>(url: string): Promise<ApiRes<T>> => {
  try {
    const res = await AxiosInstance.get<ApiRes<T>>(url);
    return { data: res.data, error: "" };
  } catch (err) {
    const AxiosError = err as AxiosError;
    return { data: null as T, error: AxiosError.message };
  }
};

export const AxiosPost = async <T>(
  url: string,
  body: any
): Promise<ApiRes<T>> => {
  try {
    const res = await AxiosInstance.post<ApiRes<T>>(url, body);
    return { data: res.data.data!, error: "" };
  } catch (err) {
    const AxiosError = err as AxiosError;
    return { data: null as T, error: AxiosError.message };
  }
};
