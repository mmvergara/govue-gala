import axios, { AxiosError, type AxiosResponse } from "axios";
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

export const AxiosRequest = async <T>(
  url: string,
  method: "GET" | "POST" | "PUT" | "DELETE" = "GET",
  body?: any
): Promise<ApiRes<T>> => {
  try {
    let res: AxiosResponse<T> | null = null;
    if (method === "GET") {
      res = await AxiosInstance.get<T>(url);
    } else if (method === "POST") {
      res = await AxiosInstance.post<T>(url, body);
    } else if (method === "PUT") {
      res = await AxiosInstance.put<T>(url, body);
    } else if (method === "DELETE") {
      res = await AxiosInstance.delete<T>(url);
    }
    if (!res) throw new Error("AxiosResponse is null");
    return { data: res.data, error: "" };
  } catch (error) {
    const AxiosError = error as AxiosError;
    return { data: null as T, error: AxiosError.message };
  }
};
