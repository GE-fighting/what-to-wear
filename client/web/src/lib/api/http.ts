import { API_BASE_URL } from "@/lib/config/env";

export interface ApiResponse<T> {
  code: number;
  message: string;
  data?: T;
}

export interface ApiPageData<T> {
  items: T[];
  total: number;
  page: number;
  size: number;
  pages?: number;
}

export class ApiError extends Error {
  code: number;
  status: number;
  constructor(message: string, code: number, status: number) {
    super(message);
    this.name = "ApiError";
    this.code = code;
    this.status = status;
  }
}

async function request<T>(path: string, init?: RequestInit): Promise<ApiResponse<T>> {
  const token = typeof window !== "undefined" ? localStorage.getItem("token") : null;
  const res = await fetch(`${API_BASE_URL}${path}`, {
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(init?.headers || {}),
    },
    ...init,
  });

  const isJson = res.headers.get("content-type")?.includes("application/json");
  const raw = isJson ? await res.json().catch(() => ({})) : await res.text();

  // If not ok, try to surface backend message
  if (!res.ok) {
    if (isJson && raw && typeof raw === "object" && "message" in raw && "code" in raw) {
      const r = raw as ApiResponse<unknown>;
      throw new ApiError(r.message || `Request failed: ${res.status}`, Number(r.code) || res.status, res.status);
    }
    throw new ApiError(typeof raw === "string" && raw ? raw : `Request failed: ${res.status}`, res.status, res.status);
  }

  // Normal success path. Server uses unified envelope { code, message, data }
  if (isJson && raw && typeof raw === "object" && "code" in raw && "message" in raw) {
    return raw as ApiResponse<T>;
  }

  // Fallback: treat raw as direct data when server didn't wrap
  return {
    code: 200,
    message: "success",
    data: (raw as unknown) as T,
  };
}

export async function http<T>(path: string, init?: RequestInit): Promise<T> {
  const resp = await request<T>(path, init);
  if (resp.code !== 200) {
    throw new ApiError(resp.message || "Request error", resp.code, 200);
  }
  return resp.data as T;
}

// Also export the low-level request in case callers need message/code
export const httpRaw = request;
