// Initialize Types

type ParsedResponse<T> =
  | { data: null; error: string }
  | {
      data: T;
      error: null;
    };
const MAIN_URL = "http://localhost:8080";

export const getRequest = async <T>(
  url: string
): Promise<ParsedResponse<T>> => {
  try {
    const res = await fetch(`${MAIN_URL}${url}`);
    const data = (await res.json()) as T;
    return { data: data, error: null };
  } catch (e) {
    return { data: null, error: "Failed to fetch data" };
  }
};

export const postRequest = async <T>(
  url: string,
  body: Record<string, unknown>
): Promise<ParsedResponse<T>> => {
  try {
    const res = await fetch(`${MAIN_URL}${url}`, {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = (await res.json()) as T;
    return { data: data, error: null };
  } catch (e) {
    return { data: null, error: "Failed to fetch data" };
  }
};

export const putRequest = async <T>(
  url: string,
  body: Record<string, unknown>
): Promise<ParsedResponse<T>> => {
  try {
    const res = await fetch(`${MAIN_URL}${url}`, {
      method: "PUT",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = (await res.json()) as T;
    return { data: data, error: null };
  } catch (e) {
    return { data: null, error: "Failed to fetch data" };
  }
};
