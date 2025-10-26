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
  payload?: unknown;

  constructor(message: string, code: number, status: number, payload?: unknown) {
    super(message);
    this.name = "ApiError";
    this.code = code;
    this.status = status;
    this.payload = payload;
  }
}

type UnauthorizedHandler = (context: { status: number; code?: number; message?: string; path: string }) => void;

let unauthorizedHandler: UnauthorizedHandler | undefined;

export function configureHttp(options: { onUnauthorized?: UnauthorizedHandler }) {
  unauthorizedHandler = options.onUnauthorized;
}

export type HttpRequestOptions = RequestInit & {
  json?: unknown;
  skipAuth?: boolean;
};

function isApiEnvelope(value: unknown): value is ApiResponse<unknown> {
  return Boolean(value) && typeof value === "object" && "code" in (value as Record<string, unknown>) && "message" in (value as Record<string, unknown>);
}

function mergeHeaders(target: Headers, extra?: HeadersInit) {
  if (!extra) return;
  const additional = new Headers(extra);
  additional.forEach((value, key) => {
    target.set(key, value);
  });
}

function isSuccessCode(code: number) {
  return code === 200 || code === 0;
}

async function request<T>(path: string, init?: HttpRequestOptions): Promise<ApiResponse<T>> {
  const { json, skipAuth, headers: initHeaders, ...rest } = init ?? {};
  const headers = new Headers({
    "Content-Type": "application/json",
  });

  if (!skipAuth && typeof window !== "undefined") {
    const token = window.localStorage.getItem("token");
    if (token) {
      headers.set("Authorization", `Bearer ${token}`);
    }
  }

  mergeHeaders(headers, initHeaders);

  const requestInit: RequestInit = {
    ...rest,
    headers,
    body: json !== undefined ? JSON.stringify(json) : rest.body,
  };

  const res = await fetch(`${API_BASE_URL}${path}`, requestInit);
  const isJson = res.headers.get("content-type")?.includes("application/json");
  const raw = isJson ? await res.json().catch(() => ({})) : await res.text();

  if (!res.ok) {
    // Surface unauthorized state through the optional handler so the UI layer decides how to respond.
    if (res.status === 401) {
      unauthorizedHandler?.({
        status: res.status,
        path,
        code: isApiEnvelope(raw) ? Number(raw.code) : undefined,
        message: isApiEnvelope(raw) ? raw.message : undefined,
      });
      throw new ApiError("Unauthorized. Please log in again.", 401, res.status, raw);
    }

    if (isApiEnvelope(raw)) {
      const payload = raw as ApiResponse<unknown>;
      throw new ApiError(payload.message || `Request failed: ${res.status}`, Number(payload.code) || res.status, res.status, raw);
    }

    const fallbackMessage = typeof raw === "string" && raw ? raw : `Request failed: ${res.status}`;
    throw new ApiError(fallbackMessage, res.status, res.status, raw);
  }

  if (isApiEnvelope(raw)) {
    return raw as ApiResponse<T>;
  }

  return {
    code: 200,
    message: "success",
    data: (raw as unknown) as T,
  };
}

async function coreHttp<T>(path: string, init?: HttpRequestOptions): Promise<T> {
  const resp = await request<T>(path, init);
  if (!isSuccessCode(resp.code)) {
    throw new ApiError(resp.message || "Request error", resp.code, 200, resp);
  }

  // When the server reports success but omits `data`, fail fast to avoid leaking undefined through the typed API.
  if (resp.data === undefined) {
    throw new ApiError("Response body is empty.", resp.code, 200, resp);
  }

  return resp.data;
}

type MethodInit = Omit<HttpRequestOptions, "method" | "body" | "json">;

export interface HttpClient {
  <T>(path: string, init?: HttpRequestOptions): Promise<T>;
  get<T>(path: string, init?: MethodInit): Promise<T>;
  post<T>(path: string, body?: unknown, init?: MethodInit): Promise<T>;
  put<T>(path: string, body?: unknown, init?: MethodInit): Promise<T>;
  patch<T>(path: string, body?: unknown, init?: MethodInit): Promise<T>;
  delete<T>(path: string, init?: MethodInit): Promise<T>;
}

export const http: HttpClient = Object.assign(
  <T>(path: string, init?: HttpRequestOptions) => coreHttp<T>(path, init),
  {
    get<T>(path: string, init?: MethodInit) {
      return coreHttp<T>(path, { ...init, method: "GET" });
    },
    post<T>(path: string, body?: unknown, init?: MethodInit) {
      return coreHttp<T>(path, { ...init, method: "POST", json: body });
    },
    put<T>(path: string, body?: unknown, init?: MethodInit) {
      return coreHttp<T>(path, { ...init, method: "PUT", json: body });
    },
    patch<T>(path: string, body?: unknown, init?: MethodInit) {
      return coreHttp<T>(path, { ...init, method: "PATCH", json: body });
    },
    delete<T>(path: string, init?: MethodInit) {
      return coreHttp<T>(path, { ...init, method: "DELETE" });
    },
  }
);

export const httpRaw = request;
