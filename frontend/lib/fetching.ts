// Initialize Types

type ParsedResponse<T> =
  | { data: null; error: string | null }
  | {
      data: T;
      error: null;
    };
const MAIN_URL = "http://localhost:8080";

const getRequest = async <T>(url: string): Promise<ParsedResponse<T>> => {
  try {
    const res = await fetch(`${MAIN_URL}${url}`);
    const data = await res.json();
    return { data, error: null };
  } catch (e) {
    return { data: null, error: "Failed to fetch data" };
  }
};
